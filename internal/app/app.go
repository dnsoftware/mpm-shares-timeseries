package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/dnsoftware/mpm-miners-processor/pkg/certmanager"
	jwtauth "github.com/dnsoftware/mpm-miners-processor/pkg/jwt"
	"github.com/dnsoftware/mpm-save-get-shares/pkg/logger"
	"github.com/dnsoftware/mpm-save-get-shares/pkg/utils"
	"github.com/golang-migrate/migrate/v4"
	"google.golang.org/grpc"

	"github.com/dnsoftware/mpm-shares-timeseries/config"
	pb "github.com/dnsoftware/mpm-shares-timeseries/internal/adapter/grpc"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/adapter/grpc/proto"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/constants"
	clickhouse2 "github.com/dnsoftware/mpm-shares-timeseries/internal/infrastructure/clickhouse"
	"github.com/dnsoftware/mpm-shares-timeseries/pkg/clickhouseconn"
)

func Run(ctx context.Context, cfg config.Config) {
	basePath, err := utils.GetProjectRoot(constants.ProjectRootAnchorFile)

	filePath, err := logger.GetLoggerMainLogPath()
	if err != nil {
		panic("Bad logger init: " + err.Error())
	}
	logger.InitLogger(logger.LogLevelDebug, filePath)

	/********* Инициализация трассировщика **********/
	/********* КОНЕЦ Инициализация трассировщика **********/

	// Подключение к базе данных ClickHouse
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: cfg.ClickhouseAddr,
		Auth: clickhouse.Auth{
			Database: "default",
			Username: cfg.ClickhouseUsername,
			Password: cfg.ClickhousePassword,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Debug: true,
	})
	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Проверка подключения
	err = conn.Ping(ctx)
	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	// Укажите путь к миграциям и строку подключения к базе данных
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s/%s", "default", "", cfg.ClickhouseAddr[0], "default")
	m, err := migrate.New(
		"file://"+basePath+"/"+constants.MigrationDir,
		dsn,
	)
	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	// Сброс миграций
	m.Force(-1)

	// Применить миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Log().Fatal(err.Error())
	}
	log.Println("Миграции успешно применены")

	conn.Close()

	// Подключаемся к mpmhouse
	conn, err = clickhouseconn.NewClickhouseConnect(clickhouseconn.Config{
		Addr:             cfg.ClickhouseAddr,
		Database:         cfg.ClickhouseDatabase,
		Username:         cfg.ClickhouseUsername,
		Password:         cfg.ClickhousePassword,
		MaxExecutionTime: 10,
	})
	if err != nil {
		logger.Log().Fatal(err.Error())
	}
	defer conn.Close()

	// Проверка подключения
	err = conn.Ping(ctx)
	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	cfgStore := clickhouse2.ShareStorageConfig{
		Conn:        conn,
		ClusterName: "clickhouse_cluster",
		Database:    "mpmhouse",
	}
	store, err := clickhouse2.NewClickhouseShareStorage(cfgStore)
	if err != nil {
		logger.Log().Fatal(err.Error())
	}

	jwt := jwtauth.NewJWTServiceSymmetric(cfg.JWTServiceName, cfg.JWTValidServices, cfg.JWTSecret, 1)

	certManager, err := certmanager.NewCertManager(basePath + "/certs")
	if err != nil {
		logger.Log().Fatal("NewCertManager error: " + err.Error())
	}

	// Запускаем сервер на определенном порту
	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		logger.Log().Fatal(fmt.Sprintf("Failed to listen: %v", err))
	}

	// Поднимаем gRPC-сервер в фоновом процессе
	serverReady := make(chan struct{})
	var grpcServer *grpc.Server
	go func() {
		serverCreds, err := certManager.GetServerCredentials()
		if err != nil {
			logger.Log().Fatal("GetServerCredentials error: " + err.Error())
		}

		interceptor := jwt.GetValidateInterceptor()
		grpcServer = grpc.NewServer(grpc.UnaryInterceptor(interceptor), grpc.Creds(*serverCreds))
		shareStorage, err := pb.NewGRPCServer(store)
		if err != nil {
			logger.Log().Fatal("NewGRPCServer error: " + err.Error())
		}

		proto.RegisterSharesServiceServer(grpcServer, shareStorage)

		close(serverReady) // Уведомляем, что сервер готов
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	<-serverReady // Ждем, пока сервер отправит сигнал готовности (вычитываем пустое значение после закрытия канала)

	// Настройка graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down gRPC server...")

	// Останавливаем сервер
	grpcServer.GracefulStop()
	logger.Log().Info("gRPC server stopped")
}

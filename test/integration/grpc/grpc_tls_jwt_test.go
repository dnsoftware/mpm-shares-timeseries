package grpc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/dnsoftware/mpm-save-get-shares/pkg/logger"
	"github.com/dnsoftware/mpm-save-get-shares/pkg/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"

	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	pb "github.com/dnsoftware/mpm-shares-timeseries/internal/adapter/grpc"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/adapter/grpc/proto"
	//"github.com/dnsoftware/mpm-shares-timeseries/internal/constants"
	"github.com/dnsoftware/mpm-miners-processor/pkg/certmanager"
	jwt2 "github.com/dnsoftware/mpm-miners-processor/pkg/jwt"

	"github.com/dnsoftware/mpm-shares-timeseries/config"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/constants"
	clickhouse2 "github.com/dnsoftware/mpm-shares-timeseries/internal/infrastructure/clickhouse"
	"github.com/dnsoftware/mpm-shares-timeseries/pkg/clickhouseconn"
)

const bufSize = 1024 * 1024 // Размер буфера для соединений в памяти

var lis *bufconn.Listener

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial() // Возвращает соединение внутри процесса
}

// тестируем mTLS соединение между клиентом и сервером с использованием JWT авторизации
// Должен быть запущен кластер Clickhouse
func TestTLSJWTTest(t *testing.T) {
	basePath, err := utils.GetProjectRoot(constants.ProjectRootAnchorFile)
	require.NoError(t, err)
	configFile := basePath + "/config_example.yaml"
	envFile := basePath + "/.env_example"

	cfg, err := config.New(configFile, envFile)
	require.NoError(t, err)

	// Создаем буферизованный listener
	lis = bufconn.Listen(bufSize)
	serverReady := make(chan struct{})

	filePath, err := logger.GetLoggerTestLogPath()
	require.NoError(t, err)
	logger.InitLogger(logger.LogLevelDebug, filePath)

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
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Проверка подключения
	err = conn.Ping(ctx)
	require.NoError(t, err)

	// Укажите путь к миграциям и строку подключения к базе данных
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s/%s", "default", "", cfg.ClickhouseAddr[0], "default")
	m, err := migrate.New(
		"file://"+basePath+"/"+constants.MigrationDir,
		dsn,
	)
	require.NoError(t, err)

	// Сброс миграций
	m.Force(-1)

	// Применить миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		require.NoError(t, err)
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
	require.NoError(t, err)
	defer conn.Close()

	// Проверка подключения
	err = conn.Ping(ctx)
	require.NoError(t, err)

	cfgStore := clickhouse2.ShareStorageConfig{
		Conn:        conn,
		ClusterName: "clickhouse_cluster",
		Database:    "mpmhouse",
	}
	store, err := clickhouse2.NewClickhouseShareStorage(cfgStore)
	require.NoError(t, err)

	jwt := jwt2.NewJWTServiceSymmetric("normalizer", []string{"normalizer"}, "jwtsecret", 60)

	path, err := utils.GetProjectRoot(".env_example")
	require.NoError(t, err)
	certManager, err := certmanager.NewCertManager(path + "/certs")
	require.NoError(t, err)

	// Поднимаем gRPC-сервер в фоновом процессе
	go func() {
		serverCreds, err := certManager.GetServerCredentials()
		require.NoError(t, err)

		interceptor := jwt.GetValidateInterceptor()
		grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor), grpc.Creds(*serverCreds))
		shareStorage, err := pb.NewGRPCServer(store)
		require.NoError(t, err)
		proto.RegisterSharesServiceServer(grpcServer, shareStorage)
		close(serverReady) // Уведомляем, что сервер готов
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	<-serverReady // Ждем, пока сервер отправит сигнал готовности (вычитываем пустое значение после закрытия канала)

	// Создаем контекст с тайм-аутом
	ctx, cancel2 := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel2()

	// Полномочия для TLS соединения
	clientCreds, err := certManager.GetClientCredentials()
	require.NoError(t, err)

	// Создаем gRPC соединение через `NewClient`
	connDial, err := grpc.DialContext(ctx,
		"bufnet",                          // Адрес символический, используется только для идентификации (т.к. используем bufconn в тестах)
		grpc.WithContextDialer(bufDialer), // Указываем кастомный диалер
		// grpc.WithTransportCredentials(insecure.NewCredentials()), // Отключаем TLS
		grpc.WithTransportCredentials(*clientCreds), // используем TLS
		grpc.WithUnaryInterceptor(jwt.GetClientInterceptor()),
	)
	if err != nil {
		t.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer connDial.Close()

	// Создаем клиента
	client := proto.NewSharesServiceClient(connDial)
	_ = client

	// Генерация токена
	token, err := jwt.GetActualToken()
	require.NoError(t, err)

	// Добавление токена в метаданные
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	numShares := 10
	shares, err := GenerateProtoShares(numShares, 100)
	require.NoError(t, err)

	start := time.Now().UnixMicro()

	resp, err := client.AddSharesBatch(ctx, &proto.AddSharesBatchRequest{
		Shares: shares,
	})
	require.NoError(t, err)

	require.Equal(t, int64(numShares), resp.AddedCount)

	end := time.Now().UnixMicro()
	require.NoError(t, err)

	fmt.Println(fmt.Sprintf("%.6f", float64(end-start)/1000000))

}

// Генерация среза шар
// Задержка в миллисекундах между генерацией шар (нужно чтобы время у шар различалось)
func GenerateProtoShares(n int, msecDelay time.Duration) ([]*proto.Share, error) {
	shares := make([]*proto.Share, 0, n)
	servers := []string{"EU-HSHP-ALPH-1", "RU-HSHP-DNX-1", "US-OVH-KAS-1"}
	coinIDs := map[string]int64{
		"EU-HSHP-ALPH-1": 4,
		"RU-HSHP-DNX-1:": 10,
		"US-OVH-KAS-1":   8,
	}
	rewardMethods := []string{"PPS", "PPLNS", "SOLO"}

	for i := 0; i < n; i++ {
		if msecDelay > 0 {
			time.Sleep(msecDelay * time.Millisecond)
		}

		// Выбираем случайный индекс
		randomIndex, err := SecureRandomInt(0, len(servers))
		if err != nil {
			return nil, err
		}
		randWorkerID, err := SecureRandomInt(1, 10000)
		if err != nil {
			return nil, err
		}
		randWalletID, err := SecureRandomInt(1, 1000)
		if err != nil {
			return nil, err
		}
		difficulty, err := SecureRandomFloat64(0.000001, 10.0)
		if err != nil {
			return nil, err
		}
		delta, _ := SecureRandomFloat64(0.000001, 100.0-difficulty)
		if err != nil {
			return nil, err
		}
		shareFloatDif := difficulty + delta
		if err != nil {
			return nil, err
		}

		nonce, err := GenerateRandomHexString(64)
		if err != nil {
			return nil, err
		}
		randomMethod, err := SecureRandomInt(0, len(rewardMethods))
		if err != nil {
			return nil, err
		}
		shareCost, err := SecureRandomFloat64(0.000001, 1)
		if err != nil {
			return nil, err
		}
		// Получаем случайный ключ и значение
		randomServer := servers[randomIndex]
		share := proto.Share{
			Uuid:         uuid.New().String(),
			ServerId:     randomServer,
			CoinId:       coinIDs[randomServer],
			WorkerId:     int64(randWorkerID),
			WalletId:     int64(randWalletID),
			ShareDate:    time.Now().UnixMilli(),
			Difficulty:   fmt.Sprintf("%v", difficulty),
			ShareDif:     fmt.Sprintf("%v", shareFloatDif),
			Nonce:        nonce,
			IsSolo:       false,
			RewardMethod: rewardMethods[randomMethod],
			Cost:         fmt.Sprintf("%v", shareCost),
		}

		shares = append(shares, &share)
	}

	return shares, nil
}

func SecureRandomInt(min, max int) (int, error) {
	if min >= max {
		return 0, fmt.Errorf("invalid range: min must be less than max")
	}

	// Вычисляем разницу между max и min
	rangeBig := big.NewInt(int64(max - min))

	// Генерируем случайное большое число в диапазоне [0, rangeBig)
	nBig, err := rand.Int(rand.Reader, rangeBig)
	if err != nil {
		return 0, err
	}

	// Преобразуем большое число в int и добавляем min, чтобы получить число в диапазоне [min, max)
	return int(nBig.Int64()) + min, nil
}

func SecureRandomFloat64(min, max float64) (float64, error) {
	// Разница между max и min
	rangeFloat := max - min
	if rangeFloat <= 0 {
		return 0, fmt.Errorf("invalid range: min must be less than max")
	}

	// Случайное большое число в диапазоне [0, 1)
	nBig, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0, err
	}

	// Преобразуем большое число в float64
	n := float64(nBig.Int64()) / (1 << 53)

	// Масштабируем в диапазон [min, max)
	return min + n*rangeFloat, nil
}

// GenerateRandomHexString генерирует случайную строку длиной n (в символах) из шестнадцатеричных символов.
func GenerateRandomHexString(length int) (string, error) {
	if length%2 != 0 {
		return "", fmt.Errorf("length must be an even number")
	}

	// Создаем срез байтов половинного размера от длины строки
	bytes := make([]byte, length/2)

	// Заполняем срез случайными байтами
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Конвертируем байты в шестнадцатеричную строку
	return hex.EncodeToString(bytes), nil
}

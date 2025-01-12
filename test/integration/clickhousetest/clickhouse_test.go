package clickhousetest

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/dnsoftware/mpm-save-get-shares/pkg/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/require"

	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/dnsoftware/mpm-shares-timeseries/config"
	clickhouse2 "github.com/dnsoftware/mpm-shares-timeseries/internal/adapter/clickhouse"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/constants"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/entity"
	"github.com/dnsoftware/mpm-shares-timeseries/pkg/clickhouseconn"
)

// Должен быть запущен кластер Clickhouse
func TestAddShare(t *testing.T) {
	basePath, err := utils.GetProjectRoot(constants.ProjectRootAnchorFile)
	require.NoError(t, err)
	configFile := basePath + "/config_example.yaml"
	envFile := basePath + "/.env_example"

	cfg, err := config.New(configFile, envFile)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

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
		ClickhouseAddr:             cfg.ClickhouseAddr,
		ClickhouseDatabase:         cfg.ClickhouseDatabase,
		ClickhouseUsername:         cfg.ClickhouseUsername,
		ClickhousePassword:         cfg.ClickhousePassword,
		ClickhouseMaxExecutionTime: 1,
	})
	require.NoError(t, err)
	defer conn.Close()

	// Проверка подключения
	err = conn.Ping(ctx)
	require.NoError(t, err)

	cfgStore := clickhouse2.ShareStorageConfig{
		Conn:        conn,
		ClusterName: "clickhouse_cluster",
	}
	store, err := clickhouse2.NewClickhouseShareStorage(cfgStore)
	require.NoError(t, err)

	ctx = context.Background()
	formattedTime := time.Now().Format("2006-01-02 15:04:05.000")
	share := entity.Share{
		UUID:         "9e7b8188-01a6-4989-8805-e4337e31195a",
		ServerID:     "EU-HSHP-ALPH-1",
		CoinID:       4,
		WorkerID:     4,
		WalletID:     2,
		ShareDate:    formattedTime,
		Difficulty:   "0.0089410000",
		Sharedif:     "5.1467700000",
		Nonce:        "9c44010001030201010202030400040402040304915711c0",
		IsSolo:       false,
		RewardMethod: "PPLNS",
		Cost:         "0.00124",
	}

	err = store.AddShare(ctx, share)
	require.NoError(t, err)

}

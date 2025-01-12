package clickhouseconn

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Config struct {
	ClickhouseAddr             []string // список нод
	ClickhouseDatabase         string   // назхвание БД
	ClickhouseUsername         string   // имя пользователя базы
	ClickhousePassword         string   // пароль пользователя
	ClickhouseMaxExecutionTime int
}

func NewClickhouseConnect(cfg Config) (driver.Conn, error) {

	maxExecutionTime := 60
	if cfg.ClickhouseMaxExecutionTime > 0 {
		maxExecutionTime = cfg.ClickhouseMaxExecutionTime
	}

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: cfg.ClickhouseAddr,
		Auth: clickhouse.Auth{
			Database: cfg.ClickhouseDatabase,
			Username: cfg.ClickhouseUsername,
			Password: cfg.ClickhousePassword,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Settings: map[string]interface{}{
			"max_execution_time": maxExecutionTime,
		},
		MaxOpenConns: 10,
		MaxIdleConns: 5,
		Debug:        true,
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

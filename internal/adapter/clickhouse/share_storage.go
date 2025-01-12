package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/dnsoftware/mpm-shares-timeseries/internal/entity"
)

type ShareStorageConfig struct {
	Conn        driver.Conn
	ClusterName string
}

type ClickhouseShareStorage struct {
	conn        driver.Conn
	clusterName string
}

func NewClickhouseShareStorage(cfg ShareStorageConfig) (*ClickhouseShareStorage, error) {
	s := &ClickhouseShareStorage{
		conn:        cfg.Conn,
		clusterName: cfg.ClusterName,
	}
	return s, nil
}

// AddShare Добавление единичной шары (для теста, в основном коде не используется, используется пакетная вставка)
func (c *ClickhouseShareStorage) AddShare(ctx context.Context, share entity.Share) error {

	query := `INSERT INTO mpmhouse.shares (uuid, server_id, coin_id, worker_id, wallet_id, share_date, difficulty, sharedif, nonce, is_solo, reward_method, cost) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	//difficulty
	//sharedif
	//cost

	if err := c.conn.Exec(ctx, query, share.UUID, share.ServerID, share.CoinID, share.WorkerID, share.WalletID, share.ShareDate, share.Difficulty, share.Sharedif, share.Nonce, share.IsSolo, share.RewardMethod, share.Cost); err != nil {
		return err
	}

	return nil
}

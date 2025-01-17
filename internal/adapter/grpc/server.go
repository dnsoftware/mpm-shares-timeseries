package grpc

import (
	"context"

	"github.com/dnsoftware/mpm-shares-timeseries/internal/adapter/grpc/proto"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/entity"
	"github.com/dnsoftware/mpm-shares-timeseries/internal/infrastructure/clickhouse"
)

type GRPCServer struct {
	proto.UnimplementedSharesServiceServer
	store *clickhouse.ClickhouseShareStorage
}

func NewGRPCServer(store *clickhouse.ClickhouseShareStorage) (*GRPCServer, error) {
	s := &GRPCServer{
		store: store,
	}

	return s, nil
}

func (s *GRPCServer) AddSharesBatch(ctx context.Context, in *proto.AddSharesBatchRequest) (*proto.AddSharesBatchResponse, error) {
	shares := make([]entity.Share, 0, len(in.Shares))
	for _, sh := range in.Shares {
		newShare := entity.Share{
			UUID:         sh.Uuid,
			ServerID:     sh.ServerId,
			CoinID:       sh.CoinId,
			WorkerID:     sh.WorkerId,
			WalletID:     sh.WalletId,
			ShareDate:    sh.ShareDate,
			Difficulty:   sh.Difficulty,
			Sharedif:     sh.ShareDif,
			Nonce:        sh.Nonce,
			IsSolo:       sh.IsSolo,
			RewardMethod: sh.RewardMethod,
			Cost:         sh.Cost,
		}

		shares = append(shares, newShare)
	}

	err := s.store.AddSharesBatch(ctx, shares)
	if err != nil {
		return &proto.AddSharesBatchResponse{AddedCount: int64(0)}, err
	}

	resp := &proto.AddSharesBatchResponse{AddedCount: int64(len(shares))}

	return resp, err
}

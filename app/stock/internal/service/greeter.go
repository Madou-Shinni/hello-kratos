package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	stockV1 "helloword/api/stock/v1"
	v1 "helloword/api/stock/v1"
	"helloword/app/stock/internal/biz"
)

var stock = 20

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedStockServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) DeductStock(ctx context.Context, req *stockV1.DeductStockRequest) (*emptypb.Empty, error) {
	if stock < int(req.Stock) {
		// 不重试
		log.Info("库存不足: %v", stock)
		return nil, errors.New(409, "库存不足", "库存不足")
	}

	// 重试
	if req.Stock > 10 {
		log.Info("超时重试: %v", stock)
		return nil, errors.New(425, "超时", "超时")
	}

	// 扣减库存成功
	stock -= int(req.Stock)
	log.Info("扣减库存成功 :%v", stock)

	return &emptypb.Empty{}, nil
}

func (s *GreeterService) AddStock(ctx context.Context, req *stockV1.IncreaseStockRequest) (*emptypb.Empty, error) {
	stock += int(req.Stock)
	log.Infof("增加库存成功: %v", stock)
	return nil, nil
}

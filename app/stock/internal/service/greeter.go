package service

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	stockV1 "helloword/api/stock/v1"
	v1 "helloword/api/stock/v1"
	"helloword/app/stock/internal/biz"
)

var stock = 20
var integral = 20

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedStockServer

	uc *biz.GreeterUsecase
	db *gorm.DB
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, db *gorm.DB) *GreeterService {
	return &GreeterService{uc: uc, db: db}
}

func (s *GreeterService) DeductStock(ctx context.Context, req *stockV1.DeductStockRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Error("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	err = grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		if stock < int(req.Stock) {
			// 不重试，库存不足，不进行补偿
			log.Info("库存不足: %v", stock)
			return errors.New(409, "发生错误", "发生错误")
		}

		// 扣减库存成功
		stock -= int(req.Stock)
		log.Info("扣减库存成功 :%v", stock)
		// 发生错误 需要补偿
		db.Exec("update stock set stock = stock - ? WHERE id = 1", req.Stock)
		return nil
	})
	if err != nil {
		log.Error("dtm error: %v", err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *GreeterService) IncreaseStock(ctx context.Context, req *stockV1.IncreaseStockRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Error("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		stock += int(req.Stock)
		log.Infof("增加库存成功: %v", stock)
		db.Exec("update stock set stock = stock + ? WHERE id = 1", req.Stock)
		return nil
	})
	return &emptypb.Empty{}, nil
}

func (s *GreeterService) DeductIntegral(ctx context.Context, req *stockV1.DeductIntegralRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Error("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	err = grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		if integral < int(req.Integral) {
			// 不重试，积分不足，不进行补偿
			log.Info("积分不足: %v", integral)
			db.Rollback()
			return errors.New(409, "发生错误", "发生错误")
		}

		// 扣减积分成功
		integral -= int(req.Integral)
		log.Info("扣减积分成功 :%v", integral)
		// 发生错误 需要补偿
		db.Exec("update integral set integral = integral - ? WHERE id = 1", req.Integral)
		return errors.New(409, "发生错误", "发生错误")
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *GreeterService) IncreaseIntegral(ctx context.Context, req *stockV1.IncreaseIntegralRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Error("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		integral += int(req.Integral)
		log.Infof("增加积分成功: %v", integral)
		db.Exec("update integral set integral = integral + ? WHERE id = 1", req.Integral)
		return nil
	})
	return &emptypb.Empty{}, nil
}

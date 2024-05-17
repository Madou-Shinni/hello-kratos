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
	"helloword/model"
)

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
		log.Errorf("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	var stock model.Stock
	err = grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		s.db.First(&stock, 1)
		if stock.Stock < int(req.Stock) {
			// 不重试，库存不足，不进行补偿
			log.Infof("库存不足: %v", stock.Stock)
			return errors.New(409, "发生错误", "发生错误")
		}

		// 扣减库存成功
		db.Exec("update stock set stock = stock - ? WHERE id = 1", req.Stock)
		return errors.New(409, "发生错误", "发生错误")
		//return nil
	})
	if err != nil {
		s.db.Model(&model.Stock{}).First(&stock, 1)
		log.Errorf("扣减库存发生错误: %v", stock.Stock)
		return nil, err
	}

	s.db.First(&stock, 1)
	log.Infof("扣减库存成功: %v", stock.Stock)

	return &emptypb.Empty{}, nil
}

func (s *GreeterService) IncreaseStock(ctx context.Context, req *stockV1.IncreaseStockRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Errorf("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		db.Exec("update stock set stock = stock + ? WHERE id = 1", req.Stock)
		return nil
	})

	var stock model.Stock
	s.db.First(&stock, 1)
	log.Infof("增加库存成功: %v", stock.Stock)

	return &emptypb.Empty{}, nil
}

func (s *GreeterService) DeductIntegral(ctx context.Context, req *stockV1.DeductIntegralRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Error("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	var integral model.Integral
	err = grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		s.db.First(&integral, 1)
		if integral.Integral < int(req.Integral) {
			// 不重试，积分不足
			log.Infof("积分不足: %v", integral)
			return errors.New(409, "发生错误", "发生错误")
		}

		// 扣减积分
		db.Exec("update integral set integral = integral - ? WHERE id = 1", req.Integral)
		//return errors.New(409, "发生错误", "发生错误")
		return nil
	})
	if err != nil {
		var first model.Integral
		s.db.Model(&model.Integral{}).First(&first, 1)
		log.Errorf("扣减积分发生错误: %v", first.Integral)
		return nil, err
	}

	s.db.First(&integral, 1)
	log.Infof("扣减积分成功: %v", integral.Integral)

	return &emptypb.Empty{}, nil
}

func (s *GreeterService) IncreaseIntegral(ctx context.Context, req *stockV1.IncreaseIntegralRequest) (*emptypb.Empty, error) {
	grpc, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		log.Errorf("dtm error: %v", err)
		return nil, err
	}
	begin := s.db.Begin()
	grpc.Call(begin.Statement.ConnPool.(*sql.Tx), func(db *sql.Tx) error {
		db.Exec("update integral set integral = integral + ? WHERE id = 1", req.Integral)
		return nil
	})

	var integral model.Integral
	s.db.First(&integral, 1)
	log.Infof("增加积分成功: %v", integral.Integral)
	return &emptypb.Empty{}, nil
}

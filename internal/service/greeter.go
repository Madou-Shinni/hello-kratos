package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "helloword/api/helloworld/v1"
	stockV1 "helloword/api/stock/v1"
	"helloword/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	return &v1.HelloReply{Message: g.Hello}, errors.New(1, "错误原因", "发生了错误")
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) DeductStock(ctx context.Context, req *stockV1.DeductStockRequest) (*emptypb.Empty, error) {
	return nil, s.uc.DeductStock(ctx, req)
}

func (s *GreeterService) AddStock(ctx context.Context, req *stockV1.IncreaseStockRequest) (*emptypb.Empty, error) {
	return nil, s.uc.AddStock(ctx, req)
}

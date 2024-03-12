package biz

import (
	"context"
	"helloword/types"

	"github.com/go-kratos/kratos/v2/log"
)

// Consumer is a Consumer model.
type Consumer struct {
	types.MessageDuobaoPayload
	Hello string
}

// ConsumerRepo is a Greater repo.
type ConsumerRepo interface {
	Save(context.Context, *Consumer) (*Consumer, error)
}

// ConsumerUsecase is a Consumer usecase.
type ConsumerUsecase struct {
	repo ConsumerRepo
	log  *log.Helper
}

// NewConsumerUsecase new a Consumer usecase.
func NewConsumerUsecase(repo ConsumerRepo, logger log.Logger) *ConsumerUsecase {
	return &ConsumerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateConsumer creates a Consumer, and returns the new Consumer.
func (uc *ConsumerUsecase) CreateConsumer(ctx context.Context, g *Consumer) (*Consumer, error) {
	uc.log.WithContext(ctx).Infof("CreateConsumer: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

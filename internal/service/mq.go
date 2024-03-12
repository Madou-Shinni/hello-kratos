package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"helloword/internal/biz"
	"helloword/types"
)

type ConsumerService struct {
	log *log.Helper
	uc  *biz.ConsumerUsecase
}

func NewConsumerService(logger log.Logger, uc *biz.ConsumerUsecase) *ConsumerService {

	l := log.NewHelper(log.With(logger, "module", "Consumer/service/logger-service"))
	return &ConsumerService{
		log: l,
		uc:  uc,
	}
}

func (s *ConsumerService) HandleConsumer(channel string, payload *types.MessageDuobaoPayload) error {
	ctx := context.Background()
	fmt.Println("################ 执行任务 #################")
	switch payload.Type {
	case 1: // 优惠券
		s.uc.CreateConsumer(ctx, &biz.Consumer{Hello: "优惠券 任务", MessageDuobaoPayload: *payload})
	case 2: // 积分
		s.uc.CreateConsumer(ctx, &biz.Consumer{Hello: "积分 任务", MessageDuobaoPayload: *payload})
	case 3: // 被邀请的好友完成注册
		s.uc.CreateConsumer(ctx, &biz.Consumer{Hello: "被邀请的好友完成注册 任务", MessageDuobaoPayload: *payload})
	}
	s.log.Errorf("channel: %s, payload: %s\n", channel, *payload)
	return nil
}

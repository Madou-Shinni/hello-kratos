package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"
	"helloword/internal/conf"
	"helloword/internal/service"
	"helloword/types"
	"math/rand"
	"time"
)

const (
	Channel = "kratos:activity::duobao"
)

// NewAsynqServer create a asynq server.
func NewAsynqServer(cfg *conf.Data, logger log.Logger, svc *service.ConsumerService) *asynq.Server {

	ctx := context.Background()

	srv := asynq.NewServer(
		asynq.WithAddress(cfg.Redis.Addr),
	)

	registerAsynqTasks(ctx, logger, srv, svc)

	// 模拟
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		intn := r.Intn(3) // 随机数 [0,3)
		srv.NewTask(Channel, types.MessageDuobaoPayload{
			Type:   intn + 1, // 随机数 [1,4)
			UserId: i + 1,
		})
	}

	return srv
}

func registerAsynqTasks(ctx context.Context, logger log.Logger, srv *asynq.Server, svc *service.ConsumerService) {
	var err error
	log := log.NewHelper(log.With(logger, "module", "Consumer/service/logger-service"))
	err = asynq.RegisterSubscriber(srv, Channel, svc.HandleConsumer)
	if err != nil {
		log.Errorf("register subscriber failed: %v", err)
		return
	}
}

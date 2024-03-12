# asynq

## 安装依赖

```shell
go get -u github.com/tx7do/kratos-transport/transport/asynq
```

## 实现步骤

### 创建server，提供ProviderSet

[mq.go](mq.go)

先创建server，然后在server.go中添加ProviderSet

```go
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

// 注册回调
func registerAsynqTasks(ctx context.Context, logger log.Logger, srv *asynq.Server, svc *service.ConsumerService) {
	var err error
	log := log.NewHelper(log.With(logger, "module", "Consumer/service/logger-service"))
    // 为每一个任务注册一个消费者
	err = asynq.RegisterSubscriber(srv, Channel, svc.HandleConsumer)
	if err != nil {
		log.Errorf("register subscriber failed: %v", err)
		return
	}
}
```

### 实现每一个任务回调

在service实现回调，你可以继续往下执行biz逻辑，再执行data逻辑。

[mq.go](..%2Fservice%2Fmq.go)

这里我只有一个任务，所以只有一个回调，这里模拟实际业务中应该调用的业务逻辑`s.uc.CreateConsumer`。

```go
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
```

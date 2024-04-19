package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	v1 "helloword/api/stock/v1"
	"helloword/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewRedis, NewConsumerRepo, NewStockClient)

// Data .
type Data struct {
	// TODO wrapped database client
	rdb redis.Cmdable
}

// NewData .
func NewData(c *conf.Data, rdb redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rdb: rdb,
	}, cleanup, nil
}

// NewRedis redis
func NewRedis(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "activity-service/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     "", // no password set
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}

	return client
}

func NewStockClient() v1.StockClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewStockClient(conn)
}

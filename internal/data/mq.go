package data

import (
	"context"
	"fmt"

	"helloword/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type consumerRepo struct {
	data *Data
	log  *log.Helper
}

// NewConsumerRepo .
func NewConsumerRepo(data *Data, logger log.Logger) biz.ConsumerRepo {
	return &consumerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *consumerRepo) Save(ctx context.Context, g *biz.Consumer) (*biz.Consumer, error) {
	r.data.rdb.Set(ctx, fmt.Sprintf("activity duobao:userId:%d", g.UserId), g.Hello, -1)
	return g, nil
}

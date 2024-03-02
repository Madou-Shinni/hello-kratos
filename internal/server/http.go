package server

import (
	v1 "helloword/api/helloworld/v1"
	"helloword/common"
	"helloword/internal/conf"
	"helloword/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// 自定义响应
	opts = append(opts, http.ResponseEncoder(common.EncoderResponse()))
	opts = append(opts, http.ErrorEncoder(common.EncoderError()))

	srv := http.NewServer(opts...)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"helloword/internal/biz"
	"helloword/internal/conf"
	"helloword/internal/data"
	"helloword/internal/server"
	"helloword/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	cmdable := data.NewRedis(confData, logger)
	dataData, cleanup, err := data.NewData(confData, cmdable, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	stockClient := data.NewStockClient()
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger, stockClient)
	greeterService := service.NewGreeterService(greeterUsecase)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	consumerRepo := data.NewConsumerRepo(dataData, logger)
	consumerUsecase := biz.NewConsumerUsecase(consumerRepo, logger)
	consumerService := service.NewConsumerService(logger, consumerUsecase)
	asynqServer := server.NewAsynqServer(confData, logger, consumerService)
	app := newApp(logger, httpServer, asynqServer)
	return app, func() {
		cleanup()
	}, nil
}

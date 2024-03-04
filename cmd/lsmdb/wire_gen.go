// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/ZLSMDB/lsmdb_server/internal/data"
	"github.com/ZLSMDB/lsmdb_server/internal/server"
	"github.com/ZLSMDB/lsmdb_server/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tsandl/skvdb/leveldb"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, confServer *conf.Server, confData *conf.Data, logger log.Logger, db *leveldb.DB) (*kratos.App, func(), error) {
	levelDBRepo := data.NewLevelDBRepo(confData, db, logger)
	levelDBUsecase := biz.NewLevelDBUsecase(levelDBRepo, logger)
	s3 := data.NewS3Cli(confData)
	client := data.NewRedisCli(confData)
	dataData, cleanup, err := data.NewData(db, s3, client, logger)
	if err != nil {
		return nil, nil, err
	}
	ossRepo := data.NewOssRepo(dataData, logger)
	ossUsecase := biz.NewOssUsecase(ossRepo, logger)
	lsmdbService := service.NewLsmdbService(levelDBUsecase, logger, ossUsecase, bootstrap)
	registerService := service.NewRegisterService(bootstrap, logger)
	grpcServer := server.NewGRPCServer(confServer, lsmdbService, registerService, logger)
	httpServer := server.NewHTTPServer(confServer, lsmdbService, registerService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

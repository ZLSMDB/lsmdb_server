//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/ZLSMDB/lsmdb_server/internal/data"
	"github.com/ZLSMDB/lsmdb_server/internal/server"
	"github.com/ZLSMDB/lsmdb_server/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tecbot/gorocksdb"
	LevelDB "github.com/tsandl/skvdb/leveldb"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, *conf.Server, *conf.Data, log.Logger, *LevelDB.DB, *gorocksdb.DB) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

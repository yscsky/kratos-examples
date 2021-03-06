//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yscsky/kratos-examples/blog/internal/biz"
	"github.com/yscsky/kratos-examples/blog/internal/conf"
	"github.com/yscsky/kratos-examples/blog/internal/data"
	"github.com/yscsky/kratos-examples/blog/internal/server"
	"github.com/yscsky/kratos-examples/blog/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

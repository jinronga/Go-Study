//go:build wireinject
// +build wireinject

package demowire

import (
	"github.com/google/wire"
	"go-wire/demo1"
)

var providerSet = wire.NewSet(
	demo1.NewController,
	demo1.NewService,
	demo1.NewDao,
)

func Initialize() (*demo1.Controller, error) {
	// 初始化代码
	panic(wire.Build(providerSet))
}

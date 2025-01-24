package demo1

import (
	"github.com/google/wire"
	"go-wire/demo1/demowire"
	"testing"
)

func TestDemo1(t *testing.T) {
	// 依赖注入
	daoDao := NewDao()
	serviceService := NewService(daoDao)
	controller := NewController(serviceService)
	controller.Run()
}

func TestWireDemo1(t *testing.T) {
	wire.Build(NewService, NewController)
}

func TestWireDemo2(t *testing.T) {
	controller, err := demowire.Initialize()
	if err != nil {
		panic(err)
	}
	controller.Run()
}

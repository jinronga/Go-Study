// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package demowire

import (
	"github.com/google/wire"
	"go-wire/demo1"
)

// Injectors from wire.go:

func Initialize() (*demo1.Controller, error) {
	dao := demo1.NewDao()
	service := demo1.NewService(dao)
	controller := demo1.NewController(service)
	return controller, nil
}

// wire.go:

var providerSet = wire.NewSet(demo1.NewController, demo1.NewService, demo1.NewDao)

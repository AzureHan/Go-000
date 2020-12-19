// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"work/internal/biz"
	"work/internal/data"
	"work/internal/service"

	"github.com/google/wire"
)

func InitializeGreeterService() service.Greeter {
	wire.Build(service.NewGreeter, biz.NewGreeter, data.NewGreeter)
	return service.Greeter{}
}

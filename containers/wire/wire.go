//go:build wireinject
// +build wireinject

package wire

import (
	"dependencyInjection/Services/service1"
	"dependencyInjection/Services/service2"

	"github.com/google/wire"
)

type Container struct {
	Service1 *service1.Service
	Service2 *service2.Service
}

func newContainer(Remoteservice *service1.Service, Remoteservice2 *service2.Service) Container {
	return Container{
		Service1: Remoteservice,
		Service2: Remoteservice2,
	}
}

func CreateContainer(time float64) Container {
	wire.Build(
		service1.NewService,
		service2.NewService2,
		newContainer,
	)
	return Container{}
}

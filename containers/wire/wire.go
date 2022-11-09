//go:build wireinject
// +build wireinject

package wire

import (
	remoteservice "dependencyInjection/remoteService"

	"github.com/google/wire"
)

type Container struct {
	Remoteservice  *remoteservice.Service1
	Remoteservice2 *remoteservice.Service2
}

func newContainer(Remoteservice *remoteservice.Service1, Remoteservice2 *remoteservice.Service2) Container {
	return Container{
		Remoteservice:  Remoteservice,
		Remoteservice2: Remoteservice2,
	}
}

func CreateContainer(time float64) Container {
	wire.Build(
		remoteservice.NewRemoteService,
		remoteservice.NewRemoteService2,
		newContainer,
	)
	return Container{}
}

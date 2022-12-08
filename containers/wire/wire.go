//go:build wireinject
// +build wireinject

// go:generate wire
package wire

import (
	"dependencyInjection/Services/service1"
	"dependencyInjection/Services/service2"

	"github.com/google/wire"
)

type Container struct {
	SleepService *service1.SleepService
	Person       *service2.Person
}

func newContainer(SleepService *service1.SleepService, Person *service2.Person) Container {
	return Container{
		SleepService: SleepService,
		Person:       Person,
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

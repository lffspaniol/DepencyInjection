package manual

import (
	"dependencyInjection/Services/service1"
	"dependencyInjection/Services/service2"
)

type Container struct {
	Remoteservice  *service1.Service
	Remoteservice2 *service2.Service
}

// CreateContainer creates a container for the remoteService service.
func CreateContainer(t float64) *Container {
	r := service1.NewService(t)
	r2 := service2.NewService2(r)
	return &Container{
		Remoteservice:  r,
		Remoteservice2: r2,
	}
}

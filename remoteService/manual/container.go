package manual

import remoteservice "dependencyInjection/remoteService"

type Container struct {
	Remoteservice  *remoteservice.Service1
	Remoteservice2 *remoteservice.Service2
}

// CreateContainer creates a container for the remoteService service.
func CreateContainer() *Container {
	r := remoteservice.NewRemoteService(1)
	r2 := remoteservice.NewRemoteService2(r)
	return &Container{
		Remoteservice:  r,
		Remoteservice2: r2,
	}
}

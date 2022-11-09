package contextinjection

import (
	"context"
	remoteservice "dependencyInjection/remoteService"
)

const ctxkey2 = key("RemoteService2")

func NewRemoteService2(ctx context.Context) context.Context {
	// Create a new RemoteService
	service1 := FromContext(ctx)
	service2 := remoteservice.Service2{Service: &service1}
	// Create a new context with the RemoteService
	return ToContext2(ctx, service2)
}

func ToContext2(ctx context.Context, r remoteservice.Service2) context.Context {
	return context.WithValue(ctx, ctxkey2, r)
}

func FromContext2(ctx context.Context) remoteservice.Service2 {
	// Get the RemoteService from the context
	r, _ := ctx.Value(ctxkey2).(remoteservice.Service2)

	return r
}

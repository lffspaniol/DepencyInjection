package contextinjection

import (
	"context"
	"dependencyInjection/Services/service1"
)

type key string

const ctxkey = key("RemoteService")

// ToContext: add a new RemoteService to the context
func ToContext(ctx context.Context, r service1.Service) context.Context {
	return context.WithValue(ctx, ctxkey, r)
}

func FromContext(ctx context.Context) service1.Service {
	// Get the RemoteService from the context
	r, _ := ctx.Value(ctxkey).(service1.Service)

	return r
}

func NewRemoteService(sleepTime float64) context.Context {
	return ToContext(context.Background(), service1.Service{SleepTime: sleepTime})
}

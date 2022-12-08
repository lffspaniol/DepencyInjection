package contextinjection

import (
	"context"
	"dependencyInjection/Services/service1"
)

type key string

const ctxkey = key("RemoteService")

// ToContext: add a new RemoteService to the context
func ToContext(ctx context.Context, r service1.SleepService) context.Context {
	return context.WithValue(ctx, ctxkey, r)
}

func FromContext(ctx context.Context) service1.SleepService {
	// Get the RemoteService from the context
	r, _ := ctx.Value(ctxkey).(service1.SleepService)

	return r
}

func NewRemoteService(sleepTime float64) context.Context {
	return ToContext(context.Background(), service1.SleepService{SleepTime: sleepTime})
}

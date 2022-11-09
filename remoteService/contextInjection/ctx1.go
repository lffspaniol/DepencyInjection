package contextinjection

import (
	"context"
	remoteservice "dependencyInjection/remoteService"
)

type key string

const ctxkey = key("RemoteService")

// ToContext: add a new RemoteService to the context
func ToContext(ctx context.Context, r remoteservice.Service1) context.Context {
	return context.WithValue(ctx, ctxkey, r)
}

func FromContext(ctx context.Context) remoteservice.Service1 {
	// Get the RemoteService from the context
	r, _ := ctx.Value(ctxkey).(remoteservice.Service1)

	return r
}

func NewRemoteService(sleepTime float64) context.Context {
	return ToContext(context.Background(), remoteservice.Service1{SleepTime: sleepTime})
}

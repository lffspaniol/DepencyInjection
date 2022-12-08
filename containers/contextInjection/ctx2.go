package contextinjection

import (
	"context"
	"dependencyInjection/Services/service2"
)

const ctxkey2 = key("RemoteService2")

func NewRemoteService2(ctx context.Context) context.Context {
	// Create a new RemoteService
	service1 := FromContext(ctx)
	service2 := service2.Person{SleepService: &service1}
	// Create a new context with the RemoteService
	return ToContext2(ctx, service2)
}

func ToContext2(ctx context.Context, r service2.Person) context.Context {
	return context.WithValue(ctx, ctxkey2, r)
}

func FromContext2(ctx context.Context) service2.Person {
	// Get the RemoteService from the context
	r, _ := ctx.Value(ctxkey2).(service2.Person)

	return r
}

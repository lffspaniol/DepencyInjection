package main

import (
	"context"
	remoteservice "dependencyInjection/remoteService"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewServeMux(sevice1 *Handler1, sevice2 *Handler2) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/service1", sevice1)
	mux.Handle("/service2", sevice2)
	return mux
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func main() {
	fx.
		New(
			fx.Provide(
				func() float64 { return 1 },
				remoteservice.NewRemoteService,
				remoteservice.NewRemoteService2,
				NewServeMux,
				NewHandlerForService1,
				NewHandlerForService2,
				zap.NewProduction,
			),
			fx.Invoke(NewHTTPServer),
		).
		Run()
}

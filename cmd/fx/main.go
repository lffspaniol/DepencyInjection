package main

import (
	"context"
	"dependencyInjection/Services/service1"
	"dependencyInjection/Services/service2"
	"fmt"
	"html"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Handler2 struct {
	remoteservice *service2.Person
	log           *zap.Logger
}

// HandlewService1 is the handler for the service1 endpoint.
func (h *Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Service2Handler.ServeHTTP")
	h.remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// NewEchoHandler builds a new EchoHandler.
func NewHandlerForService2(remoteservice *service2.Person, zap *zap.Logger) *Handler2 {
	return &Handler2{
		remoteservice: remoteservice,
		log:           zap,
	}
}

type Handler1 struct {
	remoteservice *service1.SleepService
	log           *zap.Logger
}

// HandlewService1 is the handler for the service1 endpoint.
func (h *Handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Service1Handler.ServeHTTP")
	h.remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// NewEchoHandler builds a new EchoHandler.
func NewHandlerForService1(r *service1.SleepService, log *zap.Logger) *Handler1 {
	return &Handler1{
		remoteservice: r,
		log:           log,
	}
}

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
				service2.NewService2,
				service1.NewService,
				NewServeMux,
				NewHandlerForService1,
				NewHandlerForService2,
				zap.NewProduction,
			),
			fx.Invoke(NewHTTPServer),
		).
		Run()
}

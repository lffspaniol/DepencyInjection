package main

import (
	remoteservice "dependencyInjection/remoteService"
	"fmt"
	"html"
	"net/http"

	"go.uber.org/zap"
)

type Handler1 struct {
	remoteservice *remoteservice.Service1
	log           *zap.Logger
}

// HandlewService1 is the handler for the service1 endpoint.
func (h *Handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Service1Handler.ServeHTTP")
	h.remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// NewEchoHandler builds a new EchoHandler.
func NewHandlerForService1(r *remoteservice.Service1, log *zap.Logger) *Handler1 {
	return &Handler1{
		remoteservice: r,
		log:           log,
	}
}

package main

import (
	remoteservice "dependencyInjection/remoteService"
	"fmt"
	"html"
	"net/http"

	"go.uber.org/zap"
)

type Handler2 struct {
	remoteservice *remoteservice.Service2
	log           *zap.Logger
}

// HandlewService1 is the handler for the service1 endpoint.
func (h *Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Service2Handler.ServeHTTP")
	h.remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// NewEchoHandler builds a new EchoHandler.
func NewHandlerForService2(remoteservice *remoteservice.Service2, zap *zap.Logger) *Handler2 {
	return &Handler2{
		remoteservice: remoteservice,
		log:           zap,
	}
}

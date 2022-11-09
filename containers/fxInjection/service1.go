package fxinjection

import (
	remoteservice "dependencyInjection/remoteService"
	"fmt"
	"html"
	"net/http"

	"go.uber.org/zap"
)

type Service1Handler struct {
	remoteservice *remoteservice.Service1
	log           *zap.Logger
}

// HandlewService1 is the handler for the service1 endpoint.
func (h *Service1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Service1Handler.ServeHTTP")
	h.remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// NewEchoHandler builds a new EchoHandler.
func NewService1Handler(r *remoteservice.Service1, log *zap.Logger) *Service1Handler {
	return &Service1Handler{
		remoteservice: r,
		log:           log,
	}
}

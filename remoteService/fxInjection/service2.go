package fxinjection

import (
	remoteservice "dependencyInjection/remoteService"
	"fmt"
	"html"
	"net/http"

	"go.uber.org/zap"
)

type Service2Handler struct {
	remoteservice *remoteservice.Service2
	log           *zap.Logger
}

// HandlewService1 is the handler for the service1 endpoint.
func (h *Service2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Service2Handler.ServeHTTP")
	h.remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// NewEchoHandler builds a new EchoHandler.
func NewService1Handler2(remoteservice *remoteservice.Service2, zap *zap.Logger) *Service2Handler {
	return &Service2Handler{
		remoteservice: remoteservice,
		log:           zap,
	}
}

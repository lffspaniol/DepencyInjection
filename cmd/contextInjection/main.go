package main

import (
	"context"
	contextinjection "dependencyInjection/containers/contextInjection"
	"fmt"
	"html"
	"log"
	"net/http"
)

func middleware(ctx context.Context, h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service1 := contextinjection.FromContext(ctx)
		service1.Sleep()
		h.ServeHTTP(w, r)
	})
}
func middleware2(ctx context.Context, h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service2 := contextinjection.FromContext2(ctx)
		service2.Sleep()
		h.ServeHTTP(w, r)
	})
}

func main() {
	ctx := contextinjection.NewRemoteService(1)
	ctx = contextinjection.NewRemoteService2(ctx)

	http.HandleFunc("/service1", middleware(ctx, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}))

	http.HandleFunc("/service2", middleware2(ctx, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}))

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

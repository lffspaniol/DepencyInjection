package main

import (
	"dependencyInjection/containers/manual"
	"fmt"
	"html"
	"log"
	"net/http"
)

type app struct {
	container *manual.Container
}

func (a *app) HandleService1(w http.ResponseWriter, r *http.Request) {
	a.container.Remoteservice.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func (a *app) HandleService2(w http.ResponseWriter, r *http.Request) {
	a.container.Remoteservice2.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	app := &app{
		container: manual.CreateContainer(1),
	}

	http.HandleFunc("/service1", app.HandleService1)
	http.HandleFunc("/service2", app.HandleService2)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

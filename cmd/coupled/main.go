package main

import (
	"dependencyInjection/Services/service1"
	"dependencyInjection/Services/service2"
	"fmt"
	"html"
	"log"
	"net/http"
)

func HandleService1(w http.ResponseWriter, r *http.Request) {
	s := service1.Service{
		SleepTime: 1.0,
	}
	s.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func HandleService2(w http.ResponseWriter, r *http.Request) {
	s := service1.Service{
		SleepTime: 1,
	}
	s2 := service2.Service{
		Service: &s,
	}
	s2.Sleep()
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/service1", HandleService1)
	http.HandleFunc("/service2", HandleService2)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"dependencyInjection/Services/service1"
	"dependencyInjection/Services/service2"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	s := service1.SleepService{
		SleepTime: 1,
	}
	s2 := service2.Person{
		SleepService: &s,
	}
	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		s.Sleep()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		s2.Sleep()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

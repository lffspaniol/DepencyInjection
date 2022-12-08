package service2

import (
	"dependencyInjection/Services/service1"
	"log"
)

type Sleeper interface {
	Sleep()
}

// type Person struct {
// 	ServiceSleep *service1.SleepService
// }

type Person struct {
	SleepService Sleeper
}

func (r *Person) Sleep() {
	log.Println("person service sleeping...")
	r.SleepService.Sleep()
}

func NewService2(s *service1.SleepService) *Person {
	return &Person{SleepService: s}
}

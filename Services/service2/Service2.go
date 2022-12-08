package service2

import (
	"dependencyInjection/Services/service1"
	"log"
)

type Sleeper interface {
	Sleep()
}

// type Service struct {
// 	Service *service1.Service
// }

type Service struct {
	Service Sleeper
}

func (r *Service) Sleep() {
	log.Println("sevice2")
	r.Service.Sleep()
}

func NewService2(s *service1.Service) *Service {
	return &Service{Service: s}
}

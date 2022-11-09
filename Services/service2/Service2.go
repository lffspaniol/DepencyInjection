package service2

import (
	"dependencyInjection/Services/service1"
	"log"
)

type Service struct {
	Service *service1.Service
}

func (r *Service) Sleep() {
	log.Println("sevice2")
	r.Service.Sleep()
}

func NewService2(s *service1.Service) *Service {
	return &Service{Service: s}
}

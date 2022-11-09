package remoteservice

import "log"

type Service2 struct {
	Service *Service1
}

func (r *Service2) Sleep() {
	log.Println("sevice2")
	r.Service.Sleep()
}

func NewRemoteService2(s *Service1) *Service2 {
	return &Service2{Service: s}
}

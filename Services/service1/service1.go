package service1

import (
	"fmt"
	"time"
)

type Service struct {
	SleepTime float64
}

func (r *Service) Sleep() {
	// ...
	fmt.Println("Sleeping...for:", r.SleepTime)
	// Calling Sleep method
	time.Sleep(time.Duration(r.SleepTime) * time.Second)

	// Printed after sleep is over
	fmt.Println("Sleep Over.....")

}

func NewService(t float64) *Service {
	return &Service{SleepTime: t}
}

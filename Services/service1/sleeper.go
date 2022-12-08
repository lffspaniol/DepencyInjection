package service1

import (
	"fmt"
	"time"
)

type SleepService struct {
	SleepTime float64
}

func (r *SleepService) Sleep() {
	// ...
	fmt.Println("Sleeping...for:", r.SleepTime)
	// Calling Sleep method
	time.Sleep(time.Duration(r.SleepTime) * time.Second)

	// Printed after sleep is over
	fmt.Println("Sleep Over.....")

}

func NewService(t float64) *SleepService {
	return &SleepService{SleepTime: t}
}

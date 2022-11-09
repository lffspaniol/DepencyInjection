package remoteservice

import (
	"fmt"
	"time"
)

type Service1 struct {
	SleepTime float64
}

func (r *Service1) Sleep() {
	// ...
	fmt.Println("Sleeping...for:", r.SleepTime)
	// Calling Sleep method
	time.Sleep(time.Duration(r.SleepTime) * time.Second)

	// Printed after sleep is over
	fmt.Println("Sleep Over.....")

}

func NewRemoteService(t float64) *Service1 {
	return &Service1{SleepTime: t}
}

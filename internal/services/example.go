package services

import (
	"fmt"
	"time"
)

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

type ExampleService struct {
}

func (e *ExampleService) ExampleFunction() {
	for {
		// to check restart, change here
		fmt.Println("Hello World!")
		time.Sleep(10 * time.Second)
	}
}

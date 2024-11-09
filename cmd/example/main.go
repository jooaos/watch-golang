package main

import "github.com/jooaos/watch-golang/internal/services"

func main() {
	s := services.NewExampleService()
	s.ExampleFunction()
}

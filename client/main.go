package main

import (
	"context"
	"fmt"

	example "example-micro-service/proto/example"

	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeterCli := example.NewExampleService("go.micro.srv.example", service.Client())

	// Call the greeter
	rsp, err := greeterCli.Call(context.TODO(), &example.Request{Name: "mike"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)
}

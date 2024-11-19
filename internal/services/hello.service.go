package services

import (
	"fmt"
	"test-grpc/grpc/internal/model"
	"test-grpc/grpc/internal/port"
)

type HelloService struct{}

func NewHelloService() port.HelloServicePort {
	return &HelloService{}
}

func (h *HelloService) Sayhello(name string) (model.Hello, error) {
	if name == "" {
		// return model.Hello{}, fmt.Errorf("name cannot be empty")
		return model.Hello{Message: "annonymous"}, nil
	}
	message := fmt.Sprintf("Hello, %s!", name)
	return model.Hello{Message: message}, nil
}

func (h *HelloService) SayHelloAgain(name string) (model.Hello, error) {
	if name == "" {
		// return model.Hello{}, fmt.Errorf("name cannot be empty")
		return model.Hello{Message: "annonymous"}, nil
	}
	message := fmt.Sprintf("Hello, %s Again!", name)
	return model.Hello{Message: message}, nil
}

package port

import "test-grpc/grpc/internal/model"

type HelloServicePort interface {
	Sayhello(name string) (model.Hello, error)
	SayHelloAgain(name string) (model.Hello, error)
}

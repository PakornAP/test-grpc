package adapters

import (
	"context"
	"test-grpc/grpc/internal/pb/helloworld"
	"test-grpc/grpc/internal/port"
)

type HelloGrpcAdapter struct {
	helloworld.UnimplementedHelloServiceServer
	Service port.HelloServicePort
}

func NewHelloGrpcAdapter(service port.HelloServicePort) *HelloGrpcAdapter {
	return &HelloGrpcAdapter{
		Service: service,
	}
}

func (a *HelloGrpcAdapter) Sayhello(_ context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	hello, err := a.Service.Sayhello(req.Name)
	if err != nil {
		return nil, err
	}
	return &helloworld.HelloReply{
		Message: hello.Message,
	}, nil
}

func (a *HelloGrpcAdapter) SayHelloAgain(_ context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	hello, err := a.Service.SayHelloAgain(req.Name)
	if err != nil {
		return nil, err
	}
	return &helloworld.HelloReply{
		Message: hello.Message,
	}, nil
}

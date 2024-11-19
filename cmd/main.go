package main

import (
	grpc_controller "test-grpc/grpc/http/controller/grpc"
)

func main() {
	grpc := grpc_controller.NewGRPCServer(8081)
	grpc.StartServer()
}

// ----------

package grpc_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	adapters "test-grpc/grpc/internal/adapters/grpc"
	"test-grpc/grpc/internal/model"
	"test-grpc/grpc/internal/pb/helloworld"
	"test-grpc/grpc/internal/pb/routeguide"
	"test-grpc/grpc/internal/services"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	Port int
}

func NewGRPCServer(port int) *GRPCServer {
	return &GRPCServer{
		Port: port,
	}
}

func LoadFeatures() []model.Feature {
	var data []byte
	var savedFeatures []model.Feature
	filePath := "internal/testdata/route_guide_db.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	if err := json.Unmarshal(data, &savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	return savedFeatures
}

func (gs *GRPCServer) StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gs.Port))
	if err != nil {
		log.Fatalf("can't connect listener %v", err)
	}
	server := grpc.NewServer()
	// Declare service
	helloService := services.NewHelloService()
	routeguideService := services.NewRouteGuideService(LoadFeatures())
	// bind service to each adapter
	helloAdapter := adapters.NewHelloGrpcAdapter(helloService)
	routeguideAdapter := adapters.NewRouteGuideGrpcAdapter(routeguideService)
	// regis the adapter to server
	helloworld.RegisterHelloServiceServer(server, helloAdapter)
	routeguide.RegisterRouteGuideServer(server, routeguideAdapter)
	fmt.Println("GRPC Server is Running...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}

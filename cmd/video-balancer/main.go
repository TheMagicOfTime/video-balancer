package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"video-balancer/api/proto"
	"video-balancer/internal/balancer"
	"video-balancer/internal/config"
	"video-balancer/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	videoBalancer := balancer.NewBalancer(cfg.CDNHost)

	grpcServer := grpc.NewServer()

	proto.RegisterVideoBalancerServer(
		grpcServer,
		service.NewService(videoBalancer),
	)

	reflection.Register(grpcServer)

	addr := fmt.Sprintf(":%s", cfg.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server is listening on %s", addr)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

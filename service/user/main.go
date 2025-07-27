package main

import (
	"user/handler"
	pb "user/proto"

	"go-micro.dev/v5"
)

func main() {
	// Create service
	service := micro.New("user")

	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterUserHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}

package routes

import (
	"auth-service/pkg/controllers"
	"auth-service/pkg/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) {

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic("Failed to start service:", err)
	}
	fmt.Printf("Successfully start %s service on port %v\n", "auth service", "3000")

	_controllers := controllers.New(db)
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServer(grpcServer, _controllers)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

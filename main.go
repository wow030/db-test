package main

import (
	"log"
	"net"

	"db-test/database"
	pb "db-test/pkg"
	"db-test/service"

	"google.golang.org/grpc"
)

//go:generate protoc --go_out=plugins=grpc:. ./pkg/db-service.proto
func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dao := database.CreateDAO(database.CreateConn())
	svc := service.CreateService(dao)

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	pb.RegisterDBServer(grpcServer, &svc)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

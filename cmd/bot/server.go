package main

import (
	"google.golang.org/grpc"
	apiPkg "homework/internal/api"
	pb "homework/pkg/api"
	"net"
)

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New())
	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

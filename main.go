package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = "8000"

func main() {
	listen, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("failed to net listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("start gRPC server on %s port", port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server: %s", err)
	}

}

package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"

	userpb "go-grpc-playground/protos/v2/users"
)

const (
	portNumber           = "8101"
	gRPCServerPortNumber = "8001"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := userpb.RegisterUserHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+gRPCServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)

	}
}

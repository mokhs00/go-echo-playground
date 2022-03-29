package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go-grpc-playground/data"
	userpb "go-grpc-playground/protos/v2/users"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const port = "8001"

type userServer struct {
	userpb.UserServer
}

func customMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("Requested at:", time.Now())

		res, err := handler(ctx, req)
		return res, err
	}
}

func main() {
	listen, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("failed to net listen : %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				customMiddleware(),
			)),
	)

	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server: %s", err)
	}

}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.UserId

	var userInfo *userpb.UserInfo

	for _, user := range data.Users {
		if user.UserId == userId {
			userInfo = user
			break
		}
	}

	return &userpb.GetUserResponse{User: userInfo}, nil
}

func (s *userServer) ListUser(ctx context.Context, req *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {
	userInfoList := make([]*userpb.UserInfo, len(data.Users))
	for i, user := range data.Users {
		userInfoList[i] = user
	}

	return &userpb.ListUserResponse{
		Users: userInfoList,
	}, nil

}

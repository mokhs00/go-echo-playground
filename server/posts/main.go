package main

import (
	"context"
	userClient "go-grpc-playground/client/users"
	"go-grpc-playground/data"
	postpb "go-grpc-playground/protos/v1/posts"
	userpb "go-grpc-playground/protos/v1/users"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = "8002"
const userServerHost = "localhost:8001"

type postServer struct {
	postpb.PostServer

	userCli userpb.UserClient
}

func main() {
	listen, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("failed to net listen : %v", err)
	}

	userCli := userClient.GetUserClient(userServerHost)
	grpcServer := grpc.NewServer()
	postpb.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli,
	})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server: %s", err)

	}
}

func (s *postServer) ListPost(ctx context.Context, req *postpb.ListPostRequest) (*postpb.ListPostResponse, error) {
	var postMessages []*postpb.PostMessage

	for _, up := range data.UserPosts {
		user, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: up.UserId})
		if err != nil {
			return nil, err
		}

		for _, p := range up.Posts {
			p.Author = user.User.Name
		}

		postMessages = append(postMessages, up.Posts...)
	}

	return &postpb.ListPostResponse{
		Posts: postMessages,
	}, nil
}

func (s *postServer) ListPostByUserId(ctx context.Context, req *postpb.ListPostByUserIdRequest) (*postpb.ListPostByUserIdResponse, error) {
	userId := req.UserId

	user, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userId})

	if err != nil {
		return nil, err
	}

	var postMessages []*postpb.PostMessage

	for _, up := range data.UserPosts {
		if up.UserId == userId {
			for _, p := range up.Posts {
				p.Author = user.User.Name
			}

			postMessages = up.Posts
			break
		}
	}

	return &postpb.ListPostByUserIdResponse{
		Posts: postMessages,
	}, nil

}

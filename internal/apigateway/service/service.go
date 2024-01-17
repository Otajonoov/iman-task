package service

import (
	"fmt"
	"iman-task/internal/apigateway/genproto/post"
	"iman-task/internal/pkg"
	"log"

	"google.golang.org/grpc"
)

type Service interface {
	Post() post.PostServiceClient
}

type service struct {
	post post.PostServiceClient
}

func (s *service) Post() post.PostServiceClient {
	return s.post
}

func NewService(conf *pkg.Config) (Service, error) {

	connPost, err := grpc.Dial(
		fmt.Sprintf("post%s", conf.PostServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("Error dialing post gRPC server:", err)
		return nil, err
	}

	s := &service{
		post: post.NewPostServiceClient(connPost),
	}

	return s, nil
}

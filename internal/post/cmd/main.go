package main

import (
	"iman-task/internal/post/port"
	"iman-task/internal/post/port/genproto/post"
	"log"
	"net"

	"iman-task/internal/pkg"
	"iman-task/internal/post/adapter"
	"iman-task/internal/post/usecase"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := pkg.Load()
	pgxConn, err := pkg.ConnDB()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer pgxConn.Close()

	postRepo := adapter.NewPostRepo(pgxConn)
	collector := adapter.NewCollector()
	postUsecase := usecase.NewPostUsecase(postRepo, collector, pgxConn)
	postGrpcServer := port.NewServer(postUsecase)

	lis, err := net.Listen("tcp", cfg.PostServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	post.RegisterPostServiceServer(s, postGrpcServer)

	log.Println("POST service started in port ", cfg.PostServicePort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}

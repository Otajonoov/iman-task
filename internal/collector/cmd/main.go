package main

import (
	"iman-task/internal/collector/port"
	"iman-task/internal/pkg"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"iman-task/internal/collector/adapter"
	"iman-task/internal/collector/port/genproto/collector"
	"iman-task/internal/collector/usecase"
)

func main() {
	cfg := pkg.Load()

	collectorRepo := adapter.NewCollector()
	collectorUsecase := usecase.NewCollectorUsecase(&collectorRepo)
	collectorGrpcServer := port.NewServiceServer(collectorUsecase)

	lis, err := net.Listen("tcp", cfg.CollectorServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	collector.RegisterCollectorServer(s, collectorGrpcServer)

	log.Println("collector service started in port ", cfg.CollectorServicePort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}

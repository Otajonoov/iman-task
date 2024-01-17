package service

import (
	"fmt"
	"iman-task/internal/pkg"
	"iman-task/internal/post/port/genproto/collector"
	"log"

	"google.golang.org/grpc"
)

type Service interface {
	Create() collector.CollectorClient
}

type service struct {
	create collector.CollectorClient
}

func (s *service) Create() collector.CollectorClient {
	return s.create
}

func NewService(conf *pkg.Config) (Service, error) {

	connCollector, err := grpc.Dial(
		fmt.Sprintf("collector%s", conf.CollectorServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("Error dialing collector gRPC server:", err)
		return nil, err
	}

	s := &service{
		create: collector.NewCollectorClient(connCollector),
	}

	return s, nil
}

package main

import (
	"iman-task/internal/apigateway/router"
	"iman-task/internal/apigateway/service"
	"iman-task/internal/pkg"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg := pkg.Load()

	svc, err := service.NewService(&cfg)
	if err != nil {
		log.Println("Failed to create service:", err.Error())
	}

	server := router.New(&router.Option{
		Service: svc,
	})

	if err := server.Run(cfg.HTTP_PORT); err != nil {
		log.Fatal("failed to run http server", err.Error())
		panic(err.Error())
	}
}

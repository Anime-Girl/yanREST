package app

import (
	"log"
	"testo/internal/controller"
	"testo/internal/repository"
	"testo/internal/server"
	"testo/internal/service"
)

func Run() {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	controller := controller.NewController(service)

	server := server.NewServer(controller.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatal("cant run")
	}
}

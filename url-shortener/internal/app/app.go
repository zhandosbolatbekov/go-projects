package app

import (
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func Run(configPath string) {
	// TODO: configs

	// TODO: init dependencies

	// TODO: Services, Repos & API Handlers

	repos := repository.NewRepositories()
	services := service.NewServices(repos)
	_ = services

	// TODO: HTTP Server

	// TODO: Graceful Shutdown
}

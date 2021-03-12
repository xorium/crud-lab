package main

import (
	"crud-lab/internal/infrastructure/postgres"
	"crud-lab/internal/infrastructure/web"
	"crud-lab/internal/services"
	"crud-lab/pkg/config"
)

func main() {
	cfg := config.NewEnvConfig()

	repo, err := postgres.NewRepo(cfg)
	if err != nil {
		panic(err)
	}

	svc := services.NewServices(repo)
	webApp := web.NewApp(cfg, svc)

	if err := webApp.Start(); err != nil {
		panic(err)
	}
}

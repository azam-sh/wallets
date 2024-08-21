package main

import (
	"wallets/config"
	"wallets/internal/repository"
	"wallets/internal/service"
	"wallets/internal/transport/http"
	"wallets/internal/transport/http/handlers"
	"wallets/internal/transport/http/middleware"
	"wallets/internal/transport/http/router"
	"wallets/pkg/database"
	"wallets/pkg/logger"
)

func main() {
	conf := config.NewConfig()

	logger := logger.SetupLogger(conf.Environment)
	logger.Info("starting wallets project")

	db := database.New(conf.PostgresURL, logger)
	repo := repository.New(db, logger)
	repo.AutoMigrate(logger)
	svc := service.New(repo, conf, logger)
	handlers := handlers.New(svc, logger)
	mw := middleware.New(conf, svc)
	router := router.InitRouter(handlers, mw)
	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)
	server.Run()
}

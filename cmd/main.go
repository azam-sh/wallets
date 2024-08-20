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
	repo := repository.NewRepository(db, logger)
	svc := service.NewService(repo, conf, logger)
	handlers := handlers.NewHandler(svc, logger)
	mw := middleware.NewMiddleware(conf, svc)
	router := router.InitRouter(handlers, mw)
	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)
	server.Run()
}

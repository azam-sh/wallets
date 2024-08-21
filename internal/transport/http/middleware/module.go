package middleware

import (
	"log/slog"
	"net/http"
	"wallets/config"
	"wallets/internal/service"
)

type middleware struct {
	config  *config.Config
	service service.Service
	logger  *slog.Logger
}

type Middleware interface {
	Authenticate(next http.Handler) http.Handler
	LogMiddleware(next http.Handler) http.Handler
}

func New(config *config.Config, svc service.Service, logger *slog.Logger) Middleware {
	return &middleware{
		config:  config,
		service: svc,
		logger:  logger,
	}
}

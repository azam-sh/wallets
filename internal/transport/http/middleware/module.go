package middleware

import (
	"net/http"
	"wallets/config"
	"wallets/internal/service"
)

type middleware struct {
	config  *config.Config
	service service.Service
}

type Middleware interface {
	Authenticate(next http.Handler) http.Handler
}

func New(config *config.Config, svc service.Service) Middleware {
	return &middleware{
		config:  config,
		service: svc,
	}
}

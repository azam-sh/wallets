package middleware

import (
	"wallets/config"
	"wallets/internal/service"
)

type Middleware struct {
	config  *config.Config
	service service.ServiceInterface
}

type MiddlewareInterface interface {
}

func NewMiddleware(config *config.Config, svc service.ServiceInterface) MiddlewareInterface {
	return &Middleware{
		config:  config,
		service: svc,
	}
}

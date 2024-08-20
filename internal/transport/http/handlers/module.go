package handlers

import (
	"log/slog"
	"wallets/internal/service"
)

type Handler struct {
	svc    service.ServiceInterface
	logger *slog.Logger
}

func NewHandler(svc service.ServiceInterface, logger *slog.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

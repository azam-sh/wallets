package handlers

import (
	"log/slog"
	"wallets/internal/service"
)

type Handler struct {
	svc    service.Service
	logger *slog.Logger
}

func NewHandler(svc service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

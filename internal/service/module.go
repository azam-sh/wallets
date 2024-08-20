package service

import (
	"log/slog"
	"wallets/config"
	"wallets/internal/models"
	"wallets/internal/repository"
)

type Service struct {
	Repo   repository.RepositoryInterface
	Config *config.Config
	Loger  *slog.Logger
}

type ServiceInterface interface {
	CheckAccount(phone string) (acc models.Account, err error)
}

func NewService(repo repository.RepositoryInterface, config *config.Config, log *slog.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Loger:  log,
	}
}

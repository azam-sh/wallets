package service

import (
	"log/slog"
	"wallets/config"
	"wallets/internal/models"
	"wallets/internal/repository"
)

type service struct {
	Repo   repository.Repository
	Config *config.Config
	Loger  *slog.Logger
}

type Service interface {
	CheckAccount(phone string) (acc models.Account, err error)
}

func New(repo repository.Repository, config *config.Config, log *slog.Logger) Service {
	return &service{
		Repo:   repo,
		Config: config,
		Loger:  log,
	}
}

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
	CheckAccount(phone string) (accounts models.CheckAccResp, err error)
	RefillBalance(input models.RefillBalanceReq) error
	GetMonthlyTrns(userId int64, input models.Pagination) (trns models.TrnsHistory, err error)
	GetBalance(accId int64) (balance int64, err error)
}

func New(repo repository.Repository, config *config.Config, log *slog.Logger) Service {
	return &service{
		Repo:   repo,
		Config: config,
		Loger:  log,
	}
}

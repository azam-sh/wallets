package repository

import (
	"log/slog"
	"wallets/internal/models"

	"gorm.io/gorm"
)

type repository struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

type Repository interface {
	GetAccByPhone(phone string) (acc models.CheckAccResp, err error)
	RefillBalance(amount int64, accId int64, user models.UserForBalance) (err error)
	GetUserByAccId(id int64) (user models.UserForBalance, err error)
	GetAccById(id int64) (acc models.Account, err error)
	GetMonthlyTrns(userId int64, input models.Pagination) (trns models.TrnsHistory, err error)
	GetBalance(accId int64) (balance int64, err error)
	AutoMigrate(logger *slog.Logger)
}

func New(db *gorm.DB, log *slog.Logger) Repository {
	return &repository{
		DB:     db,
		Logger: log,
	}
}

func (repo *repository) AutoMigrate(logger *slog.Logger) {
	err := repo.DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}

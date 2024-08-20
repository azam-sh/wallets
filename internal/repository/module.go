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
	GetUserByPhone(phone string) (user models.User, err error)
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

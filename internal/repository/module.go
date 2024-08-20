package repository

import (
	"log/slog"
	"wallets/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

type RepositoryInterface interface {
	GetUserByPhone(phone string) (user models.User, err error)
}

func NewRepository(db *gorm.DB, log *slog.Logger) RepositoryInterface {
	return &Repository{
		DB:     db,
		Logger: log,
	}
}

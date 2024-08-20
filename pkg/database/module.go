package database

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(storagePath string, log *slog.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(storagePath), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return db
}

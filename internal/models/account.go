package models

import "time"

type Account struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey"`
	Balance   int64     `json:"balance" gorm:"column:balance"`
	UserId    int64     `json:"user_id" gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at"`
}

func (a *Account) TableName() string {
	return "accounts"
}

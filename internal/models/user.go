package models

import "time"

type User struct {
	Id           int64        `json:"id" gorm:"column:id;primaryKey"`
	FirstName    string       `json:"first_name" gorm:"column:first_name"`
	LastName     string       `json:"last_name" gorm:"column:last_name"`
	Phone        string       `json:"phone" gorm:"column:phone;unique"`
	Identified   int          `json:"identified" gorm:"column:identified"` // 1 - identified; 2 - not identified
	LimitId      int64        `json:"limit_id" gorm:"column:limit_id"`
	BalanceLimit BalanceLimit `gorm:"foreignKey:LimitId"`
	CreatedAt    time.Time    `json:"-" gorm:"column:created_at"`
	UpdatedAt    time.Time    `json:"-" gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

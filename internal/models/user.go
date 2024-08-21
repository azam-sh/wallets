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

type CheckAccReq struct {
	Phone string `json:"phone"`
}

type UserForBalance struct {
	Id         int64  `gorm:"column:id"`
	Identified int    `gorm:"column:identified"`
	Phone      string `gorm:"column:phone"`
	MaxBalance int64  `gorm:"max_balance"`
}

type CheckAccResp struct {
	Id         int64     `json:"id" gorm:"column:id"`
	FirstName  string    `json:"first_name" gorm:"column:first_name"`
	LastName   string    `json:"last_name" gorm:"column:last_name"`
	Identified int       `json:"identified" gorm:"column:identified"`
	Accounts   []Account `json:"accounts" gorm:"-"`
}

func (u *User) TableName() string {
	return "users"
}

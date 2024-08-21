package models

import "time"

type Account struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey"`
	Balance   int64     `json:"balance" gorm:"column:balance"`
	UserId    int64     `json:"user_id" gorm:"column:user_id"`
	User      User      `json:"-" gorm:"foreignKey:UserId"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at"`
}

type TopUpBalanceReq struct {
	AccountId int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type BalanceReq struct {
	AccountId int64 `json:"account_id"`
}

func (a *Account) TableName() string {
	return "accounts"
}

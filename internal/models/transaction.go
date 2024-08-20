package models

import "time"

type Transaction struct {
	Id            int64     `json:"id" gorm:"column:id;primaryKey"`
	Amount        int64     `json:"amount" gorm:"column:amount"`
	AccountId     int64     `json:"account_id" gorm:"column:account_id"`
	BalanceBefore int64     `json:"balance_before" gorm:"column:balance_before"`
	BalanceAfter  int64     `json:"balance_after" gorm:"column:balance_after"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"-" gorm:"column:updated_at"`
}

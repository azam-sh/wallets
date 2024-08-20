package models

type BalanceLimit struct {
	Id          int64 `json:"id" gorm:"column:id;primaryKey"`
	MaxBalance  int64 `json:"max_balance" gorm:"column:max_balance"`
	UserIdentId int   `json:"user_ident_id" gorm:"column:user_ident_id"` // 1 - identified; 2 - not identified
}

func (bl *BalanceLimit) TableName() string {
	return "balance_limits"
}

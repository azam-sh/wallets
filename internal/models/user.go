package models

type User struct {
	Id         int64  `json:"id" gorm:"column:id"`
	FirstName  string `json:"first_name" gorm:"column:first_name"`
	LastName   string `json:"last_name" gorm:"column:last_name"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Identified int    `json:"identified" gorm:"column:identified"`
	LimitId    int64  `json:"limit_id" gorm:"column:limit_id"`
}

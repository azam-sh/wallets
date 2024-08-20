package models

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Identified int    `json:"identified"`
	LimitId    int64  `json:"limit_id"`
}

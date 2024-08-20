package models

type Account struct {
	Id      int64 `json:"id"`
	Balance int64 `json:"balance"`
	UserId  int64 `json:"user_id"`
}

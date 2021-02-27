package models

type Users struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Type string `json:"type"`
}
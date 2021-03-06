package models

import (
	"time"
)

type DbData struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	Key       string    `json:"key"`
	UpdatedOn time.Time `gorm:"column:updatedOn"`
}

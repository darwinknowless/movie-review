package model

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        uint16       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
	FirstName string       `json:"firstname"`
	LastName  string       `json:"lastname"`
	MovieID   uint16       `json:"movieid"`
}

package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint16       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
	FullName  string       `json:"fullname"`
	Role      string       `json:"role"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
}

package model

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID        uint16       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
	Title     string       `json:"title"`
	Desc      string       `json:"desc"`
	Rate      int8         `json:"rate"`
	Img       string       `json:"img"`
	Author    []Author     `gorm:"ForeignKey:MovieID" json:"author"`
}

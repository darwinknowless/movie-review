package controller

import (
	"gorm.io/gorm"
)

type Gorm struct {
	DB *gorm.DB // berisi koneksi DB
}

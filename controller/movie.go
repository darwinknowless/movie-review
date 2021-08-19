package controller

import (
	"movie-review/model"
	"time"

	"github.com/gin-gonic/gin"
)

func (db *Gorm) GetAllMovie(c *gin.Context) {
	var movie []model.Movie // harus dibuat slice karna kita get all data

	// select * from movie
	// db.DB.Find(&movie)

	// select * from movie where deleted_at is null
	db.DB.Where("deleted_at is null").Find(&movie) // hasil dari query akan di masukkan ke var movie
	writeResponse(c, movie, nil)                   // untuk menampilkan response yang lebih rapi
}

func (db *Gorm) PostMovie(c *gin.Context) {
	var movie model.Movie // tidak dibuat slice karna kita cuma post 1 data

	err := c.Bind(&movie) // data dari body di masukkan ke var movie
	if err != nil {
		writeResponse(c, nil, err) // ketika ada error kita juga tetap harus panggil func writeResponse
		// yang sudah dibuat agar bentuk response tetap konsisten
	} else {
		db.DB.Create(&movie) // data dari var movie di insert ke DB
		writeResponse(c, movie, nil)
	}
}

func (db *Gorm) UpdateMovie(c *gin.Context) {
	var movie model.Movie

	id := c.Query("id")

	err := c.Bind(&movie)
	if err != nil {
		writeResponse(c, nil, err)
	} else {
		db.DB.Model(&movie).Where("id=?", id).Updates(model.Movie{
			Title: movie.Title,
			Desc:  movie.Desc,
			Rate:  movie.Rate,
			Img:   movie.Img,
		}) // update multiple column menggunakan "Updates"
		writeResponse(c, movie, nil)
	}
}

func (db *Gorm) DeleteMovie(c *gin.Context) {
	var movie model.Movie
	time.Now()

	id := c.Query("id")
	// hard delete
	// delete from movie where id='id'
	// db.DB.Delete(&movie, id)

	// soft delete
	// update movie set deleted_at='2021-08-19 15:34:10' where id='id'
	db.DB.Model(&movie).Where("id=?", id).Update("deleted_at", time.Now()) // update single column menggunakan "Update"
	writeResponse(c, nil, nil)
}

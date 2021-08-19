package model

import (
	"log"
	"movie-review/connection"
	"strconv"
)

func SeedMovie() {
	db := connection.GetConnection()

	var movieArr = [3][5]string{
		// {"Title", "Desc", "Rate", "Img"}
		{"The Avengers", "Description of The avengers movie", "9", "film.com/avengers.jpg"},
		{"Black Widow", "Description of Black Widow movie", "7", "film.com/black-widow.jpg"},
		{"Spider-Man: Homecoming", "Description of Spider-Man: Homecoming movie", "9", "film.com/spiderman-homecoming.jpg"},
	}

	var movie Movie
	for _, value := range movieArr {
		// jika ingin membuat ID auto incr value id di isi 0 / zero value
		// nanti akan dihandle oleh gorm
		movie.ID = 0
		movie.Title = value[0]
		movie.Desc = value[1]
		rate, _ := strconv.ParseInt(value[2], 10, 32)
		movie.Rate = int8(rate)
		movie.Img = value[3]
		db.Create(&movie)
	}
	log.Println("Seeder movie created")
}

func SeedAuthor() {
	db := *connection.GetConnection()

	var authorArr = [...][4]string{
		//  {"MovieID", "FirstName", "LastName"}
		{"1", "Joss", "Whedon"},
		{"2", "Cate", "Shortland"},
		{"3", "Jon", "Watts"},
	}

	var author Author
	for _, value := range authorArr {
		author.ID = 0
		movieId, _ := strconv.ParseUint(value[0], 10, 32)
		author.MovieID = uint16(movieId)
		author.FirstName = value[1]
		author.LastName = value[2]
		db.Create(&author)
	}
	log.Println("Seeder author created")
}

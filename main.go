package main

import (
	"log"
	"movie-review/connection"
	"movie-review/controller"
	"movie-review/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Cannot load .env file", err)
	}
}

func main() {
	connection.Connect()
	db := controller.Gorm{DB: connection.GetConnection()}

	// Migration
	model.Migration()

	// Seeder
	model.SeedMovie()
	model.SeedAuthor()

	r := gin.Default()

	// movie
	r.GET("/movie", db.GetAllMovie)
	r.POST("/movie", db.PostMovie)
	r.PATCH("/movie", db.UpdateMovie)
	r.DELETE("/movie", db.DeleteMovie)

	// author

	r.Run(":9000")
}

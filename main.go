package main

import (
	"fmt"
	"log"
	"movie-review/connection"
	"movie-review/controller"
	"movie-review/middleware"
	"movie-review/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Cannot load .env file: ", err)
	}
}

func main() {
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Unix())
	connection.Connect()
	db := controller.Gorm{DB: connection.GetConnection()}

	// Migration
	model.Migration()

	// Seeder
	model.SeedMovie()
	model.SeedAuthor()

	r := gin.Default()

	// auth
	r.POST("/auth/register", db.Register)
	r.POST("/auth/login", db.Login)

	// movie
	r.GET("/movie", db.GetAllMovie)
	r.POST("/movie", middleware.AuthMiddleware, db.PostMovie)
	r.PATCH("/movie", middleware.AuthMiddleware, db.UpdateMovie)
	r.DELETE("/movie", middleware.AuthMiddleware, db.DeleteMovie)

	// author

	r.Run(":9000")
}

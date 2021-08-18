package main

import (
	"movie-review/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	var db controller.DB
	r := gin.Default()

	// movie
	r.GET("/movie", db.GetAllMovie)
	r.POST("/movie", db.PostMovie)

	// author

	r.Run(":9000")
}

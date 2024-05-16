package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/backend-pplbo/internal/config"
	"log"
)

func main() {

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "hello world"})
	})

	cfg := config.NewConfig(".")

	log.Println(*cfg)

	log.Fatal(r.Run(":8080"))
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "hello world"})
	})

	log.Fatal(r.Run(":8080"))
}

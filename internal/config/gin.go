package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGin(config *Config) *gin.Engine {
	app := gin.Default()
	app.Use(NewErrorHandler())
	app.Use(cors.Default())

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return app
}

func NewErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

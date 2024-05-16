package config

import "github.com/gin-gonic/gin"

func NewGin(config *Config) *gin.Engine {
	app := gin.Default()
	app.Use(NewErrorHandler())

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return app
}

func NewErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

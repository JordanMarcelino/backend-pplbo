package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/backend-pplbo/internal/api"
	"github.com/jordanmarcelino/backend-pplbo/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RouteConfig struct {
	App         *gin.Engine
	Config      *Config
	Log         *logrus.Logger
	UserHandler *api.UserHandler
}

func NewRouteConfig(app *gin.Engine, config *Config, log *logrus.Logger, userHandler *api.UserHandler) *RouteConfig {
	return &RouteConfig{App: app, Config: config, Log: log, UserHandler: userHandler}
}

func (c *RouteConfig) SetupRoutes() {
	apiGroup := c.App.Group("/api/v1")
	{
		apiGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, models.NewSuccessResponse[*string](nil, "root endpoint"))
		})
	}

	user := apiGroup.Group("/users")
	{
		user.POST("/login", c.UserHandler.Login)
		user.POST("/register", c.UserHandler.Register)
		user.GET("/profile/:id", c.UserHandler.Get)
	}
}

func (c *RouteConfig) StartServer(port int) {
	err := c.App.Run(fmt.Sprintf(":%d", port))

	if err != nil {
		c.Log.Warnf("failed to start server : %+v", err)
	}
}

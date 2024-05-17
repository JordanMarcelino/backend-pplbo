package main

import (
	"github.com/jordanmarcelino/backend-pplbo/internal/api"
	"github.com/jordanmarcelino/backend-pplbo/internal/config"
	"github.com/jordanmarcelino/backend-pplbo/internal/entity"
	"github.com/jordanmarcelino/backend-pplbo/internal/repository"
	"github.com/jordanmarcelino/backend-pplbo/internal/usecase"
)

func main() {
	cfg := config.NewConfig(".")
	logger := config.NewLogger(cfg)
	db := config.NewDatabase(cfg, logger)

	if err := db.AutoMigrate(new(entity.User)); err != nil {
		logger.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db, logger)
	userUseCase := usecase.NewUserUseCase(db, logger, userRepository)
	userHandler := api.NewUserHandler(logger, userUseCase)

	engine := config.NewGin(cfg)
	routeConfig := config.NewRouteConfig(engine, cfg, logger, userHandler)

	routeConfig.SetupRoutes()
	routeConfig.StartServer(8080)
}

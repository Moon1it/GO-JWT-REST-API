package main

import (
	"GO-JWT-REST-API/internal/config"
	"GO-JWT-REST-API/internal/handler"
	"GO-JWT-REST-API/internal/middleware"
	"GO-JWT-REST-API/internal/repository"
	"GO-JWT-REST-API/internal/routes"
	"GO-JWT-REST-API/internal/service"
	"GO-JWT-REST-API/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal("False init config: ", err)
	}

	db, err := database.ConnectToDB(&cfg.DBConfig)
	if err != nil {
		logrus.Fatal("False connect to PostgreSQL: ", err)
	}

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	newApp := fiber.New()

	// Middlewares
	middleware.DefaultMiddleware(newApp)

	// Routes
	routes.SetupRoutes(newApp, newHandler)

	newApp.Listen(":3000")
}

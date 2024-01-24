package middleware

import (
	"GO-JWT-REST-API/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Middleware struct {
	service *service.Service
}

func DefaultMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
		logger.New(),
	)
}

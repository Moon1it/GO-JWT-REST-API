package routes

import (
	"GO-JWT-REST-API/internal/handler"
	"GO-JWT-REST-API/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler *handler.Handler) {
	authRoute := app.Group("/auth")
	authRoute.Get("/refresh", handler.Refresh)

	authSignRoute := authRoute.Group("/sign")
	authSignRoute.Post("/up", handler.SignUp)
	authSignRoute.Post("/in", handler.SignIn)
	authSignRoute.Post("/out", handler.SignOut)

	// Api
	apiRoute := app.Group("/api/v1")

	cardsRoute := apiRoute.Group("/cards", middleware.AuthMiddleware)
	cardsRoute.Post("/", handler.CreateCard)
	cardsRoute.Get("/", handler.GetAllCards)
	cardsRoute.Put("/:id", handler.UpdateCard)
	cardsRoute.Delete("/:id", handler.DeleteCard)
}

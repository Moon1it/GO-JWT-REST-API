package middleware

import (
	"GO-JWT-REST-API/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "authorization header is missing",
		})
	}

	userID, err := utils.ValidateAccessToken(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "validation failed",
		})
	}

	c.Locals("userID", userID)

	return c.Next()
}

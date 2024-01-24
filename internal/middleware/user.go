package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ParseUserID(c *fiber.Ctx) (uint, error) {
	userIDInterface := c.Locals("userID")
	if userIDInterface == nil {
		return 0, errors.New("missing parameter id")
	}

	userID, ok := userIDInterface.(uint)
	if !ok || userID == 0 {
		return 0, errors.New("invalid user id")
	}

	return userID, nil
}

package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ParseCardID(c *fiber.Ctx) (uint, error) {
	paramCardID := c.Params("id")

	intCardID, err := strconv.Atoi(paramCardID)
	if err != nil || intCardID <= 0 {
		return 0, errors.New("invalid card ID")
	}

	return uint(intCardID), nil
}

package handler

import (
	"GO-JWT-REST-API/internal/models/http"
	"GO-JWT-REST-API/internal/repository"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	authRequest := &http.AuthRequest{}

	if err := c.BodyParser(authRequest); err != nil {
		logrus.Errorf(`%s: %s`, ErrBodyParse, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ErrBodyParse,
		})
	}

	if err := authRequest.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	authResponse, err := h.service.AuthService.Registration(authRequest.Email, authRequest.Password)
	if err != nil {
		if errors.Is(err, repository.ErrUserAlreadyExists) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    authResponse.Tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UserID": authResponse.UserID,
		"Email":  authResponse.Email,
		"Tokens": authResponse.Tokens,
	})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	authRequest := &http.AuthRequest{}
	if err := c.BodyParser(authRequest); err != nil {
		logrus.Errorf(`%s: %s`, ErrBodyParse, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ErrBodyParse,
		})
	}
	if err := authRequest.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	authResponse, err := h.service.AuthService.Login(authRequest.Email, authRequest.Password)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    authResponse.Tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UserID": authResponse.UserID,
		"Email":  authResponse.Email,
		"Tokens": authResponse.Tokens,
	})
}

func (h *Handler) SignOut(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "refresh token is missing",
		})
	}

	err := h.service.AuthService.Logout(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	c.ClearCookie("refreshToken")

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "refresh token is missing",
		})
	}

	authResponse, err := h.service.AuthService.Refresh(refreshToken)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    authResponse.Tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UserID": authResponse.UserID,
		"Email":  authResponse.Email,
		"Tokens": authResponse.Tokens,
	})
}

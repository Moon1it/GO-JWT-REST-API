package handler

import (
	"GO-JWT-REST-API/internal/middleware"
	"GO-JWT-REST-API/internal/models/http"
	"GO-JWT-REST-API/internal/repository"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

func (h *Handler) CreateCard(c *fiber.Ctx) error {
	cardRequest := &http.CardRequest{}
	if err := c.BodyParser(cardRequest); err != nil {
		logrus.Errorf(`%s: %s`, ErrBodyParse, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ErrBodyParse,
		})
	}
	err := cardRequest.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	userID, err := middleware.ParseUserID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cardID, err := h.service.CardService.CreateCard(userID, cardRequest.Item, cardRequest.Definition)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": cardID,
	})
}

func (h *Handler) GetAllCards(c *fiber.Ctx) error {
	userID, err := middleware.ParseUserID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cards, err := h.service.CardService.GetCards(userID)
	if err != nil {
		if errors.Is(err, repository.ErrCardsNotFound) {
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"cards": cards,
	})
}

func (h *Handler) UpdateCard(c *fiber.Ctx) error {
	userID, err := middleware.ParseUserID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cardID, err := middleware.ParseCardID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cardRequest := &http.CardRequest{}
	if err = c.BodyParser(cardRequest); err != nil {
		logrus.Errorf(`%s: %s`, ErrBodyParse, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ErrBodyParse,
		})
	}
	err = cardRequest.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = h.service.CardService.UpdateCard(cardID, userID, cardRequest)
	if err != nil {
		if strings.Contains(err.Error(), "card with card_id") {
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

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *Handler) DeleteCard(c *fiber.Ctx) error {
	userID, err := middleware.ParseUserID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cardID, err := middleware.ParseCardID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = h.service.CardService.DeleteCard(cardID, userID)
	if err != nil {
		if strings.Contains(err.Error(), "card with card_id") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

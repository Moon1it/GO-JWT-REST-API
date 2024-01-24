package handler

import (
	"GO-JWT-REST-API/internal/service"
	"errors"
)

var ErrBodyParse = errors.New("failed to parse request body")

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

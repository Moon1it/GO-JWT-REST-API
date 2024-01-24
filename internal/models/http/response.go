package http

import "GO-JWT-REST-API/internal/models"

type AuthResponse struct {
	UserID uint
	Email  string
	Tokens *models.TokenPair
}

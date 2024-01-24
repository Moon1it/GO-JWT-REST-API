package service

import (
	"GO-JWT-REST-API/internal/models"
	"GO-JWT-REST-API/internal/models/http"
	"GO-JWT-REST-API/internal/repository"
)

type AuthService interface {
	Registration(email, password string) (*http.AuthResponse, error)
	Login(email, password string) (*http.AuthResponse, error)
	Logout(refreshToken string) error
	Refresh(refreshToken string) (*http.AuthResponse, error)
}

type CardService interface {
	CreateCard(userID uint, item, definition string) (uint, error)
	GetCards(userID uint) ([]models.Card, error)
	UpdateCard(cardID, userID uint, cardReq *http.CardRequest) error
	DeleteCard(cardID, userID uint) error
}

type Service struct {
	AuthService
	CardService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthServiceImpl(repository.UserRepository, repository.TokenRepository),
		CardService: NewCardServiceImpl(repository.CardRepository),
	}
}

package repository

import (
	"GO-JWT-REST-API/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type TokenRepository interface {
	GetTokenByRefreshToken(refreshToken string) (*models.Token, error)
	UpdateToken(token *models.Token) error
	DeleteToken(token string) error
}

type CardRepository interface {
	CreateCard(card *models.Card) (uint, error)
	GetCardsByUserID(userID uint) ([]models.Card, error)
	UpdateCard(card *models.Card) error
	DeleteCard(cardID, userID uint) error
}

type Repository struct {
	UserRepository
	TokenRepository
	CardRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:  NewUserRepositoryImpl(db),
		TokenRepository: NewTokenRepositoryImpl(db),
		CardRepository:  NewCardRepositoryImpl(db),
	}
}

package service

import (
	"GO-JWT-REST-API/internal/models"
	"GO-JWT-REST-API/internal/models/http"
	"GO-JWT-REST-API/internal/repository"
)

type CardServiceImpl struct {
	CardRepository repository.CardRepository
}

func NewCardServiceImpl(repository repository.CardRepository) *CardServiceImpl {
	return &CardServiceImpl{
		CardRepository: repository,
	}
}

func (cs *CardServiceImpl) CreateCard(userID uint, item, definition string) (uint, error) {
	newCard := &models.Card{
		UserID:     userID,
		Item:       item,
		Definition: definition,
	}
	return cs.CardRepository.CreateCard(newCard)
}

func (cs *CardServiceImpl) GetCards(userID uint) ([]models.Card, error) {
	return cs.CardRepository.GetCardsByUserID(userID)
}

func (cs *CardServiceImpl) UpdateCard(cardID, userID uint, cardReq *http.CardRequest) error {
	card := &models.Card{
		ID:         cardID,
		UserID:     userID,
		Item:       cardReq.Item,
		Definition: cardReq.Definition,
	}

	return cs.CardRepository.UpdateCard(card)
}

func (cs *CardServiceImpl) DeleteCard(cardID, userID uint) error {
	return cs.CardRepository.DeleteCard(cardID, userID)
}

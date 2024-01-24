package repository

import (
	"GO-JWT-REST-API/internal/models"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ErrCardsNotFound = errors.New("no cards found for user")

type CardRepositoryImpl struct {
	db *gorm.DB
}

func NewCardRepositoryImpl(db *gorm.DB) *CardRepositoryImpl {
	return &CardRepositoryImpl{
		db: db,
	}
}

func (cr *CardRepositoryImpl) CreateCard(card *models.Card) (uint, error) {
	result := cr.db.Where("user_id = ?", card.UserID).Create(&card)
	if result.Error != nil {
		logrus.Errorf(`failed to create card: %s`, result.Error)
		return 0, errors.New("failed to create card")
	}
	return card.ID, nil
}

func (cr *CardRepositoryImpl) GetCardsByUserID(userID uint) ([]models.Card, error) {
	var cards []models.Card
	err := cr.db.Where("user_id = ?", userID).Find(&cards).Error
	if err != nil {
		logrus.Errorf("failed to get user's cards: %s", err)
		return nil, errors.New("failed to get user's cards")
	}
	if len(cards) == 0 {
		return nil, ErrCardsNotFound
	}
	return cards, nil
}

func (cr *CardRepositoryImpl) UpdateCard(card *models.Card) error {
	var existingCard models.Card
	err := cr.db.Where("id = ? AND user_id = ?", card.ID, card.UserID).First(&existingCard).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("card with card_id = %d and user_id = %d not found", card.ID, card.UserID)
		}
		logrus.Errorf("failed to find card: %s", err)
		return errors.New("failed to find card")
	}

	err = cr.db.Save(card).Error
	if err != nil {
		logrus.Errorf("failed to update card: %s", err)
		return errors.New("failed to update card")
	}
	return nil
}

func (cr *CardRepositoryImpl) DeleteCard(cardID, userID uint) error {
	var existingCard models.Card
	err := cr.db.Where("id = ? AND user_id = ?", cardID, userID).First(&existingCard).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("card with card_id = %d and user_id = %d not found", cardID, userID)
		}
		logrus.Errorf("failed to find card: %s", err)
		return errors.New("failed to find card")
	}

	err = cr.db.Where("id = ? AND user_id = ?", cardID, userID).Delete(&models.Card{}).Error
	if err != nil {
		logrus.Errorf("failed to delete card: %s", err)
		return errors.New("failed to delete card")
	}

	return nil
}

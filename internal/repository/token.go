package repository

import (
	"GO-JWT-REST-API/internal/models"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ErrTokenNotFound = errors.New("token not found")

type TokenRepositoryImpl struct {
	db *gorm.DB
}

func NewTokenRepositoryImpl(db *gorm.DB) *TokenRepositoryImpl {
	return &TokenRepositoryImpl{
		db: db,
	}
}

func (tr *TokenRepositoryImpl) UpdateToken(token *models.Token) error {
	var existsToken models.Token
	err := tr.db.Where("user_id = ?", token.UserID).First(&existsToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			createErr := tr.db.Create(&token).Error
			if createErr != nil {
				logrus.Errorf("failed to create token: %s", createErr)
				return errors.New("failed to create token")
			}
			return nil
		} else {
			logrus.Errorf("failed to get token: %s", err)
			return errors.New("failed to get token")
		}
	}

	updateErr := tr.db.Model(&existsToken).Updates(map[string]interface{}{"refresh_token": token.RefreshToken}).Error
	if updateErr != nil {
		logrus.Errorf("failed to update token: %v\n", updateErr)
		return errors.New("failed to update token")
	}

	return nil
}

func (tr *TokenRepositoryImpl) DeleteToken(token string) error {
	result := tr.db.Where("refresh_token = ?", token).Delete(&models.Token{RefreshToken: token})
	if result.Error != nil {
		logrus.Errorf("failed to delete token: %v", result.Error)
		return errors.New("failed to delete token")
	}
	if result.RowsAffected == 0 {
		logrus.Errorf("token not found for deletion")
		return errors.New("token not found for deletion")
	}
	return nil
}

func (tr *TokenRepositoryImpl) GetTokenByRefreshToken(token string) (*models.Token, error) {
	existsToken := &models.Token{}
	result := tr.db.Where("refresh_token = ?", token).First(existsToken)
	if result.Error != nil {
		logrus.Errorf("failed to get token: %s", result.Error)
		return nil, errors.New("failed to get token")
	}

	if result.RowsAffected == 0 {
		return nil, ErrTokenNotFound
	}

	return existsToken, nil
}

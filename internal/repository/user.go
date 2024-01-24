package repository

import (
	"GO-JWT-REST-API/internal/models"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (ur *UserRepositoryImpl) CreateUser(newUser *models.User) (*models.User, error) {
	err := ur.db.Create(newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, ErrUserAlreadyExists
		} else {
			logrus.Errorf(`failed to create user: %s`, err.Error())
			return nil, errors.New("failed to create user")
		}
	}
	return newUser, nil
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	existsUser := &models.User{}
	err := ur.db.First(&existsUser, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		logrus.Errorf("failed to get user by email: %s", err)
		return nil, errors.New("failed to get user by email")
	}
	return existsUser, nil
}

func (ur *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	existsUser := &models.User{}
	err := ur.db.First(existsUser, models.User{ID: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		logrus.Errorf("failed to get user by id: %s", err)
		return nil, errors.New("failed to get user by id")
	}
	return existsUser, nil
}

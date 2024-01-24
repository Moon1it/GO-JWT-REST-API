package service

import (
	"GO-JWT-REST-API/internal/models"
	"GO-JWT-REST-API/internal/models/http"
	"GO-JWT-REST-API/internal/repository"
	"GO-JWT-REST-API/pkg/utils"
	"errors"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository  repository.UserRepository
	TokenRepository repository.TokenRepository
}

func NewAuthServiceImpl(userRepository repository.UserRepository, tokenRepository repository.TokenRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (as *AuthServiceImpl) Registration(email, password string) (*http.AuthResponse, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logrus.Errorf("failed to generate hash password: %s", err)
		return nil, errors.New("failed to generate hash password")
	}

	newUser := &models.User{
		Email:          email,
		Password:       string(hashPassword),
		Role:           "user",
		IsActivated:    false,
		ActivationLink: fiberUtils.UUIDv4(),
	}

	user, err := as.UserRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	tokenPayload := &models.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
	}

	tokens, err := utils.GenerateTokens(tokenPayload)
	if err != nil {
		return nil, err
	}

	refreshToken := &models.Token{
		UserID:       user.ID,
		RefreshToken: tokens.RefreshToken,
	}

	err = as.TokenRepository.UpdateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	authResponse := &http.AuthResponse{
		UserID: user.ID,
		Email:  user.Email,
		Tokens: tokens,
	}

	return authResponse, nil
}

func (as *AuthServiceImpl) Login(email, password string) (*http.AuthResponse, error) {
	user, err := as.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logrus.Errorf("failed to compare hash and password: %s", err)
		return nil, errors.New("failed to compare hash and password")
	}

	tokenPayload := &models.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
	}

	tokens, err := utils.GenerateTokens(tokenPayload)
	if err != nil {
		return nil, err
	}

	refreshToken := &models.Token{
		UserID:       user.ID,
		RefreshToken: tokens.RefreshToken,
	}

	err = as.TokenRepository.UpdateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	authResponse := &http.AuthResponse{
		UserID: user.ID,
		Email:  user.Email,
		Tokens: tokens,
	}

	return authResponse, nil
}

func (as *AuthServiceImpl) Logout(refreshToken string) error {
	return as.TokenRepository.DeleteToken(refreshToken)
}

func (as *AuthServiceImpl) Refresh(refreshToken string) (*http.AuthResponse, error) {
	err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	tokenFromDB, err := as.TokenRepository.GetTokenByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := as.UserRepository.GetUserByID(tokenFromDB.UserID)
	if err != nil {
		// Должна быть ошбика пользователь не найден
	}

	tokenPayload := &models.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
	}

	tokens, err := utils.GenerateTokens(tokenPayload)
	if err != nil {
		return nil, err
	}

	newRefreshToken := &models.Token{
		UserID:       user.ID,
		RefreshToken: tokens.RefreshToken,
	}

	err = as.TokenRepository.UpdateToken(newRefreshToken)
	if err != nil {
		return nil, err
	}

	authResponse := &http.AuthResponse{
		UserID: user.ID,
		Email:  user.Email,
		Tokens: tokens,
	}

	return authResponse, nil
}

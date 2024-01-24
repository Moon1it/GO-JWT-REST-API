package utils

import (
	"GO-JWT-REST-API/internal/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func GenerateTokens(payload *models.TokenPayload) (*models.TokenPair, error) {
	accessTokenClaims := jwt.MapClaims{
		"id":    payload.ID,
		"email": payload.Email,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_KEY")))
	if err != nil {
		return nil, fmt.Errorf("access token generation error: %v", err)
	}

	refreshTokenClaims := jwt.MapClaims{
		"id":    payload.ID,
		"email": payload.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_KEY")))
	if err != nil {
		return nil, fmt.Errorf("refresh token generation error: %v", err)
	}

	return &models.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func ValidateAccessToken(accessToken string) (uint, error) {
	parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_ACCESS_KEY")), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if !parsedToken.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("failed to extract claims from token")
	}

	claimUserID, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("id not found in claims or not a valid numeric value")
	}

	userID := uint(claimUserID)
	return userID, nil
}

func ValidateRefreshToken(refreshToken string) error {
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_REFRESH_KEY")), nil
	})
	if err != nil {
		logrus.Errorf("failed to parse refresh token: %s", err)
		return errors.New("failed to parse refresh token")
	}
	if !parsedToken.Valid {
		return errors.New("invalid token")
	}
	return nil
}

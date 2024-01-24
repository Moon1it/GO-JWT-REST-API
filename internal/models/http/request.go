package http

import (
	"errors"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

func (ar *AuthRequest) Validate() error {
	if ar.Email == "" || ar.Password == "" {
		return errors.New("email or password is empty")
	}
	return nil
}

type CardRequest struct {
	Item       string `json:"item" validate:"required,email,lte=255"`
	Definition string `json:"definition" validate:"required,lte=255"`
}

func (cr *CardRequest) Validate() error {
	if cr.Item == "" || cr.Definition == "" {
		return errors.New("item or definition is empty")
	}
	return nil
}

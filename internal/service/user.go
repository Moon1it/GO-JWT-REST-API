package service

import (
	"GO-JWT-REST-API/internal/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

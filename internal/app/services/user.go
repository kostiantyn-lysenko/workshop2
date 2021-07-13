package services

import (
	"workshop2/internal/app/models"
)

type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
}

type UserService struct {
	Users UserRepositoryInterface
}

func (s *UserService) Create(user models.User) (models.User, error) {
	return s.Users.Create(user)
}

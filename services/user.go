package services

import (
	"workshop2/models"
)

type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	Get(username string) (models.User, error)
	Update(models.User) error
}

type UserService struct {
	Users UserRepositoryInterface
}

func (s *UserService) Create(user models.User) (models.User, error) {
	return s.Users.Create(user)
}

func (s *UserService) UpdateTimezone(username string, timezone string) error {
	user, err := s.Users.Get(username)
	if err != nil {
		return err
	}

	user.Timezone = timezone

	return s.Users.Update(user)
}

package repositories

import (
	"sync"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
)

type UserRepository struct {
	Users []models.User
	sync.RWMutex
}

func (r *UserRepository) Get(username string) (models.User, error) {
	r.RLock()
	defer r.RUnlock()
	for _, u := range r.Users {
		if u.Username == username {
			return u, nil
		}
	}

	return models.User{}, errs.NewUserNotFoundError()
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	err := user.Validate()
	if err != nil {
		return user, errs.NewUserValidationError()
	}

	r.Lock()
	defer r.Unlock()

	for _, u := range r.Users {
		if u.Username == user.Username {
			return user, errs.NewUserAlreadyExistsError()
		}
	}

	r.Users = append(r.Users, user)

	return user, nil
}

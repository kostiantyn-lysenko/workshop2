package repositories

import (
	"sync"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
	"workshop2/internal/app/utils"
)

type UserRepository struct {
	Users []models.User
	sync.RWMutex
	Validator utils.ValidatorInterface
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
	err := r.Validator.Struct(user)

	if err != nil {
		return user, errs.NewUserValidationError(err.Error())
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

func (r *UserRepository) Update(user models.User) error {
	err := r.Validator.Struct(user)

	if err != nil {
		return errs.NewUserValidationError(err.Error())
	}

	r.Lock()
	defer r.Unlock()
	for i, u := range r.Users {
		if u.Username == user.Username {
			r.Users[i] = user

			return nil
		}
	}

	return errs.NewUserNotFoundError()
}

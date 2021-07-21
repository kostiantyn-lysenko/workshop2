package repositories

import (
	"sync"
	errs2 "workshop2/errs"
	models2 "workshop2/models"
	utils2 "workshop2/utils"
)

type UserRepository struct {
	Users []models2.User
	sync.RWMutex
	Validator utils2.ValidatorInterface
}

func (r *UserRepository) Get(username string) (models2.User, error) {
	r.RLock()
	defer r.RUnlock()
	for _, u := range r.Users {
		if u.Username == username {
			return u, nil
		}
	}

	return models2.User{}, errs2.NewUserNotFoundError()
}

func (r *UserRepository) Create(user models2.User) (models2.User, error) {
	err := r.Validator.Struct(user)

	if err != nil {
		return user, errs2.NewUserValidationError(err.Error())
	}

	r.Lock()
	defer r.Unlock()

	for _, u := range r.Users {
		if u.Username == user.Username {
			return user, errs2.NewUserAlreadyExistsError()
		}
	}

	r.Users = append(r.Users, user)

	return user, nil
}

func (r *UserRepository) Update(user models2.User) error {
	err := r.Validator.Struct(user)

	if err != nil {
		return errs2.NewUserValidationError(err.Error())
	}

	r.Lock()
	defer r.Unlock()
	for i, u := range r.Users {
		if u.Username == user.Username {
			r.Users[i] = user

			return nil
		}
	}

	return errs2.NewUserNotFoundError()
}

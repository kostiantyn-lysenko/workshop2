package repositories

import (
	"workshop2/errs"
	"workshop2/models"
	"workshop2/storage"
	"workshop2/utils"
)

type UserRepository struct {
	Storage   *storage.Storage
	Validator utils.ValidatorInterface
}

func (r UserRepository) GetAll() ([]models.User, error) {
	users := []models.User{}

	r.Storage.RLock()
	defer r.Storage.RUnlock()

	err := r.Storage.DB.Select(&users, `SELECT * FROM users`)

	if err != nil {
		return users, err
	}

	return users, errs.NewUserNotFoundError()
}

func (r *UserRepository) Get(username string) (models.User, error) {
	user := models.User{}

	r.Storage.RLock()
	defer r.Storage.RUnlock()

	err := r.Storage.DB.Get(&user, `SELECT * FROM users WHERE username = $1`, username)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	err := r.Validator.Struct(user)

	if err != nil {
		return user, errs.NewUserValidationError(err.Error())
	}

	r.Storage.Lock()
	defer r.Storage.Unlock()

	err = r.Storage.DB.Get(&user, `INSERT INTO users VALUES ($1, $2, $3) RETURNING *`,
		user.Username,
		user.Password,
		user.Timezone,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(user models.User) error {
	err := r.Validator.Struct(user)

	if err != nil {
		return errs.NewUserValidationError(err.Error())
	}

	r.Storage.Lock()
	defer r.Storage.Unlock()

	err = r.Storage.DB.Get(&user, `UPDATE users SET password_hash = $1, timezone = $2 WHERE username = $3 RETURNING *`,
		user.Password,
		user.Timezone,
		user.Username,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(user models.User) error {
	r.Storage.Lock()
	defer r.Storage.Unlock()

	_, err := r.Storage.DB.Exec(`DELETE FROM users WHERE username = $1`, user.Username)

	if err != nil {
		return err
	}

	return nil
}

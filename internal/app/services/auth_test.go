package services

import (
	"testing"
	"workshop2/internal/app/models"
	"workshop2/internal/app/repositories"
	"workshop2/internal/app/utils"
)

func TestSignUp(t *testing.T) {
	auth := AuthService{
		&repositories.UserRepository{
			Users:     make([]models.User, 0),
			Validator: utils.NewValidator(),
		},
	}

	t.Run("returns validation error", func(t *testing.T) {
		request := models.SignUp{
			Username:       "",
			Password:       "adsfnsd323",
			RepeatPassword: "adsfnsd323",
		}
		err := auth.SignUp(request)
		if err == nil {
			t.Errorf("username must be specified")
		}
	})

	t.Run("returns validation error", func(t *testing.T) {
		request := models.SignUp{
			Username:       "af",
			Password:       "adsfnsd323",
			RepeatPassword: "adsfnsd323",
		}
		err := auth.SignUp(request)
		if err == nil {
			t.Errorf("username must be minimum 3 characters")
		}
	})

	t.Run("returns validation error", func(t *testing.T) {
		request := models.SignUp{
			Username:       "aaa",
			Password:       "adsfnsd!",
			RepeatPassword: "adsfnsd323!",
		}
		err := auth.SignUp(request)
		if err == nil {
			t.Errorf("passwords must be the same")
		}
	})

	t.Run("returns validation error", func(t *testing.T) {
		request := models.SignUp{
			Username:       "aaa",
			Password:       "adsfnsd323",
			RepeatPassword: "adsfnsd323",
		}
		err := auth.SignUp(request)
		if err == nil {
			t.Errorf("password must contain specific characters")
		}
	})

	t.Run("returns validation error", func(t *testing.T) {
		request := models.SignUp{
			Username:       "adsfnsd323!",
			Password:       "adsfnsd323!",
			RepeatPassword: "adsfnsd323!",
		}
		err := auth.SignUp(request)
		if err == nil {
			t.Errorf("passwords must not be the same as you username")
		}
	})

}

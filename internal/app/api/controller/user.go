package controller

import (
	"encoding/json"
	"net/http"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
)

type UserServiceInterface interface {
	Create(user models.User) (models.User, error)
	UpdateTimezone(username string, timezone string) error
}

type UserController struct {
	Users UserServiceInterface
	Auth  AuthServiceInterface
}

func (e *UserController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user, err = e.Users.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateTimezone(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	tokenString, err := GetTokenCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	claims, err := c.Auth.ExtractClaims(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.Users.UpdateTimezone(claims.Subject, user.Timezone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

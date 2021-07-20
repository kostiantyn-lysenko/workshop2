package controller

import (
	"encoding/json"
	"log"
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

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		if err != nil {
			encodeErr := json.NewEncoder(w).Encode(err.Error())
			if encodeErr != nil {
				log.Fatal(encodeErr.Error())
			}
		}
		return
	}

	user, err = c.Users.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	encodeErr := json.NewEncoder(w).Encode(user)
	if encodeErr != nil {
		log.Fatal(encodeErr.Error())
	}
}

func (c *UserController) UpdateTimezone(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}
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

	username, ok := claims["Username"].(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.Users.UpdateTimezone(username, user.Timezone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}
		return
	}

	tokens, err := c.Auth.GenerateTokens(username, user.Timezone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	SetTokenCookie(w, tokens)

	w.WriteHeader(http.StatusOK)
}

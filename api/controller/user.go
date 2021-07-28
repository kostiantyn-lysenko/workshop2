package controller

import (
	"encoding/json"
	"net/http"
	"workshop2/errs"
	"workshop2/models"
	"workshop2/tokenizer"
)

//go:generate mockgen -destination=../../mocks/services/user.go -package=mocks . UserServiceInterface
type UserServiceInterface interface {
	Create(user models.User) (models.User, error)
	UpdateTimezone(username string, timezone string) error
}

type UserController struct {
	Users     UserServiceInterface
	Tokenizer tokenizer.Tokenizer
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, errs.NewFailedRequestParsingError(), http.StatusBadRequest)
		return
	}

	user, err = c.Users.Create(user)
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	respond(w, user, http.StatusCreated)
}

func (c *UserController) UpdateTimezone(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, errs.NewFailedRequestParsingError(), http.StatusBadRequest)
		return
	}

	claims, err := GetClaimsFromToken(r, c.Tokenizer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.Users.UpdateTimezone(claims.Username, user.Timezone)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	token, err := c.Tokenizer.Generate(tokenizer.Payload{claims.Username, user.Timezone})
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	SetTokenCookie(w, token)
	w.WriteHeader(http.StatusOK)
}

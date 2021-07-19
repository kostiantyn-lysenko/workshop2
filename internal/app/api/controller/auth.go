package controller

import (
	"encoding/json"
	"net/http"
	"time"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
)

type AuthServiceInterface interface {
	SignUp(request models.SignUp) ([]models.Token, error)
	SignIn(request models.SignIn) ([]models.Token, error)
	VerifyToken(token string) error
	TokenLifetime() time.Duration
}

type AuthController struct {
	Auth AuthServiceInterface
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var signin models.SignIn

	err := json.NewDecoder(r.Body).Decode(&signin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	tokens, err := c.Auth.SignIn(signin)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = errs.NewFailedAuthenticationError(err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	c.setDefaultCookie(w, tokens)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var signup models.SignUp

	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	tokens, err := c.Auth.SignUp(signup)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = errs.NewFailedSignUpError(err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	c.setDefaultCookie(w, tokens)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)

}

func (c *AuthController) setDefaultCookie(w http.ResponseWriter, tokens []models.Token) {
	tlt := c.Auth.TokenLifetime()

	cookie := &http.Cookie{
		Name:     "tokens",
		Value:    tokens[0].Value,
		HttpOnly: true,
		Expires:  time.Now().Add(tlt),
	}
	http.SetCookie(w, cookie)
}

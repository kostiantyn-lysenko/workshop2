package controller

import (
	"encoding/json"
	"net/http"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
	"workshop2/internal/app/services"
)

type AuthServiceInterface interface {
	SignUp(request models.SignUp) ([]models.Token, error)
	SignIn(request models.SignIn) ([]models.Token, error)
	VerifyToken(token string) error
	ExtractClaims(tokenString string) (services.Claims, error)
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

	SetTokenCookie(w, tokens)

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

	SetTokenCookie(w, tokens)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)

}

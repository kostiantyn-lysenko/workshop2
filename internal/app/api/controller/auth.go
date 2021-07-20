package controller

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
)

type AuthServiceInterface interface {
	SignUp(request models.SignUp) ([]models.Token, error)
	SignIn(request models.SignIn) ([]models.Token, error)
	VerifyToken(token string) error
	ExtractClaims(tokenString string) (jwt.MapClaims, error)
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
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}

		return
	}

	tokens, err := c.Auth.SignIn(signin)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = errs.NewFailedAuthenticationError(err.Error())
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}
		return
	}

	SetTokenCookie(w, tokens)

	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(tokens)
	if encodeErr != nil {
		log.Fatal(encodeErr.Error())
	}
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var signup models.SignUp

	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}
		return
	}

	tokens, err := c.Auth.SignUp(signup)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = errs.NewFailedSignUpError(err.Error())
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(encodeErr.Error())
		}
		return
	}

	SetTokenCookie(w, tokens)

	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(tokens)
	if encodeErr != nil {
		log.Fatal(encodeErr.Error())
	}

}

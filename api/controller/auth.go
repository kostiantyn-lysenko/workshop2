package controller

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	errs2 "workshop2/errs"
	models2 "workshop2/models"
)

type AuthServiceInterface interface {
	SignUp(request models2.SignUp) ([]models2.Token, error)
	SignIn(request models2.SignIn) ([]models2.Token, error)
	VerifyToken(token string) error
	ExtractClaims(tokenString string) (jwt.MapClaims, error)
	GenerateTokens(username string, timezone string) ([]models2.Token, error)
}

type AuthController struct {
	Auth AuthServiceInterface
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var signin models2.SignIn

	err := json.NewDecoder(r.Body).Decode(&signin)
	if err != nil {
		err = errs2.NewFailedRequestParsingError()
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	tokens, err := c.Auth.SignIn(signin)
	if err != nil {
		err = errs2.NewFailedAuthenticationError(err.Error())
		respondWithError(w, err, http.StatusUnauthorized)
		return
	}

	SetTokenCookie(w, tokens)
	respond(w, tokens, http.StatusOK)
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var signup models2.SignUp

	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		err = errs2.NewFailedRequestParsingError()
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	tokens, err := c.Auth.SignUp(signup)
	if err != nil {
		err = errs2.NewFailedSignUpError(err.Error())
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	SetTokenCookie(w, tokens)
	respond(w, tokens, http.StatusOK)
}

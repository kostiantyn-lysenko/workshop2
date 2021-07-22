package controller

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
	"workshop2/errs"
	"workshop2/models"
)

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SetTokenCookie(w http.ResponseWriter, tokens []models.Token) {
	tlt := time.Hour * 6

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokens[0].Value,
		HttpOnly: true,
		Expires:  time.Now().Add(tlt),
	}
	http.SetCookie(w, cookie)
}

func GetTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func GetClaimsFromToken(r *http.Request, auth AuthServiceInterface) (jwt.MapClaims, error) {
	tokenString, err := GetTokenCookie(r)
	if err != nil {
		return jwt.MapClaims{}, err
	}
	claims, err := auth.ExtractClaims(tokenString)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	return claims, nil
}

func GetUserTimezone(r *http.Request, auth AuthServiceInterface) (*time.Location, error) {
	claims, err := GetClaimsFromToken(r, auth)
	if err != nil {
		return &time.Location{}, errs.NewFailedAuthenticationError(err.Error())
	}

	timezone, ok := claims["Timezone"].(string)
	if !ok {
		return &time.Location{}, errs.NewBadTimezoneError()
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return &time.Location{}, errs.NewBadTimezoneError()
	}
	return loc, nil
}

func respond(w http.ResponseWriter, message interface{}, status int) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	encodeErr := json.NewEncoder(w).Encode(err.Error())
	if encodeErr != nil {
		log.Fatal(encodeErr.Error())
	}
}

package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"workshop2/errs"
	"workshop2/models"
	"workshop2/tokenizer"
)

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SetTokenCookie(w http.ResponseWriter, token models.Token) {
	tlt := time.Hour * 6

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token.Value,
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

func GetClaimsFromToken(r *http.Request, t tokenizer.Tokenizer) (tokenizer.Claims, error) {
	tokenString, err := GetTokenCookie(r)
	if err != nil {
		return tokenizer.Claims{}, err
	}
	claims, err := t.ExtractClaims(tokenString)
	if err != nil {
		return tokenizer.Claims{}, err
	}

	return claims, nil
}

func GetUserTimezone(r *http.Request, t tokenizer.Tokenizer) (*time.Location, error) {
	claims, err := GetClaimsFromToken(r, t)
	if err != nil {
		return &time.Location{}, errs.NewFailedAuthenticationError(err.Error())
	}

	loc, err := time.LoadLocation(claims.Timezone)
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

package controller

import (
	"net/http"
	"time"
	"workshop2/internal/app/models"
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

package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"workshop2/internal/app/api/controller"
	"workshop2/internal/app/errs"
)

type AuthenticationMiddleware struct {
	auth controller.AuthServiceInterface
}

func (mw *AuthenticationMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/v1/sign-in", "/api/v1/sign-up"}
		requestPath := r.URL.Path

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errs.NewMalformedTokenError().Error())
		}

		token := authHeader[1]
		err := mw.auth.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errs.NewMalformedTokenError().Error())
		}

		next.ServeHTTP(w, r)
	})
}

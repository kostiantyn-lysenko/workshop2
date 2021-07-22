package api

import (
	"encoding/json"
	"log"
	"net/http"
	"workshop2/api/controller"
	"workshop2/errs"
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

		token, err := controller.GetTokenCookie(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			encodeErr := json.NewEncoder(w).Encode(errs.NewMalformedTokenError().Error())
			if encodeErr != nil {
				log.Fatal(encodeErr.Error())
			}

			return
		}

		err = mw.auth.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			encodeErr := json.NewEncoder(w).Encode(errs.NewMalformedTokenError().Error())
			if encodeErr != nil {
				log.Fatal(encodeErr.Error())
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}

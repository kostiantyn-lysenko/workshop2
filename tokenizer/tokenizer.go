package tokenizer

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
	"workshop2/errs"
	"workshop2/models"
)

type Payload struct {
	Username string
	Timezone string
}

type Claims struct {
	Username string
	Timezone string
	jwt.StandardClaims
}

//go:generate mockgen -destination=../mocks/tokenizer/tokenizer.go -package=mocks . Tokenizer
type Tokenizer interface {
	Generate(payload Payload) (models.Token, error)
	Verify(str string) error
	ExtractClaims(str string) (Claims, error)
}

type JWTTokenizer struct {
	TokenLifetime time.Duration
	SignKey       string
	Method        jwt.SigningMethod
}

func NewJWTTokenizer() Tokenizer {
	return &JWTTokenizer{
		TokenLifetime: time.Hour * 6,
		SignKey:       "secret_key",
		Method:        jwt.SigningMethodHS256,
	}
}

func (t *JWTTokenizer) Generate(payload Payload) (models.Token, error) {
	claims := Claims{
		payload.Username,
		payload.Timezone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.TokenLifetime).Unix(),
			Issuer:    payload.Username,
		},
	}

	token := jwt.NewWithClaims(t.Method, claims)

	ss, err := token.SignedString([]byte(t.SignKey))
	if err != nil {
		return models.Token{}, errors.New("internal server error")
	}

	return models.Token{Value: ss}, nil
}

func (t *JWTTokenizer) Verify(tokenString string) error {
	token, err := t.parseTokenString(tokenString)

	if err != nil || !token.Valid {
		return errs.NewFailedTokenVerificationError()
	}

	return nil
}

func (t *JWTTokenizer) parseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewFailedTokenVerificationError()
		}
		return []byte(t.SignKey), nil
	})
}

func (t *JWTTokenizer) ExtractClaims(tokenString string) (Claims, error) {
	token, err := t.parseTokenString(tokenString)
	if err != nil || !token.Valid {
		return Claims{}, errs.NewFailedTokenVerificationError()
	}

	jwtClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Claims{}, errs.NewFailedTokenVerificationError()
	}

	claims := Claims{
		Username: jwtClaims["Username"].(string),
		Timezone: jwtClaims["Timezone"].(string),
	}

	return claims, nil
}

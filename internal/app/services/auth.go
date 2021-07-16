package services

import (
	"time"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
	"workshop2/internal/app/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Users                UserRepositoryInterface
	Validator            utils.ValidatorInterface
	TokenLifetime        time.Duration
	RefreshTokenLifetime time.Duration
	SignInKey            string
}

func NewAuth(ur UserRepositoryInterface, val utils.ValidatorInterface, tlt time.Duration, rtlt time.Duration, sk string) *AuthService {
	return &AuthService{
		Users:                ur,
		Validator:            val,
		TokenLifetime:        tlt,
		RefreshTokenLifetime: rtlt,
		SignInKey:            sk,
	}
}

func (s *AuthService) SignUp(request models.SignUp) ([]models.Token, error) {
	var tokens []models.Token
	err := s.Validator.Struct(request)

	if err != nil {
		return tokens, errs.NewAuthValidationError(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.RepeatPassword), bcrypt.DefaultCost)
	if err != nil {
		return tokens, errs.NewAuthValidationError(err.Error())
	}

	user := models.User{
		Username: request.Username,
		Password: string(hash),
	}

	tokens, err = s.generateTokens(user.Username)
	if err != nil {
		return tokens, err
	}

	user, err = s.Users.Create(user)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func (s *AuthService) SignIn(request models.SignIn) ([]models.Token, error) {
	var tokens []models.Token
	err := s.Validator.Struct(request)

	if err != nil {
		return tokens, errs.NewAuthValidationError(err.Error())
	}

	user, err := s.Users.Get(request.Username)
	if err != nil {
		return tokens, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return tokens, err
	}

	tokens, err = s.generateTokens(request.Username)

	return tokens, err
}

func (s *AuthService) generateTokens(username string) ([]models.Token, error) {
	var tokens []models.Token

	t, err := s.generateToken(username, s.TokenLifetime)
	if err != nil {
		return tokens, err
	}

	rt, err := s.generateToken(username, s.RefreshTokenLifetime)
	if err != nil {
		return tokens, err
	}

	t.Type = models.TokenTypeAccess
	rt.Type = models.TokenTypeRefresh

	tokens = append(tokens, t, rt)

	return tokens, nil
}

func (s *AuthService) generateToken(username string, lifetime time.Duration) (models.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(lifetime).Unix(),
		Issuer:    username,
	})

	ss, err := token.SignedString([]byte(s.SignInKey))
	if err != nil {
		return models.Token{}, err
	}

	return models.Token{Value: ss}, nil
}

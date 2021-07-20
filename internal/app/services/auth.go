package services

import (
	"time"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
	"workshop2/internal/app/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string
	Timezone string
	jwt.StandardClaims
}

type AuthService struct {
	Users                UserRepositoryInterface
	Validator            utils.ValidatorInterface
	tokenLifetime        time.Duration
	refreshTokenLifetime time.Duration
	SignInKey            string
	method               jwt.SigningMethod
}

func NewAuth(ur UserRepositoryInterface, val utils.ValidatorInterface, tlt time.Duration, rtlt time.Duration, sk string, method jwt.SigningMethod) *AuthService {
	return &AuthService{
		Users:                ur,
		Validator:            val,
		tokenLifetime:        tlt,
		refreshTokenLifetime: rtlt,
		SignInKey:            sk,
		method:               method,
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
		Timezone: request.Timezone,
	}

	tokens, err = s.generateTokens(user.Username, user.Timezone)
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

	tokens, err = s.generateTokens(request.Username, user.Timezone)

	return tokens, err
}

func (s *AuthService) generateTokens(username string, timezone string) ([]models.Token, error) {
	var tokens []models.Token
	claims := Claims{
		username,
		timezone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenLifetime).Unix(),
			Issuer:    username,
		},
	}

	t, err := s.generateToken(username, claims)
	if err != nil {
		return tokens, err
	}

	claims.ExpiresAt = time.Now().Add(s.refreshTokenLifetime).Unix()

	rt, err := s.generateToken(username, claims)
	if err != nil {
		return tokens, err
	}

	t.Type = models.TokenTypeAccess
	rt.Type = models.TokenTypeRefresh

	tokens = append(tokens, t, rt)

	return tokens, nil
}

func (s *AuthService) generateToken(username string, claims Claims) (models.Token, error) {
	token := jwt.NewWithClaims(s.method, claims)

	ss, err := token.SignedString([]byte(s.SignInKey))
	if err != nil {
		return models.Token{}, err
	}

	return models.Token{Value: ss}, nil
}

func (s *AuthService) VerifyToken(tokenString string) error {
	token, err := s.parseTokenString(tokenString)

	if err != nil || !token.Valid {
		return errs.NewFailedTokenVerificationError()
	}

	return nil
}

func (s *AuthService) parseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewFailedTokenVerificationError()
		}
		return []byte(s.SignInKey), nil
	})
}

func (s *AuthService) ExtractClaims(tokenString string) (Claims, error) {
	token, err := s.parseTokenString(tokenString)

	if err != nil || !token.Valid {
		return Claims{}, errs.NewFailedTokenVerificationError()
	}

	claims, ok := token.Claims.(Claims)
	if !ok {
		return Claims{}, errs.NewFailedTokenVerificationError()
	}

	return claims, nil
}

package services

import (
	"time"
	errs2 "workshop2/errs"
	models2 "workshop2/models"
	utils2 "workshop2/utils"

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
	Validator            utils2.ValidatorInterface
	tokenLifetime        time.Duration
	refreshTokenLifetime time.Duration
	SignInKey            string
	method               jwt.SigningMethod
}

func NewAuth(ur UserRepositoryInterface, val utils2.ValidatorInterface, tlt time.Duration, rtlt time.Duration, sk string, method jwt.SigningMethod) *AuthService {
	return &AuthService{
		Users:                ur,
		Validator:            val,
		tokenLifetime:        tlt,
		refreshTokenLifetime: rtlt,
		SignInKey:            sk,
		method:               method,
	}
}

func (s *AuthService) SignUp(request models2.SignUp) ([]models2.Token, error) {
	var tokens []models2.Token
	err := s.Validator.Struct(request)

	if err != nil {
		return tokens, errs2.NewAuthValidationError(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.RepeatPassword), bcrypt.DefaultCost)
	if err != nil {
		return tokens, errs2.NewAuthValidationError(err.Error())
	}

	user := models2.User{
		Username: request.Username,
		Password: string(hash),
		Timezone: request.Timezone,
	}

	tokens, err = s.GenerateTokens(user.Username, user.Timezone)
	if err != nil {
		return tokens, err
	}

	user, err = s.Users.Create(user)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func (s *AuthService) SignIn(request models2.SignIn) ([]models2.Token, error) {
	var tokens []models2.Token
	err := s.Validator.Struct(request)

	if err != nil {
		return tokens, errs2.NewAuthValidationError(err.Error())
	}

	user, err := s.Users.Get(request.Username)
	if err != nil {
		return tokens, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return tokens, err
	}

	tokens, err = s.GenerateTokens(request.Username, user.Timezone)

	return tokens, err
}

func (s *AuthService) GenerateTokens(username string, timezone string) ([]models2.Token, error) {
	var tokens []models2.Token
	claims := Claims{
		username,
		timezone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenLifetime).Unix(),
			Issuer:    username,
		},
	}

	t, err := s.generateToken(claims)
	if err != nil {
		return tokens, err
	}

	claims.ExpiresAt = time.Now().Add(s.refreshTokenLifetime).Unix()

	rt, err := s.generateToken(claims)
	if err != nil {
		return tokens, err
	}

	t.Type = models2.TokenTypeAccess
	rt.Type = models2.TokenTypeRefresh

	tokens = append(tokens, t, rt)

	return tokens, nil
}

func (s *AuthService) generateToken(claims Claims) (models2.Token, error) {
	token := jwt.NewWithClaims(s.method, claims)

	ss, err := token.SignedString([]byte(s.SignInKey))
	if err != nil {
		return models2.Token{}, err
	}

	return models2.Token{Value: ss}, nil
}

func (s *AuthService) VerifyToken(tokenString string) error {
	token, err := s.parseTokenString(tokenString)

	if err != nil || !token.Valid {
		return errs2.NewFailedTokenVerificationError()
	}

	return nil
}

func (s *AuthService) parseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs2.NewFailedTokenVerificationError()
		}
		return []byte(s.SignInKey), nil
	})
}

func (s *AuthService) ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := s.parseTokenString(tokenString)
	if err != nil || !token.Valid {
		return jwt.MapClaims{}, errs2.NewFailedTokenVerificationError()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errs2.NewFailedTokenVerificationError()
	}

	return claims, nil
}

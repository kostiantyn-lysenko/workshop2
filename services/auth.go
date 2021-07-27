package services

import (
	"time"
	"workshop2/errs"
	"workshop2/models"
	"workshop2/utils"

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

func (s *AuthService) SignUp(request models.SignUp) (models.Token, error) {
	var token models.Token
	err := s.Validator.Struct(request)

	if err != nil {
		return token, errs.NewAuthValidationError(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.RepeatPassword), bcrypt.DefaultCost)
	if err != nil {
		return token, errs.NewAuthValidationError(err.Error())
	}

	user := models.User{
		Username: request.Username,
		Password: string(hash),
		Timezone: request.Timezone,
	}

	token, err = s.GenerateToken(user.Username, user.Timezone)
	if err != nil {
		return token, err
	}

	user, err = s.Users.Create(user)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthService) SignIn(request models.SignIn) (models.Token, error) {
	var token models.Token
	err := s.Validator.Struct(request)

	if err != nil {
		return token, errs.NewAuthValidationError(err.Error())
	}

	user, err := s.Users.Get(request.Username)
	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return token, err
	}

	token, err = s.GenerateToken(request.Username, user.Timezone)

	return token, err
}

func (s *AuthService) GenerateToken(username string, timezone string) (models.Token, error) {
	claims := Claims{
		username,
		timezone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenLifetime).Unix(),
			Issuer:    username,
		},
	}

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

func (s *AuthService) ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := s.parseTokenString(tokenString)
	if err != nil || !token.Valid {
		return jwt.MapClaims{}, errs.NewFailedTokenVerificationError()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errs.NewFailedTokenVerificationError()
	}

	return claims, nil
}

package services

import (
	"workshop2/errs"
	"workshop2/models"
	"workshop2/tokenizer"
	"workshop2/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Users     UserRepositoryInterface
	Validator utils.ValidatorInterface
	Tokenizer tokenizer.Tokenizer
}

func NewAuth(ur UserRepositoryInterface, val utils.ValidatorInterface, tokenizer tokenizer.Tokenizer) *AuthService {
	return &AuthService{
		Users:     ur,
		Validator: val,
		Tokenizer: tokenizer,
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

	token, err = s.Tokenizer.Generate(tokenizer.Payload{user.Username, user.Timezone})
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

	token, err = s.Tokenizer.Generate(tokenizer.Payload{request.Username, user.Timezone})

	return token, err
}

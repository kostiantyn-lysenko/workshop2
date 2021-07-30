package services

import (
	"workshop2/errs"
	"workshop2/models"
	"workshop2/tokenizer"
	"workshop2/utils"
)

type AuthService struct {
	Users     UserRepositoryInterface
	Validator utils.ValidatorInterface
	Tokenizer tokenizer.Tokenizer
	Hasher    utils.Hasher
}

func NewAuth(ur UserRepositoryInterface, val utils.ValidatorInterface, tokenizer tokenizer.Tokenizer, hash utils.Hasher) *AuthService {
	return &AuthService{
		Users:     ur,
		Validator: val,
		Tokenizer: tokenizer,
		Hasher:    hash,
	}
}

func (s *AuthService) SignUp(request models.SignUp) (models.Token, error) {
	var token models.Token

	err := s.Validator.Struct(request)
	if err != nil {
		return token, errs.NewAuthValidationError(err.Error())
	}

	hash, err := s.Hasher.Generate(request.RepeatPassword)
	if err != nil {
		return token, errs.NewAuthValidationError(err.Error())
	}

	user := models.User{
		Username: request.Username,
		Password: string(hash),
		Timezone: request.Timezone,
	}

	token, err = s.Tokenizer.Generate(tokenizer.Payload{Username: user.Username, Timezone: user.Timezone})
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

	err = s.Hasher.Compare(user.Password, request.Password)
	if err != nil {
		return token, err
	}

	token, err = s.Tokenizer.Generate(tokenizer.Payload{Username: request.Username, Timezone: user.Timezone})

	return token, err
}

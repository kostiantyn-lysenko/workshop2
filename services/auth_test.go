package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"testing"
	"workshop2/errs"
	mocks "workshop2/mocks/repositories"
	mocks2 "workshop2/mocks/tokenizer"
	. "workshop2/mocks/utils"
	"workshop2/models"
	"workshop2/tokenizer"
	"workshop2/utils"
)

func TestAuthService_SignIn(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		token models.Token
		err   error
	}

	type payload struct {
		request   models.SignIn
		validator func(mockCtrlr *gomock.Controller) utils.ValidatorInterface
		hasher    func(mockCtrlr *gomock.Controller) utils.Hasher
		users     func(mockCtrlr *gomock.Controller) UserRepositoryInterface
		tokenizer func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "invalid request",
			expected: expected{
				token: models.Token{},
				err:   errs.NewAuthValidationError(""),
			},
			payload: payload{
				request: models.SignIn{Username: "username", Password: "password"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password"}).
						Return(errs.NewAuthValidationError("")).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					return NewMockHasher(mockCtrlr)
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					return mocks.NewMockUserRepositoryInterface(mockCtrlr)
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					return mocks2.NewMockTokenizer(mockCtrlr)
				},
			},
		},
		{
			name: "user repo returns error",
			expected: expected{
				token: models.Token{},
				err:   errors.New(""),
			},
			payload: payload{
				request: models.SignIn{Username: "username", Password: "password"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password"}).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					mock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					mock.EXPECT().
						Get("username").
						Return(models.User{}, errors.New("")).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					return NewMockHasher(mockCtrlr)
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					return mocks2.NewMockTokenizer(mockCtrlr)
				},
			},
		},
		{
			name: "invalid password",
			expected: expected{
				token: models.Token{},
				err:   errors.New(""),
			},
			payload: payload{
				request: models.SignIn{Username: "username", Password: "password1"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password1"}).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					mock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					mock.EXPECT().
						Get("username").
						Return(models.User{Password: "password2"}, nil).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Compare("password2", "password1").
						Return(errors.New("")).
						Times(1)
					return mock
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					return mocks2.NewMockTokenizer(mockCtrlr)
				},
			},
		},
		{
			name: "failed token generating",
			expected: expected{
				token: models.Token{""},
				err:   errors.New(""),
			},
			payload: payload{
				request: models.SignIn{Username: "username", Password: "password1"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password1"}).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					mock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					mock.EXPECT().
						Get("username").
						Return(models.User{Password: "password2", Timezone: "t"}, nil).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Compare("password2", "password1").
						Return(nil).
						Times(1)
					return mock
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					mock := mocks2.NewMockTokenizer(mockCtrlr)
					mock.EXPECT().
						Generate(
							tokenizer.Payload{
								Username: "username",
								Timezone: "t",
							},
						).
						Return(models.Token{""}, errors.New("")).
						Times(1)
					return mock
				},
			},
		},
		{
			name: "ok",
			expected: expected{
				token: models.Token{
					Value: "value",
				},
				err: nil,
			},
			payload: payload{
				request: models.SignIn{Username: "username", Password: "password1"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password1"}).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					mock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					mock.EXPECT().
						Get("username").
						Return(models.User{Password: "password2", Timezone: "t"}, nil).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Compare("password2", "password1").
						Return(nil).
						Times(1)
					return mock
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					mock := mocks2.NewMockTokenizer(mockCtrlr)
					mock.EXPECT().
						Generate(
							tokenizer.Payload{
								Username: "username",
								Timezone: "t",
							},
						).
						Return(models.Token{"value"}, nil).
						Times(1)
					return mock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := AuthService{
			Validator: tt.payload.validator(mockCtrlr),
			Hasher:    tt.payload.hasher(mockCtrlr),
			Users:     tt.payload.users(mockCtrlr),
			Tokenizer: tt.payload.tokenizer(mockCtrlr),
		}

		token, err := s.SignIn(tt.payload.request)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "unsuccessful sign-in", tt.name)
			continue
		}

		Expect(token).To(Equal(tt.expected.token), tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestAuthService_SignUp(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		token models.Token
		err   error
	}

	type payload struct {
		request   models.SignUp
		validator func(mockCtrlr *gomock.Controller) utils.ValidatorInterface
		hasher    func(mockCtrlr *gomock.Controller) utils.Hasher
		users     func(mockCtrlr *gomock.Controller) UserRepositoryInterface
		tokenizer func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "invalid request",
			expected: expected{
				token: models.Token{},
				err:   errs.NewAuthValidationError(""),
			},
			payload: payload{
				request: models.SignUp{Username: "username", RepeatPassword: "password"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignUp{Username: "username", RepeatPassword: "password"}).
						Return(errs.NewAuthValidationError("")).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					return NewMockHasher(mockCtrlr)
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					return mocks.NewMockUserRepositoryInterface(mockCtrlr)
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					return mocks2.NewMockTokenizer(mockCtrlr)
				},
			},
		},
		{
			name: "hasher returns error",
			expected: expected{
				token: models.Token{},
				err:   errs.NewAuthValidationError(""),
			},
			payload: payload{
				request: models.SignUp{
					Username:       "username",
					RepeatPassword: "password",
					Password:       "password",
					Timezone:       "t",
				},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(
							models.SignUp{
								Username:       "username",
								RepeatPassword: "password",
								Password:       "password",
								Timezone:       "t",
							},
						).
						Return(nil).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Generate("password").
						Return([]byte(""), errors.New("")).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					return mocks.NewMockUserRepositoryInterface(mockCtrlr)
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					return mocks2.NewMockTokenizer(mockCtrlr)
				},
			},
		},
		{
			name: "failed token generating",
			expected: expected{
				token: models.Token{""},
				err:   errors.New(""),
			},
			payload: payload{
				request: models.SignUp{Username: "username", RepeatPassword: "password1", Timezone: "t"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignUp{Username: "username", RepeatPassword: "password1", Timezone: "t"}).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					return mocks.NewMockUserRepositoryInterface(mockCtrlr)
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Generate("password1").
						Return([]byte(""), nil).
						Times(1)
					return mock
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					mock := mocks2.NewMockTokenizer(mockCtrlr)
					mock.EXPECT().
						Generate(
							tokenizer.Payload{
								Username: "username",
								Timezone: "t",
							},
						).
						Return(models.Token{""}, errors.New("")).
						Times(1)
					return mock
				},
			},
		},
		{
			name: "failed create user",
			expected: expected{
				token: models.Token{"value"},
				err:   errors.New(""),
			},
			payload: payload{
				request: models.SignUp{
					Username:       "username",
					Password:       "password1",
					RepeatPassword: "password1",
					Timezone:       "t",
				},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(
							models.SignUp{
								Username:       "username",
								Password:       "password1",
								RepeatPassword: "password1",
								Timezone:       "t",
							},
						).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					mock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					mock.EXPECT().
						Create(
							models.User{
								Username: "username",
								Password: "hash",
								Timezone: "t",
							},
						).
						Return(models.User{}, errors.New("")).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Generate("password1").
						Return([]byte("hash"), nil).
						Times(1)
					return mock
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					mock := mocks2.NewMockTokenizer(mockCtrlr)
					mock.EXPECT().
						Generate(
							tokenizer.Payload{
								Username: "username",
								Timezone: "t",
							},
						).
						Return(models.Token{"value"}, nil).
						Times(1)
					return mock
				},
			},
		},
		{
			name: "ok",
			expected: expected{
				token: models.Token{"value"},
				err:   nil,
			},
			payload: payload{
				request: models.SignUp{
					Username:       "username",
					Password:       "password1",
					RepeatPassword: "password1",
					Timezone:       "t",
				},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(
							models.SignUp{
								Username:       "username",
								Password:       "password1",
								RepeatPassword: "password1",
								Timezone:       "t",
							},
						).
						Return(nil).
						Times(1)
					return mock
				},
				users: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					mock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					mock.EXPECT().
						Create(
							models.User{
								Username: "username",
								Password: "hash",
								Timezone: "t",
							},
						).
						Return(models.User{}, nil).
						Times(1)
					return mock
				},
				hasher: func(mockCtrlr *gomock.Controller) utils.Hasher {
					mock := NewMockHasher(mockCtrlr)
					mock.EXPECT().
						Generate("password1").
						Return([]byte("hash"), nil).
						Times(1)
					return mock
				},
				tokenizer: func(mockCtrlr *gomock.Controller) tokenizer.Tokenizer {
					mock := mocks2.NewMockTokenizer(mockCtrlr)
					mock.EXPECT().
						Generate(
							tokenizer.Payload{
								Username: "username",
								Timezone: "t",
							},
						).
						Return(models.Token{"value"}, nil).
						Times(1)
					return mock
				},
			},
		},
	}
	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := AuthService{
			Validator: tt.payload.validator(mockCtrlr),
			Hasher:    tt.payload.hasher(mockCtrlr),
			Users:     tt.payload.users(mockCtrlr),
			Tokenizer: tt.payload.tokenizer(mockCtrlr),
		}

		token, err := s.SignUp(tt.payload.request)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "unsuccessful sign-up", tt.name)
			continue
		}

		Expect(token).To(Equal(tt.expected.token), tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

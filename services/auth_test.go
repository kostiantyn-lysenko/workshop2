package services

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"testing"
	"workshop2/errs"
	mocks "workshop2/mocks/utils"
	"workshop2/models"
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
					mock := mocks.NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password"}).
						Return(errs.NewAuthValidationError("")).
						Times(1)
					return mock
				},
			},
		},
		{
			name: "invalid request",
			expected: expected{
				token: models.Token{},
				err:   errs.NewAuthValidationError(""),
			},
			payload: payload{
				request: models.SignIn{Username: "username", Password: "password"},
				validator: func(mockCtrlr *gomock.Controller) utils.ValidatorInterface {
					mock := mocks.NewMockValidatorInterface(mockCtrlr)
					mock.EXPECT().
						Struct(models.SignIn{Username: "username", Password: "password"}).
						Return(errs.NewAuthValidationError("")).
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

}

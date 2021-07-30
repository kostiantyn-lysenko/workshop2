package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"testing"
	"workshop2/errs"
	mocks "workshop2/mocks/repositories"
	"workshop2/models"
)

func TestUserService_Create(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		user models.User
		err  error
	}

	type payload struct {
		user     models.User
		userRepo func(mockCtrlr *gomock.Controller) UserRepositoryInterface
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "user created, error is a nil",
			expected: expected{
				user: models.User{
					Username: "username",
					Timezone: "timezone",
					Password: "p",
				},
				err: nil,
			},
			payload: payload{
				user: models.User{
					Username: "username",
					Timezone: "timezone",
					Password: "p",
				},
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().
						Create(
							models.User{
								Username: "username",
								Timezone: "timezone",
								Password: "p",
							},
						).
						Return(
							models.User{
								Username: "username",
								Timezone: "timezone",
								Password: "p",
							},
							nil,
						).
						Times(1)
					return userRepoMock
				},
			},
		},
		{
			name: "user not created, error is a non-nil",
			expected: expected{
				user: models.User{
					Username: "username",
					Timezone: "timezone",
					Password: "p",
				},
				err: errors.New("err"),
			},
			payload: payload{
				user: models.User{
					Username: "username",
					Timezone: "oldTimezone",
					Password: "p",
				},
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().
						Create(
							models.User{
								Username: "username",
								Timezone: "oldTimezone",
								Password: "p",
							},
						).
						Return(models.User{}, errors.New("err")).
						Times(1)
					return userRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := UserService{
			Users: tt.payload.userRepo(mockCtrlr),
		}

		user, err := s.Create(tt.payload.user)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create user", tt.name)
			continue
		}

		Expect(user).To(Equal(tt.expected.user), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestUserService_UpdateTimezone(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		user models.User
		err  error
	}

	type payload struct {
		username string
		timezone string
		userRepo func(mockCtrlr *gomock.Controller) UserRepositoryInterface
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{

			name: "user not found, error is non-nil",
			expected: expected{
				user: models.User{
					Username: "user",
					Timezone: "timezone",
					Password: "p",
				},
				err: errs.NewUserNotFoundError(),
			},
			payload: payload{
				username: "user",
				timezone: "timezone",
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().
						Get(gomock.Any()).
						Return(
							models.User{},
							errs.NewUserNotFoundError()).
						Times(1)
					return userRepoMock
				},
			},
		},
		{
			name: "user is found, user is updated, error is a nil",
			expected: expected{
				user: models.User{
					Username: "user",
					Timezone: "newTimezone",
					Password: "p",
				},
				err: nil,
			},
			payload: payload{
				username: "user",
				timezone: "newTimezone",
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().
						Get(gomock.Any()).
						Return(
							models.User{
								Username: "user",
								Timezone: "oldTimezone",
								Password: "p",
							},
							nil,
						).
						Times(1)
					userRepoMock.EXPECT().
						Update(
							models.User{
								Username: "user",
								Timezone: "newTimezone",
								Password: "p",
							},
						).
						Return(nil).
						Times(1)
					return userRepoMock
				},
			},
		},
		{
			name: "user is found, user isn't updated, error is a non-nil",
			expected: expected{
				user: models.User{
					Username: "user",
					Timezone: "timezone",
					Password: "p",
				},
				err: errors.New("err"),
			},
			payload: payload{
				username: "user",
				timezone: "timezone",
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().
						Get("user").
						Return(
							models.User{
								Username: "user",
								Timezone: "timezone",
								Password: "p",
							},
							nil,
						).
						Times(1)
					userRepoMock.EXPECT().
						Update(
							models.User{
								Username: "user",
								Timezone: "timezone",
								Password: "p",
							},
						).
						Return(errors.New("err")).
						Times(1)
					return userRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := UserService{
			Users: tt.payload.userRepo(mockCtrlr),
		}

		err := s.UpdateTimezone(tt.payload.username, tt.payload.timezone)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create user", tt.name)
			continue
		}
	}
}

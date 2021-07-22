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

	user := models.User{}
	err := errors.New("")

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "user created, error is a nil",
			expected: expected{
				user: user,
				err:  err,
			},
			payload: payload{
				user: user,
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().Create(user).Return(user, nil).Times(1)
					return userRepoMock
				},
			},
		},
		{
			name: "user not created, error is a non-nil",
			expected: expected{
				user: user,
				err:  err,
			},
			payload: payload{
				user: user,
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().Create(user).Return(user, err).Times(1)
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

		Expect(user).To(Equal(user), "message", tt.name)
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

	username := "test"
	emptyUser := models.User{}
	defaultTimezone := "timezone"

	tc := []struct {
		name string
		expected
		payload
	}{
		{

			name: "user not found, error is non-nil",
			expected: expected{
				user: emptyUser,
				err:  errs.NewUserNotFoundError(),
			},
			payload: payload{
				username: username,
				timezone: defaultTimezone,
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().Get(username).Return(emptyUser, errs.NewUserNotFoundError()).Times(1)
					userRepoMock.EXPECT().Update(emptyUser).Times(0)
					return userRepoMock
				},
			},
		},
		{
			name: "user is found, user is updated, error is a nil",
			expected: expected{
				user: emptyUser,
				err:  nil,
			},
			payload: payload{
				username: username,
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().Get(username).Return(emptyUser, nil).Times(1)
					userRepoMock.EXPECT().Update(emptyUser).Return(nil).Times(1)
					return userRepoMock
				},
			},
		},
		{
			name: "user is found, user isn't updated, error is a non-nil",
			expected: expected{
				user: emptyUser,
				err:  errors.New(""),
			},
			payload: payload{
				username: username,
				userRepo: func(mockCtrlr *gomock.Controller) UserRepositoryInterface {
					userRepoMock := mocks.NewMockUserRepositoryInterface(mockCtrlr)
					userRepoMock.EXPECT().Get(username).Return(emptyUser, nil).Times(1)
					userRepoMock.EXPECT().Update(emptyUser).Return(errors.New("")).Times(1)
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

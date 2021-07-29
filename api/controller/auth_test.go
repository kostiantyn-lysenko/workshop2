package controller

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	mocks "workshop2/mocks/services"
	"workshop2/models"
)

func TestAuthController_SignIn(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		code int
		body string
	}

	type payload struct {
		method string
		url    string
		body   string
		auth   func(mockCtrlr *gomock.Controller) AuthServiceInterface
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "invalid body",
			expected: expected{
				code: http.StatusBadRequest,
				body: `"Provided info is invalid."`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   "",
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					return mocks.NewMockAuthServiceInterface(mockCtrlr)
				},
			},
		},
		{
			name: "empty request",
			expected: expected{
				code: http.StatusBadRequest,
				body: `"Provided info is invalid."`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   `{"username": }`,
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					return mocks.NewMockAuthServiceInterface(mockCtrlr)
				},
			},
		},
		{
			name: "failed token generating",
			expected: expected{
				code: http.StatusUnauthorized,
				body: `""`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   `{"username": "username", "password": "password"}`,
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					mock := mocks.NewMockAuthServiceInterface(mockCtrlr)
					mock.EXPECT().
						SignIn(
							models.SignIn{
								Username: "username",
								Password: "password",
							},
						).
						Return(models.Token{}, errors.New("")).
						Times(1)
					return mock
				},
			},
		},
		{
			name: "ok",
			expected: expected{
				code: http.StatusOK,
				body: `{"Value": "token_string"}`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   `{"username": "username", "password": "password"}`,
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					mock := mocks.NewMockAuthServiceInterface(mockCtrlr)
					mock.EXPECT().
						SignIn(
							models.SignIn{
								Username: "username",
								Password: "password",
							},
						).
						Return(models.Token{Value: "token_string"}, nil).
						Times(1)
					return mock
				},
			},
		},
	}

	for _, tt := range tc {
		req, err := http.NewRequest(
			tt.payload.method,
			"",
			strings.NewReader(tt.payload.body),
		)

		Expect(err).To(BeNil(), err, tt.name)

		rw := httptest.NewRecorder()
		mockCtrlr := gomock.NewController(t)

		handler := AuthController{Auth: tt.payload.auth(mockCtrlr)}
		handler.SignIn(rw, req)
		mockCtrlr.Finish()

		Expect(err).To(BeNil(), rw.Code, tt.name)
		Expect(rw.Code).To(Equal(tt.expected.code))

		if bodyBytes := rw.Body.Bytes(); len(bodyBytes) != 0 {
			Expect(bodyBytes).To(MatchJSON(tt.expected.body), tt.name)
		}
	}
}

func TestAuthController_SignUp(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		code int
		body string
	}

	type payload struct {
		method string
		url    string
		body   string
		auth   func(mockCtrlr *gomock.Controller) AuthServiceInterface
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "invalid body",
			expected: expected{
				code: http.StatusBadRequest,
				body: `"Provided info is invalid."`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   "",
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					return mocks.NewMockAuthServiceInterface(mockCtrlr)
				},
			},
		},
		{
			name: "empty request",
			expected: expected{
				code: http.StatusBadRequest,
				body: `"Provided info is invalid."`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   `{"username": }`,
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					return mocks.NewMockAuthServiceInterface(mockCtrlr)
				},
			},
		},
		{
			name: "failed token generating",
			expected: expected{
				code: http.StatusInternalServerError,
				body: `""`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   `{"username": "username", "password": "password", "repeat_password": "password"}`,
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					mock := mocks.NewMockAuthServiceInterface(mockCtrlr)
					mock.EXPECT().
						SignUp(
							models.SignUp{
								Username:       "username",
								Password:       "password",
								RepeatPassword: "password",
							},
						).
						Return(models.Token{}, errors.New("")).
						Times(1)
					return mock
				},
			},
		},
		{
			name: "ok",
			expected: expected{
				code: http.StatusOK,
				body: `{"Value": "token_string"}`,
			},
			payload: payload{
				method: http.MethodPost,
				body:   `{"username": "username", "password": "password", "repeat_password": "password"}`,
				auth: func(mockCtrlr *gomock.Controller) AuthServiceInterface {
					mock := mocks.NewMockAuthServiceInterface(mockCtrlr)
					mock.EXPECT().
						SignUp(
							models.SignUp{
								Username:       "username",
								Password:       "password",
								RepeatPassword: "password",
							},
						).
						Return(models.Token{Value: "token_string"}, nil).
						Times(1)
					return mock
				},
			},
		},
	}

	for _, tt := range tc {
		req, err := http.NewRequest(
			tt.payload.method,
			"",
			strings.NewReader(tt.payload.body),
		)

		Expect(err).To(BeNil(), err, tt.name)

		rw := httptest.NewRecorder()
		mockCtrlr := gomock.NewController(t)

		handler := AuthController{Auth: tt.payload.auth(mockCtrlr)}
		handler.SignUp(rw, req)
		mockCtrlr.Finish()

		Expect(err).To(BeNil(), rw.Code, tt.name)
		Expect(rw.Code).To(Equal(tt.expected.code))

		if bodyBytes := rw.Body.Bytes(); len(bodyBytes) != 0 {
			Expect(bodyBytes).To(MatchJSON(tt.expected.body), tt.name)
		}
	}
}

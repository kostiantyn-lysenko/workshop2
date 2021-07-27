package services

import (
	"github.com/golang-jwt/jwt"
	. "github.com/onsi/gomega"
	"log"
	"reflect"
	"testing"
	"time"
	"workshop2/errs"

	_ "encoding/base64"
	//mocks "workshop2/mocks/repositories"
	"workshop2/models"
	"workshop2/utils"
)

const signInKey string = "secret_key"

var method jwt.SigningMethodHMAC

func getValidClaims() Claims {
	return Claims{
		"username",
		"timezone",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
			Issuer:    "username",
		},
	}
}

func getTokenString() string {
	token := jwt.NewWithClaims(&method, getValidClaims())

	ss, err := token.SignedString([]byte(signInKey))
	if err != nil {
		log.Fatal("Can't sign a token.")
	}

	return ss
}

func TestAuthService_ExtractClaims(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		claims jwt.Claims
		err    error
	}

	type payload struct {
		token string
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "failed token parsing",
			expected: expected{
				claims: getValidClaims(),
				err:    errs.NewFailedTokenVerificationError(),
			},
			payload: payload{
				token: getTokenString(),
			},
		},
	}

	for _, tt := range tc {
		s := AuthService{
			SignInKey: signInKey,
		}
		claims, err := s.ExtractClaims(tt.payload.token)

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to extract claims", tt.name)
			continue
		}

		Expect(claims).To(Equal(tt.expected.claims), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestAuthService_GenerateTokens(t *testing.T) {
	type fields struct {
		Users                UserRepositoryInterface
		Validator            utils.ValidatorInterface
		tokenLifetime        time.Duration
		refreshTokenLifetime time.Duration
		SignInKey            string
		method               jwt.SigningMethod
	}
	type args struct {
		username string
		timezone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				Users:                tt.fields.Users,
				Validator:            tt.fields.Validator,
				tokenLifetime:        tt.fields.tokenLifetime,
				refreshTokenLifetime: tt.fields.refreshTokenLifetime,
				SignInKey:            tt.fields.SignInKey,
				method:               tt.fields.method,
			}
			got, err := s.GenerateTokens(tt.args.username, tt.args.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTokens() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_SignIn(t *testing.T) {
	type fields struct {
		Users                UserRepositoryInterface
		Validator            utils.ValidatorInterface
		tokenLifetime        time.Duration
		refreshTokenLifetime time.Duration
		SignInKey            string
		method               jwt.SigningMethod
	}
	type args struct {
		request models.SignIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				Users:                tt.fields.Users,
				Validator:            tt.fields.Validator,
				tokenLifetime:        tt.fields.tokenLifetime,
				refreshTokenLifetime: tt.fields.refreshTokenLifetime,
				SignInKey:            tt.fields.SignInKey,
				method:               tt.fields.method,
			}
			got, err := s.SignIn(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_SignUp(t *testing.T) {
	type fields struct {
		Users                UserRepositoryInterface
		Validator            utils.ValidatorInterface
		tokenLifetime        time.Duration
		refreshTokenLifetime time.Duration
		SignInKey            string
		method               jwt.SigningMethod
	}
	type args struct {
		request models.SignUp
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				Users:                tt.fields.Users,
				Validator:            tt.fields.Validator,
				tokenLifetime:        tt.fields.tokenLifetime,
				refreshTokenLifetime: tt.fields.refreshTokenLifetime,
				SignInKey:            tt.fields.SignInKey,
				method:               tt.fields.method,
			}
			got, err := s.SignUp(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_VerifyToken(t *testing.T) {
	type fields struct {
		Users                UserRepositoryInterface
		Validator            utils.ValidatorInterface
		tokenLifetime        time.Duration
		refreshTokenLifetime time.Duration
		SignInKey            string
		method               jwt.SigningMethod
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				Users:                tt.fields.Users,
				Validator:            tt.fields.Validator,
				tokenLifetime:        tt.fields.tokenLifetime,
				refreshTokenLifetime: tt.fields.refreshTokenLifetime,
				SignInKey:            tt.fields.SignInKey,
				method:               tt.fields.method,
			}
			if err := s.VerifyToken(tt.args.tokenString); (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthService_generateToken(t *testing.T) {
	type fields struct {
		Users                UserRepositoryInterface
		Validator            utils.ValidatorInterface
		tokenLifetime        time.Duration
		refreshTokenLifetime time.Duration
		SignInKey            string
		method               jwt.SigningMethod
	}
	type args struct {
		claims Claims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				Users:                tt.fields.Users,
				Validator:            tt.fields.Validator,
				tokenLifetime:        tt.fields.tokenLifetime,
				refreshTokenLifetime: tt.fields.refreshTokenLifetime,
				SignInKey:            tt.fields.SignInKey,
				method:               tt.fields.method,
			}
			got, err := s.generateToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_parseTokenString(t *testing.T) {
	type fields struct {
		Users                UserRepositoryInterface
		Validator            utils.ValidatorInterface
		tokenLifetime        time.Duration
		refreshTokenLifetime time.Duration
		SignInKey            string
		method               jwt.SigningMethod
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				Users:                tt.fields.Users,
				Validator:            tt.fields.Validator,
				tokenLifetime:        tt.fields.tokenLifetime,
				refreshTokenLifetime: tt.fields.refreshTokenLifetime,
				SignInKey:            tt.fields.SignInKey,
				method:               tt.fields.method,
			}
			got, err := s.parseTokenString(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTokenString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTokenString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuth(t *testing.T) {
	type args struct {
		ur     UserRepositoryInterface
		val    utils.ValidatorInterface
		tlt    time.Duration
		rtlt   time.Duration
		sk     string
		method jwt.SigningMethod
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuth(tt.args.ur, tt.args.val, tt.args.tlt, tt.args.rtlt, tt.args.sk, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

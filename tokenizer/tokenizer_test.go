package tokenizer

//
//import (
//	_ "encoding/base64"
//	"errors"
//	"github.com/golang-jwt/jwt"
//	. "github.com/onsi/gomega"
//	"log"
//	"testing"
//	"workshop2/errs"
//	"workshop2/services"
//
//	//mocks "workshop2/mocks/repositories"
//	"workshop2/models"
//)
//
//const signInKey string = "secret_key"
//var randomSignMethod = jwt.SigningMethodHS256
//
//func getValidClaims() Claims {
//	return Claims{
//		Username: "username",
//		Timezone: "timezone",
//	}
//}
//
//func getTokenString() string {
//	token := jwt.NewWithClaims(randomSignMethod, getValidClaims())
//
//	ss, err := token.SignedString([]byte(signInKey))
//	if err != nil {
//		log.Fatal("Can't sign a token.")
//	}
//
//	return ss
//}
//
//func TestAuthService_ExtractClaims(t *testing.T) {
//	RegisterTestingT(t)
//
//	type expected struct {
//		claims jwt.Claims
//		err    error
//	}
//
//	type payload struct {
//		token string
//	}
//
//	tc := []struct {
//		name string
//		expected
//		payload
//	}{
//		{
//			name: "invalid token string",
//			expected: expected{
//				claims: jwt.MapClaims{},
//				err:    errs.NewFailedTokenVerificationError(),
//			},
//			payload: payload{
//				token: "invalid.token.string",
//			},
//		},
//		{
//			name: "ok",
//			expected: expected{
//				claims: jwt.MapClaims{
//					"Username": "username",
//					"Timezone": "timezone",
//				},
//				err:    nil,
//			},
//			payload: payload{
//				token: getTokenString(),
//			},
//		},
//	}
//
//	for _, tt := range tc {
//		s := services.AuthService{
//			SignInKey: signInKey,
//			method:    randomSignMethod,
//		}
//		claims, err := s.ExtractClaims(tt.payload.token)
//
//		if err != nil {
//			Expect(err).To(Equal(tt.expected.err), "failed to extract claims", tt.name)
//			continue
//		}
//
//		Expect(claims).To(Equal(tt.expected.claims), "message", tt.name)
//		Expect(err).To(BeNil(), tt.name)
//	}
//}
//
//func TestAuthService_GenerateTokens(t *testing.T) {
//	RegisterTestingT(t)
//
//	type expected struct {
//		token models.Token
//		err    error
//	}
//
//	type payload struct {
//		username string
//		signKey string
//	}
//
//	tc := []struct {
//		name string
//		expected
//		payload
//	}{
//		{
//			name: "invalid sign key",
//			expected: expected{
//				token: models.Token{Value: ""},
//				err: errors.New("internal server error"),
//			},
//			payload: payload{
//				username: "username",
//				timezone: "timezone",
//				signKey: "",
//			},
//		},
//	}
//
//	for _, tt := range tc {
//		s := services.AuthService{
//			SignInKey: tt.payload.signKey,
//			method:    randomSignMethod,
//		}
//		token, err := s.GenerateToken(tt.payload.username, tt.payload.timezone)
//		jwt.Parse(token.Value)
//
//		if err != nil {
//			Expect(err).To(Equal(tt.expected.err), "failed to extract claims", tt.name)
//			continue
//		}
//
//		Expect(token).To(Equal(tt.expected.token), "message", tt.name)
//		Expect(err).To(BeNil(), tt.name)
//	}
//}
//
//func TestAuthService_SignIn(t *testing.T) {
//
//}
//
//func TestAuthService_SignUp(t *testing.T) {
//
//}
//
//func TestAuthService_VerifyToken(t *testing.T) {
//
//}
//
//func TestAuthService_parseTokenString(t *testing.T) {
//
//}
//
//func TestNewAuth(t *testing.T) {
//
//}

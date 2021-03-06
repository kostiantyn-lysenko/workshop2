package errs

type AuthValidationError struct {
	Message string
}

func (e *AuthValidationError) Error() string {
	return e.Message
}

func NewAuthValidationError(message string) error {
	return &AuthValidationError{Message: message}
}

type FailedAuthenticationError struct {
	Message string
}

func (e *FailedAuthenticationError) Error() string {
	return e.Message
}

func NewFailedAuthenticationError(message string) error {
	return &FailedAuthenticationError{Message: message}
}

type FailedSignUpError struct {
	Message string
}

func (e *FailedSignUpError) Error() string {
	return e.Message
}

func NewFailedSignUpError(message string) error {
	return &FailedSignUpError{Message: message}
}

type FailedTokenVerificationError struct{}

func (e *FailedTokenVerificationError) Error() string {
	return "Failed token verification."
}

func NewFailedTokenVerificationError() error {
	return &FailedTokenVerificationError{}
}

type MalformedAuthTokenError struct{}

func (e *MalformedAuthTokenError) Error() string {
	return "Your token is malformed."
}

func NewMalformedTokenError() error {
	return &MalformedAuthTokenError{}
}

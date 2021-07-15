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

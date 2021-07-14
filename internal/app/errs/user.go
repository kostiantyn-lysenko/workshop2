package errs

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "User with that username does not exists in database."
}

func NewUserNotFoundError() error {
	return &UserNotFoundError{}
}

type UserAlreadyExistsError struct{}

func (e *UserAlreadyExistsError) Error() string {
	return "User with that username already exists in database."
}

func NewUserAlreadyExistsError() error {
	return &UserAlreadyExistsError{}
}

type UserValidationError struct {
	Message string
}

func (e *UserValidationError) Error() string {
	return e.Message
}

func NewUserValidationError(message string) error {
	return &UserValidationError{Message: message}
}

type BadUsernameLengthError struct{}

func (e *BadUsernameLengthError) Error() string {
	return "Username shoul be between 3 - 40 characters."
}

func NewBadUsernameLengthError() error {
	return &BadUsernameLengthError{}
}

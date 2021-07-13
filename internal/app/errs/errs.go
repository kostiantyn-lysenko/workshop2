package errs

type EventNotFoundError struct{}

func (e *EventNotFoundError) Error() string {
	return "Event with that ID does not exists in database."
}

type NotificationNotFoundError struct{}

func (e *NotificationNotFoundError) Error() string {
	return "Notification with that ID does not exists in database."
}

type IdNotNumericError struct{}

func (e *IdNotNumericError) Error() string {
	return "ID should be numeric."
}

type FailedRequestParsingError struct{}

func (e *FailedRequestParsingError) Error() string {
	return "Provided info is invalid."
}

func NewFailedRequestParsingError() error {
	return &FailedRequestParsingError{}
}

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

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

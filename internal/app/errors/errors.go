package errors

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Event with that ID does not exists in database."
}

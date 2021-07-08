package errors

type EventNotFoundError struct{}

func (e *EventNotFoundError) Error() string {
	return "Event with that ID does not exists in database."
}

type NotificationNotFoundError struct{}

func (e *NotificationNotFoundError) Error() string {
	return "Notification with that ID does not exists in database."
}

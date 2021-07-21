package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	TimeUTC     time.Time `json:"time_utc"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}

func (e *Event) ConvertInTimezone(loc time.Location) Event {
	e.Time = e.TimeUTC.In(&loc)
	return *e
}

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

package models

import (
	"time"
)

type Event struct {
	ID          int
	Title       string
	Time        time.Time
	Description string
}

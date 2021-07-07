package models

import (
	"time"
)

type Notification struct {
	ID          int
	Title       string
	Time        time.Time
	Description string
}

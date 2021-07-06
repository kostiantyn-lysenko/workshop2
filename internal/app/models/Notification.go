package models

import (
	"time"
)

type Notification struct {
	Title       string
	Time        time.Time
	Description string
}

package models

import (
	"time"
)

type Notification struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}

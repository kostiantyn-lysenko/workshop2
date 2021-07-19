package models

import (
	"time"
)

type Notification struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	TimeLocal   time.Time `json:"time_local"`
	TimeUTC     time.Time `json:"time_utc"`
	Timezone    string    `json:"timezone"`
	Description string    `json:"description"`
}

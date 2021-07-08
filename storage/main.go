package storage

import (
	"sync"
	"workshop2/internal/app/models"
)

var DB Storage

type Storage struct {
	Events        []models.Event
	Notifications []models.Notification
	sync.RWMutex
}

func init() {
	DB.Events = make([]models.Event, 0)
	DB.Notifications = make([]models.Notification, 0)
}

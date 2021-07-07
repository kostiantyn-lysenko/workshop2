package storage

import (
	"sync"
	"workshop2/internal/app/models"
)

type DB struct {
	Events        []models.Event
	Notifications []models.Notification
	sync.Mutex
	sync.RWMutex
}

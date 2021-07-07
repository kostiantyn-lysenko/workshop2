package storage

import (
	"sync"
	"workshop2/internal/app/models"
)

var DB struct {
	Events        []models.Event
	Notifications []models.Notification
	mu            sync.Mutex
	rwmu          sync.RWMutex
}

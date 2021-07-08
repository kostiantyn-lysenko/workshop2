package storage

import (
	"sync"
	"workshop2/internal/app/models"
)

type Storage struct {
	Events        []models.Event
	Notifications []models.Notification
	sync.RWMutex
}

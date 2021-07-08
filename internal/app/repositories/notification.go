package repositories

import (
	"errors"
	"workshop2/internal/app/models"
	"workshop2/storage"
)

type NotificationRepository struct {
}

func (r *NotificationRepository) GetAll() []models.Notification {

	return storage.DB.Notifications
}

func (r *NotificationRepository) get(id int) (*models.Notification, bool) {

	var notification models.Notification
	var found bool

	storage.DB.Lock()
	for _, e := range storage.DB.Notifications {
		if e.ID == id {
			notification = e
			found = true
			break
		}
	}
	storage.DB.Unlock()

	return &notification, found
}

func (r *NotificationRepository) Create(notification *models.Notification) *[]models.Notification {
	storage.DB.Lock()
	id := len(storage.DB.Notifications) + 1
	storage.DB.Unlock()
	notification.ID = id
	storage.DB.Notifications = append(storage.DB.Notifications, *notification)

	return &storage.DB.Notifications
}

func (r *NotificationRepository) Update(id int, notification *models.Notification) (*models.Notification, bool, error) {

	var changed bool

	notification, ok := r.get(id)
	if !ok {
		return nil, changed, errors.New("notification not found")
	}

	notification.ID = id
	storage.DB.Lock()
	for i, e := range storage.DB.Notifications {
		if e.ID == notification.ID {
			storage.DB.Notifications[i] = *notification

			changed = true
		}
	}
	storage.DB.Unlock()

	return notification, changed, nil
}

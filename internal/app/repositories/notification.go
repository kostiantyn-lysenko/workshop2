package repositories

import (
	"sync"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"
)

type NotificationRepository struct {
	Notifications []models.Notification
	sync.RWMutex
}

func (r *NotificationRepository) GetAll() []models.Notification {
	r.RLock()
	defer r.RUnlock()
	return r.Notifications
}

func (r *NotificationRepository) Create(notification models.Notification) models.Notification {
	r.RLock()
	id := len(r.Notifications) + 1
	r.RUnlock()
	notification.ID = id

	r.Lock()
	defer r.Unlock()
	r.Notifications = append(r.Notifications, notification)

	return notification
}

func (r *NotificationRepository) Update(id int, newNotification models.Notification) (models.Notification, error) {

	newNotification.ID = id
	r.Lock()
	defer r.Unlock()
	for i, n := range r.Notifications {
		if n.ID == newNotification.ID {
			r.Notifications[i] = newNotification

			return newNotification, nil
		}
	}

	return newNotification, &errs.NotificationNotFoundError{}
}

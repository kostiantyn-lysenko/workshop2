package repositories

import (
	"sync"
	"workshop2/errs"
	"workshop2/models"
)

type NotificationRepository struct {
	Notifications []models.Notification
	sync.RWMutex
}

func (r *NotificationRepository) GetAll() ([]models.Notification, error) {
	r.RLock()
	defer r.RUnlock()
	return r.Notifications, nil
}

func (r *NotificationRepository) Create(notification models.Notification) (models.Notification, error) {
	r.Lock()
	defer r.Unlock()
	id := 1
	if len(r.Notifications) > 0 {
		id = (r.Notifications[len(r.Notifications)-1]).ID + 1
	}

	notification.ID = id

	r.Notifications = append(r.Notifications, notification)

	return notification, nil
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

package services

import (
	"time"
	models2 "workshop2/models"
)

type NotificationRepositoryInterface interface {
	GetAll() ([]models2.Notification, error)
	Create(notification models2.Notification) (models2.Notification, error)
	Update(id int, notification models2.Notification) (models2.Notification, error)
}

type NotificationService struct {
	Notifications NotificationRepositoryInterface
}

func (s *NotificationService) GetAll(interval string, timezone time.Location) ([]models2.Notification, error) {
	var suitableNotifications = make([]models2.Notification, 0)
	notifications, _ := s.Notifications.GetAll()

	for i, n := range notifications {
		notifications[i] = n.ConvertInTimezone(timezone)
	}

	if !isInterval(intervals, interval) {
		return notifications, nil
	}

	var limit time.Time = identifyLimit(interval)
	now := time.Now().UTC()

	for _, e := range notifications {
		if now.After(e.TimeUTC) && limit.Before(e.TimeUTC) {
			suitableNotifications = append(suitableNotifications, e)
		}
	}

	return suitableNotifications, nil
}

func (s *NotificationService) Create(notification models2.Notification) (models2.Notification, error) {
	notification.TimeUTC = notification.Time.UTC()
	return s.Notifications.Create(notification)
}

func (s *NotificationService) Update(id int, notification models2.Notification) (models2.Notification, error) {
	return s.Notifications.Update(id, notification)
}

package services

import (
	"time"
	"workshop2/models"
)

//go:generate mockgen -destination=../mocks/repositories/notification.go -package=mocks . NotificationRepositoryInterface
type NotificationRepositoryInterface interface {
	GetAll() ([]models.Notification, error)
	Create(notification models.Notification) (models.Notification, error)
	Update(id int, notification models.Notification) (models.Notification, error)
}

type NotificationService struct {
	Notifications NotificationRepositoryInterface
}

func (s *NotificationService) GetAll(interval string, timezone time.Location) ([]models.Notification, error) {
	var suitableNotifications = make([]models.Notification, 0)
	notifications, err := s.Notifications.GetAll()
	if err != nil {
		return notifications, err
	}

	for i, n := range notifications {
		notifications[i] = n.ConvertInTimezone(timezone)
	}

	now := time.Now().UTC()
	limit, err := identifyLimit(interval, now)
	if err != nil {
		return notifications, nil
	}

	for _, e := range notifications {
		if now.After(e.TimeUTC) && limit.Before(e.TimeUTC) {
			suitableNotifications = append(suitableNotifications, e)
		}
	}

	return suitableNotifications, nil
}

func (s *NotificationService) Create(notification models.Notification) (models.Notification, error) {
	notification.TimeUTC = notification.Time.UTC()
	return s.Notifications.Create(notification)
}

func (s *NotificationService) Update(id int, notification models.Notification) (models.Notification, error) {
	return s.Notifications.Update(id, notification)
}

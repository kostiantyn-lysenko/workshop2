package services

import (
	"time"
	"workshop2/internal/app/models"
)

type NotificationRepositoryInterface interface {
	GetAll() []models.Notification
	Create(notification models.Notification) models.Notification
	Update(id int, notification models.Notification) (models.Notification, error)
}

type NotificationService struct {
	Notifications NotificationRepositoryInterface
}

func (s *NotificationService) GetAll(interval string) []models.Notification {
	var suitableNotifications = make([]models.Notification, 0)
	notifications := s.Notifications.GetAll()

	if !isInterval(intervals, interval) {
		return notifications
	}

	var limit time.Time = identifyLimit()
	now := time.Now()

	for _, e := range notifications {
		if now.After(e.Time) && limit.Before(e.Time) {
			suitableNotifications = append(suitableNotifications, e)
		}
	}

	return suitableNotifications
}

func (s *NotificationService) Create(notification models.Notification) models.Notification {
	return s.Notifications.Create(notification)
}

func (s *NotificationService) Update(id int, notification models.Notification) (models.Notification, error) {
	return s.Notifications.Update(id, notification)
}

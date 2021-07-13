package services

import (
	"time"
	"workshop2/internal/app/models"
)

type NotificationRepositoryInterface interface {
	GetAll() ([]models.Notification, error)
	Create(notification models.Notification) (models.Notification, error)
	Update(id int, notification models.Notification) (models.Notification, error)
}

type NotificationService struct {
	Notifications NotificationRepositoryInterface
}

func (s *NotificationService) GetAll(interval string) ([]models.Notification, error) {
	var suitableNotifications = make([]models.Notification, 0)
	notifications, _ := s.Notifications.GetAll()

	if !isInterval(intervals, interval) {
		return notifications, nil
	}

	var limit time.Time = identifyLimit(interval)
	now := time.Now()

	for _, e := range notifications {
		if now.After(e.Time) && limit.Before(e.Time) {
			suitableNotifications = append(suitableNotifications, e)
		}
	}

	return suitableNotifications, nil
}

func (s *NotificationService) Create(notification models.Notification) (models.Notification, error) {
	return s.Notifications.Create(notification)
}

func (s *NotificationService) Update(id int, notification models.Notification) (models.Notification, error) {
	return s.Notifications.Update(id, notification)
}

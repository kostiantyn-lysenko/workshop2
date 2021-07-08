package services

import (
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

func (s *NotificationService) GetAll() []models.Notification {
	return s.Notifications.GetAll()
}

func (s *NotificationService) Create(notification models.Notification) models.Notification {
	return s.Notifications.Create(notification)
}

func (s *NotificationService) Update(id int, notification models.Notification) (models.Notification, error) {
	return s.Notifications.Update(id, notification)
}

package services

import (
	"workshop2/internal/app/models"
	"workshop2/internal/app/repositories"
)

type NotificationService struct {
	notifications repositories.NotificationRepository
}

func (s *NotificationService) GetAll() []models.Notification {
	return s.notifications.GetAll()
}

func (s *NotificationService) Create(notification *models.Notification) *[]models.Notification {
	return s.notifications.Create(notification)
}

func (s *NotificationService) Update(id int, notification *models.Notification) (*models.Notification, bool, error) {
	return s.notifications.Update(id, notification)
}

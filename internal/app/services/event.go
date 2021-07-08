package services

import (
	"workshop2/internal/app/models"
)

type EventRepositoryInterface interface {
	GetAll() []models.Event
	Get(id int) (models.Event, error)
	Create(event models.Event) models.Event
	Update(id int, newEvent models.Event) (models.Event, error)
	Delete(id int) error
}

type EventService struct {
	Events EventRepositoryInterface
}

func (s *EventService) GetAll() []models.Event {
	return s.Events.GetAll()
}

func (s *EventService) Get(id int) (models.Event, error) {
	return s.Events.Get(id)
}

func (s *EventService) Create(event models.Event) models.Event {
	return s.Events.Create(event)
}

func (s *EventService) Update(id int, event models.Event) (models.Event, error) {
	return s.Events.Update(id, event)
}

func (s *EventService) Delete(id int) error {
	return s.Events.Delete(id)
}

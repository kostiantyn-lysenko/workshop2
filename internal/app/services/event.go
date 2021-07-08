package services

import (
	"errors"
	"workshop2/internal/app/models"
	"workshop2/internal/app/repositories"
)

type EventService struct {
	events repositories.EventRepository
}

func (s *EventService) GetAll() []models.Event {
	return s.events.GetAll()
}

func (s *EventService) Get(id int) (*models.Event, error) {
	event, ok := s.events.Get(id)

	if !ok {
		return nil, errors.New("book with that ID does not exists in database")
	}
	return event, nil
}

func (s *EventService) Create(event *models.Event) *[]models.Event {
	return s.events.Create(event)
}

func (s *EventService) Update(id int, event *models.Event) (*models.Event, bool, error) {
	return s.events.Update(id, event)
}

func (s *EventService) Delete(id int) (bool, error) {
	return s.events.Delete(id)
}

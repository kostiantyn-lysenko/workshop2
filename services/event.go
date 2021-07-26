package services

import (
	"time"
	"workshop2/models"
)

//go:generate mockgen -destination=../mocks/repositories/event.go -package=mocks . EventRepositoryInterface
type EventRepositoryInterface interface {
	GetAll() ([]models.Event, error)
	Get(id int) (models.Event, error)
	Create(event models.Event) (models.Event, error)
	Update(id int, newEvent models.Event) (models.Event, error)
	Delete(id int) error
}

type EventService struct {
	Events EventRepositoryInterface
	Users  UserRepositoryInterface
}

func (s *EventService) GetAll(interval string, timezone time.Location) ([]models.Event, error) {
	var suitableEvents = make([]models.Event, 0)
	events, err := s.Events.GetAll()
	if err != nil {
		return events, err
	}

	for i, e := range events {
		events[i] = e.ConvertInTimezone(timezone)
	}

	now := time.Now().UTC()
	limit, err := identifyLimit(interval, now)
	if err != nil {
		return events, nil
	}

	for _, e := range events {
		if now.After(e.TimeUTC) && limit.Before(e.TimeUTC) {
			suitableEvents = append(suitableEvents, e)
		}
	}

	return suitableEvents, nil
}

func (s *EventService) Get(id int) (models.Event, error) {
	return s.Events.Get(id)
}

func (s *EventService) Create(event models.Event) (models.Event, error) {
	event.TimeUTC = event.Time.UTC()
	return s.Events.Create(event)
}

func (s *EventService) Update(id int, event models.Event) (models.Event, error) {
	return s.Events.Update(id, event)
}

func (s *EventService) Delete(id int) error {
	return s.Events.Delete(id)
}

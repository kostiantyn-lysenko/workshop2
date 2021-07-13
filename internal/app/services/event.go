package services

import (
	"time"
	"workshop2/internal/app/models"
)

type EventRepositoryInterface interface {
	GetAll() ([]models.Event, error)
	Get(id int) (models.Event, error)
	Create(event models.Event) (models.Event, error)
	Update(id int, newEvent models.Event) (models.Event, error)
	Delete(id int) error
}

type EventService struct {
	Events EventRepositoryInterface
}

func (s *EventService) GetAll(interval string) ([]models.Event, error) {
	var suitableEvents = make([]models.Event, 0)
	events, _ := s.Events.GetAll()

	if !isInterval(intervals, interval) {
		return events, nil
	}

	var limit time.Time = identifyLimit(interval)
	now := time.Now()

	for _, e := range events {
		if now.After(e.Time) && limit.Before(e.Time) {
			suitableEvents = append(suitableEvents, e)
		}
	}

	return suitableEvents, nil
}

func (s *EventService) Get(id int) (models.Event, error) {
	return s.Events.Get(id)
}

func (s *EventService) Create(event models.Event) (models.Event, error) {
	return s.Events.Create(event)
}

func (s *EventService) Update(id int, event models.Event) (models.Event, error) {
	return s.Events.Update(id, event)
}

func (s *EventService) Delete(id int) error {
	return s.Events.Delete(id)
}

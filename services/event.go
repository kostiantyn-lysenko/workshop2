package services

import (
	"time"
	models2 "workshop2/models"
)

type EventRepositoryInterface interface {
	GetAll() ([]models2.Event, error)
	Get(id int) (models2.Event, error)
	Create(event models2.Event) (models2.Event, error)
	Update(id int, newEvent models2.Event) (models2.Event, error)
	Delete(id int) error
}

type EventService struct {
	Events EventRepositoryInterface
	Users  UserRepositoryInterface
}

func (s *EventService) GetAll(interval string, timezone time.Location) ([]models2.Event, error) {
	var suitableEvents = make([]models2.Event, 0)
	events, _ := s.Events.GetAll()

	for i, e := range events {
		events[i] = e.ConvertInTimezone(timezone)
	}

	if !isInterval(intervals, interval) {
		return events, nil
	}

	now := time.Now().UTC()
	var limit, err = identifyLimit(interval, now)
	if err != nil {
		return suitableEvents, nil
	}

	for _, e := range events {
		if now.After(e.TimeUTC) && limit.Before(e.TimeUTC) {
			suitableEvents = append(suitableEvents, e)
		}
	}

	return suitableEvents, nil
}

func (s *EventService) Get(id int) (models2.Event, error) {
	return s.Events.Get(id)
}

func (s *EventService) Create(event models2.Event) (models2.Event, error) {
	event.TimeUTC = event.Time.UTC()
	return s.Events.Create(event)
}

func (s *EventService) Update(id int, event models2.Event) (models2.Event, error) {
	return s.Events.Update(id, event)
}

func (s *EventService) Delete(id int) error {
	return s.Events.Delete(id)
}

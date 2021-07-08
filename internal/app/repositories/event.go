package repositories

import (
	"sync"
	"workshop2/internal/app/errors"
	"workshop2/internal/app/models"
)

type EventRepository struct {
	Events []models.Event
	sync.RWMutex
}

func (r *EventRepository) GetAll() []models.Event {
	r.RLock()
	defer r.RUnlock()
	return r.Events
}

func (r *EventRepository) Get(id int) (models.Event, error) {
	var event models.Event

	r.RLock()
	defer r.RUnlock()
	for _, e := range r.Events {
		if e.ID == id {
			event = e
			return event, nil
		}
	}

	return event, &errors.NotFoundError{}
}

func (r *EventRepository) Create(event models.Event) models.Event {
	r.RLock()
	id := len(r.Events) + 1
	r.RUnlock()
	event.ID = id

	r.Lock()
	defer r.Unlock()
	r.Events = append(r.Events, event)

	return event
}

func (r *EventRepository) Update(id int, newEvent models.Event) (models.Event, error) {

	newEvent.ID = id
	r.Lock()
	defer r.Unlock()
	for i, e := range r.Events {
		if e.ID == newEvent.ID {
			r.Events[i] = newEvent

			return newEvent, nil
		}
	}

	return newEvent, &errors.NotFoundError{}
}

func (r *EventRepository) Delete(id int) error {

	r.Lock()
	defer r.Unlock()
	for i, e := range r.Events {
		if e.ID == id {
			r.Events = append(r.Events[:i], r.Events[i+1:]...)
			return nil
		}
	}

	return &errors.NotFoundError{}
}

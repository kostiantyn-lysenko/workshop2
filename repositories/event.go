package repositories

import (
	"sync"
	errs2 "workshop2/errs"
	models2 "workshop2/models"
)

type EventRepository struct {
	Events []models2.Event
	sync.RWMutex
}

func (r *EventRepository) GetAll() ([]models2.Event, error) {
	r.RLock()
	defer r.RUnlock()
	return r.Events, nil
}

func (r *EventRepository) Get(id int) (models2.Event, error) {
	r.RLock()
	defer r.RUnlock()
	for _, e := range r.Events {
		if e.ID == id {
			return e, nil
		}
	}

	return models2.Event{}, &errs2.EventNotFoundError{}
}

func (r *EventRepository) Create(event models2.Event) (models2.Event, error) {
	r.Lock()
	defer r.Unlock()

	id := 1
	if len(r.Events) > 0 {
		id = (r.Events[len(r.Events)-1]).ID + 1
	}
	event.ID = id

	r.Events = append(r.Events, event)

	return event, nil
}

func (r *EventRepository) Update(id int, newEvent models2.Event) (models2.Event, error) {

	newEvent.ID = id
	r.Lock()
	defer r.Unlock()
	for i, e := range r.Events {
		if e.ID == newEvent.ID {
			r.Events[i] = newEvent

			return newEvent, nil
		}
	}

	return newEvent, &errs2.EventNotFoundError{}
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

	return &errs2.EventNotFoundError{}
}

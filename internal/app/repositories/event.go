package repositories

import (
	"errors"
	"workshop2/internal/app/models"
	"workshop2/storage"
)

type EventRepository struct {
}

func (r *EventRepository) GetAll() []models.Event {

	return storage.DB.Events
}

func (r *EventRepository) Get(id int) (*models.Event, bool) {

	var event models.Event
	var found bool

	storage.DB.Lock()
	for _, e := range storage.DB.Events {
		if e.ID == id {
			event = e
			found = true
			break
		}
	}
	storage.DB.Unlock()

	return &event, found
}

func (r *EventRepository) Create(event *models.Event) *[]models.Event {
	storage.DB.Lock()
	id := len(storage.DB.Events) + 1
	storage.DB.Unlock()
	event.ID = id
	storage.DB.Events = append(storage.DB.Events, *event)

	return &storage.DB.Events
}

func (r *EventRepository) Update(id int, event *models.Event) (*models.Event, bool, error) {

	var changed bool

	event, ok := r.Get(id)
	if !ok {
		return nil, changed, errors.New("event not found")
	}

	event.ID = id
	storage.DB.Lock()
	for i, e := range storage.DB.Events {
		if e.ID == event.ID {
			storage.DB.Events[i] = *event

			changed = true
		}
	}
	storage.DB.Unlock()

	return event, changed, nil
}

func (r *EventRepository) Delete(id int) (bool, error) {
	var deleted bool

	_, ok := r.Get(id)
	if !ok {
		return deleted, errors.New("event not found")
	}

	storage.DB.Lock()
	for i, e := range storage.DB.Events {
		if e.ID == id {
			storage.DB.Events = append(storage.DB.Events[:i], storage.DB.Events[i+1:]...)
			deleted = true
		}
	}
	storage.DB.Unlock()

	return deleted, nil
}

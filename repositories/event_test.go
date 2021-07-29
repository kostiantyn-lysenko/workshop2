package repositories

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
	"workshop2/errs"
	"workshop2/models"
)

func TestEventRepository_Create(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		event models.Event
		err   error
	}

	type payload struct {
		event  models.Event
		events []models.Event
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "the new id should be one more than the last event",
			expected: expected{
				event: models.Event{
					ID:          11,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: nil,
			},
			payload: payload{
				event: models.Event{
					ID:          0,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          10,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
		{
			name: "ok",
			expected: expected{
				event: models.Event{
					ID:          1,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: nil,
			},
			payload: payload{
				event: models.Event{
					ID:          0,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
	}

	for _, tt := range tc {
		s := EventRepository{
			Events: tt.payload.events,
		}

		event, err := s.Create(tt.payload.event)

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create event", tt.name)
			continue
		}

		Expect(event).To(Equal(tt.expected.event), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventRepository_Delete(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		err error
	}

	type payload struct {
		id     int
		events []models.Event
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "event not found",
			expected: expected{
				err: &errs.EventNotFoundError{},
			},
			payload: payload{
				id: 12,
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          10,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          123,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
		{
			name: "event not found",
			expected: expected{
				err: &errs.EventNotFoundError{},
			},
			payload: payload{
				id:     0,
				events: []models.Event{},
			},
		},
		{
			name: "ok",
			expected: expected{
				err: nil,
			},
			payload: payload{
				id: 12,
				events: []models.Event{
					{
						ID:          12,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
	}

	for _, tt := range tc {
		s := EventRepository{
			Events: tt.payload.events,
		}

		err := s.Delete(tt.payload.id)

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to delete event", tt.name)
			continue
		}

		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventRepository_Get(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		event models.Event
		err   error
	}

	type payload struct {
		id     int
		events []models.Event
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "event not found",
			expected: expected{
				event: models.Event{},
				err:   &errs.EventNotFoundError{},
			},
			payload: payload{
				id: 12,
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          10,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          123,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
		{
			name: "event not found, negative id",
			expected: expected{
				event: models.Event{},
				err:   &errs.EventNotFoundError{},
			},
			payload: payload{
				id: -2,
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          10,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          123,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
		{
			name: "event not found in empty slice",
			expected: expected{
				event: models.Event{},
				err:   &errs.EventNotFoundError{},
			},
			payload: payload{
				id:     1,
				events: []models.Event{},
			},
		},
		{
			name: "ok",
			expected: expected{
				event: models.Event{
					ID:          12,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: nil,
			},
			payload: payload{
				id: 12,
				events: []models.Event{
					{
						ID:          12,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          1,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          2,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
	}

	for _, tt := range tc {
		s := EventRepository{
			Events: tt.payload.events,
		}

		event, err := s.Get(tt.payload.id)

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to get event", tt.name)
			continue
		}

		Expect(event).To(Equal(tt.expected.event), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventRepository_GetAll(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		events []models.Event
		err    error
	}

	type payload struct {
		events []models.Event
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "events not found",
			expected: expected{
				events: []models.Event{},
				err:    nil,
			},
			payload: payload{
				events: []models.Event{},
			},
		},
		{
			name: "ok",
			expected: expected{
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          10,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
				err: nil,
			},
			payload: payload{
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          10,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
	}

	for _, tt := range tc {
		s := EventRepository{
			Events: tt.payload.events,
		}

		events, err := s.GetAll()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to get all events", tt.name)
			continue
		}

		Expect(events).To(Equal(tt.expected.events), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventRepository_Update(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		event models.Event
		err   error
	}

	type payload struct {
		id     int
		event  models.Event
		events []models.Event
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "event not found",
			expected: expected{
				event: models.Event{
					ID:          1,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: &errs.EventNotFoundError{},
			},
			payload: payload{
				id: 1,
				events: []models.Event{
					{
						ID:          0,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
		{
			name: "event not found in empty slice",
			expected: expected{
				event: models.Event{
					ID:          1,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: &errs.EventNotFoundError{},
			},
			payload: payload{
				id: 1,
				event: models.Event{
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				events: []models.Event{},
			},
		},
		{
			name: "ok",
			expected: expected{
				event: models.Event{
					ID:          12,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: nil,
			},
			payload: payload{
				id: 12,
				event: models.Event{
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				events: []models.Event{
					{
						ID:          12,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          1,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
					{
						ID:          2,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now,
						Description: "desc",
					},
				},
			},
		},
	}

	for _, tt := range tc {
		s := EventRepository{
			Events: tt.payload.events,
		}

		event, err := s.Update(tt.payload.id, tt.payload.event)

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to update event", tt.name)
			continue
		}

		Expect(event).To(Equal(tt.expected.event), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

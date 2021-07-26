package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"log"
	"testing"
	"time"
	mocks "workshop2/mocks/repositories"
	"workshop2/models"
)

func TestEventService_Create(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		event models.Event
		err   error
	}

	type payload struct {
		event models.Event
		eRepo func(mockCtrlr *gomock.Controller) EventRepositoryInterface
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "repo returned an error",
			expected: expected{
				event: models.Event{
					ID:          11,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: errors.New("err"),
			},
			payload: payload{
				event: models.Event{
					ID:          11,
					Title:       "title",
					TimeUTC:     now,
					Time:        now,
					Description: "desc",
				},
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Create(
							models.Event{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(models.Event{}, errors.New("err")).
						Times(1)
					return eRepoMock
				},
			},
		},
		{
			name: "successful creating",
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
					ID:          11,
					Title:       "title",
					TimeUTC:     now,
					Time:        now,
					Description: "desc",
				},
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Create(
							models.Event{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(
							models.Event{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := EventService{
			Events: tt.payload.eRepo(mockCtrlr),
		}

		event, err := s.Create(tt.payload.event)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create event", tt.name)
			continue
		}

		Expect(event).To(Equal(tt.expected.event), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventService_Delete(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		err error
	}

	type payload struct {
		id    int
		eRepo func(mockCtrlr *gomock.Controller) EventRepositoryInterface
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "repo returned an error",
			expected: expected{
				err: errors.New("err"),
			},
			payload: payload{
				id: 1,
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Delete(1).
						Return(errors.New("err")).
						Times(1)
					return eRepoMock
				},
			},
		},
		{
			name: "ok",
			expected: expected{
				err: nil,
			},
			payload: payload{
				id: 1,
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Delete(1).
						Return(nil).
						Times(1)
					return eRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := EventService{
			Events: tt.payload.eRepo(mockCtrlr),
		}

		err := s.Delete(tt.payload.id)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create event", tt.name)
			continue
		}

		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventService_Get(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		event models.Event
		err   error
	}

	type payload struct {
		id    int
		eRepo func(mockCtrlr *gomock.Controller) EventRepositoryInterface
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "repo returned an error",
			expected: expected{
				event: models.Event{},
				err:   errors.New("err"),
			},
			payload: payload{
				id: 1,
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Get(1).
						Return(models.Event{}, errors.New("err")).
						Times(1)
					return eRepoMock
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
				id: 1,
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Get(1).
						Return(
							models.Event{
								ID:          1,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := EventService{
			Events: tt.payload.eRepo(mockCtrlr),
		}

		event, err := s.Get(tt.payload.id)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create event", tt.name)
			continue
		}

		Expect(event).To(Equal(tt.expected.event), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventService_GetAll(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		events []models.Event
		err    error
	}

	type payload struct {
		interval  string
		timezone  time.Location
		eventRepo func(mockCtrlr *gomock.Controller) EventRepositoryInterface
	}

	now := time.Now()
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println("Bad timezone.")
		return
	}

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "repo returned an error",
			expected: expected{
				events: make([]models.Event, 0),
				err:    errors.New("err"),
			},
			payload: payload{
				interval: "interval",
				timezone: time.Location{},
				eventRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						GetAll().
						Return(make([]models.Event, 0), errors.New("err")).
						Times(1)
					return eRepoMock
				},
			},
		},
		{
			name: "invalid interval",
			expected: expected{
				events: []models.Event{
					{
						ID:          1,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now.UTC().In(loc),
						Description: "desc",
					},
				},
				err: nil,
			},
			payload: payload{
				interval: "notValidInterval",
				timezone: *loc,
				eventRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						GetAll().
						Return(
							[]models.Event{
								{
									ID:          1,
									Title:       "title",
									TimeUTC:     now.UTC(),
									Time:        now,
									Description: "desc",
								},
							},
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
		{
			name: "no interval",
			expected: expected{
				events: []models.Event{
					{
						ID:          1,
						Title:       "title",
						TimeUTC:     now.UTC(),
						Time:        now.UTC().In(loc),
						Description: "desc",
					},
				},
				err: nil,
			},
			payload: payload{
				interval: "",
				timezone: *loc,
				eventRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						GetAll().
						Return(
							[]models.Event{
								{
									ID:          1,
									Title:       "title",
									TimeUTC:     now.UTC(),
									Time:        now,
									Description: "desc",
								},
							},
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
		{
			name: "no events found",
			expected: expected{
				events: make([]models.Event, 0),
				err:    nil,
			},
			payload: payload{
				interval: "interval",
				timezone: *loc,
				eventRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						GetAll().
						Return(
							make([]models.Event, 0),
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
		{
			name: "okay, with no extra events",
			expected: expected{
				events: []models.Event{
					{
						ID:          1,
						Title:       "title1",
						TimeUTC:     now.UTC(),
						Time:        now.UTC().In(loc),
						Description: "desc",
					},
				},
				err: nil,
			},
			payload: payload{
				interval: "week",
				timezone: *loc,
				eventRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						GetAll().
						Return(
							[]models.Event{
								{
									ID:          1,
									Title:       "title1",
									TimeUTC:     now.UTC(),
									Time:        now,
									Description: "desc",
								},
								{
									ID:          2,
									Title:       "title2",
									TimeUTC:     now.UTC().AddDate(0, 0, 8),
									Time:        now.AddDate(0, 0, 1),
									Description: "desc",
								},
								{
									ID:          3,
									Title:       "title3",
									TimeUTC:     now.UTC().AddDate(0, 0, -8),
									Time:        now,
									Description: "desc",
								},
							},
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := EventService{
			Events: tt.payload.eventRepo(mockCtrlr),
		}

		events, err := s.GetAll(tt.payload.interval, tt.payload.timezone)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create events", tt.name)
			continue
		}

		Expect(events).To(Equal(tt.expected.events), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestEventService_Update(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		event models.Event
		err   error
	}

	type payload struct {
		id    int
		event models.Event
		eRepo func(mockCtrlr *gomock.Controller) EventRepositoryInterface
	}

	now := time.Now()

	tc := []struct {
		name string
		expected
		payload
	}{
		{
			name: "repo returned an error",
			expected: expected{
				event: models.Event{
					ID:          1,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: errors.New("err"),
			},
			payload: payload{
				id: 1,
				event: models.Event{
					ID:          0,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Update(
							1,
							models.Event{
								ID:          0,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(
							models.Event{
								ID:          1,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
							errors.New("err"),
						).
						Times(1)
					return eRepoMock
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
				id: 1,
				event: models.Event{
					ID:          0,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				eRepo: func(mockCtrlr *gomock.Controller) EventRepositoryInterface {
					eRepoMock := mocks.NewMockEventRepositoryInterface(mockCtrlr)
					eRepoMock.EXPECT().
						Update(
							1,
							models.Event{
								ID:          0,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(
							models.Event{
								ID:          1,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
							nil,
						).
						Times(1)
					return eRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := EventService{
			Events: tt.payload.eRepo(mockCtrlr),
		}

		event, err := s.Update(tt.payload.id, tt.payload.event)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create event", tt.name)
			continue
		}

		Expect(event).To(Equal(tt.expected.event), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

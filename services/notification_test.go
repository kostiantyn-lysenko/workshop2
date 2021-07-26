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

func TestNotificationService_Create(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		notif models.Notification
		err   error
	}

	type payload struct {
		notif  models.Notification
		nsRepo func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface
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
				notif: models.Notification{
					ID:          11,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: errors.New("err"),
			},
			payload: payload{
				notif: models.Notification{
					ID:          11,
					Title:       "title",
					TimeUTC:     now,
					Time:        now,
					Description: "desc",
				},
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						Create(
							models.Notification{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(models.Notification{}, errors.New("err")).
						Times(1)
					return nsRepoMock
				},
			},
		},
		{
			name: "successful creating",
			expected: expected{
				notif: models.Notification{
					ID:          11,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: nil,
			},
			payload: payload{
				notif: models.Notification{
					ID:          11,
					Title:       "title",
					TimeUTC:     now,
					Time:        now,
					Description: "desc",
				},
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						Create(
							models.Notification{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(
							models.Notification{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
							nil,
						).
						Times(1)
					return nsRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := NotificationService{
			Notifications: tt.payload.nsRepo(mockCtrlr),
		}

		notif, err := s.Create(tt.payload.notif)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create notification", tt.name)
			continue
		}

		Expect(notif).To(Equal(tt.expected.notif), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestNotificationService_GetAll(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		ns  []models.Notification
		err error
	}

	type payload struct {
		interval string
		timezone time.Location
		nsRepo   func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface
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
				ns:  make([]models.Notification, 0),
				err: errors.New("err"),
			},
			payload: payload{
				interval: "interval",
				timezone: time.Location{},
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						GetAll().
						Return(make([]models.Notification, 0), errors.New("err")).
						Times(1)
					return nsRepoMock
				},
			},
		},
		{
			name: "invalid interval",
			expected: expected{
				ns: []models.Notification{
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
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						GetAll().
						Return(
							[]models.Notification{
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
					return nsRepoMock
				},
			},
		},
		{
			name: "no interval",
			expected: expected{
				ns: []models.Notification{
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
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						GetAll().
						Return(
							[]models.Notification{
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
					return nsRepoMock
				},
			},
		},
		{
			name: "no notifications found",
			expected: expected{
				ns:  make([]models.Notification, 0),
				err: nil,
			},
			payload: payload{
				interval: "interval",
				timezone: *loc,
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						GetAll().
						Return(
							make([]models.Notification, 0),
							nil,
						).
						Times(1)
					return nsRepoMock
				},
			},
		},
		{
			name: "okay, with no extra notifications",
			expected: expected{
				ns: []models.Notification{
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
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						GetAll().
						Return(
							[]models.Notification{
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
					return nsRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := NotificationService{
			Notifications: tt.payload.nsRepo(mockCtrlr),
		}

		ns, err := s.GetAll(tt.payload.interval, tt.payload.timezone)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create notification", tt.name)
			continue
		}

		Expect(ns).To(Equal(tt.expected.ns), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

func TestNotificationService_Update(t *testing.T) {
	RegisterTestingT(t)

	type expected struct {
		notif models.Notification
		err   error
	}

	type payload struct {
		id     int
		notif  models.Notification
		nsRepo func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface
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
				notif: models.Notification{},
				err:   errors.New("err"),
			},
			payload: payload{
				id: 11,
				notif: models.Notification{
					ID:          1,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						Update(
							11,
							models.Notification{
								ID:          1,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(models.Notification{}, errors.New("err")).
						Times(1)
					return nsRepoMock
				},
			},
		},
		{
			name: "successful updated",
			expected: expected{
				notif: models.Notification{
					ID:          11,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				err: nil,
			},
			payload: payload{
				id: 11,
				notif: models.Notification{
					ID:          1,
					Title:       "title",
					TimeUTC:     now.UTC(),
					Time:        now,
					Description: "desc",
				},
				nsRepo: func(mockCtrlr *gomock.Controller) NotificationRepositoryInterface {
					nsRepoMock := mocks.NewMockNotificationRepositoryInterface(mockCtrlr)
					nsRepoMock.EXPECT().
						Update(
							11,
							models.Notification{
								ID:          1,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
						).
						Return(
							models.Notification{
								ID:          11,
								Title:       "title",
								TimeUTC:     now.UTC(),
								Time:        now,
								Description: "desc",
							},
							nil,
						).
						Times(1)
					return nsRepoMock
				},
			},
		},
	}

	for _, tt := range tc {
		mockCtrlr := gomock.NewController(t)
		s := NotificationService{
			Notifications: tt.payload.nsRepo(mockCtrlr),
		}

		notif, err := s.Update(tt.payload.id, tt.payload.notif)
		mockCtrlr.Finish()

		if err != nil {
			Expect(err).To(Equal(tt.expected.err), "failed to create notification", tt.name)
			continue
		}

		Expect(notif).To(Equal(tt.expected.notif), "message", tt.name)
		Expect(err).To(BeNil(), tt.name)
	}
}

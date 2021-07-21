package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	errs2 "workshop2/errs"
	models2 "workshop2/models"

	"github.com/gorilla/mux"
)

type NotificationServiceInterface interface {
	GetAll(interval string, timezone time.Location) ([]models2.Notification, error)
	Create(notification models2.Notification) (models2.Notification, error)
	Update(id int, notification models2.Notification) (models2.Notification, error)
}

type NotificationController struct {
	Notifications NotificationServiceInterface
	Auth          AuthServiceInterface
}

func (c *NotificationController) GetAll(w http.ResponseWriter, r *http.Request) {
	interval := r.FormValue("interval")
	initHeaders(w)

	loc, err := GetUserTimezone(r, c.Auth)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	notifications, _ := c.Notifications.GetAll(interval, *loc)
	respond(w, notifications, http.StatusOK)
}

func (c *NotificationController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	var notification models2.Notification

	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		respondWithError(w, errs2.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	notification, _ = c.Notifications.Create(notification)
	respond(w, notification, http.StatusOK)
}

func (c *NotificationController) Update(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, errs2.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	var notification models2.Notification
	err = json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		respondWithError(w, errs2.NewFailedRequestParsingError(), http.StatusBadRequest)
		return
	}

	updatedEvent, err := c.Notifications.Update(id, notification)
	if err != nil {
		respondWithError(w, err, http.StatusUnprocessableEntity)
		return
	}

	respond(w, updatedEvent, http.StatusOK)
}

package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"

	"github.com/gorilla/mux"
)

type NotificationServiceInterface interface {
	GetAll(interval string) ([]models.Notification, error)
	Create(notification models.Notification) (models.Notification, error)
	Update(id int, notification models.Notification) (models.Notification, error)
}

type NotificationController struct {
	Notifications NotificationServiceInterface
}

func (c *NotificationController) GetAll(w http.ResponseWriter, r *http.Request) {
	interval := string(r.FormValue("interval"))

	initHeaders(w)
	w.WriteHeader(http.StatusOK)

	notifications, _ := c.Notifications.GetAll(interval)
	json.NewEncoder(w).Encode(notifications)
}

func (c *NotificationController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var notification models.Notification

	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.FailedRequestParsingError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	notification, _ = c.Notifications.Create(notification)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(notification)
}

func (c *NotificationController) Update(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.IdNotNumericError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var notification models.Notification
	err = json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.FailedRequestParsingError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	updatedEvent, err := c.Notifications.Update(id, notification)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}

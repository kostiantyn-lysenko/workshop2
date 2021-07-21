package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"

	"github.com/gorilla/mux"
)

type EventServiceInterface interface {
	GetAll(interval string, timezone time.Location) ([]models.Event, error)
	Get(id int) (models.Event, error)
	Create(event models.Event) (models.Event, error)
	Update(id int, newEvent models.Event) (models.Event, error)
	Delete(id int) error
}

type EventController struct {
	Events EventServiceInterface
	Auth   AuthServiceInterface
}

func (c *EventController) GetAll(w http.ResponseWriter, r *http.Request) {
	interval := r.FormValue("interval")

	initHeaders(w)
	w.WriteHeader(http.StatusOK)

	loc, err := GetUserTimezone(r, c.Auth)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeErr := json.NewEncoder(w).Encode(err.Error())
		if encodeErr != nil {
			log.Fatal(err)
		}
		return
	}

	events, _ := c.Events.GetAll(interval, *loc)
	json.NewEncoder(w).Encode(events)
}

func (c *EventController) Get(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.IdNotNumericError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	event, err := c.Events.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func (c *EventController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	event, _ = c.Events.Create(event)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (c *EventController) Update(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.IdNotNumericError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var event models.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = errs.NewFailedRequestParsingError()
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	updatedEvent, err := c.Events.Update(id, event)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}

func (c *EventController) Delete(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.IdNotNumericError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.Events.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

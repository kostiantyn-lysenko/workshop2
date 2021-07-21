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

type EventServiceInterface interface {
	GetAll(interval string, timezone time.Location) ([]models2.Event, error)
	Get(id int) (models2.Event, error)
	Create(event models2.Event) (models2.Event, error)
	Update(id int, newEvent models2.Event) (models2.Event, error)
	Delete(id int) error
}

type EventController struct {
	Events EventServiceInterface
	Auth   AuthServiceInterface
}

func (c *EventController) GetAll(w http.ResponseWriter, r *http.Request) {
	interval := r.FormValue("interval")
	initHeaders(w)

	loc, err := GetUserTimezone(r, c.Auth)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	events, _ := c.Events.GetAll(interval, *loc)
	respond(w, events, http.StatusOK)
}

func (c *EventController) Get(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, errs2.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	event, err := c.Events.Get(id)
	if err != nil {
		respondWithError(w, err, http.StatusNotFound)
		return
	}

	respond(w, event, http.StatusOK)
}

func (c *EventController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	var event models2.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, errs2.NewFailedRequestParsingError(), http.StatusBadRequest)
		return
	}

	event, _ = c.Events.Create(event)
	respond(w, event, http.StatusCreated)
}

func (c *EventController) Update(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, errs2.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	var event models2.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, errs2.NewFailedRequestParsingError(), http.StatusBadRequest)
		return
	}

	updatedEvent, err := c.Events.Update(id, event)
	if err != nil {
		respondWithError(w, err, http.StatusUnprocessableEntity)
		return
	}

	respond(w, updatedEvent, http.StatusOK)
}

func (c *EventController) Delete(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, errs2.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	err = c.Events.Delete(id)
	if err != nil {
		respondWithError(w, err, http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"workshop2/errs"
	"workshop2/models"
	"workshop2/tokenizer"

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
	Events    EventServiceInterface
	Tokenizer tokenizer.Tokenizer
}

func (c *EventController) GetAll(w http.ResponseWriter, r *http.Request) {
	interval := r.FormValue("interval")
	initHeaders(w)

	loc, err := GetUserTimezone(r, c.Tokenizer)
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
		respondWithError(w, errs.NewIdNotNumericError(), http.StatusBadRequest)
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
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, errs.NewFailedRequestParsingError(), http.StatusBadRequest)
		return
	}

	event, _ = c.Events.Create(event)
	respond(w, event, http.StatusCreated)
}

func (c *EventController) Update(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, errs.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	var event models.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, errs.NewFailedRequestParsingError(), http.StatusBadRequest)
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
		respondWithError(w, errs.NewIdNotNumericError(), http.StatusBadRequest)
		return
	}

	err = c.Events.Delete(id)
	if err != nil {
		respondWithError(w, err, http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
}

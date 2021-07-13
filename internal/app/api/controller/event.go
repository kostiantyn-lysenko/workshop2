package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"workshop2/internal/app/errs"
	"workshop2/internal/app/models"

	"github.com/gorilla/mux"
)

type EventServiceInterface interface {
	GetAll(interval string) ([]models.Event, error)
	Get(id int) (models.Event, error)
	Create(event models.Event) (models.Event, error)
	Update(id int, newEvent models.Event) (models.Event, error)
	Delete(id int) error
}

type EventController struct {
	Events EventServiceInterface
}

func (e *EventController) GetAll(w http.ResponseWriter, r *http.Request) {
	interval := string(r.FormValue("interval"))

	initHeaders(w)
	w.WriteHeader(http.StatusOK)
	events, _ := e.Events.GetAll(interval)
	json.NewEncoder(w).Encode(events)
}

func (e *EventController) Get(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.IdNotNumericError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	event, err := e.Events.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func (e *EventController) Create(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.FailedRequestParsingError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	event, _ = e.Events.Create(event)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (e *EventController) Update(w http.ResponseWriter, r *http.Request) {
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
		err = &errs.FailedRequestParsingError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	updatedEvent, err := e.Events.Update(id, event)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}

func (e *EventController) Delete(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = &errs.IdNotNumericError{}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = e.Events.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

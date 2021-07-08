package api

import (
	"net/http"
	controller "workshop2/internal/app/api/controller"

	"github.com/gorilla/mux"
)

type API struct {
	port   string
	router *mux.Router
	prefix string
}

func New() *API {
	return &API{
		port:   ":8001",
		router: mux.NewRouter(),
		prefix: "/api/v1",
	}
}

func (api *API) Start() error {
	api.configureRoutes()
	return http.ListenAndServe(api.port, api.router)
}

func (api *API) configureRoutes() {
	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is Workshop2 API!"))
	}).Methods(http.MethodGet)

	api.router.HandleFunc(api.prefix+"/events", controller.GetAllEvents).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/events/{id}", controller.GetEvent).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/events", controller.CreateEvent).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/events/{id}", controller.UpdateEvent).Methods(http.MethodPut)
	api.router.HandleFunc(api.prefix+"/events/{id}", controller.DeleteEvent).Methods(http.MethodDelete)

	api.router.HandleFunc(api.prefix+"/notifications", controller.GetAllNotifications).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/notifications", controller.CreateNotification).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/notifications/{id}", controller.UpdateNotification).Methods(http.MethodPut)
}

package api

import (
	"net/http"
	"workshop2/internal/app/api/controller"
	"workshop2/internal/app/models"
	"workshop2/internal/app/repositories"
	"workshop2/internal/app/services"

	"github.com/gorilla/mux"
)

type API struct {
	port          string
	router        *mux.Router
	prefix        string
	events        controller.EventController
	notifications controller.NotificationController
}

func New() *API {
	return &API{
		port:   ":8001",
		router: mux.NewRouter(),
		prefix: "/api/v1",
		events: controller.EventController{
			Events: &services.EventService{
				Events: &repositories.EventRepository{
					Events: make([]models.Event, 0),
				},
			},
		},
		notifications: controller.NotificationController{
			Notifications: &services.NotificationService{
				Notifications: &repositories.NotificationRepository{
					Notifications: make([]models.Notification, 0),
				},
			},
		},
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

	api.router.HandleFunc(api.prefix+"/events", api.events.GetAll).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/events/{id}", api.events.Get).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/events", api.events.Create).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/events/{id}", api.events.Update).Methods(http.MethodPut)
	api.router.HandleFunc(api.prefix+"/events/{id}", api.events.Delete).Methods(http.MethodDelete)

	api.router.HandleFunc(api.prefix+"/notifications", api.notifications.GetAll).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/notifications", api.notifications.Create).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/notifications/{id}", api.notifications.Update).Methods(http.MethodPut)
}

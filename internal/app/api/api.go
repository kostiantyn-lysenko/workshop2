package api

import (
	"net/http"
	"time"
	"workshop2/internal/app/api/controller"
	"workshop2/internal/app/models"
	"workshop2/internal/app/repositories"
	"workshop2/internal/app/services"
	"workshop2/internal/app/utils"

	"github.com/gorilla/mux"
)

type API struct {
	port          string
	router        *mux.Router
	prefix        string
	events        controller.EventController
	notifications controller.NotificationController
	users         controller.UserController
	auth          controller.AuthController
}

func New() *API {
	validator := utils.NewValidator()

	return &API{
		port:   ":8002",
		router: mux.NewRouter(),
		prefix: "/api/v1",
		events: controller.EventController{
			Events: &services.EventService{
				Events: &repositories.EventRepository{
					Events: make([]models.Event, 0),
				},
			},
		},
		users: controller.UserController{
			Users: &services.UserService{
				Users: &repositories.UserRepository{
					Users:     make([]models.User, 0),
					Validator: validator,
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
		auth: controller.AuthController{
			Auth: services.NewAuth(
				&repositories.UserRepository{
					Users:     make([]models.User, 0),
					Validator: validator,
				},
				&utils.Validator{},
				time.Hour*6,
				time.Hour*24*31,
				"keyyt",
			),
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
	api.router.HandleFunc(api.prefix+"/events", api.events.GetAll).Queries("interval", "{interval}").Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/events/{id}", api.events.Get).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/events", api.events.Create).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/events/{id}", api.events.Update).Methods(http.MethodPut)
	api.router.HandleFunc(api.prefix+"/events/{id}", api.events.Delete).Methods(http.MethodDelete)

	api.router.HandleFunc(api.prefix+"/notifications", api.notifications.GetAll).Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/notifications", api.notifications.GetAll).Queries("interval", "{interval}").Methods(http.MethodGet)
	api.router.HandleFunc(api.prefix+"/notifications", api.notifications.Create).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/notifications/{id}", api.notifications.Update).Methods(http.MethodPut)

	api.router.HandleFunc(api.prefix+"/users", api.users.Create).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/sign-in", api.auth.SignIn).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/sign-up", api.auth.SignUp).Methods(http.MethodPost)
}

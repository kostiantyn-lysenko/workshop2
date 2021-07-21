package api

import (
	"log"
	"net/http"
	"time"
	controller2 "workshop2/api/controller"
	models2 "workshop2/models"
	repositories2 "workshop2/repositories"
	services2 "workshop2/services"
	utils2 "workshop2/utils"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type API struct {
	port          string
	router        *mux.Router
	prefix        string
	events        controller2.EventController
	notifications controller2.NotificationController
	users         controller2.UserController
	auth          controller2.AuthController
}

func New() *API {
	validator := utils2.NewValidator()

	userRepository := &repositories2.UserRepository{
		Users:     make([]models2.User, 0),
		Validator: validator,
	}

	authService := services2.NewAuth(
		userRepository,
		validator,
		time.Hour*6,
		time.Hour*24*31,
		"keyyt",
		jwt.SigningMethodHS256,
	)

	return &API{
		port:   ":8002",
		router: mux.NewRouter(),
		prefix: "/api/v1",
		events: controller2.EventController{
			Events: &services2.EventService{
				Events: &repositories2.EventRepository{
					Events: make([]models2.Event, 0),
				},
			},
			Auth: authService,
		},
		users: controller2.UserController{
			Users: &services2.UserService{
				userRepository,
			},
			Auth: authService,
		},
		notifications: controller2.NotificationController{
			Notifications: &services2.NotificationService{
				Notifications: &repositories2.NotificationRepository{
					Notifications: make([]models2.Notification, 0),
				},
			},
			Auth: authService,
		},
		auth: controller2.AuthController{
			Auth: authService,
		},
	}
}

func (api *API) Start() error {
	api.configureRoutes()
	return http.ListenAndServe(api.port, api.router)
}

func (api *API) configureRoutes() {
	authMiddleware := AuthenticationMiddleware{api.auth.Auth}
	api.router.Use(authMiddleware.Handle)

	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello! This is Workshop2 API!"))
		if err != nil {
			log.Fatal(err.Error())
		}

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

	api.router.HandleFunc(api.prefix+"/sign-in", api.auth.SignIn).Methods(http.MethodPost)
	api.router.HandleFunc(api.prefix+"/sign-up", api.auth.SignUp).Methods(http.MethodPost)

	api.router.HandleFunc(api.prefix+"/timezone", api.users.UpdateTimezone).Methods(http.MethodPut)
}

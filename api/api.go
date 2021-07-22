package api

import (
	"log"
	"net/http"
	"time"
	"workshop2/api/controller"
	"workshop2/models"
	"workshop2/repositories"
	"workshop2/services"
	"workshop2/utils"

	"github.com/golang-jwt/jwt"
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

	userRepository := &repositories.UserRepository{
		Users:     make([]models.User, 0),
		Validator: validator,
	}

	authService := services.NewAuth(
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
		events: controller.EventController{
			Events: &services.EventService{
				Events: &repositories.EventRepository{
					Events: make([]models.Event, 0),
				},
			},
			Auth: authService,
		},
		users: controller.UserController{
			Users: &services.UserService{
				userRepository,
			},
			Auth: authService,
		},
		notifications: controller.NotificationController{
			Notifications: &services.NotificationService{
				Notifications: &repositories.NotificationRepository{
					Notifications: make([]models.Notification, 0),
				},
			},
			Auth: authService,
		},
		auth: controller.AuthController{
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

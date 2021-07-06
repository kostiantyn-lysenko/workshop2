package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	port   string
	router *mux.Router
}

func New() *API {
	return &API{port: ":8001", router: mux.NewRouter()}
}

func (api *API) Start() error {
	api.configureRoutes()
	return http.ListenAndServe(api.port, api.router)
}

func (api *API) configureRoutes() {
	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is Workshop2 API!"))
	})
}

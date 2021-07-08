package controller

import "net/http"

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s Server) initializeRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", s.home)

	return http.Handler(r)
}

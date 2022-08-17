package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s Server) initializeRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", s.home).Methods(http.MethodGet)
	r.HandleFunc("/post/{id}", s.findPost).Methods(http.MethodGet)

	return http.Handler(r)
}

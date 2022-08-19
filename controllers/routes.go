package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s Server) initializeRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", s.home).Methods(http.MethodGet)
	r.HandleFunc("/post/{id:[0-9]+}", s.getOnePost).Methods(http.MethodGet)
	r.HandleFunc("/posts/", s.getAllPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/", s.createPost).Methods(http.MethodPost)
	r.HandleFunc("/post/{id:[0-9]+}", s.updatePost).Methods(http.MethodPut)
	r.HandleFunc("/post/{id:[0-9]+}", s.deletePost).Methods(http.MethodDelete)

	r.HandleFunc("/tag/{id:[0-9]+}", s.getOneTag).Methods(http.MethodGet)
	r.HandleFunc("/tags/", s.getAllTags).Methods(http.MethodGet)
	r.HandleFunc("/tags/", s.createTag).Methods(http.MethodPost)
	r.HandleFunc("/tag/{id:[0-9]+}", s.updateTag).Methods(http.MethodPut)
	r.HandleFunc("/tag/{id:[0-9]+}", s.deleteTag).Methods(http.MethodDelete)

	return r
}

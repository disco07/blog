package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) initializeRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	return http.Handler(r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "Running")
}
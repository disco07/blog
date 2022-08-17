package controllers

import (
	"github.com/disco07/blog/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func (s Server) findPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.JSON(w, http.StatusOK, vars["id"])
}

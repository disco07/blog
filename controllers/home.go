package controllers

import (
	"github.com/disco07/blog/utils"
	"net/http"
)

func (s Server) home(w http.ResponseWriter, _ *http.Request) {
	utils.JSON(w, http.StatusOK, "Welcome Gophers")
}

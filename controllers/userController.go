package controllers

import (
	"encoding/json"
	"github.com/disco07/blog/models"
	"github.com/disco07/blog/utils"
	"net/http"
)

func (s Server) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	user.Password = password

	if err := s.repo.CreateUser(user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "user created successfully",
	}
	utils.JSON(w, http.StatusCreated, res)
}

func (s Server) login(w http.ResponseWriter, r *http.Request) {

}

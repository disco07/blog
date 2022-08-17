package controllers

import (
	"github.com/disco07/blog/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s Server) getOnePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
	}
	post, err := s.repo.FindPostById(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
	}
	utils.JSON(w, http.StatusOK, post)
}

func (s Server) getAllPosts(w http.ResponseWriter, _ *http.Request) {
	posts, err := s.repo.FindAllPost()
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
	}
	utils.JSON(w, http.StatusOK, posts)
}

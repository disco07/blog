package controllers

import (
	"encoding/json"
	"github.com/disco07/blog/models"
	"github.com/disco07/blog/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (s Server) getOnePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	post, err := s.repo.FindPostById(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, post)
}

func (s Server) getAllPosts(w http.ResponseWriter, _ *http.Request) {
	posts, err := s.repo.FindAllPost()
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, posts)
}

func (s Server) createPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	post.PublishedAt = time.Now()

	if err := s.repo.InsertPost(post); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "successfully added",
	}
	utils.JSON(w, http.StatusCreated, res)
}

func (s Server) updatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	post.ID = id
	if err := s.repo.UpdatePost(post); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "successfully updated",
	}
	utils.JSON(w, http.StatusCreated, res)
}

func (s Server) deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := s.repo.DeletePost(id); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "successfully deleted",
	}
	utils.JSON(w, http.StatusCreated, res)
}

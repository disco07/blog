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

func (s Server) getOneTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	tag, err := s.repo.FindTagById(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, tag)
}

func (s Server) getAllTags(w http.ResponseWriter, _ *http.Request) {
	tags, err := s.repo.FindAllTag()
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, tags)
}

func (s Server) createTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	tag.PublishedAt = time.Now()

	if err := s.repo.InsertTag(tag); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "successfully added",
	}
	utils.JSON(w, http.StatusCreated, res)
}

func (s Server) updateTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	tag.ID = id
	if err := s.repo.UpdateTag(tag); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "successfully updated",
	}
	utils.JSON(w, http.StatusCreated, res)
}

func (s Server) deleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := s.repo.DeleteTag(id); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "successfully deleted",
	}
	utils.JSON(w, http.StatusCreated, res)
}

package controllers

import (
	"encoding/json"
	"gocoop-server/models"
	"log"
	"net/http"
)

func (s *Server) GetCategories(w http.ResponseWriter, req *http.Request) {
	log.Println("> GET request to /categories")

	category := models.Categories{}
	categories, err := category.GetCategories(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*categories)
}

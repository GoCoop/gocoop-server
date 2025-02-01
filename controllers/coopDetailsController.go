package controllers

import (
	"encoding/json"
	"gocoop-server/models"
	"log"
	"net/http"
)

func (s *Server) GetCoopDetails(w http.ResponseWriter, req *http.Request) {
	pv := req.PathValue("name")
	log.Printf("> GET request to /coops/%s\n", pv)

	details := models.CoopDetails{}
	coopDetails, err := details.GetCoopDetails(s.DB, pv)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coopDetails)
}

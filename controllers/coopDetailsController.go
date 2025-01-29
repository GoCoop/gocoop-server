package controllers

import (
	"encoding/json"
	"fmt"
	"gocoop-server/models"
	"log"
	"net/http"
)

func (s *Server) GetCoopDetails(w http.ResponseWriter, req *http.Request) {
	log.Println("> GET requests to /coopDetails")
	pv := req.PathValue("name")
	if pv != "" {
		fmt.Println(pv)
	} else {
		fmt.Println("Empty PathValue name")
	}

	details := models.CoopDetails{}
	coopDetails, err := details.GetCoopDetails(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*coopDetails)
}

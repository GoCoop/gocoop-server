package controllers

import (
	"encoding/json"
	"gocoop-server/middleware"
	"gocoop-server/models"
	"log"
	"net/http"
)

func (s *Server) GetCoopDetails(w http.ResponseWriter, req *http.Request) {
	slug := req.PathValue("slug")
	lang, _ := req.Context().Value(middleware.LangKey).(middleware.Locale)

	params := models.DetailsParams{Slug: slug, LangId: lang.Id}

	log.Printf("> GET request to /coops/%s\n", slug)

	details := models.CoopDetails{}
	coopDetails, err := details.GetCoopDetails(s.DB, params)
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

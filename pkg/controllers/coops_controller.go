package controllers

import (
	"encoding/json"
	"gocoop-server/pkg/middleware"
	"gocoop-server/pkg/models"
	"gocoop-server/pkg/services"
	"log"
	"net/http"
	"strings"
)

func (s *Server) GetCoops(w http.ResponseWriter, req *http.Request) {
	log.Println("> GET request to /coops")

	lang, _ := req.Context().Value(middleware.LangKey).(middleware.Locale)

	params := models.SearchParams{
		Query:    strings.ToLower(req.FormValue("query")),
		Category: req.FormValue("category"),
		LangId:   lang.Id,
	}

	coops, err := services.GetCoops(s.DB, params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coops)
}

package controllers

import (
	"encoding/json"
	"gocoop-server/models"
	"log"
	"net/http"

	"github.com/timewasted/go-accept-headers"
)

func (s *Server) GetCategories(w http.ResponseWriter, req *http.Request) {
	log.Println("> GET request to /categories")
	acceptLang := req.Header.Get("Accept-Language")

	defaultLang := "en-US"

	if acceptLang != "" {
		l := accept.Parse(acceptLang)
		defaultLang = l[0].Type
	}

	langId := returnLangId(defaultLang)

	categories, err := models.GetCategories(s.DB, langId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

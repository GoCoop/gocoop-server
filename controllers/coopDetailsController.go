package controllers

import (
	"encoding/json"
	"gocoop-server/models"
	"log"
	"net/http"

	"github.com/timewasted/go-accept-headers"
)

func (s *Server) GetCoopDetails(w http.ResponseWriter, req *http.Request) {
	slug := req.PathValue("slug")
	acceptLang := req.Header.Get(("Accept-Language"))

	defaultLang := "en-US"

	if acceptLang != "" {
		l := accept.Parse(acceptLang)
		defaultLang = l[0].Type
	}

	langId := returnLangId(defaultLang)
	params := models.Params{Slug: slug, LangId: langId}

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

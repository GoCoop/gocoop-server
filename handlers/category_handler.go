package handlers

import (
	"encoding/json"
	"net/http"
)

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var categories = []Categories{
	{ID: 1, Name: "Alimentos", Icon: "food"},
	{ID: 2, Name: "Café", Icon: "coffee"},
	{ID: 3, Name: "Cerveja", Icon: "beer"},
	{ID: 4, Name: "Crédito", Icon: "banking"},
	{ID: 5, Name: "Indústria", Icon: "industry"},
	{ID: 6, Name: "Serviços", Icon: "services"},
	{ID: 7, Name: "Tecnologia", Icon: "tech"},
}

func GetCategories(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

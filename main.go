package main

import (
	"gocoop-server/config"
	"gocoop-server/controllers"
	"gocoop-server/middleware"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	dbPool := config.ConnectToDatabase()
	defer dbPool.Close()

	s := controllers.Server{DB: dbPool}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /categories", s.GetCategories)
	mux.HandleFunc("GET /coops", s.GetCoops)
	mux.HandleFunc("GET /coops/{slug}", s.GetCoopDetails)

	handler := middleware.HandleAcceptLang(mux)

	log.Println("> Server started! Running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

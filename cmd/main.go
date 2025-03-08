package main

import (
	"gocoop-server/pkg/controllers"
	"gocoop-server/pkg/database"
	"gocoop-server/pkg/middleware"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	dbPool := database.Connect()
	defer dbPool.Close()

	s := controllers.Server{DB: dbPool}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /categories", s.GetCategories)
	mux.HandleFunc("GET /coops", s.GetCoops)
	mux.HandleFunc("GET /coops/{slug}", s.GetCoopDetails)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        middleware.Wrapper(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("> Server started! Running on port 8080.")
	log.Fatal(server.ListenAndServe())
}

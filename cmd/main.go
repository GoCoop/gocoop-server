package main

import (
	"gocoop-server/pkg/controllers"
	"gocoop-server/pkg/database"
	"gocoop-server/pkg/middleware"
	"log"
	"net/http"
	"os"
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

	fs := http.FileServer(http.Dir("../media"))
	mux.Handle("/media/", http.StripPrefix("/media/", fs))

	mux.HandleFunc("GET /categories", s.GetCategories)
	mux.HandleFunc("GET /coops", s.GetCoops)
	mux.HandleFunc("GET /coops/{slug}", s.GetCoopDetails)

	server := &http.Server{
		Addr:           os.Getenv("APP_PORT"),
		Handler:        middleware.Wrapper(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("> Server started! Running on port %s.\n", os.Getenv("APP_PORT"))
	log.Fatal(server.ListenAndServe())
}

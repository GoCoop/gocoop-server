package main

import (
	"context"
	"fmt"
	"gocoop-server/config"
	"gocoop-server/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	log.Println("> Attempting to connect to postgres database...")
	dbPool := config.ConnectToDatabase()
	defer dbPool.Close()

	var greeting string
	err = dbPool.QueryRow(context.Background(), "SELECT 'Hello, World!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /categories", handlers.GetCategories)

	log.Println("> Server started! Running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"fmt"
	"gocoop-server/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /categories", handlers.GetCategories)

	fmt.Println("Server started! Running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

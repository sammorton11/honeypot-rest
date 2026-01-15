package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Attempt struct {
	Address string `json:"address"`
	Network string `json:"network"`
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()
	origins := []string{"*"}

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}))

	r.Post("/attempt", AttemptHandler)

	log.Printf("Server running on %s\n", ":8080")
	http.ListenAndServe(":8080", r)
}

func AttemptHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Parse attempt into struct
	var attempt Attempt
	err := json.NewDecoder(r.Body).Decode(&attempt)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
 	err = json.NewEncoder(w).Encode(&attempt)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	log.Println(attempt)
}


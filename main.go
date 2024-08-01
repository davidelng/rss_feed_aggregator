package main

import (
	"database/sql"
	"github.com/davidelng/rssfeedaggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbURL := os.Getenv("CONN")
	if dbURL == "" {
		log.Fatal("Missing connection string")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not connect to database: %s", err)
	}

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUserCreate)

	log.Printf("Starting server on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

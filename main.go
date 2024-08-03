package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davidelng/rssfeedaggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	feedsToBeFetched := 10
	go startScraping(db, feedsToBeFetched, time.Minute)

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})

	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUserGetByAPIKey))
	mux.HandleFunc("POST /v1/users", apiCfg.handlerUserCreate)

	mux.HandleFunc("GET /v1/feeds", apiCfg.handlerFeedsGetAll)
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))

	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowGetByUser))
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

	mux.HandleFunc("GET /v1/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	log.Printf("Starting server on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/GavinDevelops/blog-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("POSTGRESQL_URL")
	db, dbErr := sql.Open("postgres", dbURL)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	dbQueries := database.New(db)
	config := apiConfig{DB: dbQueries}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", readinessCheck)
	mux.HandleFunc("GET /v1/err", errorCheck)
	mux.HandleFunc("POST /v1/users", config.createUser)
	mux.HandleFunc("GET /v1/users", config.middlewareAuth(config.getUser))
	mux.HandleFunc("POST /v1/feeds", config.middlewareAuth(config.createFeed))
	mux.HandleFunc("GET /v1/feeds", config.getFeeds)
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	log.Printf("Starting server on port: %v", port)
	server.ListenAndServe()
}

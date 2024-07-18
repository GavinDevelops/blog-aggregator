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

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

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
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	log.Printf("Starting server on port: %v", port)
	server.ListenAndServe()
}

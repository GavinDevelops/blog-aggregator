package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func readinessCheck(w http.ResponseWriter, _ *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, resp{Status: "ok"})
}

func errorCheck(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", readinessCheck)
	mux.HandleFunc("GET /v1/err", errorCheck)
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	log.Printf("Starting server on port: %v", port)
	server.ListenAndServe()
}

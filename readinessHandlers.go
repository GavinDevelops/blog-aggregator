package main

import "net/http"

func readinessCheck(w http.ResponseWriter, _ *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, resp{Status: "ok"})
}

func errorCheck(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

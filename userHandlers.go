package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/GavinDevelops/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		apiKey := strings.TrimPrefix(authHeader, "ApiKey ")
		user, getErr := config.DB.GetUserByApiKey(r.Context(), apiKey)
		if getErr != nil {
			respondWithError(w, http.StatusNotFound, getErr.Error())
			return
		}
		handler(w, r, user)
	}
}

func (config *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, http.StatusBadRequest, decodeErr.Error())
		return
	}
	user, createErr := config.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      params.Name,
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
	if createErr != nil {
		respondWithError(w, http.StatusInternalServerError, createErr.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

func (config *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, user)
}

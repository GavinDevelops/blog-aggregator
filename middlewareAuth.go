package main

import (
	"net/http"
	"strings"

	"github.com/GavinDevelops/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

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

package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GavinDevelops/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, http.StatusBadRequest, decodeErr.Error())
		return
	}
	feedUUID, parseErr := uuid.Parse(params.FeedID)
	if parseErr != nil {
		respondWithError(w, http.StatusBadRequest, parseErr.Error())
		return
	}
	feedFollow, createErr := c.DB.CreateFeedFollows(
		r.Context(),
		database.CreateFeedFollowsParams{
			ID:        uuid.New(),
			FeedID:    feedUUID,
			UserID:    user.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
	if createErr != nil {
		respondWithError(w, http.StatusInternalServerError, createErr.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (c *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedUUID, parseErr := uuid.Parse(r.PathValue("feedFollowID"))
	if parseErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Feed Follow ID")
		return
	}
	feedFollow, deleteErr := c.DB.DeleteFeedFollow(
		r.Context(),
		database.DeleteFeedFollowParams{
			ID:     feedUUID,
			UserID: user.ID,
		})
	if deleteErr != nil {
		respondWithError(w, http.StatusInternalServerError, deleteErr.Error())
		return
	}
	respondWithJson(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (c *apiConfig) getFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	userFeedFollows, getErr := c.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if getErr != nil {
		respondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	respondWithJson(w, http.StatusOK, databaseFeedFollowsToFeedFollows(userFeedFollows))
}

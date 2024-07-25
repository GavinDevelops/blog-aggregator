package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GavinDevelops/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		respondWithError(w, http.StatusBadRequest, decodeErr.Error())
		return
	}
	feed, createErr := config.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if createErr != nil {
		respondWithError(w, http.StatusInternalServerError, createErr.Error())
		return
	}
	feedFollow, createFollowErr := config.DB.CreateFeedFollows(
		r.Context(),
		database.CreateFeedFollowsParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			FeedID:    feed.ID,
			UserID:    user.ID,
		})
	if createFollowErr != nil {
		respondWithError(w, http.StatusInternalServerError, createFollowErr.Error())
		return
	}
	type responseStruct struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}
	respondWithJson(w, http.StatusCreated, responseStruct{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFollow(feedFollow),
	})
}

func (config *apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, getErr := config.DB.GetFeeds(r.Context())
	if getErr != nil {
		respondWithError(w, http.StatusInternalServerError, getErr.Error())
		return
	}
	respondWithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/deepakraj-p28/RSS_Aggregator/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) handlerCreateFeedFollows(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedfollow, err := apiCnfg.DB.CreateFeedFollows(req.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("Couldn't create feed follow: %s", err))
		return
	}

	respondWithJson(respWriter, 200, models.DatabaseFeedFollowToFeedFollow(feedfollow))
}

func (apiCnfg *apiConfig) handlerGetFeedFollowsForUser(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	feedFollows, err := apiCnfg.DB.GetFeedFollowsForUser(req.Context(), user.ID)
	if err != nil {
		respondWithError(respWriter, 400, "Error fetching feed follows for user")
	}

	respondWithJson(respWriter, 200, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCnfg *apiConfig) handlerGetUsersForFeed(respWriter http.ResponseWriter, req *http.Request) {
	feedID, err := uuid.Parse(chi.URLParam(req, "feedFollowID"))
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("couldn't parse feed follow ID: %v", err))
	}

	feedUsers, err := apiCnfg.DB.GetUsersForFeed(req.Context(), feedID)
	if err != nil {
		respondWithError(respWriter, 400, "Error fetching feed follows for user")
	}

	respondWithJson(respWriter, 200, models.DatabaseFeedFollowsToFeedFollows(feedUsers))
}

func (apiCnfg *apiConfig) handlerDeleteFeedFollow(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	// type parameters struct {
	// 	ID     uuid.UUID `json:"feed_id"`
	// 	UserID uuid.UUID `json:"user_id"`
	// }
	feedID, err := uuid.Parse(chi.URLParam(req, "feedFollowID"))

	err = apiCnfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		ID:     feedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(respWriter, 400, "Error Deleting feed follow for user")
	}

	respondWithJson(respWriter, 200, struct{}{})
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/deepakraj-p28/RSS_Aggregator/internal/models"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) handlerCreateFeed(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, err := apiCnfg.DB.CreateFeed(req.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}

	respondWithJson(respWriter, 200, models.DatabaseFeedToFeed(feed))
}

func (apiCnfg *apiConfig) handlerGetFeed(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	feed, err := apiCnfg.DB.GetFeed(req.Context(), user.ID)
	if err != nil {
		respondWithError(respWriter, 400, "Error fetching feed")
		return
	}

	respondWithJson(respWriter, 200, models.DatabaseFeedToFeed(feed))
}

func (apiCnfg *apiConfig) handlerGetFeeds(respWriter http.ResponseWriter, req *http.Request) {
	feeds, err := apiCnfg.DB.GetFeeds(req.Context())
	if err != nil {
		respondWithError(respWriter, 400, "Error fetching feed")
		return
	}

	respondWithJson(respWriter, 200, models.DatabaseFeedsToFeeds(feeds))
}

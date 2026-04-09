package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/deepakraj-p28/RSS_Aggregator/internal/models"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) handlerCreateUser(respWriter http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCnfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(respWriter, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}

	respondWithJson(respWriter, 201, models.DatabaseUserToUser(user))
}

func (apiCnfg *apiConfig) handlerGetUser(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	//Fetching authenticated user is being done by middleware_auth
	respondWithJson(respWriter, 200, models.DatabaseUserToUser(user))
}

func (apiCnfg *apiConfig) handlerGetPosts(respWriter http.ResponseWriter, req *http.Request, user database.User) {
	posts, err := apiCnfg.DB.GetPosts(req.Context(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  5,
	})
	if err != nil {
		log.Println("Couldn't fetch posts for user:", err)
	}
	respondWithJson(respWriter, 200, posts)
}

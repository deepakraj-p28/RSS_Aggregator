package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) HandlerCreateUser(respWriter http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(respWriter, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCnfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(respWriter, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}

	RespondWithJson(respWriter, 200, model.DatabaseUserToUser(user))
}

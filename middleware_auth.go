package main

import (
	"net/http"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/auth"
	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
)

type AuthedUser func(http.ResponseWriter, *http.Request, database.User)

func (apiCnfg *apiConfig) middleware_Auth(handler AuthedUser) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		apiKey, err := auth.GetAPIKey(req.Header)
		if err != nil {
			respondWithError(respWriter, 403, "could not fetch api key")
			return
		}

		user, err := apiCnfg.DB.GetUserByApiKey(req.Context(), apiKey)
		if err != nil {
			respondWithError(respWriter, 400, "auth error: could not get user")
			return
		}

		handler(respWriter, req, user)
	}
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(respWriter http.ResponseWriter, respCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		respWriter.WriteHeader(500)
		return
	}
	respWriter.Header().Add("Content-Type", "application/json")
	respWriter.WriteHeader(respCode)
	respWriter.Write(data)
}

func RespondWithError(respWriter http.ResponseWriter, respCode int, errMsg string) {
	if respCode > 499 {
		log.Printf("Responding with error: %s\n", errMsg)
	}

	//marshal the field to a json object so we can pass the struct as payload to respondWithJson
	//key will be error: msg will be Error field
	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(respWriter, respCode, errorResponse{
		Error: errMsg,
	})
}

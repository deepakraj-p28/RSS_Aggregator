package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(respWriter http.ResponseWriter, respCode int, payload interface{}) {
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

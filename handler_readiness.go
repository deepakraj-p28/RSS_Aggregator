package main

import "net/http"

//Define Handler in a way http go standard lib expects

func HandlerReadiness(respWriter http.ResponseWriter, request *http.Request) {
	respondWithJson(respWriter, 200, struct{}{})
}

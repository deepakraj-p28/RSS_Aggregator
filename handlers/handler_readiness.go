package handlers

import "net/http"

//Define Handler in a way http go standard lib expects

func HandlerReadiness(respWriter http.ResponseWriter, request *http.Request) {
	RespondWithJson(respWriter, 200, struct{}{})
}

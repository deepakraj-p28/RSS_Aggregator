package main

import (
	"net/http"
)

//Define Handler in a way http go standard lib expects

func HandlerError(respWriter http.ResponseWriter, request *http.Request) {
	//(respWriter, 400, "Something went wrong")
	respondWithError(respWriter, 400, "Something went wrong")

}

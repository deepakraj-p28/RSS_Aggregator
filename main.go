package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Overload()
	port := os.Getenv("PORT")

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

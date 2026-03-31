package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deepakraj-p28/RSS_Aggregator/src/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Overload()
	port := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthcheck", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerError)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Starting server at port:%s\n", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		log.Fatalf(err.Error(), err)
	}
}

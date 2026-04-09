package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// type apiConfig struct {
// 	DB *database.Queries
// }

func main() {
	godotenv.Overload()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in env")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apicnf := DBConnection()

	startScraping(apicnf.DB, 10, 2*time.Minute)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthcheck", HandlerReadiness)
	v1Router.Get("/err", HandlerError)
	v1Router.Post("/user", apicnf.handlerCreateUser)
	v1Router.Get("/user", apicnf.middleware_Auth(apicnf.handlerGetUser))

	v1Router.Post("/feed", apicnf.middleware_Auth(apicnf.handlerCreateFeed))
	v1Router.Get("/feed", apicnf.middleware_Auth(apicnf.handlerGetFeed))
	v1Router.Get("/feed", apicnf.handlerGetFeeds)

	v1Router.Post("/feedfollows", apicnf.middleware_Auth(apicnf.handlerCreateFeedFollows))
	v1Router.Get("/feedfollows", apicnf.middleware_Auth(apicnf.handlerGetFeedFollowsForUser))
	v1Router.Get("/feedfollows/{feedFollowID}", apicnf.handlerGetUsersForFeed)
	v1Router.Delete("/feedfollows/{feedFollowID}", apicnf.middleware_Auth(apicnf.handlerDeleteFeedFollow))

	v1Router.Get("/posts", apicnf.middleware_Auth(apicnf.handlerGetPosts))

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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	// dbUrl := os.Getenv("DB_URL")
	// if dbUrl == "" {
	// 	log.Fatal("DB_URL not found in env")
	// }

	// conn, err := sql.Open("postgres", dbUrl)
	// if err != nil {
	// 	log.Fatal("Cannot connect to DB")
	// }

	// apiCnfg := apiConfig{
	// 	DB: database.New(conn),
	// }

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

	v1Router := chi.NewRouter()
	v1Router.Get("/healthcheck", HandlerReadiness)
	v1Router.Get("/err", HandlerError)
	v1Router.Post("/user", apicnf.handlerCreateUser)
	v1Router.Get("/user", apicnf.middleware_Auth(apicnf.handlerGetUser))

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

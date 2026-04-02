package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

var apiCnfg apiConfig

func DBConnection() *apiConfig {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL not found in env")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	apiCnfg = apiConfig{
		DB: database.New(conn),
	}
	fmt.Printf("Returning pointer to api Config: %v", &apiCnfg)
	return &apiCnfg

}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/mhv2408/my-blog/internal/database"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	fmt.Println("My Blog!!")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	// Open DB Connection
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
		return
	}
	apiCfg := &apiConfig{
		db: database.New(dbConn),
	}

	http.HandleFunc("/", corsMiddleware(homePageHandler))
	http.HandleFunc("POST /login", corsMiddleware(loginHandler))
	http.HandleFunc("GET /editor/post", corsMiddleware(loginMiddleware(writePostHandler)))
	http.HandleFunc("POST editor/post", corsMiddleware(loginMiddleware(apiCfg.posts)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

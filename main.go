package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mhv2408/my-blog/internal/database"
)

type apiConfig struct {
	db *database.Queries
}

type Post struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Post        string    `json:"post"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
	PublishedAt string    `json:"published_at"`
}

func main() {
	fmt.Println("My Blog!!")
	godotenv.Load(".env")

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

	http.HandleFunc("GET /get-blogs", corsMiddleware(homePageHandler))
	http.HandleFunc("POST /login", corsMiddleware(loginHandler))
	//http.HandleFunc("GET editor/")
	http.HandleFunc("GET /editor/post", corsMiddleware(loginMiddleware(writePostHandler)))
	http.HandleFunc("POST /editor/post", corsMiddleware(loginMiddleware(apiCfg.posts)))
	http.HandleFunc("GET /editor/post/{postId}", corsMiddleware(loginMiddleware(apiCfg.postByIdHandler)))

	http.HandleFunc("GET /editor/dashboard", corsMiddleware(loginMiddleware(apiCfg.dashboardHandler)))

	http.HandleFunc("DELETE /editor/post/{postId}", corsMiddleware(loginMiddleware(apiCfg.deletePostHandler)))
	http.HandleFunc("PATCH /editor/post/{postId}/status", corsMiddleware(loginMiddleware(apiCfg.updateStatusHandler)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

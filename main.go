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
	"github.com/mhv2408/my-blog/middleware"
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
	Slug        string    `json:"slug"`
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
	mux := http.NewServeMux()

	mux.HandleFunc("GET /get-blogs", apiCfg.postsHandler)
	mux.HandleFunc("GET /get-post/{slug}", apiCfg.postBySlugHandler)

	mux.HandleFunc("POST /login", loginHandler)
	//http.HandleFunc("GET editor/")
	mux.HandleFunc("GET /editor/post", middleware.LoginMiddleware(writePostHandler))
	mux.HandleFunc("POST /editor/post", middleware.LoginMiddleware(apiCfg.posts))
	mux.HandleFunc("GET /editor/post/{postId}", middleware.LoginMiddleware(apiCfg.postByIdHandler))

	mux.HandleFunc("GET /editor/dashboard", middleware.LoginMiddleware(apiCfg.dashboardHandler))

	mux.HandleFunc("DELETE /editor/post/{postId}", middleware.LoginMiddleware(apiCfg.deletePostHandler))
	mux.HandleFunc("PATCH /editor/post/{postId}/status", middleware.LoginMiddleware(apiCfg.updateStatusHandler))

	log.Fatal(http.ListenAndServe(":8080", middleware.CorsMiddleware(mux)))
}

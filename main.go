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

type Blog struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
	PublishedAt time.Time `json:"published_at"`
	Slug        string    `json:"slug"`
}

func main() {
	fmt.Println("My Blog!!")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("unable to load the DB URL")
	}

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

	mux.HandleFunc("GET /get-blogs", apiCfg.blogsHandler)
	mux.HandleFunc("GET /get-blogs/{slug}", apiCfg.blogByIdHandler)

	mux.HandleFunc("POST /login", loginHandler)
	//http.HandleFunc("GET editor/")
	mux.HandleFunc("GET /editor/blog", middleware.LoginMiddleware(writeBlogHandler))
	mux.HandleFunc("POST /editor/blog", middleware.LoginMiddleware(apiCfg.blogs))
	mux.HandleFunc("GET /editor/post/{postId}", middleware.LoginMiddleware(apiCfg.blogByIdHandler))
	mux.HandleFunc("PUT /editor/post/{postId}", middleware.LoginMiddleware(apiCfg.updateBlogHandler))

	mux.HandleFunc("GET /editor/dashboard", middleware.LoginMiddleware(apiCfg.dashboardHandler))

	mux.HandleFunc("DELETE /editor/post/{postId}", middleware.LoginMiddleware(apiCfg.deleteBlogHandler))
	mux.HandleFunc("PATCH /editor/post/{postId}/status", middleware.LoginMiddleware(apiCfg.updateStatusHandler))

	srv := &http.Server{
		Addr:              "localhost:8080",
		Handler:           middleware.CorsMiddleware(mux),
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

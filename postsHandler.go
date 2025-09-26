package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

type BlogPost struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	PublishedAt time.Time `json:"publisheded_at"`
	Slug        string    `json:"slug"`
}

func getPosts(dbPosts []database.GetBlogsRow) []BlogPost {
	res := make([]BlogPost, 0, len(dbPosts))
	for _, post := range dbPosts {
		res = append(res, BlogPost{
			Id:          post.BlogID,
			Title:       post.Title,
			Summary:     post.Summary,
			PublishedAt: post.PublishedAt.Time,
			Slug:        post.Slug,
		})
	}
	return res
}

func (cfg *apiConfig) blogsHandler(w http.ResponseWriter, r *http.Request) {
	db_blogs, err := cfg.db.GetBlogs(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot get the blogs", err)
		return
	}

	respondWithJson(w, http.StatusOK, getPosts(db_blogs))

}

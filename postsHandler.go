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
	PublishedAt time.Time `json:"published_at"`
	Slug        string    `json:"slug"`
}

func getBlogs(dbBlogs []database.GetBlogsRow) []BlogPost {
	res := make([]BlogPost, 0, len(dbBlogs))
	for _, blog := range dbBlogs {
		res = append(res, BlogPost{
			Id:          blog.BlogID,
			Title:       blog.Title,
			Summary:     blog.Summary,
			PublishedAt: blog.PublishedAt.Time,
			Slug:        blog.Slug,
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

	respondWithJson(w, http.StatusOK, getBlogs(db_blogs))

}

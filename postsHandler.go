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

func getPosts(dbPosts []database.GetPostsRow) []BlogPost {
	res := make([]BlogPost, 0, len(dbPosts))
	for _, post := range dbPosts {
		res = append(res, BlogPost{
			Id:          post.PostsID,
			Title:       post.Title,
			Summary:     post.Summary,
			PublishedAt: post.PublishedAt.Time,
			Slug:        post.Slug,
		})
	}
	return res
}

func (cfg *apiConfig) postsHandler(w http.ResponseWriter, r *http.Request) {
	db_posts, err := cfg.db.GetPosts(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot get the posts", err)
		return
	}

	respondWithJson(w, http.StatusOK, getPosts(db_posts))

}

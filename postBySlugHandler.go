package main

import (
	"net/http"
)

func (cfg *apiConfig) postBySlugHandler(w http.ResponseWriter, r *http.Request) {
	post_slug := r.PathValue("slug")

	dbPost, err := cfg.db.GetPostBySlug(r.Context(), post_slug)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot retrieve post from slug", err)
		return
	}
	respondWithJson(w, http.StatusOK, Post{
		Id:          dbPost.PostsID,
		Title:       dbPost.Title,
		Post:        dbPost.Post,
		PublishedAt: dbPost.PublishedAt.Time.Format("YYYY-MM-DD"),
		Slug:        dbPost.Slug,
	})
}

package main

import (
	"net/http"
)

func (cfg *apiConfig) postBySlugHandler(w http.ResponseWriter, r *http.Request) {
	post_slug := r.PathValue("slug")

	dbPost, err := cfg.db.GetBlogBySlug(r.Context(), post_slug)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot retrieve post from slug", err)
		return
	}
	respondWithJson(w, http.StatusOK, Blog{
		Id:          dbPost.BlogID,
		Title:       dbPost.Title,
		Post:        dbPost.Content,
		PublishedAt: dbPost.PublishedAt.Time,
		Slug:        dbPost.Slug,
	})
}

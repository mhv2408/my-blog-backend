package main

import (
	"net/http"
)

func (cfg *apiConfig) blogBySlugHandler(w http.ResponseWriter, r *http.Request) {
	blog_slug := r.PathValue("slug")

	dbBlog, err := cfg.db.GetBlogBySlug(r.Context(), blog_slug)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot retrieve blog from slug", err)
		return
	}
	respondWithJson(w, http.StatusOK, Blog{
		Id:          dbBlog.BlogID,
		Title:       dbBlog.Title,
		Content:     dbBlog.Content,
		PublishedAt: dbBlog.PublishedAt.Time,
		Slug:        dbBlog.Slug,
	})
}

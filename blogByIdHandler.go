package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) blogByIdHandler(w http.ResponseWriter, r *http.Request) {
	BlogId, err := uuid.Parse(r.PathValue("blogId"))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot parse the blog id", err)
		return
	}

	dbBlog, err := cfg.db.GetBlogByID(r.Context(), BlogId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot find the blog with the id", err)
		return
	}
	respondWithJson(w, http.StatusOK, Blog{
		Id:      dbBlog.BlogID,
		Title:   dbBlog.Title,
		Summary: dbBlog.Summary,
		Content: dbBlog.Content,
	})
}

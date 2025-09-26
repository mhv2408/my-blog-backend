package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) postByIdHandler(w http.ResponseWriter, r *http.Request) {
	PostId, err := uuid.Parse(r.PathValue("postId"))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot parse the post id", err)
		return
	}

	dbPost, err := cfg.db.GetBlogByID(r.Context(), PostId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot find the post with the id", err)
		return
	}
	respondWithJson(w, http.StatusOK, Blog{
		Id:      dbPost.BlogID,
		Title:   dbPost.Title,
		Summary: dbPost.Summary,
		Post:    dbPost.Content,
	})
}

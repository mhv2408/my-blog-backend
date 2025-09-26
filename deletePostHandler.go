package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := uuid.Parse(r.PathValue("postId"))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to parse the postId", err)
		return
	}
	err = cfg.db.DeleteBlogById(r.Context(), postId)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to delete the post", err)
		return
	}

	respondWithJson(w, http.StatusOK, nil)

}

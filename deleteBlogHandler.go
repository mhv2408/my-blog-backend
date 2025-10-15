package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	blogId, err := uuid.Parse(r.PathValue("blogId"))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to parse the blogId", err)
		return
	}
	err = cfg.db.DeleteBlogById(r.Context(), blogId)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to delete the blog", err)
		return
	}

	respondWithJson(w, http.StatusOK, nil)

}

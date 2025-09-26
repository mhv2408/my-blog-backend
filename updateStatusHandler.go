package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

func (cfg *apiConfig) updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	blogId, err := uuid.Parse(r.PathValue("blogId"))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot parse blog_id from url", err)
		return
	}

	type payload struct {
		Status string `json:"status"`
	}
	decoder := json.NewDecoder(r.Body)
	var data payload
	err = decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to parse the status from json", err)
		return
	}

	err = cfg.db.UpdateBlogStatus(r.Context(), database.UpdateBlogStatusParams{
		BlogID: blogId,
		Status: data.Status,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to update the status in DB", err)
		return
	}
	respondWithJson(w, http.StatusOK, nil)
}

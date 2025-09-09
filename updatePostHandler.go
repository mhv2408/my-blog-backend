package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

func (cfg *apiConfig) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Id      uuid.UUID `json:"id"`
		Title   string    `json:"title"`
		Summary string    `json:"summary"`
		Post    string    `json:"post"`
		Status  string    `json:"status"`
	}
	var data payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot parse the json payload", err)
		return
	}

	cfg.db.UpdatePost(r.Context(), database.UpdatePostParams{
		PostsID: data.Id,
		Title:   data.Title,
		Summary: data.Summary,
		Post:    data.Post,
		Status:  data.Status,
	})
	respondWithJson(w, http.StatusOK, nil)
}

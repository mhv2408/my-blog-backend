package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

func (cfg *apiConfig) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post_id, err := uuid.Parse(r.PathValue("postId"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot find the post_id from path value URL", err)
		return
	}

	type payload struct {
		Title   string `json:"title"`
		Summary string `json:"summary"`
		Post    string `json:"post"`
		Status  string `json:"status"`
	}
	var data payload
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot parse the json payload", err)
		return
	}

	err = cfg.db.UpdateBlog(r.Context(), database.UpdateBlogParams{
		BlogID:  post_id,
		Title:   data.Title,
		Summary: data.Summary,
		Content: data.Post,
		Status:  data.Status,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to update the post", err)
		return
	}
	respondWithJson(w, http.StatusOK, nil)
}

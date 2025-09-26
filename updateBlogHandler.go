package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

func (cfg *apiConfig) updateBlogHandler(w http.ResponseWriter, r *http.Request) {
	blog_id, err := uuid.Parse(r.PathValue("blogId"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot find the blog from path value URL", err)
		return
	}

	type payload struct {
		Title   string `json:"title"`
		Summary string `json:"summary"`
		Content string `json:"content"`
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
		BlogID:  blog_id,
		Title:   data.Title,
		Summary: data.Summary,
		Content: data.Content,
		Status:  data.Status,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to update the post", err)
		return
	}
	respondWithJson(w, http.StatusOK, nil)
}

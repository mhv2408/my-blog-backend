package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

type SummaryPost struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Summary string    `json:"summary"`
	Status  string    `json:"status"`
}

func GetApiPosts(dbPosts []database.GetPostsDashboardRow) []SummaryPost {
	res := make([]SummaryPost, 0, len(dbPosts))

	for _, post := range dbPosts {
		res = append(res, SummaryPost{
			Id:      post.PostsID,
			Title:   post.Title,
			Summary: post.Summary,
			Status:  post.Status,
		})
	}
	return res
}

func (cfg *apiConfig) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	dbPosts, err := cfg.db.GetPostsDashboard(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to retrive the posts", err)
		return
	}
	apiPosts := GetApiPosts(dbPosts)

	respondWithJson(w, http.StatusOK, apiPosts)

}

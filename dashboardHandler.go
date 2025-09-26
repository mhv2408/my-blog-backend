package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mhv2408/my-blog/internal/database"
)

type SummaryBlog struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Summary string    `json:"summary"`
	Status  string    `json:"status"`
}

func GetApiBlogs(dbBlogs []database.GetBlogsDashboardRow) []SummaryBlog {
	res := make([]SummaryBlog, 0, len(dbBlogs))

	for _, blog := range dbBlogs {
		res = append(res, SummaryBlog{
			Id:      blog.BlogID,
			Title:   blog.Title,
			Summary: blog.Summary,
			Status:  blog.Status,
		})
	}
	return res
}

func (cfg *apiConfig) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	dbBlogs, err := cfg.db.GetBlogsDashboard(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to retrive the blogs", err)
		return
	}
	apiBlogs := GetApiBlogs(dbBlogs)

	respondWithJson(w, http.StatusOK, apiBlogs)

}

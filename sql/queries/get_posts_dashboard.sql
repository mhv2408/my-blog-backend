-- name: GetPostsDashboard :many

SELECT posts_id, title, summary, status
FROM posts;
-- name: GetPosts :many

SELECT posts_id, title, summary, status
FROM posts;
-- name: GetPostBySlug :one
SELECT *
FROM posts
WHERE slug=$1;
-- name: GetBlogBySlug :one
SELECT *
FROM blogs
WHERE slug=$1;
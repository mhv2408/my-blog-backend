-- name: GetPosts :many

SELECT posts_id, title, summary, published_at, slug
FROM posts
WHERE status='publish';
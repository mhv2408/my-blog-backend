-- name: GetPosts :many

SELECT posts_id, title, summary, published_at
FROM posts
WHERE status='publish';
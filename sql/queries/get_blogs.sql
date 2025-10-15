-- name: GetBlogs :many

SELECT blog_id, title, summary, published_at, slug
FROM blogs
WHERE status='published';
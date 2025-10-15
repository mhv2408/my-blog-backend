-- name: UpdateBlogStatus :exec
UPDATE blogs
SET status = $2, published_at = CASE WHEN $2='published' THEN NOW() ELSE NULL END
WHERE blog_id = $1;
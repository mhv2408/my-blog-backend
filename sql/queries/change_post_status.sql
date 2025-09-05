-- name: UpdatePostStatus :exec
UPDATE posts
SET status = $2, published_at = CASE WHEN $2='publish' THEN NOW() ELSE NULL END
WHERE posts_id = $1;
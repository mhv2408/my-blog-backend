-- name: DeletePostById :exec

DELETE FROM posts
WHERE posts_id = $1;
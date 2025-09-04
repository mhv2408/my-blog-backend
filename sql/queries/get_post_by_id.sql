-- name: GetPostByID :one

SELECT * 
FROM posts
WHERE posts_id=$1;
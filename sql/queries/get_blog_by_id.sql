-- name: GetBlogByID :one

SELECT * 
FROM blogs
WHERE blog_id=$1;
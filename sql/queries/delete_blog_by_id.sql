-- name: DeleteBlogById :exec

DELETE FROM blogs
WHERE blog_id = $1;
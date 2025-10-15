-- name: GetBlogsDashboard :many

SELECT blog_id, title, summary, status
FROM blogs;
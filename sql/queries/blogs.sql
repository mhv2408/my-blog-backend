-- name: CreateBlog :one

INSERT INTO blogs(title, summary, content,status, published_at,slug)
VALUES(
    $1,
    $2,
    $3,
    $4,
    CASE WHEN $4='published' THEN NOW() ELSE NULL END,
    $5
)
RETURNING *;
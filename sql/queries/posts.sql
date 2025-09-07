-- name: CreatePost :one

INSERT INTO posts(title, summary, post,status, published_at,slug)
VALUES(
    $1,
    $2,
    $3,
    $4,
    CASE WHEN $4='publish' THEN NOW() ELSE NULL END,
    $5
)
RETURNING *;
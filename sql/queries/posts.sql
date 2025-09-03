-- name: CreatePost :one

INSERT INTO posts(title, summary, post,status, published_at)
VALUES(
    $1,
    $2,
    $3,
    $4,
    CASE WHEN status='published' THEN NOW() ELSE NULL END
)
RETURNING *;
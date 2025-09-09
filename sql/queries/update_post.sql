-- name: UpdatePost :exec
UPDATE posts
SET 
title=$2, 
summary=$3, 
post=$4, 
status=$5,
published_at = CASE WHEN $5='publish' AND published_at IS NULL THEN NOW() ELSE published_at END
WHERE posts_id=$1;

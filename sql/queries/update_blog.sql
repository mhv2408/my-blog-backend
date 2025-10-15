-- name: UpdateBlog :exec
UPDATE blogs
SET 
title=$2, 
summary=$3, 
content=$4, 
status=$5,
published_at = CASE WHEN $5='published' AND published_at IS NULL THEN NOW() ELSE published_at END
WHERE blog_id=$1;

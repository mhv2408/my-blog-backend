-- +goose UP

ALTER TABLE posts
ADD COLUMN slug TEXT NOT NULL;

-- +goose DOWN
ALTER TABLE posts
DROP COLUMN slug;
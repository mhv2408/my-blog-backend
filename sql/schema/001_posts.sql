-- +goose UP
CREATE TABLE posts (
    posts_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title        TEXT NOT NULL,
    summary      TEXT NOT NULL,
    post         TEXT NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW(),
    status       TEXT NOT NULL CHECK (status IN ('draft', 'published')),
    published_at TIMESTAMP NULL
);

-- +goose Down

DROP TABLE posts;

-- +goose Up
CREATE TABLE posts (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title  TEXT NOT NULL,
    description TEXT,
    language TEXT,
    publishedAt TIMESTAMP NOT NULL,
    link TEXT UNIQUE NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
-- +migrate Up
CREATE TABLE posts
(
    id         uuid PRIMARY KEY     DEFAULT uuid_generate_v4(),
    user_id    uuid        NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    image_url  TEXT        NOT NULL,
    text       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE posts;

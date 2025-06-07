-- +migrate Up
CREATE TABLE followers
(
    id           uuid PRIMARY KEY     DEFAULT uuid_generate_v4(),
    follower_id  uuid        NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    following_id uuid        NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (follower_id, following_id)
);

-- +migrate Down

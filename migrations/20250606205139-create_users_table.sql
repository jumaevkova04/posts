-- +migrate Up
CREATE TABLE users
(
    id         uuid PRIMARY KEY      DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    surname    VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE users;

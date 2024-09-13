
-- +migrate Up
CREATE TABLE users (
    id BYTEA PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE users;


-- +migrate Up
CREATE TABLE users (
    id BYTEA PRIMARY KEY UNIQUE,
    name VARCHAR(255) NOT NULL,
    password BYTEA NOT NULL,
);

CREATE TABLE user_profiles (
    id BYTEA PRIMARY KEY UNIQUE,
    userId BYTEA UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    description VARCHAR(511) NOT NULL,
    CONSTRAINT fk_users FOREGIN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE users;

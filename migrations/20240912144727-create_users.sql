
-- +migrate Up
CREATE TABLE users (
    id BYTEA PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password BYTEA NOT NULL
);

CREATE TABLE user_profiles (
    id BYTEA PRIMARY KEY,
    user_id BYTEA UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    description VARCHAR(511) NOT NULL,
    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE user_profiles;
DROP TABLE users;
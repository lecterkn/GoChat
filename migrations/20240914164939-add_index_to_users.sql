
-- +migrate Up
ALTER TABLE users ADD UNIQUE(id);

-- +migrate Down

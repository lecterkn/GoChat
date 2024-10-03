
-- +migrate Up
ALTER TABLE channels ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT FALSE;

-- +migrate Down
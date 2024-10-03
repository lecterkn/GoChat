
-- +migrate Up
CREATE TABLE messages(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    channel_id UUID NOT NULL,
    message VARCHAR(511) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_m_users FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_m_channels FOREIGN KEY (channel_id) REFERENCES channels (id) ON DELETE CASCADE
);
CREATE INDEX idx_channel ON messages (channel_id);

-- +migrate Down
DROP TABLE messages;

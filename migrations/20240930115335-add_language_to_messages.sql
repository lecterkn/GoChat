
-- +migrate Up
ALTER TABLE users ADD COLUMN language SMALLINT NOT NULL;

CREATE TABLE channel_languages(
    channel_id UUID NOT NULL,
    language SMALLINT NOT NULL,
    UNIQUE (channel_id, language),
    CONSTRAINT fk_cl_channels FOREIGN KEY (channel_id) REFERENCES channels (id) ON DELETE CASCADE
);

CREATE TABLE message_english_contents(
    channel_id UUID NOT NULL,
    message_id UUID NOT NULL,
    content VARCHAR(511) NOT NULL
);
CREATE UNIQUE INDEX idx_mec_ci_mi ON message_english_contents (channel_id, message_id);

CREATE TABLE message_japanese_contents(
    channel_id UUID NOT NULL,
    message_id UUID NOT NULL,
    content VARCHAR(511) NOT NULL
);
CREATE UNIQUE INDEX idx_mjc_ci_mi ON message_japanese_contents (channel_id, message_id);

CREATE TABLE message_chinese_contents(
    channel_id UUID NOT NULL,
    message_id UUID NOT NULL,
    content VARCHAR(511) NOT NULL
);
CREATE UNIQUE INDEX idx_mcc_ci_mi ON message_chinese_contents (channel_id, message_id);

-- +migrate Down
DROP TABLE message_chinese_contents;
DROP TABLE message_english_contents;
DROP TABLE message_japanese_contents;
DROP TABLE channel_languages;
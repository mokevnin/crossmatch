-- +goose Up
CREATE TABLE channels (
    id BIGINT PRIMARY KEY,
    owner_id BIGINT NOT NULL,
    username TEXT NOT NULL,
    topic TEXT NOT NULL,
    language TEXT NOT NULL,
    subscribers INTEGER NOT NULL CHECK (subscribers >= 0),
    avg_reach INTEGER NOT NULL CHECK (avg_reach >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX channels_username_idx ON channels (username);
CREATE INDEX channels_owner_idx ON channels (owner_id);
CREATE INDEX channels_topic_language_idx ON channels (topic, language);

CREATE TABLE decisions (
    from_channel_id BIGINT NOT NULL REFERENCES channels (id) ON DELETE CASCADE,
    to_channel_id BIGINT NOT NULL REFERENCES channels (id) ON DELETE CASCADE,
    interested BOOLEAN NOT NULL,
    decided_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (from_channel_id, to_channel_id),
    CHECK (from_channel_id <> to_channel_id)
);

CREATE INDEX decisions_to_channel_idx ON decisions (to_channel_id);

-- +goose Down
DROP TABLE decisions;
DROP TABLE channels;

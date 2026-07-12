CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE tactical_event (
    id                INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    match_id          INTEGER NOT NULL,
    player_id         INTEGER,
    team_id           INTEGER,
    event_type        TEXT NOT NULL,
    pitch_x           FLOAT,
    pitch_y           FLOAT,
    timestamp_seconds FLOAT,
    confidence_score  FLOAT CHECK (confidence_score >= 0 AND confidence_score <= 1),
    metadata          JSONB,
    embedding         vector(768),
    embedding_model   TEXT,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_tactical_event_match
        FOREIGN KEY (match_id) REFERENCES match (id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_tactical_event_player
        FOREIGN KEY (player_id) REFERENCES player (id)
        ON DELETE SET NULL,

    CONSTRAINT fk_tactical_event_team
        FOREIGN KEY (team_id) REFERENCES team (id)
        ON DELETE SET NULL,

    CONSTRAINT chk_tactical_event_embedding_model
        CHECK (embedding IS NULL OR embedding_model IS NOT NULL)
);

CREATE INDEX idx_tactical_event_match ON tactical_event (match_id);

CREATE INDEX idx_tactical_event_embedding
    ON tactical_event USING hnsw (embedding vector_cosine_ops);
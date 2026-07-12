CREATE TABLE memory_entry (
    id               INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    entry_type       TEXT NOT NULL,
    match_id         INTEGER,
    team_id          INTEGER,
    player_id        INTEGER,
    content          TEXT NOT NULL,
    embedding        vector(768),
    embedding_model  TEXT,
    source_metadata  JSONB,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_memory_entry_match
        FOREIGN KEY (match_id) REFERENCES match (id)
        ON DELETE SET NULL,

    CONSTRAINT fk_memory_entry_team
        FOREIGN KEY (team_id) REFERENCES team (id)
        ON DELETE SET NULL,

    CONSTRAINT fk_memory_entry_player
        FOREIGN KEY (player_id) REFERENCES player (id)
        ON DELETE SET NULL,

    CONSTRAINT chk_memory_entry_embedding_model
        CHECK (embedding IS NULL OR embedding_model IS NOT NULL)
);

CREATE INDEX idx_memory_entry_embedding
    ON memory_entry USING hnsw (embedding vector_cosine_ops);
ALTER TABLE player_match ADD CONSTRAINT uq_player_match_id_match
    UNIQUE (id, match_id);

CREATE TABLE player_position_sample (
    id                 INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    player_match_id    INTEGER NOT NULL,
    match_id           INTEGER NOT NULL,
    pitch_x            FLOAT,
    pitch_y            FLOAT,
    timestamp_seconds  FLOAT,
    frame_number       INTEGER NOT NULL,
    confidence_score   FLOAT CHECK (confidence_score >= 0 AND confidence_score <= 1),
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_player_position_sample_player_match
        FOREIGN KEY (player_match_id)
        REFERENCES player_match (id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_player_position_sample_match
        FOREIGN KEY (player_match_id, match_id)
        REFERENCES player_match (id, match_id)
        ON DELETE RESTRICT,

    CONSTRAINT uq_player_position_sample_frame
        UNIQUE (player_match_id, frame_number)
);

COMMENT ON COLUMN player_position_sample.confidence_score IS 'ByteTrack detection confidence';
COMMENT ON COLUMN player_position_sample.player_match_id IS 'Source of truth for player/team/match identity — see player_match';
COMMENT ON TABLE player_position_sample IS 'Append-only CV output. updated_at exists for schema consistency but rows should not be mutated in practice — insert corrections as new rows.';
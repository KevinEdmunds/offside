CREATE TABLE formation_snapshot (
    id                 INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    match_id           INTEGER NOT NULL,
    team_id            INTEGER NOT NULL,
    timestamp_seconds  FLOAT,
    formation_label    TEXT NOT NULL,
    compactness_score  FLOAT,
    confidence_score   FLOAT CHECK (confidence_score >= 0 AND confidence_score <= 1),
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_formation_snapshot_match
        FOREIGN KEY (match_id) REFERENCES match (id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_formation_snapshot_team
        FOREIGN KEY (team_id) REFERENCES team (id)
        ON DELETE RESTRICT
);

CREATE INDEX idx_formation_snapshot_match ON formation_snapshot (match_id);
CREATE INDEX idx_formation_snapshot_team ON formation_snapshot (team_id);
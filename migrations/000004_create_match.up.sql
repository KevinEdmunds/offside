CREATE TABLE match (
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id     TEXT NOT NULL,
    competition_id  INTEGER NOT NULL,
    home_team_id    INTEGER NOT NULL,
    away_team_id    INTEGER NOT NULL,
    kickoff_time    TIMESTAMPTZ,
    home_score      INTEGER,
    away_score      INTEGER,
    status          TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT match_external_id_unique UNIQUE (external_id),
    CONSTRAINT match_status_check
        CHECK (status IN ('scheduled', 'live', 'finished', 'postponed', 'cancelled')),
    CONSTRAINT fk_match_competition
        FOREIGN KEY (competition_id)
        REFERENCES competition (id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_match_home_team
        FOREIGN KEY (home_team_id)
        REFERENCES team (id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_match_away_team
        FOREIGN KEY (away_team_id)
        REFERENCES team (id)
        ON DELETE RESTRICT
);

COMMENT ON COLUMN match.status IS 'Allowed values: scheduled|live|finished|postponed|cancelled — enforced by match_status_check constraint';
COMMENT ON COLUMN match.home_score IS 'NULL until match is finished';
COMMENT ON COLUMN match.away_score IS 'NULL until match is finished';
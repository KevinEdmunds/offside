CREATE TABLE player (
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id     TEXT NOT NULL,
    first_name      TEXT NOT NULL,
    second_name      TEXT NOT NULL,
    nationality     TEXT,
    primary_team_id INTEGER,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT player_external_id_unique UNIQUE (external_id),
    CONSTRAINT fk_player_primary_team
        FOREIGN KEY (primary_team_id)
        REFERENCES team (id)
        ON DELETE SET NULL
);

COMMENT ON COLUMN player.primary_team_id IS 'Convenience snapshot only — use PLAYER_MATCH.team_id for historical accuracy';
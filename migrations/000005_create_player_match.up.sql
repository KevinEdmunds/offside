CREATE TABLE player_match (
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    player_id       INTEGER NOT NULL,
    match_id        INTEGER NOT NULL,
    team_id         INTEGER NOT NULL,
    position_played TEXT,
    minutes_played  INTEGER,
    started         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT player_match_unique UNIQUE (player_id, match_id),
    CONSTRAINT fk_player_match_player
        FOREIGN KEY (player_id)
        REFERENCES player (id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_player_match_match
        FOREIGN KEY (match_id)
        REFERENCES match (id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_player_match_team
        FOREIGN KEY (team_id)
        REFERENCES team (id)
        ON DELETE RESTRICT
);

COMMENT ON COLUMN player_match.team_id IS 'The team this player represented in this specific match — source of truth for historical team membership, not player.primary_team_id';
COMMENT ON COLUMN player_match.position_played IS 'Position played in this specific match — may differ from player primary position';
COMMENT ON COLUMN player_match.minutes_played IS 'NULL if match not yet played or data not yet available';
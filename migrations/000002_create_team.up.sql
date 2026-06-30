CREATE TABLE team (
    id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id TEXT NOT NULL,
    name        TEXT NOT NULL,
    short_name  TEXT NOT NULL,
    country     TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT team_external_id_unique UNIQUE (external_id)
);
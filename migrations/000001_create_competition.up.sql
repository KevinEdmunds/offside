CREATE TABLE competition (
    id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_id TEXT NOT NULL,
    name        TEXT NOT NULL,
    country     TEXT,
    season      TEXT,
    type        TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT competition_external_id_unique UNIQUE (external_id)
);

COMMENT ON COLUMN competition.country IS 'NULL = international competition not tied to a single country (e.g. World Cup, Champions League)';
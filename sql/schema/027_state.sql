-- +goose Up
CREATE TABLE db_state (
    id INTEGER PRIMARY KEY DEFAULT 1,
    data_hash TEXT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT single_row_check CHECK (id = 1)
);

INSERT INTO db_state(id, data_hash) VALUES (1, 'initial_hash');


-- +goose Down
DROP TABLE IF EXISTS db_state;
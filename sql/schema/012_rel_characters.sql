-- +goose Up
CREATE TABLE j_character_base_stat (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    character_id INTEGER NOT NULL REFERENCES characters(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);




-- +goose Down
DROP TABLE IF EXISTS j_character_base_stat;
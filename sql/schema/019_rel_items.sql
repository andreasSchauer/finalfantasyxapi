-- +goose Up
CREATE TABLE mix_combinations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    first_item_id INTEGER NOT NULL REFERENCES items(id),
    second_item_id INTEGER NOT NULL REFERENCES items(id),

    UNIQUE(first_item_id, second_item_id)
);


CREATE TABLE j_mixes_combinations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    mix_id INTEGER NOT NULL REFERENCES mixes(id),
    combo_id INTEGER NOT NULL REFERENCES mix_combinations(id),
    is_best_combo BOOLEAN NOT NULL,

    UNIQUE (mix_id, combo_id)
);



-- +goose Down
DROP TABLE IF EXISTS j_mixes_combinations;
DROP TABLE IF EXISTS mix_combinations;
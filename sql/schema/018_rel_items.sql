-- +goose Up
CREATE TABLE mix_combinations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    mix_id INTEGER NOT NULL REFERENCES mixes(id),
    first_item_id INTEGER NOT NULL REFERENCES items(id),
    second_item_id INTEGER NOT NULL REFERENCES items(id),
    is_best_combo BOOLEAN NOT NULL,

    UNIQUE(first_item_id, second_item_id)
);


CREATE TABLE j_items_related_stats(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    item_id INTEGER NOT NULL REFERENCES items(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_items_available_menus(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    item_id INTEGER NOT NULL REFERENCES items(id),
    submenu_id INTEGER NOT NULL REFERENCES submenus(id)
);



-- +goose Down
DROP TABLE IF EXISTS j_items_available_menus;
DROP TABLE IF EXISTS j_items_related_stats;
DROP TABLE IF EXISTS mix_combinations;
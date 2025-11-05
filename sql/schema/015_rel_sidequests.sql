-- +goose Up
CREATE TABLE blitzball_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    position_id INTEGER NOT NULL REFERENCES blitzball_positions(id),
    possible_item_id INTEGER NOT NULL REFERENCES possible_items(id)
);


CREATE TABLE quest_completions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    quest_id INTEGER NOT NULL REFERENCES quests(id),
    condition TEXT NOT NULL,
    item_amount_id INTEGER NOT NULL REFERENCES item_amounts(id)
);


CREATE TABLE completion_locations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    completion_id INTEGER NOT NULL REFERENCES quest_completions(id),
    area_id INTEGER NOT NULL REFERENCES areas(id),
    notes TEXT
);


-- +goose Down
DROP TABLE IF EXISTS completion_locations;
DROP TABLE IF EXISTS quest_completions;
DROP TABLE IF EXISTS blitzball_items;
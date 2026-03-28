-- +goose Up
ALTER TABLE monster_arena_creations
ADD COLUMN monster_id INTEGER REFERENCES monsters(id);

CREATE TABLE blitzball_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    position_id INTEGER NOT NULL REFERENCES blitzball_positions(id),
    possible_item_id INTEGER NOT NULL REFERENCES possible_items(id)
);


CREATE TABLE quest_completions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    condition TEXT,
    item_amount_id INTEGER NOT NULL REFERENCES item_amounts(id)
);


CREATE TABLE completion_areas (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    completion_id INTEGER NOT NULL REFERENCES quest_completions(id),
    area_id INTEGER NOT NULL REFERENCES areas(id),
    notes TEXT
);

ALTER TABLE quests
ADD COLUMN completion_id INTEGER REFERENCES quest_completions(id);


-- +goose Down
ALTER TABLE quests
DROP COLUMN IF EXISTS completion_id;

DROP TABLE IF EXISTS completion_areas;
DROP TABLE IF EXISTS quest_completions;
DROP TABLE IF EXISTS blitzball_items;

ALTER TABLE monster_arena_creations
DROP COLUMN IF EXISTS monster_id;
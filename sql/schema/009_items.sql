-- +goose Up
CREATE TYPE item_type AS ENUM ('item', 'key-item');
CREATE TYPE item_category AS ENUM ('distiller', 'healing', 'offensive', 'other', 'sphere', 'support');
CREATE TYPE item_usability AS ENUM ('always', 'in-battle', 'outside-battle');
CREATE TYPE key_item_category AS ENUM ('celestial', 'jecht-sphere', 'other', 'primer', 'story');

CREATE TABLE master_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    type item_type NOT NULL,

    UNIQUE(name, type)
);


CREATE TABLE items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    master_items_id INTEGER UNIQUE NOT NULL REFERENCES master_items(id),
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    sphere_grid_description TEXT,
    category item_category NOT NULL,
    usability item_usability,
    base_price INTEGER,
    sell_value INTEGER NOT NULL
);


CREATE TABLE item_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    item_id INTEGER NOT NULL REFERENCES items(id),
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    cursor target_type NOT NULL,

    UNIQUE (item_id, ability_id)
);


CREATE TABLE key_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    master_items_id INTEGER UNIQUE NOT NULL REFERENCES master_items(id),
    category key_item_category NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS key_items;
DROP TABLE IF EXISTS item_abilities;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS master_items;
DROP TYPE IF EXISTS item_type;
DROP TYPE IF EXISTS item_category;
DROP TYPE IF EXISTS item_usability;
DROP TYPE IF EXISTS key_item_category;
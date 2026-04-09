-- +goose Up
CREATE TYPE item_type AS ENUM ('item', 'key-item');


CREATE TYPE item_category AS ENUM ('distiller', 'healing', 'offensive', 'other', 'sphere', 'support');


CREATE TYPE item_usability AS ENUM ('always', 'in-battle', 'outside-battle', 'unusable');


CREATE TYPE sphere_color AS ENUM ('red', 'yellow', 'black', 'purple', 'blue', 'white');


CREATE TYPE sphere_effect AS ENUM('activation', 'removal', 'creation', 'teleportation');


CREATE TYPE node_position AS ENUM('neighboring', 'ally-position', 'any');


CREATE TYPE node_state AS ENUM('active-self', 'active-ally', 'active-any', 'inactive', 'any');


CREATE TYPE node_type AS ENUM('hp', 'mp', 'strength', 'defense', 'magic', 'magic-defense', 'agility', 'luck', 'evasion', 'accuracy', 'skill', 'special', 'wht-magic', 'blk-magic', 'lv-1-lock', 'lv-2-lock', 'lv-3-lock', 'lv-4-lock', 'empty');


CREATE TYPE key_item_category AS ENUM ('celestial', 'jecht-sphere', 'other', 'primer', 'story');


CREATE TYPE mix_category AS ENUM ('9999-damage', 'critical-hits', 'fire-elemental', 'gravity-based', 'hp-mp', 'ice-elemental', 'lightning-elemental', 'non-elemental', 'overdrive-speed', 'positive-status', 'recovery', 'water-elemental');




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
    master_item_id INTEGER UNIQUE NOT NULL REFERENCES master_items(id),
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    category item_category NOT NULL,
    usability item_usability NOT NULL,
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


CREATE TABLE spheres (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    item_id INTEGER NOT NULL REFERENCES items(id),
    sphere_grid_description TEXT NOT NULL,
    sphere_color sphere_color NOT NULL,
    sphere_effect sphere_effect NOT NULL,
    target_node_position node_position NOT NULL,
    target_node_state node_state
);


CREATE TABLE key_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    master_item_id INTEGER UNIQUE NOT NULL REFERENCES master_items(id),
    category key_item_category NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL
);


CREATE TABLE primers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    key_item_id INTEGER UNIQUE NOT NULL REFERENCES key_items(id),
    al_bhed_letter TEXT UNIQUE NOT NULL,
    english_letter TEXT UNIQUE NOT NULL
);



CREATE TABLE mixes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_id INTEGER UNIQUE NOT NULL REFERENCES overdrives(id),
    category mix_category NOT NULL
);


CREATE TABLE item_amounts (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    master_item_id INTEGER NOT NULL REFERENCES master_items(id),
    amount INTEGER NOT NULL
);



CREATE TABLE possible_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    item_amount_id INTEGER NOT NULL REFERENCES item_amounts(id),
    chance percent NOT NULL
);




-- +goose Down
DROP TABLE IF EXISTS possible_items;
DROP TABLE IF EXISTS item_amounts;
DROP TABLE IF EXISTS mixes;
DROP TABLE IF EXISTS primers;
DROP TABLE IF EXISTS key_items;
DROP TABLE IF EXISTS spheres;
DROP TABLE IF EXISTS item_abilities;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS master_items;
DROP TYPE IF EXISTS mix_category;
DROP TYPE IF EXISTS key_item_category;
DROP TYPE IF EXISTS node_type;
DROP TYPE IF EXISTS node_state;
DROP TYPE IF EXISTS node_position;
DROP TYPE IF EXISTS sphere_effect;
DROP TYPE IF EXISTS sphere_color;
DROP TYPE IF EXISTS item_type;
DROP TYPE IF EXISTS item_category;
DROP TYPE IF EXISTS item_usability;
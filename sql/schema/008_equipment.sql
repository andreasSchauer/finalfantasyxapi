-- +goose Up
CREATE TYPE key_item_base AS ENUM ('sun', 'moon', 'jupiter', 'venus', 'saturn', 'mars', 'mercury');
CREATE TYPE celestial_formula AS ENUM ('hp-high', 'hp-low', 'mp-high');

CREATE TABLE celestial_weapons (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    key_item_base key_item_base UNIQUE NOT NULL,
    formula celestial_formula NOT NULL
);


CREATE TYPE equip_type AS ENUM ('weapon', 'armor');
CREATE TYPE auto_ability_category AS ENUM ('ap-overdrive', 'auto-cure', 'auto-status', 'break-limit', 'counter', 'elemental-protection', 'elemental-strike', 'other', 'sos-status', 'stat-+x%', 'status-infliction', 'status-protection');
CREATE TYPE aa_activation_condition AS ENUM ('always', 'active-party', 'hp-critical', 'outside-battle');
CREATE TYPE counter_type AS ENUM ('physical', 'magical');



CREATE TABLE auto_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    effect TEXT NOT NULL,
    type equip_type,
    category auto_ability_category NOT NULL,
    ability_value INTEGER,
    activation_condition aa_activation_condition,
    counter counter_type
);


CREATE TYPE equip_class AS ENUM ('standard', 'unique', 'celestial-weapon');

CREATE TABLE equipment_tables (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    type equip_type NOT NULL,
    classification equip_class NOT NULL,
    specific_character_id INTEGER REFERENCES characters(id),
    version INTEGER,
    priority INTEGER,
    pool_1_amt INTEGER,
    pool_2_amt INTEGER,
    empty_slots_amt INTEGER NOT NULL,
    UNIQUE(type, classification, specific_character_id, version, priority)
);



-- +goose Down
DROP TABLE IF EXISTS equipment_tables;
DROP TYPE IF EXISTS equip_class;
DROP TABLE IF EXISTS auto_abilities;
DROP TYPE IF EXISTS counter_type;
DROP TYPE IF EXISTS aa_activation_condition;
DROP TYPE IF EXISTS auto_ability_category;
DROP TYPE IF EXISTS equip_type;
DROP TABLE IF EXISTS celestial_weapons;
DROP TYPE IF EXISTS celestial_formula;
DROP TYPE IF EXISTS key_item_base;
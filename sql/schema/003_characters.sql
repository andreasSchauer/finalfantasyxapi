-- +goose Up
CREATE TYPE unit_type AS ENUM ('aeon', 'character');


CREATE TYPE weapon_type AS ENUM ('sword', 'staff', 'blitzball', 'doll', 'spear', 'blade', 'claw', 'seymour-staff');


CREATE TYPE armor_type AS ENUM ('shield', 'ring', 'armguard', 'bangle', 'armlet', 'bracer', 'targe', 'seymour-armor');


CREATE TYPE accuracy_source AS ENUM ('accuracy', 'rate');


CREATE TABLE player_units (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    type unit_type NOT NULL
);


CREATE TABLE characters (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    unit_id INTEGER NOT NULL REFERENCES player_units(id),
    story_only BOOLEAN NOT NULL,
    weapon_type weapon_type NOT NULL,
    armor_type armor_type NOT NULL,
    physical_attack_range distance NOT NULL,
    can_fight_underwater BOOLEAN NOT NULL
);




CREATE TABLE aeons (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    unit_id INTEGER NOT NULL REFERENCES player_units(id),
    unlock_condition TEXT NOT NULL,
    is_optional BOOLEAN NOT NULL,
    battles_to_regenerate INTEGER NOT NULL,
    phys_atk_damage_constant INTEGER,
    phys_atk_range distance,
    phys_atk_shatter_rate uint8
);


CREATE TABLE character_classes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL
);



-- +goose Down
DROP TABLE IF EXISTS character_classes;
DROP TABLE IF EXISTS aeons;
DROP TABLE IF EXISTS characters;
DROP TABLE IF EXISTS player_units;
DROP TYPE IF EXISTS accuracy_source;
DROP TYPE IF EXISTS weapon_type;
DROP TYPE IF EXISTS armor_type;
DROP TYPE IF EXISTS unit_type;
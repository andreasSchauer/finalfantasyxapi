-- +goose Up
CREATE TYPE unit_type AS ENUM ('aeon', 'character');

CREATE TYPE weapon_type AS ENUM ('sword', 'staff', 'blitzball', 'doll', 'spear', 'blade', 'claw', 'seymour-staff');

CREATE TYPE armor_type AS ENUM ('shield', 'ring', 'armguard', 'bangle', 'armlet', 'bracer', 'targe', 'seymour-armor');

CREATE TYPE accuracy_source AS ENUM ('accuracy', 'rate');

CREATE TYPE character_class_category AS ENUM ('character', 'aeon', 'group');

CREATE TYPE sphere_grid_type AS ENUM ('standard', 'expert');


-- 1
CREATE TABLE player_units (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    type unit_type NOT NULL
);


-- 1
CREATE TABLE sphere_grids (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    type sphere_grid_type NOT NULL,
    hp INTEGER NOT NULL,
    mp INTEGER NOT NULL,
    strength INTEGER NOT NULL,
    defense INTEGER NOT NULL,
    magic INTEGER NOT NULL,
    magic_defense INTEGER NOT NULL,
    agility INTEGER NOT NULL,
    luck INTEGER NOT NULL,
    evasion INTEGER NOT NULL,
    accuracy INTEGER NOT NULL,
    lv_1_locks INTEGER NOT NULL,
    lv_2_locks INTEGER NOT NULL,
    lv_3_locks INTEGER NOT NULL,
    lv_4_locks INTEGER NOT NULL,
    empty_nodes INTEGER NOT NULL
);


-- 4
CREATE TABLE characters (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    unit_id INTEGER NOT NULL REFERENCES player_units(id),
    is_story_based BOOLEAN NOT NULL,
    weapon_type weapon_type NOT NULL,
    armor_type armor_type NOT NULL,
    physical_attack_range distance NOT NULL,
    can_fight_underwater BOOLEAN NOT NULL,
    std_sphere_grid_id INTEGER REFERENCES sphere_grids(id),
    exp_sphere_grid_id INTEGER REFERENCES sphere_grids(id)
);


-- 4
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


-- 1
CREATE TABLE character_classes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    category character_class_category NOT NULL
);



-- +goose Down
DROP TABLE IF EXISTS character_classes;
DROP TABLE IF EXISTS aeons;
DROP TABLE IF EXISTS characters;
DROP TABLE IF EXISTS sphere_grids;
DROP TABLE IF EXISTS player_units;
DROP TYPE IF EXISTS sphere_grid_type;
DROP TYPE IF EXISTS character_class_category;
DROP TYPE IF EXISTS accuracy_source;
DROP TYPE IF EXISTS weapon_type;
DROP TYPE IF EXISTS armor_type;
DROP TYPE IF EXISTS unit_type;
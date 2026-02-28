-- +goose Up
CREATE TYPE ability_type AS ENUM ('other-ability', 'player-ability', 'enemy-ability', 'overdrive-ability', 'trigger-command', 'item-ability');


CREATE TYPE player_ability_category AS ENUM ('skill', 'special', 'white-magic', 'black-magic', 'aeon');


CREATE TABLE ability_attributes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    rank INTEGER,
    appears_in_help_bar BOOLEAN NOT NULL,
    can_copycat BOOLEAN NOT NULL,
    UNIQUE(rank, appears_in_help_bar, can_copycat)
);


CREATE TABLE abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    version INTEGER,
    specification TEXT,
    attributes_id INTEGER REFERENCES ability_attributes(id),
    type ability_type NOT NULL,

    UNIQUE(name, version, type),
    CHECK(
        (type != 'overdrive-ability' AND attributes_id IS NOT NULL) OR
        (type = 'overdrive-ability' AND attributes_id IS NULL)
    )
);


CREATE TABLE other_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id),
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    cursor target_type
);


CREATE TABLE player_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id),
    description TEXT,
    effect TEXT NOT NULL,
    category player_ability_category NOT NULL,
    can_use_outside_battle BOOLEAN NOT NULL,
    mp_cost INTEGER NOT NULL,
    cursor target_type
);


CREATE TABLE enemy_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id),
    effect TEXT
);


CREATE TABLE overdrive_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id)
);


CREATE TABLE trigger_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id),
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    cursor target_type NOT NULL
);


CREATE TABLE overdrives (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    version INTEGER,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    attributes_id INTEGER NOT NULL REFERENCES ability_attributes(id),
    unlock_condition TEXT,
    countdown_in_sec INTEGER,
    cursor target_type,

    UNIQUE(name, version)
);


CREATE TABLE ronso_rages (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_id INTEGER UNIQUE NOT NULL REFERENCES overdrives(id)
);



-- +goose Down
DROP TABLE IF EXISTS ronso_rages;
DROP TABLE IF EXISTS overdrives;
DROP TABLE IF EXISTS trigger_commands;
DROP TABLE IF EXISTS overdrive_abilities;
DROP TABLE IF EXISTS enemy_abilities;
DROP TABLE IF EXISTS player_abilities;
DROP TABLE IF EXISTS other_abilities;
DROP TABLE IF EXISTS abilities;
DROP TABLE IF EXISTS ability_attributes;
DROP TYPE IF EXISTS player_ability_category;
DROP TYPE IF EXISTS ability_type;
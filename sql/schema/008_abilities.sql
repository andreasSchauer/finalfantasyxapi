-- +goose Up
CREATE TYPE ability_type AS ENUM ('player-ability', 'enemy-ability', 'overdrive-ability', 'trigger-command', 'item');

CREATE TABLE abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    version INTEGER,
    specification TEXT,
    type ability_type NOT NULL,

    UNIQUE(name, version, type)
);


CREATE TYPE submenu_type AS ENUM ('blk-magic', 'skill', 'special', 'summon', 'wht-magic');

CREATE TABLE player_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id),
    description TEXT,
    effect TEXT NOT NULL,
    submenu submenu_type,
    can_use_outside_battle BOOLEAN NOT NULL,
    mp_cost INTEGER,
    rank INTEGER,
    appears_in_help_bar BOOLEAN NOT NULL,
    can_copycat BOOLEAN NOT NULL
);


CREATE TABLE enemy_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER UNIQUE NOT NULL REFERENCES abilities(id),
    effect TEXT,
    rank INTEGER NOT NULL,
    appears_in_help_bar BOOLEAN NOT NULL,
    can_copycat BOOLEAN NOT NULL
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
    rank INTEGER NOT NULL,
    appears_in_help_bar BOOLEAN NOT NULL,
    can_copycat BOOLEAN NOT NULL
);



-- +goose Down
DROP TABLE IF EXISTS trigger_commands;
DROP TABLE IF EXISTS overdrive_abilities;
DROP TABLE IF EXISTS enemy_abilities;
DROP TABLE IF EXISTS player_abilities;
DROP TYPE IF EXISTS submenu_type;
DROP TABLE IF EXISTS abilities;
DROP TYPE IF EXISTS ability_type;
-- +goose Up
CREATE TYPE command_category AS ENUM ('menu', 'random-ability');
CREATE TYPE target_type AS ENUM ('self', 'single-character', 'single-enemy', 'single-target', 'random-enemy', 'all-characters', 'all-enemies', 'target-party', 'everyone');

CREATE TABLE commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    category command_category NOT NULL,
    cursor target_type
);


CREATE TYPE submenu_type AS ENUM ('blk-magic', 'skill', 'special', 'summon', 'wht-magic', 'item', 'weapon', 'armor', 'characters', 'use', 'overdrive');


CREATE TABLE overdrive_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    rank INTEGER NOT NULL,
    open_menu submenu_type
);



CREATE TABLE overdrives (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    od_command_id INTEGER references overdrive_commands(id),
    name TEXT NOT NULL,
    version INTEGER,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    rank INTEGER NOT NULL,
    appears_in_help_bar BOOLEAN NOT NULL,
    can_copycat BOOLEAN NOT NULL,
    unlock_condition TEXT,
    countdown_in_sec INTEGER,
    cursor target_type,

    UNIQUE(name, version)
);


-- +goose Down
DROP TABLE IF EXISTS overdrives;
DROP TABLE IF EXISTS overdrive_commands;
DROP TYPE IF EXISTS submenu_type;
DROP TABLE IF EXISTS commands;
DROP TYPE IF EXISTS command_category;
DROP TYPE IF EXISTS target_type;
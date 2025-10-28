-- +goose Up
CREATE TYPE target_type AS ENUM ('self', 'single-character', 'single-enemy', 'single-target', 'random-enemy', 'all-characters', 'all-enemies', 'target-party', 'everyone');


CREATE TYPE topmenu_type AS ENUM ('main', 'left', 'right', 'hidden');


CREATE TABLE aeon_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    topmenu topmenu_type NOT NULL,
    cursor target_type
);


CREATE TABLE submenus (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    topmenu topmenu_type
);


CREATE TABLE overdrive_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    rank INTEGER NOT NULL,
    topmenu topmenu_type NOT NULL
); 




-- +goose Down
DROP TABLE IF EXISTS overdrive_commands;
DROP TABLE IF EXISTS submenus;
DROP TABLE IF EXISTS aeon_commands;
DROP TYPE IF EXISTS topmenu_type;
DROP TYPE IF EXISTS target_type;
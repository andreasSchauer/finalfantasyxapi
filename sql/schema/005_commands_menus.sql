-- +goose Up
CREATE TYPE target_type AS ENUM ('self', 'single-character', 'single-enemy', 'single-target', 'random-character', 'random-enemy', 'all-characters', 'all-enemies', 'target-party', 'n-targets', 'everyone');



CREATE TABLE topmenus (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL
);


CREATE TABLE aeon_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    cursor target_type
);


CREATE TABLE submenus (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL
);


CREATE TABLE overdrive_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    rank INTEGER NOT NULL
); 




-- +goose Down
DROP TABLE IF EXISTS overdrive_commands;
DROP TABLE IF EXISTS submenus;
DROP TABLE IF EXISTS aeon_commands;
DROP TABLE IF EXISTS topmenus;
DROP TYPE IF EXISTS target_type;
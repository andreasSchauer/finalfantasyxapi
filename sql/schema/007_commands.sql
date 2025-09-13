-- +goose Up
CREATE TYPE target_type AS ENUM ('self', 'single-character', 'single-enemy', 'single-target', 'random-enemy', 'all-characters', 'all-enemies', 'target-party', 'everyone');

CREATE TABLE aeon_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    cursor target_type
);


CREATE TABLE menu_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL
);






-- +goose Down
DROP TABLE IF EXISTS menu_commands;
DROP TABLE IF EXISTS aeon_commands;
DROP TYPE IF EXISTS target_type;
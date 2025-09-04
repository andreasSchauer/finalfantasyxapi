-- +goose Up
CREATE TABLE stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    effect TEXT NOT NULL,
    min_val INTEGER NOT NULL,
    max_val INTEGER NOT NULL,
    max_val_2 INTEGER
);


CREATE TABLE elements (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);


CREATE TABLE affinities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    damage_factor REAL
);


CREATE TABLE agility_tiers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    min_agility INTEGER NOT NULL,
    max_agility INTEGER NOT NULL,
    tick_speed INTEGER NOT NULL,
    monster_min_icv INTEGER,
    monster_max_icv INTEGER,
    character_max_icv INTEGER
);


CREATE TABLE agility_subtiers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    agility_tier_id INTEGER NOT NULL REFERENCES agility_tiers(id),
    subtier_min_agility INTEGER NOT NULL,
    subtier_max_agility INTEGER NOT NULL,
    character_min_icv INTEGER
);


CREATE TYPE overdrive_type AS ENUM ('formula', 'per-action');

CREATE TABLE overdrive_modes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    type overdrive_type NOT NULL,
    fill_rate REAL
);


CREATE TABLE status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    effect TEXT NOT NULL
);


CREATE TABLE properties (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    effect TEXT NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS properties;
DROP TABLE IF EXISTS status_conditions;
DROP TABLE IF EXISTS overdrive_modes;
DROP TYPE IF EXISTS overdrive_type;
DROP TABLE IF EXISTS agility_subtiers;
DROP TABLE IF EXISTS agility_tiers;
DROP TABLE IF EXISTS affinities;
DROP TABLE IF EXISTS elements;
DROP TABLE IF EXISTS stats;
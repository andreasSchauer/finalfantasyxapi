-- +goose Up
CREATE TYPE overdrive_mode_type AS ENUM ('formula', 'per-action');


CREATE TYPE nullify_armored AS ENUM ('target', 'bearer');


CREATE TYPE modifier_type AS ENUM ('battle-dependent', 'factor', 'fixed-value', 'percentage');


CREATE TABLE stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    effect TEXT NOT NULL,
    min_val INTEGER NOT NULL,
    max_val INTEGER NOT NULL,
    max_val_2 INTEGER
);


CREATE TABLE base_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    stat_id INTEGER NOT NULL REFERENCES stats(id),
    value INTEGER NOT NULL
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


CREATE TABLE elemental_resists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    element_id INTEGER NOT NULL REFERENCES elements(id),
    affinity_id INTEGER NOT NULL REFERENCES affinities(id)
);



CREATE TABLE agility_tiers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    min_agility uint8 NOT NULL,
    max_agility uint8 NOT NULL,
    tick_speed INTEGER NOT NULL,
    monster_min_icv INTEGER,
    monster_max_icv INTEGER,
    character_max_icv INTEGER
);


CREATE TABLE agility_subtiers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    agility_tier_id INTEGER NOT NULL REFERENCES agility_tiers(id),
    subtier_min_agility uint8 NOT NULL,
    subtier_max_agility uint8 NOT NULL,
    character_min_icv INTEGER
);


CREATE TABLE overdrive_modes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    type overdrive_mode_type NOT NULL,
    fill_rate percentage
);


CREATE TABLE status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    effect TEXT NOT NULL,
    visualization TEXT,
    nullify_armored nullify_armored
);


CREATE TABLE status_resists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    resistance uint8 NOT NULL
);


CREATE TABLE properties (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    effect TEXT NOT NULL,
    nullify_armored nullify_armored
);


CREATE TABLE modifiers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    effect TEXT NOT NULL,
    type modifier_type NOT NULL,
    default_value REAL
);





-- +goose Down
DROP TABLE IF EXISTS modifiers;
DROP TABLE IF EXISTS properties;
DROP TABLE IF EXISTS status_resists;
DROP TABLE IF EXISTS status_conditions;
DROP TABLE IF EXISTS overdrive_modes;
DROP TABLE IF EXISTS agility_subtiers;
DROP TABLE IF EXISTS agility_tiers;
DROP TABLE IF EXISTS elemental_resists;
DROP TABLE IF EXISTS affinities;
DROP TABLE IF EXISTS elements;
DROP TABLE IF EXISTS base_stats;
DROP TABLE IF EXISTS stats;
DROP TYPE IF EXISTS modifier_type;
DROP TYPE IF EXISTS nullify_armored;
DROP TYPE IF EXISTS overdrive_mode_type;
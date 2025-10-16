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


CREATE TABLE elemental_affinities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    element_id INTEGER NOT NULL REFERENCES elements(id),
    affinity_id INTEGER NOT NULL REFERENCES affinities(id)
);


CREATE DOMAIN uint8 AS INTEGER
    CHECK (VALUE >= 0 AND VALUE <= 255);

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


CREATE TYPE overdrive_type AS ENUM ('formula', 'per-action');
CREATE DOMAIN percentage AS REAL
    CHECK (VALUE >= 0 AND VALUE <= 1);

CREATE TABLE overdrive_modes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    effect TEXT NOT NULL,
    type overdrive_type NOT NULL,
    fill_rate percentage
);


CREATE TYPE nullify_armored AS ENUM ('target', 'bearer');

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


CREATE TYPE duration_type AS ENUM ('blocks', 'endless', 'instant', 'turns', 'user-turns');

CREATE TABLE inflicted_statusses (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    probability uint8 NOT NULL,
    duration_type duration_type NOT NULL,
    amount INTEGER
);


CREATE TABLE properties (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    effect TEXT NOT NULL,
    nullify_armored nullify_armored
);


CREATE TYPE modifier_type AS ENUM ('battle-dependent', 'factor','fixed-value', 'percentage');

CREATE TABLE modifiers (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    effect TEXT NOT NULL,
    type modifier_type NOT NULL,
    default_value REAL
);


CREATE TYPE calculation_type AS ENUM ('added-percentage', 'added-value', 'multiply', 'multiply-highest', 'set-value');

CREATE TABLE stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    stat_id INTEGER NOT NULL REFERENCES stats(id),
    calculation_type calculation_type NOT NULL,
    value REAL NOT NULL
);


CREATE TABLE modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    modifier_id INTEGER NOT NULL REFERENCES modifiers(id),
    calculation_type calculation_type NOT NULL,
    value REAL NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS modifier_changes;
DROP TABLE IF EXISTS stat_changes;
DROP TYPE IF EXISTS calculation_type;
DROP TABLE IF EXISTS modifiers;
DROP TYPE IF EXISTS modifier_type;
DROP TABLE IF EXISTS properties;
DROP TABLE IF EXISTS inflicted_statusses;
DROP TYPE IF EXISTS duration_type;
DROP TABLE IF EXISTS status_resists;
DROP TABLE IF EXISTS status_conditions;
DROP TYPE IF EXISTS nullify_armored;
DROP TABLE IF EXISTS overdrive_modes;
DROP TYPE IF EXISTS overdrive_type;
DROP DOMAIN IF EXISTS percentage;
DROP TABLE IF EXISTS agility_subtiers;
DROP TABLE IF EXISTS agility_tiers;
DROP DOMAIN IF EXISTS uint8;
DROP TABLE IF EXISTS elemental_affinities;
DROP TABLE IF EXISTS affinities;
DROP TABLE IF EXISTS elements;
DROP TABLE IF EXISTS stats;
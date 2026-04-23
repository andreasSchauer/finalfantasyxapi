-- +goose Up
CREATE TYPE blitzball_tournament_category AS ENUM ('league', 'tournament');
CREATE TYPE blitzball_position_slot AS ENUM ('1st', '2nd', '3rd', 'top-scorer');
CREATE TYPE quest_type AS ENUM ('sidequest', 'subquest');
CREATE TYPE ma_creation_category AS ENUM ('area', 'species', 'original');
CREATE DOMAIN null_ma_creation_category AS ma_creation_category;
CREATE TYPE ma_creation_area AS ENUM ('besaid', 'kilika', 'mi''ihen-highroad', 'mushroom-rock-road', 'djose', 'thunder-plains', 'macalania', 'bikanel', 'calm-lands', 'cavern-of-the-stolen-fayth', 'mount-gagazet', 'sin', 'omega-ruins');
CREATE DOMAIN null_ma_creation_area AS ma_creation_area;
CREATE TYPE ma_creation_species AS ENUM ('bird', 'bomb', 'drake', 'elemental', 'evil-eye', 'flan', 'fungus', 'helm', 'imp', 'iron-giant', 'lupine', 'reptile', 'ruminant', 'wasp');
CREATE DOMAIN null_ma_creation_species AS ma_creation_species;
CREATE TYPE creations_unlocked_category AS ENUM ('area', 'species');
CREATE DOMAIN null_creations_unlocked_category AS creations_unlocked_category;
CREATE TYPE availability_type AS ENUM ('always', 'pre-story', 'post', 'post-story', 'post-game', 'story');
CREATE DOMAIN null_availability_type AS availability_type;


-- 1
CREATE TABLE blitzball_positions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    category blitzball_tournament_category NOT NULL,
    slot blitzball_position_slot NOT NULL,
    UNIQUE(category, slot)
);


-- 4
CREATE TABLE quests (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    type quest_type NOT NULL,
    availability availability_type NOT NULL,
    is_repeatable BOOLEAN NOT NULL,
    UNIQUE(name, type)
);


-- 5
CREATE TABLE sidequests (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    quest_id INTEGER UNIQUE NOT NULL REFERENCES quests(id)
);


-- 6
CREATE TABLE subquests (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    quest_id INTEGER NOT NULL REFERENCES quests(id),
    sidequest_id INTEGER NOT NULL REFERENCES sidequests(id),
    UNIQUE(quest_id, sidequest_id)
);


-- 7
CREATE TABLE monster_arena_creations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    subquest_id INTEGER NOT NULL REFERENCES subquests(id),
    category ma_creation_category NOT NULL,
    required_area ma_creation_area,
    required_species ma_creation_species,
    underwater_only BOOLEAN NOT NULL,
    creations_unlocked_category creations_unlocked_category,
    amount INTEGER NOT NULL,

    CHECK (
        (creations_unlocked_category IS NULL) OR
        (required_area IS NULL AND required_species IS NULL AND underwater_only = false)
    ),
    CHECK (amount >= 0 AND amount <= 10)
);


-- +goose Down
DROP TABLE IF EXISTS monster_arena_creations;
DROP TABLE IF EXISTS subquests;
DROP TABLE IF EXISTS sidequests;
DROP TABLE IF EXISTS quests;
DROP TABLE IF EXISTS blitzball_positions;
DROP DOMAIN IF EXISTS null_availability_type;
DROP TYPE IF EXISTS availability_type;
DROP DOMAIN IF EXISTS null_ma_creation_species;
DROP TYPE IF EXISTS ma_creation_species;
DROP DOMAIN IF EXISTS null_ma_creation_area;
DROP TYPE IF EXISTS ma_creation_area;
DROP DOMAIN IF EXISTS null_ma_creation_category;
DROP TYPE IF EXISTS ma_creation_category;
DROP DOMAIN IF EXISTS null_creations_unlocked_category;
DROP TYPE IF EXISTS creations_unlocked_category;
DROP TYPE IF EXISTS quest_type;
DROP TYPE IF EXISTS blitzball_tournament_category;
DROP TYPE IF EXISTS blitzball_position_slot;
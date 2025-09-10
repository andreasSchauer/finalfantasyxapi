-- +goose Up
CREATE TABLE treasure_lists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY
);


CREATE TYPE treasure_type AS ENUM ('chest', 'gift', 'object');
CREATE TYPE loot_type AS ENUM ('item', 'equipment', 'gil');

CREATE TABLE treasures (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    treasure_list_id INTEGER NOT NULL REFERENCES treasure_lists(id),
    version INTEGER NOT NULL,
    treasure_type treasure_type NOT NULL,
    loot_type loot_type NOT NULL,
    is_post_airship BOOLEAN NOT NULL,
    is_anima_treasure BOOLEAN NOT NULL,
    notes TEXT,
    gil_amount INTEGER
);


CREATE TYPE shop_category AS ENUM ('standard', 'oaka', 'travel-agency', 'wantz');

CREATE TABLE shops (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    version INTEGER,
    notes TEXT,
    category shop_category NOT NULL
);


CREATE TYPE blitzball_tournament_category AS ENUM ('league', 'tournament');
CREATE TYPE blitzball_item_slot AS ENUM ('1st', '2nd', '3rd', 'top-scorer');

CREATE TABLE blitzball_items_lists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    category blitzball_tournament_category NOT NULL,
    slot blitzball_item_slot NOT NULL,
    UNIQUE(category, slot)
);


CREATE TYPE quest_type AS ENUM ('sidequest', 'subquest');

CREATE TABLE quests (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    type quest_type NOT NULL,
    UNIQUE(name, type)
);


CREATE TABLE sidequests (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    quest_id INTEGER UNIQUE NOT NULL REFERENCES quests(id)
);


CREATE TABLE subquests (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    quest_id INTEGER NOT NULL REFERENCES quests(id),
    parent_sidequest_id INTEGER NOT NULL REFERENCES sidequests(id),
    UNIQUE(quest_id, parent_sidequest_id)
);


CREATE TYPE ma_creation_category AS ENUM ('area', 'species', 'original');
CREATE TYPE ma_creation_area AS ENUM ('besaid', 'kilika', 'mi''ihen-highroad', 'mushroom-rock-road', 'djose', 'thunder-plains', 'macalania', 'bikanel', 'calm-lands', 'cavern-of-the-stolen-fayth', 'mount-gagazet', 'sin', 'omega-ruins');
CREATE TYPE ma_creation_species AS ENUM ('bird', 'bomb', 'drake', 'elemental', 'evil-eye', 'flan', 'fungus', 'helm', 'imp', 'iron-giant', 'lupine', 'reptile', 'ruminant', 'wasp');
CREATE TYPE creations_unlocked_category AS ENUM ('area', 'species');


CREATE TABLE monster_arena_creations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
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
DROP TYPE IF EXISTS ma_creation_area;
DROP TYPE IF EXISTS ma_creation_species;
DROP TYPE IF EXISTS ma_creation_category;
DROP TYPE IF EXISTS creations_unlocked_category;
DROP TABLE IF EXISTS subquests;
DROP TABLE IF EXISTS sidequests;
DROP TABLE IF EXISTS quests;
DROP TYPE IF EXISTS quest_type;
DROP TABLE IF EXISTS blitzball_items_lists;
DROP TYPE IF EXISTS blitzball_tournament_category;
DROP TYPE IF EXISTS blitzball_item_slot;
DROP TABLE IF EXISTS shops;
DROP TYPE IF EXISTS shop_category;
DROP TABLE IF EXISTS treasures;
DROP TYPE IF EXISTS treasure_type;
DROP TYPE IF EXISTS loot_type;
DROP TABLE IF EXISTS treasure_lists;
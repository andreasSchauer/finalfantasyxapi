-- +goose Up
CREATE TYPE treasure_type AS ENUM ('chest', 'gift', 'object');


CREATE TYPE loot_type AS ENUM ('item', 'equipment', 'gil');


CREATE TYPE shop_category AS ENUM ('standard', 'oaka', 'travel-agency', 'wantz');


CREATE TABLE locations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL
);


CREATE TABLE sublocations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    location_id INTEGER NOT NULL REFERENCES locations(id),
    name TEXT UNIQUE NOT NULL,
    specification TEXT
);


CREATE TABLE areas (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    sublocation_id INTEGER NOT NULL REFERENCES sublocations(id),
    name TEXT NOT NULL,
    version INTEGER,
    specification TEXT,
    story_only BOOLEAN NOT NULL,
    has_save_sphere BOOLEAN NOT NULL,
    airship_drop_off BOOLEAN NOT NULL,
    has_compilation_sphere BOOLEAN NOT NULL,
    can_ride_chocobo BOOLEAN NOT NULL,
    
    UNIQUE(sublocation_id, name, version)
);


CREATE TABLE treasures (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    version INTEGER NOT NULL,
    treasure_type treasure_type NOT NULL,
    loot_type loot_type NOT NULL,
    is_post_airship BOOLEAN NOT NULL,
    is_anima_treasure BOOLEAN NOT NULL,
    notes TEXT,
    gil_amount INTEGER,

    UNIQUE(area_id, version)
);


CREATE TABLE shops (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    version INTEGER,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    notes TEXT,
    category shop_category NOT NULL,

    UNIQUE(area_id, version)
);


CREATE TABLE encounter_locations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    version INTEGER,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    notes TEXT,

    UNIQUE(area_id, version)
);


-- +goose Down
DROP TABLE IF EXISTS encounter_locations;
DROP TABLE IF EXISTS shops;
DROP TABLE IF EXISTS treasures;
DROP TABLE IF EXISTS areas;
DROP TABLE IF EXISTS sublocations;
DROP TABLE IF EXISTS locations;
DROP TYPE IF EXISTS shop_category;
DROP TYPE IF EXISTS treasure_type;
DROP TYPE IF EXISTS loot_type;
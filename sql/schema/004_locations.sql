-- +goose Up
CREATE TABLE locations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);


CREATE TABLE sub_locations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    location_id INTEGER NOT NULL REFERENCES locations(id),
    name TEXT NOT NULL,
    specification TEXT
);


CREATE TABLE areas (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    sub_location_id INTEGER NOT NULL REFERENCES sub_locations(id),
    name TEXT NOT NULL,
    section TEXT,
    can_revisit BOOLEAN NOT NULL,
    has_save_sphere BOOLEAN NOT NULL,
    airship_drop_off BOOLEAN NOT NULL,
    has_compilation_sphere BOOLEAN NOT NULL,
    UNIQUE(sub_location_id, name, section)
);


-- +goose Down
DROP TABLE IF EXISTS areas;
DROP TABLE IF EXISTS sub_locations;
DROP TABLE IF EXISTS locations;
-- +goose Up
CREATE TYPE area_connection_type AS ENUM('both-directions', 'one-direction', 'warp');

CREATE TABLE area_connections (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    connection_type area_connection_type NOT NULL,
    story_only BOOLEAN NOT NULL,
    notes TEXT
);


CREATE TABLE j_area_connection (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    connection_id INTEGER NOT NULL REFERENCES area_connections(id)
);


CREATE TABLE formation_boss_songs (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    song_id INTEGER NOT NULL REFERENCES songs(id),
    celebrate_victory BOOLEAN NOT NULL
);


CREATE TYPE monster_formation_category AS ENUM ('boss-fight', 'on-demand-fight', 'random-encounter', 'static-encounter', 'story-fight', 'tutorial');

CREATE TABLE monster_formations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    formation_location_id INTEGER NOT NULL REFERENCES formation_locations(id),
    category monster_formation_category NOT NULL,
    is_forced_ambush BOOLEAN NOT NULL,
    can_escape BOOLEAN NOT NULL,
    boss_song_id INTEGER REFERENCES formation_boss_songs(id),
    notes TEXT
);


CREATE TABLE j_location_monster_formation (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    formation_location_id INTEGER NOT NULL REFERENCES formation_locations(id),
    monster_formation_id INTEGER NOT NULL REFERENCES monster_formations(id)
);


CREATE TABLE j_monster_formation_monster_amount (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_formation_id INTEGER NOT NULL REFERENCES monster_formations(id),
    monster_amount_id INTEGER NOT NULL REFERENCES monster_amounts(id)
);


CREATE TABLE j_monster_formation_trigger_command (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_formation_id INTEGER NOT NULL REFERENCES monster_formations(id),
    trigger_command_id INTEGER NOT NULL REFERENCES trigger_commands(id)
);



-- +goose Down
DROP TABLE IF EXISTS j_monster_formation_trigger_command;
DROP TABLE IF EXISTS j_monster_formation_monster_amount;
DROP TABLE IF EXISTS j_location_monster_formation;
DROP TABLE IF EXISTS monster_formations;
DROP TYPE IF EXISTS monster_formation_category;
DROP TABLE IF EXISTS formation_boss_songs;
DROP TABLE IF EXISTS j_area_connection;
DROP TABLE IF EXISTS area_connections;
DROP TYPE IF EXISTS area_connection_type;
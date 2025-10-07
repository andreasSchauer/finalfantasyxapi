-- +goose Up
CREATE DOMAIN ost_disc as INTEGER
    CHECK (VALUE IN (1, 2, 3, 4));


CREATE TABLE song_credits (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    composer TEXT,
    arranger TEXT,
    performer TEXT,
    lyricist TEXT,

    UNIQUE(composer, arranger, performer, lyricist)
);


CREATE TYPE music_use_case AS ENUM('blitzball-game', 'blitzball-menu', 'boss-battle-default', 'chocobo', 'game-over', 'main-menu', 'random-encounter-default', 'victory');


CREATE TABLE songs (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    streaming_name TEXT,
    in_game_name TEXT,
    ost_name TEXT,
    translation TEXT,
    streaming_track_number INTEGER,
    music_sphere_id INTEGER,
    ost_disc ost_disc,
    ost_track_number INTEGER,
    duration_in_seconds INTEGER NOT NULL,
    can_loop BOOLEAN NOT NULL,
    special_use_case music_use_case,
    credits_id INTEGER REFERENCES song_credits(id)
);



CREATE TABLE fmvs (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    translation TEXT,
    cutscene_description TEXT NOT NULL,
    song_id INTEGER REFERENCES songs(id),
    area_id INTEGER NOT NULL REFERENCES areas(id)
);



-- +goose Down
DROP TABLE IF EXISTS fmvs;
DROP TABLE IF EXISTS songs;
DROP TYPE IF EXISTS music_use_case;
DROP TABLE IF EXISTS song_credits;
DROP DOMAIN IF EXISTS ost_disc;
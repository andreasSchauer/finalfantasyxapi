-- +goose Up
CREATE TYPE bg_replacement_type AS ENUM ('until-trigger', 'until-zone-change');


CREATE TABLE background_music (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    condition TEXT,
    replaces_encounter_music BOOLEAN NOT NULL
);


CREATE TABLE j_songs_background_music (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    song_id INTEGER NOT NULL REFERENCES songs(id),
    bm_id INTEGER NOT NULL REFERENCES background_music(id),
    area_id INTEGER NOT NULL REFERENCES areas(id)
);


CREATE TABLE cues (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    scene_description TEXT NOT NULL,
    area_id INTEGER REFERENCES areas(id),
    replaces_bg_music bg_replacement_type,
    end_trigger TEXT,
    replaces_encounter_music bool NOT NULL
);


CREATE TABLE j_songs_cues(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    song_id INTEGER NOT NULL REFERENCES songs(id),
    cue_id INTEGER NOT NULL REFERENCES cues(id),
    area_id INTEGER NOT NULL REFERENCES areas(id)
);


CREATE TABLE song_credits (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    composer TEXT,
    arranger TEXT,
    performer TEXT,
    lyricist TEXT,

    UNIQUE(composer, arranger, performer, lyricist)
);


ALTER TABLE songs
ADD COLUMN credits_id INTEGER REFERENCES song_credits(id);



-- +goose Down
ALTER TABLE songs
DROP COLUMN IF EXISTS credits_id;

DROP TABLE IF EXISTS song_credits;
DROP TABLE IF EXISTS j_songs_cues;
DROP TABLE IF EXISTS cues;
DROP TABLE IF EXISTS j_songs_background_music;
DROP TABLE IF EXISTS background_music;
DROP TYPE IF EXISTS bg_replacement_type;
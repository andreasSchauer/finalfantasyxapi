-- +goose Up
CREATE TYPE monster_species AS ENUM ('adamantoise', 'aeon', 'armor', 'basilisk', 'blade', 'behemoth', 'bird', 'bomb', 'cactuar', 'cephalopod', 'chest', 'chimera', 'coeurl', 'defender', 'dinofish', 'doomstone', 'drake', 'eater', 'elemental', 'evil-eye', 'flan', 'fungus', 'gel', 'geo', 'haizhe', 'helm', 'hermit', 'humanoid', 'imp', 'iron-giant', 'larva', 'lupine', 'machina', 'malboro', 'mech', 'mimic', 'ochu', 'ogre', 'phantom', 'piranha', 'plant', 'reptile', 'roc', 'ruminant', 'sacred-beast', 'sahagin', 'sin', 'sinspawn', 'spellspinner', 'spirit-beast', 'tonberry', 'unspecified', 'wasp', 'weapon', 'worm', 'wyrm');


CREATE TYPE ctb_icon_type AS ENUM ('monster', 'boss', 'boss-numbered', 'summon', 'cid');


CREATE TYPE monster_formation_category AS ENUM ('boss-fight', 'on-demand-fight', 'random-encounter', 'static-encounter', 'story-fight', 'tutorial');


CREATE TABLE monsters (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    version INTEGER,
    specification TEXT,
    notes TEXT,
    species monster_species NOT NULL,
    is_story_based BOOLEAN NOT NULL,
    is_repeatable BOOLEAN NOT NULL,
    can_be_captured BOOLEAN NOT NULL,
    area_conquest_location ma_creation_area,
    ctb_icon_type ctb_icon_type NOT NULL,
    has_overdrive BOOLEAN NOT NULL,
    is_underwater BOOLEAN NOT NULL,
    is_zombie BOOLEAN NOT NULL,
    distance distance NOT NULL,
    ap INTEGER NOT NULL,
    ap_overkill INTEGER NOT NULL,
    overkill_damage INTEGER NOT NULL,
    gil INTEGER NOT NULL,
    steal_gil INTEGER,
    doom_countdown uint8,
    poison_rate percentage,
    threaten_chance uint8,
    zanmato_level zanmato_level NOT NULL,
    monster_arena_price INTEGER,
    sensor_text TEXT,
    scan_text TEXT,

    UNIQUE(name, version)
);


CREATE TABLE monster_amounts (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    amount INTEGER NOT NULL
);


CREATE TABLE monster_selections (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL
);


CREATE TABLE formation_boss_songs (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    song_id INTEGER NOT NULL REFERENCES songs(id),
    celebrate_victory BOOLEAN NOT NULL
);


CREATE TABLE formation_data (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    category monster_formation_category NOT NULL,
    is_forced_ambush BOOLEAN NOT NULL,
    can_escape BOOLEAN NOT NULL,
    boss_song_id INTEGER REFERENCES formation_boss_songs(id),
    notes TEXT
);


CREATE TABLE monster_formations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    version INTEGER,
    monster_selection_id INTEGER NOT NULL REFERENCES monster_selections(id),
    formation_data_id INTEGER NOT NULL REFERENCES formation_data(id),
    UNIQUE(version, monster_selection_id, formation_data_id)
);


CREATE TABLE encounter_locations (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    specification TEXT
);


CREATE TABLE formation_trigger_commands (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    trigger_command_id INTEGER NOT NULL REFERENCES trigger_commands(id),
    condition TEXT,
    use_amount INTEGER
);


-- +goose Down
DROP TABLE IF EXISTS formation_trigger_commands;
DROP TABLE IF EXISTS encounter_locations;
DROP TABLE IF EXISTS monster_formations;
DROP TABLE IF EXISTS formation_data;
DROP TABLE IF EXISTS formation_boss_songs;
DROP TABLE IF EXISTS monster_selections;
DROP TABLE IF EXISTS monster_amounts;
DROP TABLE IF EXISTS monsters;
DROP TYPE IF EXISTS monster_formation_category;
DROP TYPE IF EXISTS ctb_icon_type;
DROP TYPE IF EXISTS monster_species;
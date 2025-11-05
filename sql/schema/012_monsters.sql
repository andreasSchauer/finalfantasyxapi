-- +goose Up
CREATE TYPE monster_species AS ENUM ('adamantoise', 'aeon', 'armor', 'basilisk', 'blade', 'behemoth', 'bird', 'bomb', 'cactuar', 'cephalopod', 'chest', 'chimera', 'coeurl', 'defender', 'dinofish', 'doomstone', 'drake', 'eater', 'elemental', 'evil-eye', 'flan', 'fungus', 'gel', 'geo', 'haizhe', 'helm', 'hermit', 'humanoid', 'imp', 'iron-giant', 'larva', 'lupine', 'machina', 'malboro', 'mech', 'mimic', 'ochu', 'ogre', 'phantom', 'piranha', 'plant', 'reptile', 'roc', 'ruminant', 'sacred-beast', 'sahagin', 'sin', 'sinspawn', 'spellspinner', 'spirit-beast', 'tonberry', 'unspecified', 'wasp', 'weapon', 'worm', 'wyrm');


CREATE TYPE ctb_icon_type AS ENUM ('monster', 'boss', 'boss-numbered', 'summon');


CREATE TABLE monsters (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    version INTEGER,
    specification TEXT,
    notes TEXT,
    species monster_species NOT NULL,
    is_story_based BOOLEAN NOT NULL,
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
    sensor_text TEXT NOT NULL,
    scan_text TEXT,

    UNIQUE(name, version)
);


CREATE TABLE monster_amounts (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    amount INTEGER NOT NULL
);



-- +goose Down
DROP TABLE IF EXISTS monster_amounts;
DROP TABLE IF EXISTS monsters;
DROP TYPE IF EXISTS ctb_icon_type;
DROP TYPE IF EXISTS monster_species;
-- +goose Up
CREATE TABLE j_character_class_player_units (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES character_classes(id),
    unit_id INTEGER NOT NULL REFERENCES player_units(id)
);


CREATE TABLE j_characters_base_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    character_id INTEGER NOT NULL REFERENCES characters(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE aeon_equipment (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    celestial_wpn BOOLEAN NOT NULL,
    equip_type equip_type NOT NULL
);


CREATE TABLE j_aeons_weapon_armor (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_id INTEGER NOT NULL REFERENCES aeons(id),
    aeon_equipment_id INTEGER NOT NULL REFERENCES aeon_equipment(id)
);


CREATE TABLE j_aeons_base_stats_a (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_id INTEGER NOT NULL REFERENCES aeons(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE j_aeons_base_stats_b (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_id INTEGER NOT NULL REFERENCES aeons(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE j_aeons_base_stats_x (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_id INTEGER NOT NULL REFERENCES aeons(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id),
    battles INTEGER NOT NULL
);


CREATE TABLE default_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES character_classes(id),
    ability_id INTEGER NOT NULL REFERENCES abilities(id)
);


CREATE TABLE default_overdrive_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES character_classes(id),
    ability_id INTEGER NOT NULL REFERENCES overdrive_abilities(id)
);


ALTER TABLE aeons
ADD COLUMN area_id INTEGER REFERENCES areas(id),
ADD COLUMN accuracy_id INTEGER REFERENCES ability_accuracies(id);


ALTER TABLE characters
ADD COLUMN area_id INTEGER REFERENCES areas(id);


-- +goose Down
ALTER TABLE characters
DROP COLUMN IF EXISTS area_id;

ALTER TABLE aeons
DROP COLUMN IF EXISTS accuracy_id,
DROP COLUMN IF EXISTS area_id;

DROP TABLE IF EXISTS default_overdrive_abilities;
DROP TABLE IF EXISTS default_abilities;
DROP TABLE IF EXISTS j_aeons_weapon_armor;
DROP TABLE IF EXISTS aeon_equipment;
DROP TABLE IF EXISTS j_aeons_base_stats_x;
DROP TABLE IF EXISTS j_aeons_base_stats_b;
DROP TABLE IF EXISTS j_aeons_base_stats_a;
DROP TABLE IF EXISTS j_characters_base_stats;
DROP TABLE IF EXISTS j_character_class_player_units;
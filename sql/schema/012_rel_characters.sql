-- +goose Up
CREATE TABLE j_unit_character_class (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    unit_id INTEGER NOT NULL REFERENCES player_units(id),
    class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


CREATE TABLE j_character_base_stat (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    character_id INTEGER NOT NULL REFERENCES characters(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE aeon_equipments (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    celestial_wpn BOOLEAN NOT NULL,
    equip_type equip_type NOT NULL
);


CREATE TABLE j_aeon_equipment (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_id INTEGER NOT NULL REFERENCES aeons(id),
    aeon_equipment_id INTEGER NOT NULL REFERENCES aeon_equipments(id)
);


CREATE TABLE j_aeon_base_stat (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_id INTEGER NOT NULL REFERENCES aeons(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE default_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES character_classes(id),
    ability_id INTEGER NOT NULL REFERENCES player_abilities(id)
);


CREATE TABLE default_overdrive_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES character_classes(id),
    ability_id INTEGER NOT NULL REFERENCES overdrive_abilities(id)
);


ALTER TABLE aeons
ADD COLUMN area_id INTEGER REFERENCES areas(id);



-- +goose Down
ALTER TABLE aeons
DROP COLUMN IF EXISTS area_id;

DROP TABLE IF EXISTS default_overdrive_abilities;
DROP TABLE IF EXISTS default_abilities;
DROP TABLE IF EXISTS j_aeon_equipment;
DROP TABLE IF EXISTS aeon_equipments;
DROP TABLE IF EXISTS j_aeon_base_stat;
DROP TABLE IF EXISTS j_character_base_stat;
DROP TABLE IF EXISTS j_unit_character_class;
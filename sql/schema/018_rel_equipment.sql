-- +goose Up
ALTER TABLE celestial_weapons
ADD COLUMN character_id INTEGER REFERENCES characters(id),
ADD COLUMN aeon_id INTEGER REFERENCES aeons(id);


ALTER TABLE auto_abilities
ADD COLUMN required_item_amount_id INTEGER REFERENCES item_amounts(id),
ADD COLUMN grad_rcvry_stat_id INTEGER REFERENCES stats(id),
ADD COLUMN on_hit_element_id INTEGER REFERENCES elements(id),
ADD COLUMN added_elem_affinity_id INTEGER REFERENCES elemental_affinities(id),
ADD COLUMN on_hit_status_id INTEGER REFERENCES inflicted_statusses(id),
ADD COLUMN added_property_id INTEGER REFERENCES properties(id),
ADD COLUMN cnvrsn_from_mod_id INTEGER REFERENCES modifiers(id),
ADD COLUMN cnvrsn_to_mod_id INTEGER REFERENCES modifiers(id),
ADD CONSTRAINT aa_fk_stat_id CHECK (grad_rcvry_stat_id <= 2),
ADD CONSTRAINT aa_fk_element_id CHECK (on_hit_element_id <= 4);


CREATE TABLE j_auto_abilities_related_stats(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_auto_abilities_locked_out (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    parent_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    child_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id)
);


CREATE TABLE j_auto_abilities_required_item (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    item_id INTEGER NOT NULL REFERENCES items(id)
);


CREATE TABLE j_auto_abilities_added_statusses (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_auto_abilities_added_status_resists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    status_resist_id INTEGER NOT NULL REFERENCES status_resists(id)
);


CREATE TABLE j_auto_abilities_stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    stat_change_id INTEGER NOT NULL REFERENCES stat_changes(id)
);


CREATE TABLE j_auto_abilities_modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    modifier_change_id INTEGER NOT NULL REFERENCES modifier_changes(id)
);


CREATE TABLE equipment_names (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    character_id INTEGER NOT NULL REFERENCES characters(id),
    name TEXT NOT NULL
);


CREATE TABLE j_equipment_tables_names (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    equipment_table_id INTEGER NOT NULL REFERENCES equipment_tables(id),
    equipment_name_id INTEGER NOT NULL REFERENCES equipment_names(id),
    celestial_weapon_id INTEGER REFERENCES celestial_weapons(id)
);



CREATE TYPE auto_ability_pool AS ENUM ('required', 'one', 'two');

CREATE TABLE j_equipment_tables_ability_pool (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    equipment_table_id INTEGER NOT NULL REFERENCES equipment_tables(id),
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    ability_pool auto_ability_pool NOT NULL
);





-- +goose Down
DROP TABLE IF EXISTS j_equipment_tables_ability_pool;
DROP TYPE IF EXISTS auto_ability_pool;
DROP TABLE IF EXISTS j_equipment_tables_names;
DROP TABLE IF EXISTS equipment_names;
DROP TABLE IF EXISTS j_auto_abilities_modifier_changes;
DROP TABLE IF EXISTS j_auto_abilities_stat_changes;
DROP TABLE IF EXISTS j_auto_abilities_added_statusses;
DROP TABLE IF EXISTS j_auto_abilities_added_status_resists;
DROP TABLE IF EXISTS j_auto_abilities_required_item;
DROP TABLE IF EXISTS j_auto_abilities_locked_out;
DROP TABLE IF EXISTS j_auto_abilities_related_stats


ALTER TABLE auto_abilities
DROP CONSTRAINT IF EXISTS aa_fk_element_id,
DROP CONSTRAINT IF EXISTS aa_fk_stat_id,
DROP COLUMN IF EXISTS cnvrsn_to_mod_id,
DROP COLUMN IF EXISTS cnvrsn_from_mod_id,
DROP COLUMN IF EXISTS added_property_id,
DROP COLUMN IF EXISTS on_hit_status_id,
DROP COLUMN IF EXISTS added_elem_affinity_id,
DROP COLUMN IF EXISTS on_hit_element_id,
DROP COLUMN IF EXISTS grad_rcvry_stat_id,
DROP COLUMN IF EXISTS required_item_amount_id;


ALTER TABLE celestial_weapons
DROP COLUMN IF EXISTS character_id,
DROP COLUMN IF EXISTS aeon_id;
-- +goose Up
CREATE TYPE alteration_type AS ENUM ('change', 'gain', 'loss');
CREATE TYPE equipment_slots_type AS ENUM ('ability-slots', 'attached-abilities');


CREATE TABLE monster_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER UNIQUE NOT NULL REFERENCES monsters(id),
    drop_chance uint8 NOT NULL,
    drop_condition TEXT,
    other_items_condition TEXT,
    steal_common_id INTEGER REFERENCES item_amounts(id),
    steal_rare_id INTEGER REFERENCES item_amounts(id),
    drop_common_id INTEGER REFERENCES item_amounts(id),
    drop_rare_id INTEGER REFERENCES item_amounts(id),
    secondary_drop_common_id INTEGER REFERENCES item_amounts(id),
    secondary_drop_rare_id INTEGER REFERENCES item_amounts(id),
    bribe_id INTEGER REFERENCES item_amounts(id)
);


CREATE TABLE equipment_slots_chances (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    amount equipment_rolls NOT NULL,
    chance percent NOT NULL
);


CREATE TABLE monster_equipment (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER UNIQUE NOT NULL REFERENCES monsters(id),
    drop_chance uint8 NOT NULL,
    power uint8 NOT NULL,
    critical_plus INTEGER NOT NULL
);


CREATE TABLE monster_equipment_slots (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_equipment_id INTEGER NOT NULL REFERENCES monster_equipment(id),
    min_amount equipment_slots NOT NULL,
    max_amount equipment_slots NOT NULL,
    type equipment_slots_type NOT NULL,
    UNIQUE(monster_equipment_id, type)
);


CREATE TABLE equipment_drops (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id),
    is_forced BOOLEAN NOT NULL,
    probability auto_ability_probability,
    type equip_type NOT NULL
);


CREATE TABLE altered_states (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    condition TEXT NOT NULL,
    is_temporary BOOLEAN NOT NULL,
    UNIQUE(monster_id, condition)
);


CREATE TABLE alt_state_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    altered_state_id INTEGER NOT NULL REFERENCES altered_states(id),
    alteration_type alteration_type NOT NULL,
    distance distance,
    UNIQUE(altered_state_id, alteration_type)
);


CREATE TABLE monster_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    is_forced BOOLEAN NOT NULL,
    is_unused BOOLEAN NOT NULL
);


CREATE TABLE j_monsters_properties (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    property_id INTEGER NOT NULL REFERENCES properties(id)
);


CREATE TABLE j_monsters_auto_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id)
);


CREATE TABLE j_monsters_ronso_rages (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    overdrive_id INTEGER NOT NULL REFERENCES overdrives(id)
);


CREATE TABLE j_monsters_base_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE j_monsters_elem_resists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    elem_resist_id INTEGER NOT NULL REFERENCES elemental_resists(id)
);


CREATE TABLE j_monsters_immunities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_monsters_status_resists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    status_resist_id INTEGER NOT NULL REFERENCES status_resists(id)
);


CREATE TABLE j_monsters_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_id INTEGER NOT NULL REFERENCES monsters(id),
    monster_ability_id INTEGER NOT NULL REFERENCES monster_abilities(id)
);


CREATE TABLE j_monster_items_other_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_items_id INTEGER NOT NULL REFERENCES monster_items(id),
    possible_item_id INTEGER NOT NULL REFERENCES possible_items(id)
);


CREATE TABLE j_monster_equipment_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_equipment_id INTEGER NOT NULL REFERENCES monster_equipment(id),
    equipment_drop_id INTEGER NOT NULL REFERENCES equipment_drops(id)
);


CREATE TABLE j_monster_equipment_slots_chances (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_equipment_id INTEGER NOT NULL REFERENCES monster_equipment(id),
    equipment_slots_id INTEGER NOT NULL REFERENCES monster_equipment_slots(id),
    slots_chance_id INTEGER NOT NULL REFERENCES equipment_slots_chances(id)
);


CREATE TABLE j_equipment_drops_characters (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    monster_equipment_id INTEGER NOT NULL REFERENCES monster_equipment(id),
    equipment_drop_id INTEGER NOT NULL REFERENCES equipment_drops(id),
    character_id INTEGER NOT NULL REFERENCES characters(id)
);


CREATE TABLE j_alt_state_changes_properties (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    alt_state_change_id INTEGER NOT NULL REFERENCES alt_state_changes(id),
    property_id INTEGER NOT NULL REFERENCES properties(id)
);


CREATE TABLE j_alt_state_changes_auto_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    alt_state_change_id INTEGER NOT NULL REFERENCES alt_state_changes(id),
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id)
);


CREATE TABLE j_alt_state_changes_base_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    alt_state_change_id INTEGER NOT NULL REFERENCES alt_state_changes(id),
    base_stat_id INTEGER NOT NULL REFERENCES base_stats(id)
);


CREATE TABLE j_alt_state_changes_elem_resists (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    alt_state_change_id INTEGER NOT NULL REFERENCES alt_state_changes(id),
    elem_resist_id INTEGER NOT NULL REFERENCES elemental_resists(id)
);


CREATE TABLE j_alt_state_changes_status_immunities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    alt_state_change_id INTEGER NOT NULL REFERENCES alt_state_changes(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_alt_state_changes_added_statusses (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    alt_state_change_id INTEGER NOT NULL REFERENCES alt_state_changes(id),
    inflicted_status_id INTEGER NOT NULL REFERENCES inflicted_statusses(id)
);



-- +goose Down
DROP TABLE IF EXISTS j_alt_state_changes_added_statusses;
DROP TABLE IF EXISTS j_alt_state_changes_status_immunities;
DROP TABLE IF EXISTS j_alt_state_changes_elem_resists;
DROP TABLE IF EXISTS j_alt_state_changes_base_stats;
DROP TABLE IF EXISTS j_alt_state_changes_auto_abilities;
DROP TABLE IF EXISTS j_alt_state_changes_properties;
DROP TABLE IF EXISTS j_equipment_drops_characters;
DROP TABLE IF EXISTS j_monster_equipment_slots_chances;
DROP TABLE IF EXISTS j_monster_equipment_abilities;
DROP TABLE IF EXISTS j_monster_items_other_items;
DROP TABLE IF EXISTS j_monsters_abilities;
DROP TABLE IF EXISTS j_monsters_status_resists;
DROP TABLE IF EXISTS j_monsters_immunities;
DROP TABLE IF EXISTS j_monsters_elem_resists;
DROP TABLE IF EXISTS j_monsters_base_stats;
DROP TABLE IF EXISTS j_monsters_ronso_rages;
DROP TABLE IF EXISTS j_monsters_auto_abilities;
DROP TABLE IF EXISTS j_monsters_properties;
DROP TABLE IF EXISTS monster_abilities;
DROP TABLE IF EXISTS alt_state_changes;
DROP TABLE IF EXISTS altered_states;
DROP TABLE IF EXISTS equipment_drops;
DROP TABLE IF EXISTS monster_equipment_slots;
DROP TABLE IF EXISTS monster_equipment;
DROP TABLE IF EXISTS equipment_slots_chances;
DROP TABLE IF EXISTS monster_items;
DROP TYPE IF EXISTS equipment_slots_type;
DROP TYPE IF EXISTS alteration_type;
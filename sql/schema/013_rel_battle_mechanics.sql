-- +goose Up
ALTER TABLE stats
ADD COLUMN sphere_id INTEGER REFERENCES items(id);


ALTER TABLE elements
ADD COLUMN opposite_element_id INTEGER REFERENCES elements(id);


ALTER TABLE status_conditions
ADD COLUMN added_elem_resist_id INTEGER REFERENCES elemental_resists(id);


CREATE TABLE od_mode_actions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES characters(id),
    amount INTEGER NOT NULL
);


CREATE TABLE j_overdrive_modes_actions_to_learn (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_mode_id INTEGER NOT NULL REFERENCES overdrive_modes(id),
    action_id INTEGER NOT NULL REFERENCES od_mode_actions(id)
);


CREATE TABLE j_status_conditions_related_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_status_conditions_removed_status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    parent_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    child_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_status_conditions_stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    stat_change_id INTEGER NOT NULL REFERENCES stat_changes(id)
);


CREATE TABLE j_status_conditions_modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    modifier_change_id INTEGER NOT NULL REFERENCES modifier_changes(id)
);


CREATE TABLE j_properties_related_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_properties_removed_status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_properties_stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    stat_change_id INTEGER NOT NULL REFERENCES stat_changes(id)
);


CREATE TABLE j_properties_modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    modifier_change_id INTEGER NOT NULL REFERENCES modifier_changes(id)
);


-- +goose Down
DROP TABLE IF EXISTS j_properties_modifier_changes;
DROP TABLE IF EXISTS j_properties_stat_changes;
DROP TABLE IF EXISTS j_properties_removed_status_conditions;
DROP TABLE IF EXISTS j_properties_related_stats;
DROP TABLE IF EXISTS j_status_conditions_modifier_changes;
DROP TABLE IF EXISTS j_status_conditions_stat_changes;
DROP TABLE IF EXISTS j_status_conditions_removed_status_conditions;
DROP TABLE IF EXISTS j_status_conditions_related_stats;
DROP TABLE IF EXISTS j_overdrive_modes_actions_to_learn;
DROP TABLE IF EXISTS od_mode_actions;


ALTER TABLE status_conditions
DROP COLUMN IF EXISTS added_elem_resist_id;


ALTER TABLE elements
DROP COLUMN IF EXISTS opposite_element_id;


ALTER TABLE stats
DROP COLUMN IF EXISTS sphere_id;
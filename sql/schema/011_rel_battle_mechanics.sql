-- +goose Up
ALTER TABLE stats
ADD COLUMN sphere_id INTEGER REFERENCES items(id);

ALTER TABLE elements
ADD COLUMN opposite_element_id INTEGER REFERENCES elements(id);


CREATE TABLE actions_to_learn (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES characters(id),
    amount INTEGER NOT NULL
);


CREATE TABLE j_od_mode_action (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_mode_id INTEGER NOT NULL REFERENCES overdrive_modes(id),
    action_id INTEGER NOT NULL REFERENCES actions_to_learn(id)
);


CREATE TABLE j_status_condition_stat (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_status_condition_self (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    parent_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    child_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TYPE calculation_type AS ENUM ('added-percentage', 'added-value', 'multiply', 'multiply-highest', 'set-value');

CREATE TABLE stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    stat_id INTEGER NOT NULL REFERENCES stats(id),
    calculation_type calculation_type NOT NULL,
    value REAL NOT NULL
);


CREATE TABLE j_status_condition_stat_change (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    stat_change_id INTEGER NOT NULL REFERENCES stat_changes(id)
);


CREATE TABLE modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    modifier_id INTEGER NOT NULL REFERENCES modifiers(id),
    calculation_type calculation_type NOT NULL,
    value REAL NOT NULL
);


CREATE TABLE j_status_condition_modifier_change (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    modifier_change_id INTEGER NOT NULL REFERENCES modifier_changes(id)
);


CREATE TABLE j_property_stat (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_property_status_condition (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_property_stat_change (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    stat_change_id INTEGER NOT NULL REFERENCES stat_changes(id)
);


CREATE TABLE j_property_modifier_change (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    property_id INTEGER NOT NULL REFERENCES properties(id),
    modifier_change_id INTEGER NOT NULL REFERENCES modifier_changes(id)
);


-- +goose Down
DROP TABLE IF EXISTS j_property_modifier_change;
DROP TABLE IF EXISTS j_property_stat_change;
DROP TABLE IF EXISTS j_property_status_condition;
DROP TABLE IF EXISTS j_property_stat;
DROP TABLE IF EXISTS j_status_condition_modifier_change;
DROP TABLE IF EXISTS modifier_changes;
DROP TABLE IF EXISTS j_status_condition_stat_change;
DROP TABLE IF EXISTS stat_changes;
DROP TYPE IF EXISTS calculation_type;
DROP TABLE IF EXISTS j_status_condition_self;
DROP TABLE IF EXISTS j_status_condition_stat;
DROP TABLE IF EXISTS j_od_mode_action;
DROP TABLE IF EXISTS actions_to_learn;

ALTER TABLE elements
DROP COLUMN IF EXISTS opposite_element_id;

ALTER TABLE stats
DROP COLUMN IF EXISTS sphere_id;
-- +goose Up
CREATE TABLE j_battle_interactions_affected_by (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    CHECK (status_condition_id IN (4, 13, 28))
);


CREATE TABLE j_battle_interactions_inflicted_delay (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    inflicted_delay_id INTEGER NOT NULL REFERENCES inflicted_delays(id)
);


CREATE TABLE j_battle_interactions_inflicted_status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    inflicted_status_id INTEGER NOT NULL REFERENCES inflicted_statusses(id)
);


CREATE TABLE j_battle_interactions_removed_status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id)
);


CREATE TABLE j_battle_interactions_copied_status_conditions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    inflicted_status_id INTEGER NOT NULL REFERENCES inflicted_statusses(id)
);


CREATE TABLE j_battle_interactions_stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    stat_change_id INTEGER NOT NULL REFERENCES stat_changes(id)
);


CREATE TABLE j_battle_interactions_modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id),
    modifier_change_id INTEGER NOT NULL REFERENCES modifier_changes(id)
);


CREATE TABLE j_damages_damage_calc (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    damage_id INTEGER NOT NULL REFERENCES damages(id),
    ability_damage_id INTEGER NOT NULL REFERENCES ability_damages(id)
);




-- +goose Down
DROP TABLE IF EXISTS j_damages_damage_calc;
DROP TABLE IF EXISTS j_battle_interactions_modifier_changes;
DROP TABLE IF EXISTS j_battle_interactions_stat_changes;
DROP TABLE IF EXISTS j_battle_interactions_copied_status_conditions;
DROP TABLE IF EXISTS j_battle_interactions_removed_status_conditions;
DROP TABLE IF EXISTS j_battle_interactions_inflicted_status_conditions;
DROP TABLE IF EXISTS j_battle_interactions_inflicted_delay;
DROP TABLE IF EXISTS j_battle_interactions_affected_by;

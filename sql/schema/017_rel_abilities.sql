-- +goose Up
CREATE TABLE j_abilities_battle_interactions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id)
);


CREATE TABLE j_other_abilities_related_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    other_ability_id INTEGER NOT NULL REFERENCES other_abilities(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_other_abilities_learned_by (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    other_ability_id INTEGER NOT NULL REFERENCES other_abilities(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


CREATE TABLE j_player_abilities_related_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    player_ability_id INTEGER NOT NULL REFERENCES player_abilities(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_player_abilities_learned_by (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    player_ability_id INTEGER NOT NULL REFERENCES player_abilities(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


CREATE TABLE j_overdrive_abilities_related_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_ability_id INTEGER NOT NULL REFERENCES overdrive_abilities(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_trigger_commands_related_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    trigger_command_id INTEGER NOT NULL REFERENCES trigger_commands(id),
    stat_id INTEGER NOT NULL REFERENCES stats(id)
);


CREATE TABLE j_overdrives_overdrive_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_id INTEGER NOT NULL REFERENCES overdrives(id),
    overdrive_ability_id INTEGER NOT NULL REFERENCES overdrive_abilities(id)
);


ALTER TABLE other_abilities
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id),
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id),
ADD COLUMN open_submenu_id INTEGER REFERENCES submenus(id);


ALTER TABLE player_abilities
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id),
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id),
ADD COLUMN open_submenu_id INTEGER REFERENCES submenus(id),
ADD COLUMN standard_grid_char_id INTEGER REFERENCES characters(id),
ADD COLUMN expert_grid_char_id INTEGER REFERENCES characters(id),
ADD COLUMN aeon_learn_item_id INTEGER REFERENCES item_amounts(id);


ALTER TABLE overdrives
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id),
ADD COLUMN od_command_id INTEGER REFERENCES overdrive_commands(id),
ADD COLUMN character_class_id INTEGER REFERENCES character_classes(id);


ALTER TABLE trigger_commands
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id);


-- +goose Down
ALTER TABLE trigger_commands
DROP COLUMN IF EXISTS topmenu_id;


ALTER TABLE overdrives
DROP COLUMN IF EXISTS character_class_id,
DROP COLUMN IF EXISTS od_command_id,
DROP COLUMN IF EXISTS topmenu_id;


ALTER TABLE player_abilities
DROP COLUMN IF EXISTS aeon_learn_item_id,
DROP COLUMN IF EXISTS expert_grid_char_id,
DROP COLUMN IF EXISTS standard_grid_char_id,
DROP COLUMN IF EXISTS open_submenu_id,
DROP COLUMN IF EXISTS submenu_id,
DROP COLUMN IF EXISTS topmenu_id;


ALTER TABLE other_abilities
DROP COLUMN IF EXISTS open_submenu_id,
DROP COLUMN IF EXISTS submenu_id,
DROP COLUMN IF EXISTS topmenu_id;


DROP TABLE IF EXISTS j_overdrives_overdrive_abilities;
DROP TABLE IF EXISTS j_trigger_commands_related_stats;
DROP TABLE IF EXISTS j_overdrive_abilities_related_stats;
DROP TABLE IF EXISTS j_player_abilities_learned_by;
DROP TABLE IF EXISTS j_player_abilities_related_stats;
DROP TABLE IF EXISTS j_other_abilities_learned_by;
DROP TABLE IF EXISTS j_other_abilities_related_stats;
DROP TABLE IF EXISTS j_abilities_battle_interactions;
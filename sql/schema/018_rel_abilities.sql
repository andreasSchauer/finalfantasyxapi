-- +goose Up
ALTER TABLE overdrives
ADD COLUMN od_command_id INTEGER references overdrive_commands(id),
ADD COLUMN character_class_id INTEGER REFERENCES character_classes(id);


CREATE TABLE j_overdrives_overdrive_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_id INTEGER NOT NULL REFERENCES overdrives(id),
    overdrive_ability_id INTEGER NOT NULL REFERENCES overdrive_abilities(id)
);


CREATE TABLE j_abilities_battle_interactions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    battle_interaction_id INTEGER NOT NULL REFERENCES battle_interactions(id)
);


-- +goose Down
DROP TABLE IF EXISTS j_abilities_battle_interactions;
DROP TABLE IF EXISTS j_overdrives_overdrive_abilities;

ALTER TABLE overdrives
DROP COLUMN IF EXISTS character_class_id,
DROP COLUMN IF EXISTS od_command_id;
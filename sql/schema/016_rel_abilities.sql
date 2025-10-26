-- +goose Up
ALTER TABLE overdrives
ADD COLUMN od_command_id INTEGER references overdrive_commands(id),
ADD COLUMN character_class_id INTEGER REFERENCES character_classes(id);


CREATE TABLE j_overdrive_ability (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    overdrive_id INTEGER NOT NULL REFERENCES overdrives(id),
    overdrive_ability_id INTEGER NOT NULL REFERENCES overdrive_abilities(id)
);


-- +goose Down
DROP TABLE IF EXISTS j_overdrive_ability;

ALTER TABLE overdrives
DROP COLUMN IF EXISTS character_class_id,
DROP COLUMN IF EXISTS od_command_id;
-- +goose Up
ALTER TABLE aeon_commands
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id);


ALTER TABLE overdrive_commands
ADD COLUMN character_class_id INTEGER REFERENCES character_classes(id),
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id);


CREATE TABLE j_submenu_character_class (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    submenu_id INTEGER NOT NULL REFERENCES submenus(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


CREATE TABLE j_aeon_command_ability (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_command_id INTEGER NOT NULL REFERENCES aeon_commands(id),
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);




-- +goose DOWN
DROP TABLE IF EXISTS j_aeon_command_ability;
DROP TABLE IF EXISTS j_submenu_character_class;

ALTER TABLE overdrive_commands
DROP COLUMN IF EXISTS submenu_id,
DROP COLUMN IF EXISTS character_class_id;

ALTER TABLE aeon_commands
DROP COLUMN IF EXISTS submenu_id;
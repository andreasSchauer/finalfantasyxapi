-- +goose Up
ALTER TABLE aeon_commands
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id),
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id);


ALTER TABLE submenus
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id);


ALTER TABLE overdrive_commands
ADD COLUMN character_class_id INTEGER REFERENCES character_classes(id),
ADD COLUMN topmenu_id INTEGER REFERENCES topmenus(id),
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id);


CREATE TABLE j_submenus_users (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    submenu_id INTEGER NOT NULL REFERENCES submenus(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


CREATE TABLE j_aeon_commands_possible_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    aeon_command_id INTEGER NOT NULL REFERENCES aeon_commands(id),
    ability_id INTEGER NOT NULL REFERENCES abilities(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);




-- +goose DOWN
DROP TABLE IF EXISTS j_aeon_commands_possible_abilities;
DROP TABLE IF EXISTS j_submenus_users;

ALTER TABLE overdrive_commands
DROP COLUMN IF EXISTS submenu_id,
DROP COLUMN IF EXISTS topmenu_id,
DROP COLUMN IF EXISTS character_class_id;

ALTER TABLE submenus
DROP COLUMN IF EXISTS topmenu_id;

ALTER TABLE aeon_commands
DROP COLUMN IF EXISTS topmenu_id,
DROP COLUMN IF EXISTS submenu_id;
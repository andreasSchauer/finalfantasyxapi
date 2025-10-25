-- +goose Up
CREATE TABLE j_submenu_character_class (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    submenu_id INTEGER NOT NULL REFERENCES submenus(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


ALTER TABLE overdrive_commands
ADD COLUMN character_class_id INTEGER REFERENCES character_classes(id),
ADD COLUMN submenu_id INTEGER REFERENCES submenus(id);


-- +goose DOWN
ALTER TABLE overdrive_commands
DROP COLUMN IF EXISTS submenu_id,
DROP COLUMN IF EXISTS character_class_id;

DROP TABLE IF EXISTS j_submenu_character_class;
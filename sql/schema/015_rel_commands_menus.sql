-- +goose Up
CREATE TABLE j_submenu_character_class (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    submenu_id INTEGER NOT NULL REFERENCES submenus(id),
    character_class_id INTEGER NOT NULL REFERENCES character_classes(id)
);


-- +goose DOWN
DROP TABLE IF EXISTS j_submenu_character_class;
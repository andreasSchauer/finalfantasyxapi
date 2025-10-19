-- +goose Up
CREATE TABLE j_character_class_overdrive (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    class_id INTEGER NOT NULL REFERENCES character_classes(id),
    overdrive_id INTEGER NOT NULL REFERENCES overdrives(id)
);



-- +goose Down
DROP TABLE IF EXISTS j_character_class_overdrive;
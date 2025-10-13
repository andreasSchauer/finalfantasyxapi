-- +goose Up
ALTER TABLE celestial_weapons
ADD COLUMN character_id INTEGER REFERENCES characters(id),
ADD COLUMN aeon_id INTEGER REFERENCES aeons(id);



-- +goose Down
ALTER TABLE celestial_weapons
DROP COLUMN character_id,
DROP COLUMN aeon_id;
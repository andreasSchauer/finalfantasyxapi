-- +goose Up
ALTER TABLE stats
ADD COLUMN sphere_id INTEGER REFERENCES items(id);



-- +goose Down
ALTER TABLE stats
DROP COLUMN IF EXISTS sphere_id;
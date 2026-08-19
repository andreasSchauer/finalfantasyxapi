-- +goose Up
CREATE TYPE translation_direction AS ENUM ('to-al-bhed', 'to-english');


-- +goose Down
DROP TYPE IF EXISTS translation_direction;
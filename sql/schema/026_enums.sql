-- +goose Up
CREATE TYPE battle_start AS ENUM ('normal', 'preemptive','ambush');
CREATE TYPE haste_status AS ENUM ('haste', 'slow', 'auto-haste');
CREATE TYPE translation_direction AS ENUM ('to-al-bhed', 'to-english');
CREATE TYPE turn_order_rng AS ENUM ('best', 'worst', 'median');


-- +goose Down
DROP TYPE IF EXISTS turn_order_rng;
DROP TYPE IF EXISTS translation_direction;
DROP TYPE IF EXISTS haste_status;
DROP TYPE IF EXISTS battle_start;
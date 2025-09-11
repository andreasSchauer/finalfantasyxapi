-- name: CreateAbility :one
INSERT INTO abilities (data_hash, name, version, specification, type)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = abilities.data_hash
RETURNING *;


-- name: CreatePlayerAbility :exec
INSERT INTO player_abilities (data_hash, ability_id, description, effect, submenu, can_use_outside_battle, mp_cost, rank, appears_in_help_bar, can_copycat)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEnemyAbility :exec
INSERT INTO enemy_abilities (data_hash, ability_id, effect, rank, appears_in_help_bar, can_copycat)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveAbility :exec
INSERT INTO overdrive_abilities (data_hash, ability_id)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTriggerCommand :exec
INSERT INTO trigger_commands (data_hash, ability_id, description, effect, rank, appears_in_help_bar, can_copycat)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(data_hash) DO NOTHING;
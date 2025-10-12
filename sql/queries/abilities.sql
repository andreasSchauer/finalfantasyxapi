-- name: CreateAbility :one
INSERT INTO abilities (data_hash, name, version, specification, attributes_id, type)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = abilities.data_hash
RETURNING *;


-- name: CreateAbilityAttributes :one
INSERT INTO ability_attributes (data_hash, rank, appears_in_help_bar, can_copycat)
VALUES ($1, $2, $3, $4)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = ability_attributes.data_hash
RETURNING *;


-- name: CreatePlayerAbility :exec
INSERT INTO player_abilities (data_hash, ability_id, description, effect, topmenu, can_use_outside_battle, mp_cost, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEnemyAbility :exec
INSERT INTO enemy_abilities (data_hash, ability_id, effect)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveAbility :exec
INSERT INTO overdrive_abilities (data_hash, ability_id)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTriggerCommand :exec
INSERT INTO trigger_commands (data_hash, ability_id, description, effect, topmenu, cursor)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveCommand :one
INSERT INTO overdrive_commands (data_hash, name, description, rank, topmenu)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = overdrive_commands.data_hash
RETURNING *;


-- name: CreateOverdrive :one
INSERT INTO overdrives (data_hash, od_command_id, name, version, description, effect, topmenu, attributes_id, unlock_condition, countdown_in_sec, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = overdrives.data_hash
RETURNING *;
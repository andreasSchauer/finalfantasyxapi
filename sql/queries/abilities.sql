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


-- name: CreatePlayerAbility :one
INSERT INTO player_abilities (data_hash, ability_id, description, effect, topmenu, can_use_outside_battle, mp_cost, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = player_abilities.data_hash
RETURNING *;


-- name: CreateEnemyAbility :one
INSERT INTO enemy_abilities (data_hash, ability_id, effect)
VALUES ($1, $2, $3)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = enemy_abilities.data_hash
RETURNING *;


-- name: CreateOverdriveAbility :one
INSERT INTO overdrive_abilities (data_hash, ability_id)
VALUES ($1, $2)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = overdrive_abilities.data_hash
RETURNING *;


-- name: CreateTriggerCommand :one
INSERT INTO trigger_commands (data_hash, ability_id, description, effect, topmenu, cursor)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = trigger_commands.data_hash
RETURNING *;


-- name: CreateOverdrive :one
INSERT INTO overdrives (data_hash, name, version, description, effect, topmenu, attributes_id, unlock_condition, countdown_in_sec, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = overdrives.data_hash
RETURNING *;


-- name: UpdateOverdrive :exec
UPDATE overdrives
SET data_hash = $1,
    od_command_id = $2,
    character_class_id = $3
WHERE id = $4;


-- name: CreateOverdriveAbilityJunction :exec
INSERT INTO j_overdrive_ability (data_hash, overdrive_id, overdrive_ability_id)
VALUES($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
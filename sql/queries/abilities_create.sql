-- name: CreateAbility :one
INSERT INTO abilities (data_hash, name, version, specification, attributes_id, type)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = abilities.data_hash
RETURNING *;


-- name: CreateAbilityAttributes :one
INSERT INTO ability_attributes (data_hash, rank, appears_in_help_bar, can_copycat)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = ability_attributes.data_hash
RETURNING *;


-- name: CreateOtherAbility :one
INSERT INTO other_abilities (data_hash, ability_id, description, effect, cursor)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = other_abilities.data_hash
RETURNING *;


-- name: UpdateOtherAbility :exec
UPDATE other_abilities
SET data_hash = $1,
    topmenu_id = $2,
    submenu_id = $3,
    open_submenu_id = $4
WHERE id = $5;


-- name: CreatePlayerAbility :one
INSERT INTO player_abilities (data_hash, ability_id, description, effect, category, can_use_outside_battle, mp_cost, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = player_abilities.data_hash
RETURNING *;


-- name: UpdatePlayerAbility :exec
UPDATE player_abilities
SET data_hash = $1,
    topmenu_id = $2,
    submenu_id = $3,
    open_submenu_id = $4,
    standard_grid_char_id = $5,
    expert_grid_char_id = $6,
    aeon_learn_item_id = $7
WHERE id = $8;


-- name: CreateotherAbilitiesRelatedStatsJunction :exec
INSERT INTO j_other_abilities_related_stats (data_hash, other_ability_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateotherAbilitiesLearnedByJunction :exec
INSERT INTO j_other_abilities_learned_by (data_hash, other_ability_id, character_class_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePlayerAbilitiesRelatedStatsJunction :exec
INSERT INTO j_player_abilities_related_stats (data_hash, player_ability_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePlayerAbilitiesLearnedByJunction :exec
INSERT INTO j_player_abilities_learned_by (data_hash, player_ability_id, character_class_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEnemyAbility :one
INSERT INTO enemy_abilities (data_hash, ability_id, effect)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = enemy_abilities.data_hash
RETURNING *;


-- name: CreateOverdriveAbility :one
INSERT INTO overdrive_abilities (data_hash, ability_id)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = overdrive_abilities.data_hash
RETURNING *;


-- name: CreateOverdriveAbilitiesRelatedStatsJunction :exec
INSERT INTO j_overdrive_abilities_related_stats (data_hash, overdrive_ability_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTriggerCommand :one
INSERT INTO trigger_commands (data_hash, ability_id, description, effect, cursor)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = trigger_commands.data_hash
RETURNING *;


-- name: UpdateTriggerCommand :exec
UPDATE trigger_commands
SET data_hash = $1,
    topmenu_id = $2
WHERE id = $3;


-- name: CreateTriggerCommandsRelatedStatsJunction :exec
INSERT INTO j_trigger_commands_related_stats (data_hash, trigger_command_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdrive :one
INSERT INTO overdrives (data_hash, name, version, description, effect, attributes_id, unlock_condition, countdown_in_sec, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = overdrives.data_hash
RETURNING *;


-- name: UpdateOverdrive :exec
UPDATE overdrives
SET data_hash = $1,
    topmenu_id = $2,
    od_command_id = $3,
    character_class_id = $4
WHERE id = $5;


-- name: CreateRonsoRage :one
INSERT INTO ronso_rages (data_hash, overdrive_id)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = ronso_rages.data_hash
RETURNING *;


-- name: CreateOverdrivesOverdriveAbilitiesJunction :exec
INSERT INTO j_overdrives_overdrive_abilities (data_hash, overdrive_id, overdrive_ability_id)
VALUES($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAbilitiesBattleInteractionsJunction :exec
INSERT INTO j_abilities_battle_interactions (data_hash, ability_id, battle_interaction_id)
VALUES($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
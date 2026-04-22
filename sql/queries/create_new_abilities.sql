-- name: CreateAbilityBulk :many
INSERT INTO abilities (data_hash, name, version, specification, attributes_id, type)
SELECT 
    unnest(sqlc.arg('data_hash')::text[]), 
    unnest(sqlc.arg('names')::text[]), 
    unnest(sqlc.arg('version')::null_int[]), 
    unnest(sqlc.arg('specification')::null_string[]),
    unnest(sqlc.arg('attributes_id')::null_int[]),
    unnest(sqlc.arg('type')::ability_type[])
ON CONFLICT (data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAbilityAttributesBulk :many
INSERT INTO ability_attributes (data_hash, rank, appears_in_help_bar, can_copycat)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('rank')::null_int[]),
    unnest(sqlc.arg('appears_in_help_bar')::boolean[]),
    unnest(sqlc.arg('can_copycat')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateUnspecifiedAbilityBulk :many
INSERT INTO unspecified_abilities (data_hash, ability_id, description, effect, cursor, topmenu_id, submenu_id, open_submenu_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('cursor')::null_target_type[]),
    unnest(sqlc.arg('topmenu_id')::null_int[]),
    unnest(sqlc.arg('submenu_id')::null_int[]),
    unnest(sqlc.arg('open_submenu_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreatePlayerAbilityBulk :many
INSERT INTO player_abilities (data_hash, ability_id, description, effect, category, can_use_outside_battle, mp_cost, cursor, topmenu_id, submenu_id, open_submenu_id, standard_grid_char_id, exp_grid_char_ids, aeon_learn_item_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('category')::player_ability_category[]),
    unnest(sqlc.arg('can_use_outside_battle')::boolean[]),
    unnest(sqlc.arg('mp_cost')::int[]),
    unnest(sqlc.arg('cursor')::null_target_type[]),
    unnest(sqlc.arg('topmenu_id')::null_int[]),
    unnest(sqlc.arg('submenu_id')::null_int[]),
    unnest(sqlc.arg('open_submenu_id')::null_int[]),
    unnest(sqlc.arg('std_grid_char_id')::null_int[]),
    unnest(sqlc.arg('exp_grid_char_id')::null_int[]),
    unnest(sqlc.arg('aeon_learn_item_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateEnemyAbilityBulk :many
INSERT INTO enemy_abilities (data_hash, ability_id, effect)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('effect')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateOverdriveAbilityBulk :many
INSERT INTO overdrive_abilities (data_hash, ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateTriggerCommandBulk :many
INSERT INTO trigger_commands (data_hash, ability_id, description, effect, cursor, topmenu_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('cursor')::target_type[]),
    unnest(sqlc.arg('topmenu_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateOverdriveBulk :many
INSERT INTO overdrives (data_hash, name, version, description, effect, attributes_id, unlock_condition, countdown_in_sec, cursor, topmenu_id, od_command_id, character_class_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('attributes_id')::int[]),
    unnest(sqlc.arg('unlock_condition')::null_string[]),
    unnest(sqlc.arg('countdown_in_sec')::null_int[]),
    unnest(sqlc.arg('cursor')::null_target_type[]),
    unnest(sqlc.arg('topmenu_id')::null_int[]),
    unnest(sqlc.arg('od_command_id')::null_int[]),
    unnest(sqlc.arg('character_class_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateRonsoRageBulk :many
INSERT INTO ronso_rages (data_hash, overdrive_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('overdrive_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;
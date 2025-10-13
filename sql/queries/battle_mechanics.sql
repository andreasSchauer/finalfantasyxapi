-- name: CreateStat :one
INSERT INTO stats (data_hash, name, effect, min_val, max_val, max_val_2)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = stats.data_hash
RETURNING *;


-- name: UpdateStat :exec
UPDATE stats
SET data_hash = $1,
    name = $2,
    effect = $3,
    min_val = $4,
    max_val = $5,
    max_val_2 = $6,
    sphere_id = $7
WHERE id = $8;


-- name: CreateBaseStat :one
INSERT INTO base_stats (data_hash, stat_id, value)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = base_stats.data_hash
RETURNING *;


-- name: CreateElement :one
INSERT INTO elements (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = elements.data_hash
RETURNING *;


-- name: UpdateElement :exec
UPDATE elements
SET data_hash = $1,
    name = $2,
    opposite_element_id = $3
WHERE id = $4;


-- name: CreateAffinity :exec
INSERT INTO affinities (data_hash, name, damage_factor)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAgilityTier :one
INSERT INTO agility_tiers (data_hash, min_agility, max_agility, tick_speed, monster_min_icv, monster_max_icv, character_max_icv)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = agility_tiers.data_hash
RETURNING *;


-- name: CreateAgilitySubtier :exec
INSERT INTO agility_subtiers (data_hash, agility_tier_id, subtier_min_agility, subtier_max_agility, character_min_icv)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveMode :one
INSERT INTO overdrive_modes (data_hash, name, description, effect, type, fill_rate)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = overdrive_modes.data_hash
RETURNING *;


-- name: CreateODModeAction :one
INSERT INTO od_mode_actions (data_hash, user_id, amount)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = od_mode_actions.data_hash
RETURNING *;


-- name: CreateODModeActionJunction :exec
INSERT INTO j_od_mode_action (data_hash, overdrive_mode_id, action_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusCondition :one
INSERT INTO status_conditions (data_hash, name, effect, nullify_armored)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = status_conditions.data_hash
RETURNING *;


-- name: CreateStatusConditionStatJunction :exec
INSERT INTO j_status_condition_stat (data_hash, status_condition_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionSelfJunction :exec
INSERT INTO j_status_condition_self (data_hash, parent_condition_id, child_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateProperty :one
INSERT INTO properties (data_hash, name, effect, nullify_armored)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = properties.data_hash
RETURNING *;


-- name: CreatePropertyStatJunction :exec
INSERT INTO j_property_stat (data_hash, property_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertyStatusConditionJunction :exec
INSERT INTO j_property_status_condition (data_hash, property_id, status_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateModifier :one
INSERT INTO modifiers (data_hash, name, effect, type, default_value)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = modifiers.data_hash
RETURNING *;


-- name: CreateStatChange :one
INSERT INTO stat_changes (data_hash, stat_id, calculation_type, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = stat_changes.data_hash
RETURNING *;


-- name: CreateStatusConditionStatChangeJunction :exec
INSERT INTO j_status_condition_stat_change (data_hash, status_condition_id, stat_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertyStatChangeJunction :exec
INSERT INTO j_property_stat_change (data_hash, property_id, stat_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateModifierChange :one
INSERT INTO modifier_changes (data_hash, modifier_id, calculation_type, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = modifier_changes.data_hash
RETURNING *;


-- name: CreateStatusConditionModifierChangeJunction :exec
INSERT INTO j_status_condition_modifier_change (data_hash, status_condition_id, modifier_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertyModifierChangeJunction :exec
INSERT INTO j_property_modifier_change (data_hash, property_id, modifier_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
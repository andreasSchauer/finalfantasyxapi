-- name: CreateStat :one
INSERT INTO stats (data_hash, name, effect, min_val, max_val, max_val_2)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = stats.data_hash
RETURNING *;


-- name: UpdateStat :exec
UPDATE stats
SET data_hash = $1,
    sphere_id = $2
WHERE id = $3;


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
    opposite_element_id = $2
WHERE id = $3;


-- name: CreateAffinity :one
INSERT INTO affinities (data_hash, name, damage_factor)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = affinities.data_hash
RETURNING *;


-- name: CreateElementalResist :one
INSERT INTO elemental_resists (data_hash, element_id, affinity_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = elemental_resists.data_hash
RETURNING *;


-- name: CreateAgilityTier :one
INSERT INTO agility_tiers (data_hash, min_agility, max_agility, tick_speed, monster_min_icv, monster_max_icv, character_max_icv)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = agility_tiers.data_hash
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


-- name: GetOverdriveMode :one
SELECT * FROM overdrive_modes WHERE id = $1;


-- name: GetOverdriveModes :many
SELECT * FROM overdrive_modes ORDER BY id;


-- name: GetOverdriveModesByType :many
SELECT * FROM overdrive_modes WHERE type = $1 ORDER BY id;


-- name: CreateODModeAction :one
INSERT INTO od_mode_actions (data_hash, user_id, amount)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = od_mode_actions.data_hash
RETURNING *;


-- name: CreateOverdriveModesActionsToLearnJunction :exec
INSERT INTO j_overdrive_modes_actions_to_learn (data_hash, overdrive_mode_id, action_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: GetOverdriveModeActions :many
SELECT
    om.id AS overdrive_mode_id,
    om.name AS overdrive_mode,
    c.id AS character_id,
    pu.name AS character,
    a.user_id AS user_id,
    a.amount AS amount
FROM od_mode_actions a
LEFT JOIN characters c ON a.user_id = c.id
LEFT JOIN player_units pu ON c.unit_id = pu.id
LEFT JOIN j_overdrive_modes_actions_to_learn j ON j.action_id = a.id
LEFT JOIN overdrive_modes om ON j.overdrive_mode_id = om.id
WHERE om.id = $1
ORDER BY c.id;


-- name: CreateStatusCondition :one
INSERT INTO status_conditions (data_hash, name, effect, visualization, nullify_armored)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = status_conditions.data_hash
RETURNING *;


-- name: UpdateStatusCondition :exec
UPDATE status_conditions
SET data_hash = $1,
    added_elem_resist_id = $2
WHERE id = $3;


-- name: CreateStatusConditionsRelatedStatsJunction :exec
INSERT INTO j_status_conditions_related_stats(data_hash, status_condition_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsRemovedStatusConditionsJunction :exec
INSERT INTO j_status_conditions_removed_status_conditions (data_hash, parent_condition_id, child_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsStatChangesJunction :exec
INSERT INTO j_status_conditions_stat_changes (data_hash, status_condition_id, stat_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsModifierChangesJunction :exec
INSERT INTO j_status_conditions_modifier_changes (data_hash, status_condition_id, modifier_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusResist :one
INSERT INTO status_resists (data_hash, status_condition_id, resistance)
VALUES ( $1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = status_resists.data_hash
RETURNING *;


-- name: CreateProperty :one
INSERT INTO properties (data_hash, name, effect, nullify_armored)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = properties.data_hash
RETURNING *;


-- name: CreatePropertiesRelatedStatsJunction :exec
INSERT INTO j_properties_related_stats(data_hash, property_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertiesRemovedStatusConditionsJunction :exec
INSERT INTO j_properties_removed_status_conditions (data_hash, property_id, status_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertiesStatChangesJunction :exec
INSERT INTO j_properties_stat_changes (data_hash, property_id, stat_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertiesModifierChangesJunction :exec
INSERT INTO j_properties_modifier_changes (data_hash, property_id, modifier_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateModifier :one
INSERT INTO modifiers (data_hash, name, effect, type, default_value)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = modifiers.data_hash
RETURNING *;
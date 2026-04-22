-- name: CreateStatBulk :many
INSERT INTO stats (data_hash, name, effect, min_val, max_val, max_val_2, sphere_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('min_val')::int[]),
    unnest(sqlc.arg('max_val')::int[]),
    unnest(sqlc.arg('max_val_2')::null_int[]),
    unnest(sqlc.arg('sphere_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;



-- name: CreateBaseStatBulk :many
INSERT INTO base_stats (data_hash, stat_id, value)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('stat_id')::int[]),
    unnest(sqlc.arg('value')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateElementBulk :many
INSERT INTO elements (data_hash, name)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: UpdateElementBulk :exec
UPDATE elements AS e
SET opposite_element_id = u.opposite_id
FROM (
    SELECT 
        unnest(sqlc.arg('id')::int[]) AS id,
        unnest(sqlc.arg('opposite_id')::null_int[])
) AS u
WHERE e.id = u.id;


-- name: CreateElementalResistBulk :many
INSERT INTO elemental_resists (data_hash, element_id, affinity)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('element_id')::int[]),
    unnest(sqlc.arg('affinity')::elemental_affinity[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAgilityTierBulk :many
INSERT INTO agility_tiers (data_hash, min_agility, max_agility, tick_speed, monster_min_icv, monster_max_icv, character_max_icv)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('min_agility')::int[]),
    unnest(sqlc.arg('max_agility')::int[]),
    unnest(sqlc.arg('tick_speed')::int[]),
    unnest(sqlc.arg('monster_min_icv')::null_int[]),
    unnest(sqlc.arg('monster_max_icv')::null_int[]),
    unnest(sqlc.arg('character_max_icv')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAgilitySubtierBulk :many
INSERT INTO agility_subtiers (data_hash, agility_tier_id, subtier_min_agility, subtier_max_agility, character_min_icv)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('agility_tier_id')::int[]),
    unnest(sqlc.arg('min_agility')::int[]),
    unnest(sqlc.arg('max_agility')::int[]),
    unnest(sqlc.arg('character_min_icv')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;




-- name: CreateOverdriveModeBulk :many
INSERT INTO overdrive_modes (data_hash, name, description, effect, type, fill_rate)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('type')::overdrive_mode_type[]),
    unnest(sqlc.arg('fill_rate')::null_float[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateODModeActionBulk :many
INSERT INTO od_mode_actions (data_hash, user_id, amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('user_id')::int[]),
    unnest(sqlc.arg('amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;




-- name: CreateStatusConditionBulk :many
INSERT INTO status_conditions (data_hash, name, category, is_permanent, effect, visualization, nullify_armored)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('category')::status_condition_category[]),
    unnest(sqlc.arg('is_permanent')::boolean[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('visualization')::null_string[]),
    unnest(sqlc.arg('nullify_armored')::null_nullify_armored[]),
    unnest(sqlc.arg('added_elem_resist_id')::null_int[]),
    unnest(sqlc.arg('inflicted_delay_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateStatusResistBulk :many
INSERT INTO status_resists (data_hash, status_condition_id, resistance)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('status_condition_id')::int[]),
    unnest(sqlc.arg('resistance')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreatePropertyBulk :many
INSERT INTO properties (data_hash, name, effect, nullify_armored)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('nullify_armored')::null_nullify_armored[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateModifierBulk :many
INSERT INTO modifiers (data_hash, name, effect, category, default_value)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('category')::modifier_category[]),
    unnest(sqlc.arg('default_value')::null_float[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;







-- name: CreateOverdriveModesActionsToLearnJunctionBulk :exec
INSERT INTO j_overdrive_modes_actions_to_learn (data_hash, overdrive_mode_id, action_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('overdrive_mode_id')::int[]),
    unnest(sqlc.arg('action_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsRelatedStatsJunctionBulk :exec
INSERT INTO j_status_conditions_related_stats(data_hash, status_condition_id, stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('status_condition_id')::int[]),
    unnest(sqlc.arg('stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsRemovedStatusConditionsJunctionBulk :exec
INSERT INTO j_status_conditions_removed_status_conditions (data_hash, parent_condition_id, child_condition_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('parent_condition_id')::int[]),
    unnest(sqlc.arg('child_condition_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsStatChangesJunctionBulk :exec
INSERT INTO j_status_conditions_stat_changes (data_hash, status_condition_id, stat_change_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('status_condition_id')::int[]),
    unnest(sqlc.arg('stat_change_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusConditionsModifierChangesJunctionBulk :exec
INSERT INTO j_status_conditions_modifier_changes (data_hash, status_condition_id, modifier_change_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('status_condition_id')::int[]),
    unnest(sqlc.arg('modifier_change_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertiesRelatedStatsJunctionBulk :exec
INSERT INTO j_properties_related_stats(data_hash, property_id, stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('property_id')::int[]),
    unnest(sqlc.arg('stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertiesStatChangesJunctionBulk :exec
INSERT INTO j_properties_stat_changes (data_hash, property_id, stat_change_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('property_id')::int[]),
    unnest(sqlc.arg('stat_change_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreatePropertiesModifierChangesJunctionBulk :exec
INSERT INTO j_properties_modifier_changes (data_hash, property_id, modifier_change_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('property_id')::int[]),
    unnest(sqlc.arg('modifier_change_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;
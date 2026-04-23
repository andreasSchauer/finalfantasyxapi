-- name: CreateCelestialWeaponBulk :many
INSERT INTO celestial_weapons (data_hash, name, key_item_base, formula)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('key_item_base')::key_item_base[]),
    unnest(sqlc.arg('formula')::celestial_formula[]),
    unnest(sqlc.arg('character_id')::null_int[]),
    unnest(sqlc.arg('aeon_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAutoAbilityBulk :many
INSERT INTO auto_abilities (data_hash, name, description, effect, type, category, ability_value, activation_condition, counter)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('description')::null_string[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('type')::equip_type[]),
    unnest(sqlc.arg('category')::auto_ability_category[]),
    unnest(sqlc.arg('ability_value')::null_int[]),
    unnest(sqlc.arg('activation_condition')::aa_activation_condition[]),
    unnest(sqlc.arg('counter')::null_counter_type[]),
    unnest(sqlc.arg('required_item_amount_id')::null_int[]),
    unnest(sqlc.arg('grad_rcvry_stat_id')::null_int[]),
    unnest(sqlc.arg('on_hit_element_id')::null_int[]),
    unnest(sqlc.arg('added_elem_resist_id')::null_int[]),
    unnest(sqlc.arg('on_hit_status_id')::null_int[]),
    unnest(sqlc.arg('added_property_id')::null_int[]),
    unnest(sqlc.arg('cnvrsn_from_mod_id')::null_int[]),
    unnest(sqlc.arg('cnvrsn_to_mod_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateEquipmentTableBulk :many
INSERT INTO equipment_tables (data_hash, type, classification, specific_character_id, version, priority, required_slots)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('type')::equip_type[]),
    unnest(sqlc.arg('classification')::equip_class[]),
    unnest(sqlc.arg('specific_character_id')::null_int[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('priority')::null_int[]),
    unnest(sqlc.arg('required_slots')::null_equipment_slots[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateEquipmentNameBulk :many
INSERT INTO equipment_names (data_hash, character_id, name)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('character_id')::int[]),
    unnest(sqlc.arg('name')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAbilityPoolBulk :many
INSERT INTO ability_pools (data_hash, equipment_table_id, pool_idx, req_amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('equipment_table_id')::int[]),
    unnest(sqlc.arg('pool_idx')::int[]),
    unnest(sqlc.arg('req_amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id, data_hash;




-- name: CreateAutoAbilitiesRelatedStatsJunctionBulk :exec
INSERT INTO j_auto_abilities_related_stats(data_hash, auto_ability_id, stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesLockedOutJunctionBulk :exec
INSERT INTO j_auto_abilities_locked_out (data_hash, parent_ability_id, child_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('parent_ability_id')::int[]),
    unnest(sqlc.arg('child_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesRequiredItemJunctionBulk :exec
INSERT INTO j_auto_abilities_required_item (data_hash, auto_ability_id, item_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('item_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesAddedStatussesJunctionBulk :exec
INSERT INTO j_auto_abilities_added_statusses (data_hash, auto_ability_id, status_condition_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('status_condition_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesAddedStatusResistsJunctionBulk :exec
INSERT INTO j_auto_abilities_added_status_resists (data_hash, auto_ability_id, status_resist_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('status_resist_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesStatChangesJunctionBulk :exec
INSERT INTO j_auto_abilities_stat_changes (data_hash, auto_ability_id, stat_change_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('stat_change_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesModifierChangesJunctionBulk :exec
INSERT INTO j_auto_abilities_modifier_changes (data_hash, auto_ability_id, modifier_change_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('modifier_change_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentTablesNamesJunctionBulk :exec
INSERT INTO j_equipment_tables_names (data_hash, equipment_table_id, equipment_name_id, celestial_weapon_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('equipment_table_id')::int[]),
    unnest(sqlc.arg('equipment_name_id')::int[]),
    unnest(sqlc.arg('celestial_weapon_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulk :exec
INSERT INTO j_equipment_tables_required_auto_abilities (data_hash, equipment_table_id, auto_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('equipment_table_id')::int[]),
    unnest(sqlc.arg('auto_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAbilityPoolsAutoAbilitiesJunctionBulk :exec
INSERT INTO j_ability_pools_auto_abilities (data_hash, ability_pool_id, auto_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_pool_id')::int[]),
    unnest(sqlc.arg('auto_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;
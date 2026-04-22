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
RETURNING id;


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
RETURNING id;


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
RETURNING id;


-- name: CreateEquipmentNameBulk :many
INSERT INTO equipment_names (data_hash, character_id, name)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('character_id')::int[]),
    unnest(sqlc.arg('name')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id;


-- name: CreateAbilityPoolBulk :many
INSERT INTO ability_pools (data_hash, equipment_table_id, pool_idx, req_amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('equipment_table_id')::int[]),
    unnest(sqlc.arg('pool_idx')::int[]),
    unnest(sqlc.arg('req_amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXLUDED.data_hash
RETURNING id;
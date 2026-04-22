-- name: CreatePlayerUnitBulk :many
INSERT INTO player_units (data_hash, name, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('type')::unit_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateCharacterBulk :many
INSERT INTO characters (data_hash, unit_id, is_story_based, weapon_type, armor_type, physical_attack_range, can_fight_underwater)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('unit_id')::int[]),
    unnest(sqlc.arg('is_story_based')::boolean[]),
    unnest(sqlc.arg('weapon_type')::weapon_type[]),
    unnest(sqlc.arg('armor_type')::armor_type[]),
    unnest(sqlc.arg('physical_attack_range')::int[]),
    unnest(sqlc.arg('can_fight_underwater')::boolean[]),
    unnest(sqlc.arg('area_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAeonBulk :many
INSERT INTO aeons (data_hash, unit_id, unlock_condition, is_optional, battles_to_regenerate, phys_atk_damage_constant, phys_atk_range, phys_atk_shatter_rate)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('unit_id')::int[]),
    unnest(sqlc.arg('unlock_condition')::text[]),
    unnest(sqlc.arg('is_optional')::boolean[]),
    unnest(sqlc.arg('battles_to_regenerate')::int[]),
    unnest(sqlc.arg('phys_atk_damage_constant')::null_int[]),
    unnest(sqlc.arg('phys_atk_range')::null_int[]),
    unnest(sqlc.arg('phys_atk_shatter_rate')::null_int[]),
    unnest(sqlc.arg('area_id')::null_int[]),
    unnest(sqlc.arg('accuracy_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateCharacterClassBulk :many
INSERT INTO character_classes (data_hash, name, category)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('category')::character_class_category[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAeonEquipmentBulk :many
INSERT INTO aeon_equipment (data_hash, auto_ability_id, celestial_wpn, equip_type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('celestial_wpn')::boolean[]),
    unnest(sqlc.arg('equip_type')::equip_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;
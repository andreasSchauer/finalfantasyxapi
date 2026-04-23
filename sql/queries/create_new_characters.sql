-- name: CreatePlayerUnitBulk :many
INSERT INTO player_units (data_hash, name, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('type')::unit_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


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
RETURNING id, data_hash;


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
RETURNING id, data_hash;


-- name: CreateCharacterClassBulk :many
INSERT INTO character_classes (data_hash, name, category)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('category')::character_class_category[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAeonEquipmentBulk :many
INSERT INTO aeon_equipment (data_hash, auto_ability_id, celestial_wpn, equip_type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('celestial_wpn')::boolean[]),
    unnest(sqlc.arg('equip_type')::equip_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;






-- name: CreateCharactersBaseStatsJunctionBulk :exec
INSERT INTO j_characters_base_stats (data_hash, character_id, base_stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('character_id')::int[]),
    unnest(sqlc.arg('base_stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateCharacterClassPlayerUnitsJunctionBulk :exec
INSERT INTO j_character_class_player_units (data_hash, class_id, unit_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('class_id')::int[]),
    unnest(sqlc.arg('unit_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateDefaultAbilityJunctionBulk :exec
INSERT INTO j_default_abilities (data_hash, class_id, ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('class_id')::int[]),
    unnest(sqlc.arg('ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateDefaultOverdriveAbilityJunctionBulk :exec
INSERT INTO j_default_overdrive_abilities (data_hash, class_id, ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('class_id')::int[]),
    unnest(sqlc.arg('ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeonsBaseStatAJunctionBulk :exec
INSERT INTO j_aeons_base_stats_a (data_hash, aeon_id, base_stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('aeon_id')::int[]),
    unnest(sqlc.arg('base_stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeonsBaseStatBJunctionBulk :exec
INSERT INTO j_aeons_base_stats_b (data_hash, aeon_id, base_stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('aeon_id')::int[]),
    unnest(sqlc.arg('base_stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeonsBaseStatXJunctionBulk :exec
INSERT INTO j_aeons_base_stats_x (data_hash, aeon_id, base_stat_id, battles)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('aeon_id')::int[]),
    unnest(sqlc.arg('base_stat_id')::int[]),
    unnest(sqlc.arg('battles')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeonsWeaponArmorJunctionBulk :exec
INSERT INTO j_aeons_weapon_armor (data_hash, aeon_id, aeon_equipment_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('aeon_id')::int[]),
    unnest(sqlc.arg('aeon_equipment_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;
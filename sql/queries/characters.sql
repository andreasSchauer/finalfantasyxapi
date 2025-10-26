-- name: CreatePlayerUnit :one
INSERT INTO player_units (data_hash, name, type)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = player_units.data_hash
RETURNING *;


-- name: CreateCharacter :one
INSERT INTO characters (data_hash, unit_id, story_only, weapon_type, armor_type, physical_attack_range, can_fight_underwater)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = characters.data_hash
RETURNING *;


-- name: CreateCharacterBaseStatJunction :exec
INSERT INTO j_character_base_stat (data_hash, character_id, base_stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeon :one
INSERT INTO aeons (data_hash, unit_id, unlock_condition, is_optional, battles_to_regenerate, phys_atk_damage_constant, phys_atk_range, phys_atk_shatter_rate, phys_atk_acc_source, phys_atk_hit_chance, phys_atk_acc_modifier)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = aeons.data_hash
RETURNING *;


-- name: UpdateAeon :exec
UPDATE aeons
SET data_hash = $1,
    area_id = $2
WHERE id = $3;


-- name: CreateCharacterClass :one
INSERT INTO character_classes (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = character_classes.data_hash
RETURNING *;


-- name: CreateUnitsCharClassesJunction :exec
INSERT INTO j_unit_character_class (data_hash, unit_id, class_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateDefaultAbility :exec
INSERT INTO default_abilities (data_hash, class_id, ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateDefaultOverdriveAbility :exec
INSERT INTO default_overdrive_abilities (data_hash, class_id, ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeonEquipment :one
INSERT INTO aeon_equipments (data_hash, auto_ability_id, celestial_wpn, equip_type)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = aeon_equipments.data_hash
RETURNING *;


-- name: CreateAeonBaseStatJunction :exec
INSERT INTO j_aeon_base_stat (data_hash, aeon_id, base_stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeonEquipmentJunction :exec
INSERT INTO j_aeon_equipment (data_hash, aeon_id, aeon_equipment_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
-- name: CreateCelestialWeapon :one
INSERT INTO celestial_weapons (data_hash, name, key_item_base, formula)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = celestial_weapons.data_hash
RETURNING *;


-- name: UpdateCelestialWeapon :exec
UPDATE celestial_weapons
SET data_hash = $1,
    character_id = $2,
    aeon_id = $3
WHERE id = $4;


-- name: CreateAutoAbility :one
INSERT INTO auto_abilities (data_hash, name, description, effect, type, category, ability_value, activation_condition, counter)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = auto_abilities.data_hash
RETURNING *;


-- name: UpdateAutoAbility :exec
UPDATE auto_abilities
SET data_hash = $1,
    grad_rcvry_stat_id = $2,
    on_hit_element_id = $3,
    added_elem_resist_id = $4,
    on_hit_status_id = $5,
    added_property_id = $6,
    cnvrsn_from_mod_id = $7,
    cnvrsn_to_mod_id = $8
WHERE id = $9;


-- name: CreateAutoAbilitiesRelatedStatsJunction :exec
INSERT INTO j_auto_abilities_related_stats(data_hash, auto_ability_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesLockedOutJunction :exec
INSERT INTO j_auto_abilities_locked_out (data_hash, parent_ability_id, child_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesRequiredItemJunction :exec
INSERT INTO j_auto_abilities_required_item (data_hash, auto_ability_id, item_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesAddedStatussesJunction :exec
INSERT INTO j_auto_abilities_added_statusses (data_hash, auto_ability_id, status_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesAddedStatusResistsJunction :exec
INSERT INTO j_auto_abilities_added_status_resists (data_hash, auto_ability_id, status_resist_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesStatChangesJunction :exec
INSERT INTO j_auto_abilities_stat_changes (data_hash, auto_ability_id, stat_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbilitiesModifierChangesJunction :exec
INSERT INTO j_auto_abilities_modifier_changes (data_hash, auto_ability_id, modifier_change_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentTable :one
INSERT INTO equipment_tables (data_hash, type, classification, specific_character_id, version, priority, pool_1_amt, pool_2_amt, empty_slots_amt)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = equipment_tables.data_hash
RETURNING *;


-- name: CreateEquipmentName :one
INSERT INTO equipment_names (data_hash, character_id, name)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = equipment_names.data_hash
RETURNING *;


-- name: CreateEquipmentTablesNamesJunction :exec
INSERT INTO j_equipment_tables_names (data_hash, equipment_table_id, equipment_name_id, celestial_weapon_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentTablesAbilityPoolJunction :exec
INSERT INTO j_equipment_tables_ability_pool (data_hash, equipment_table_id, auto_ability_id, ability_pool)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
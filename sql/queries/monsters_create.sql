-- name: CreateMonster :one
INSERT INTO monsters (data_hash, name, version, specification, notes, species, is_story_based, is_repeatable, can_be_captured, area_conquest_location, ctb_icon_type, has_overdrive, is_underwater, is_zombie, distance, ap, ap_overkill, overkill_damage, gil, steal_gil, doom_countdown, poison_rate, threaten_chance, zanmato_level, monster_arena_price, sensor_text, scan_text)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monsters.data_hash
RETURNING *;



-- name: CreateMonsterAmount :one
INSERT INTO monster_amounts (data_hash, monster_id, amount)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_amounts.data_hash
RETURNING *;


-- name: CreateMonsterItem :one
INSERT INTO monster_items (data_hash, monster_id, drop_chance, drop_condition, other_items_condition, steal_common_id, steal_rare_id, drop_common_id, drop_rare_id, secondary_drop_common_id, secondary_drop_rare_id, bribe_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_items.data_hash
RETURNING *;


-- name: CreateMonsterEquipment :one
INSERT INTO monster_equipment (data_hash, monster_id, drop_chance, power, critical_plus)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_equipment.data_hash
RETURNING *;


-- name: CreateMonsterEquipmentSlots :one
INSERT INTO monster_equipment_slots (data_hash, monster_equipment_id, min_amount, max_amount, type)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_equipment_slots.data_hash
RETURNING *;


-- name: CreateEquipmentSlotsChance :one
INSERT INTO equipment_slots_chances (data_hash, amount, chance)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = equipment_slots_chances.data_hash
RETURNING *;


-- name: CreateEquipmentDrop :one
INSERT INTO equipment_drops (data_hash, auto_ability_id, is_forced, probability, type)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = equipment_drops.data_hash
RETURNING *;


-- name: CreateAlteredState :one
INSERT INTO altered_states (data_hash, monster_id, condition, is_temporary)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = altered_states.data_hash
RETURNING *;


-- name: CreateAltStateChange :one
INSERT INTO alt_state_changes (data_hash, altered_state_id, alteration_type, distance, added_status_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = alt_state_changes.data_hash
RETURNING *;


-- name: CreateMonsterAbility :one
INSERT INTO monster_abilities (data_hash, ability_id, is_forced, is_unused)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_abilities.data_hash
RETURNING *;


-- name: CreateMonstersPropertiesJunction :exec
INSERT INTO j_monsters_properties (data_hash, monster_id, property_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersAutoAbilitiesJunction :exec
INSERT INTO j_monsters_auto_abilities (data_hash, monster_id, auto_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersRonsoRagesJunction :exec
INSERT INTO j_monsters_ronso_rages (data_hash, monster_id, ronso_rage_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersBaseStatsJunction :exec
INSERT INTO j_monsters_base_stats (data_hash, monster_id, base_stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersElemResistsJunction :exec
INSERT INTO j_monsters_elem_resists (data_hash, monster_id, elem_resist_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersImmunitiesJunction :exec
INSERT INTO j_monsters_immunities (data_hash, monster_id, status_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersStatusResistsJunction :exec
INSERT INTO j_monsters_status_resists (data_hash, monster_id, status_resist_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersAbilitiesJunction :exec
INSERT INTO j_monsters_abilities (data_hash, monster_id, monster_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterItemsOtherItemsJunction :exec
INSERT INTO j_monster_items_other_items (data_hash, monster_items_id, possible_item_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterEquipmentAbilitiesJunction :exec
INSERT INTO j_monster_equipment_abilities (data_hash, monster_equipment_id, equipment_drop_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterEquipmentSlotsChancesJunction :exec
INSERT INTO j_monster_equipment_slots_chances(data_hash, monster_equipment_id, equipment_slots_id, slots_chance_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentDropsCharactersJunction :exec
INSERT INTO j_equipment_drops_characters (data_hash, monster_equipment_id, equipment_drop_id, character_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesPropertiesJunction :exec
INSERT INTO j_alt_state_changes_properties (data_hash, alt_state_change_id, property_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesAutoAbilitiesJunction :exec
INSERT INTO j_alt_state_changes_auto_abilities (data_hash, alt_state_change_id, auto_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesBaseStatsJunction :exec
INSERT INTO j_alt_state_changes_base_stats (data_hash, alt_state_change_id, base_stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesElemResistsJunction :exec
INSERT INTO j_alt_state_changes_elem_resists (data_hash, alt_state_change_id, elem_resist_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesStatusImmunitiesJunction :exec
INSERT INTO j_alt_state_changes_status_immunities (data_hash, alt_state_change_id, status_condition_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
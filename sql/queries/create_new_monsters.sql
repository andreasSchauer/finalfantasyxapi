-- name: CreateMonsterBulk :many
INSERT INTO monsters (data_hash, name, version, specification, notes, species, availability, is_repeatable, can_be_captured, area_conquest_location, category, ctb_icon_type, has_overdrive, is_underwater, is_zombie, distance, ap, ap_overkill, overkill_damage, gil, steal_gil, doom_countdown, poison_rate, threaten_chance, zanmato_level, monster_arena_price, sensor_text, scan_text)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('specification')::null_string[]),
    unnest(sqlc.arg('notes')::null_string[]),
    unnest(sqlc.arg('species')::monster_species[]),
    unnest(sqlc.arg('availability')::availability_type[]),
    unnest(sqlc.arg('is_repeatable')::boolean[]),
    unnest(sqlc.arg('can_be_captured')::boolean[]),
    unnest(sqlc.arg('area_conquest_location')::null_ma_creation_area[]),
    unnest(sqlc.arg('category')::monster_category[]),
    unnest(sqlc.arg('ctb_icon_type')::ctb_icon_type[]),
    unnest(sqlc.arg('has_overdrive')::boolean[]),
    unnest(sqlc.arg('is_underwater')::boolean[]),
    unnest(sqlc.arg('is_zombie')::boolean[]),
    unnest(sqlc.arg('distance')::int[]),
    unnest(sqlc.arg('ap')::int[]),
    unnest(sqlc.arg('ap_overkill')::int[]),
    unnest(sqlc.arg('overkill_damage')::int[]),
    unnest(sqlc.arg('gil')::int[]),
    unnest(sqlc.arg('steal_gil')::null_int[]),
    unnest(sqlc.arg('doom_countdown')::null_int[]),
    unnest(sqlc.arg('poison_rate')::null_float[]),
    unnest(sqlc.arg('threaten_chance')::null_int[]),
    unnest(sqlc.arg('zanmato_level')::int[]),
    unnest(sqlc.arg('monster_arena_price')::null_int[]),
    unnest(sqlc.arg('sensor_text')::null_string[]),
    unnest(sqlc.arg('scan_text')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterAmountBulk :many
INSERT INTO monster_amounts (data_hash, monster_id, amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterFormationBulk :many
INSERT INTO monster_formations (data_hash, version, monster_selection_id, formation_data_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('monster_selection_id')::int[]),
    unnest(sqlc.arg('formation_data_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterSelectionBulk :many
INSERT INTO monster_selections (data_hash)
SELECT
    unnest(sqlc.arg('data_hash')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateEncounterAreaBulk :many
INSERT INTO encounter_areas (data_hash, area_id, specification)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('specification')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateFormationBossSongBulk :many
INSERT INTO formation_boss_songs (data_hash, song_id, celebrate_victory)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('song_id')::int[]),
    unnest(sqlc.arg('celebrate_victory')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateFormationDataBulk :many
INSERT INTO formation_data (data_hash, category, availability, is_forced_ambush, can_escape, boss_song_id, notes)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('category')::monster_formation_category[]),
    unnest(sqlc.arg('availability')::availability_type[]),
    unnest(sqlc.arg('is_forced_ambush')::boolean[]),
    unnest(sqlc.arg('can_escape')::boolean[]),
    unnest(sqlc.arg('boss_song_id')::null_int[]),
    unnest(sqlc.arg('notes')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateFormationTriggerCommandBulk :many
INSERT INTO formation_trigger_commands (data_hash, trigger_command_id, condition, use_amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('trigger_command_id')::int[]),
    unnest(sqlc.arg('condition')::null_string[]),
    unnest(sqlc.arg('use_amount')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterItemBulk :many
INSERT INTO monster_items (data_hash, monster_id, drop_chance, drop_condition, other_items_condition, steal_common_id, steal_rare_id, drop_common_id, drop_rare_id, secondary_drop_common_id, secondary_drop_rare_id, bribe_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('drop_chance')::int[]),
    unnest(sqlc.arg('drop_condition')::null_string[]),
    unnest(sqlc.arg('other_items_condition')::null_string[]),
    unnest(sqlc.arg('steal_common_id')::null_int[]),
    unnest(sqlc.arg('steal_rare_id')::null_int[]),
    unnest(sqlc.arg('drop_common_id')::null_int[]),
    unnest(sqlc.arg('drop_rare_id')::null_int[]),
    unnest(sqlc.arg('secondary_drop_common_id')::null_int[]),
    unnest(sqlc.arg('secondary_drop_rare_id')::null_int[]),
    unnest(sqlc.arg('bribe_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterEquipmentBulk :many
INSERT INTO monster_equipment (data_hash, monster_id, drop_chance, power, critical_plus)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('drop_chance')::int[]),
    unnest(sqlc.arg('power')::int[]),
    unnest(sqlc.arg('critical_plus')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterEquipmentSlotsBulk :many
INSERT INTO monster_equipment_slots (data_hash, monster_equipment_id, min_amount, max_amount, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_equipment_id')::int[]),
    unnest(sqlc.arg('min_amount')::int[]),
    unnest(sqlc.arg('max_amount')::int[]),
    unnest(sqlc.arg('type')::equipment_slots_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateEquipmentSlotsChanceBulk :many
INSERT INTO equipment_slots_chances (data_hash, amount, chance)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('amount')::int[]),
    unnest(sqlc.arg('chance')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateEquipmentDropBulk :many
INSERT INTO equipment_drops (data_hash, auto_ability_id, is_forced, probability, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('is_forced')::boolean[]),
    unnest(sqlc.arg('probability')::null_auto_ability_probability[]),
    unnest(sqlc.arg('type')::equip_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAlteredStateBulk :many
INSERT INTO altered_states (data_hash, monster_id, condition, is_temporary)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('condition')::text[]),
    unnest(sqlc.arg('is_temporary')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAltStateChangeBulk :many
INSERT INTO alt_state_changes (data_hash, altered_state_id, alteration_type, distance, added_status_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('altered_state_id')::int[]),
    unnest(sqlc.arg('alteration_type')::alteration_type[]),
    unnest(sqlc.arg('distance')::null_int[]),
    unnest(sqlc.arg('added_status_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMonsterAbilityBulk :many
INSERT INTO monster_abilities (data_hash, ability_id, is_forced, is_unused)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('is_forced')::boolean[]),
    unnest(sqlc.arg('is_unused')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;







-- name: CreateMonsterFormationsEncounterAreasJunctionBulk :exec
INSERT INTO j_monster_formations_encounter_areas (data_hash, monster_formation_id, encounter_area_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_formation_id')::int[]),
    unnest(sqlc.arg('encounter_area_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterSelectionsMonstersJunctionBulk :exec
INSERT INTO j_monster_selections_monsters (data_hash, monster_selection_id, monster_amount_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_selection_id')::int[]),
    unnest(sqlc.arg('monster_amount_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationsTriggerCommandsJunctionBulk :exec
INSERT INTO j_monster_formations_trigger_commands (data_hash, monster_formation_id, trigger_command_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_formation_id')::int[]),
    unnest(sqlc.arg('trigger_command_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateFormationTriggerCommandsUsersJunctionBulk :exec
INSERT INTO j_formation_trigger_commands_users (data_hash, trigger_command_id, character_class_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('trigger_command_id')::int[]),
    unnest(sqlc.arg('character_class_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersPropertiesJunctionBulk :exec
INSERT INTO j_monsters_properties (data_hash, monster_id, property_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('property_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersAutoAbilitiesJunctionBulk :exec
INSERT INTO j_monsters_auto_abilities (data_hash, monster_id, auto_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('auto_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersRonsoRagesJunctionBulk :exec
INSERT INTO j_monsters_ronso_rages (data_hash, monster_id, ronso_rage_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('ronso_rage_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersBaseStatsJunctionBulk :exec
INSERT INTO j_monsters_base_stats (data_hash, monster_id, base_stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('base_stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersElemResistsJunctionBulk :exec
INSERT INTO j_monsters_elem_resists (data_hash, monster_id, elem_resist_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('elem_resist_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersImmunitiesJunctionBulk :exec
INSERT INTO j_monsters_immunities (data_hash, monster_id, status_condition_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('status_condition_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersStatusResistsJunctionBulk :exec
INSERT INTO j_monsters_status_resists (data_hash, monster_id, status_resist_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('status_resist_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonstersAbilitiesJunctionBulk :exec
INSERT INTO j_monsters_abilities (data_hash, monster_id, monster_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('monster_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterItemsOtherItemsJunctionBulk :exec
INSERT INTO j_monster_items_other_items (data_hash, monster_items_id, possible_item_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_items_id')::int[]),
    unnest(sqlc.arg('possible_item_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterEquipmentAbilitiesJunctionBulk :exec
INSERT INTO j_monster_equipment_abilities (data_hash, monster_equipment_id, equipment_drop_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_equipment_id')::int[]),
    unnest(sqlc.arg('equipment_drop_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterEquipmentSlotsChancesJunctionBulk :exec
INSERT INTO j_monster_equipment_slots_chances(data_hash, monster_equipment_id, equipment_slots_id, slots_chance_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_equipment_id')::int[]),
    unnest(sqlc.arg('equipment_slots_id')::int[]),
    unnest(sqlc.arg('slots_chance_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentDropsCharactersJunctionBulk :exec
INSERT INTO j_equipment_drops_characters (data_hash, monster_equipment_id, equipment_drop_id, character_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_equipment_id')::int[]),
    unnest(sqlc.arg('equipment_drop_id')::int[]),
    unnest(sqlc.arg('character_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesPropertiesJunctionBulk :exec
INSERT INTO j_alt_state_changes_properties (data_hash, alt_state_change_id, property_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('alt_state_change_id')::int[]),
    unnest(sqlc.arg('property_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesAutoAbilitiesJunctionBulk :exec
INSERT INTO j_alt_state_changes_auto_abilities (data_hash, alt_state_change_id, auto_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('alt_state_change_id')::int[]),
    unnest(sqlc.arg('auto_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesBaseStatsJunctionBulk :exec
INSERT INTO j_alt_state_changes_base_stats (data_hash, alt_state_change_id, base_stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('alt_state_change_id')::int[]),
    unnest(sqlc.arg('base_stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesElemResistsJunctionBulk :exec
INSERT INTO j_alt_state_changes_elem_resists (data_hash, alt_state_change_id, elem_resist_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('alt_state_change_id')::int[]),
    unnest(sqlc.arg('elem_resist_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAltStateChangesStatusImmunitiesJunctionBulk :exec
INSERT INTO j_alt_state_changes_status_immunities (data_hash, alt_state_change_id, status_condition_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('alt_state_change_id')::int[]),
    unnest(sqlc.arg('status_condition_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;
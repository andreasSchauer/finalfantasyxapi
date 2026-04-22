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
RETURNING id;


-- name: CreateMonsterAmountBulk :many
INSERT INTO monster_amounts (data_hash, monster_id, amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateMonsterFormationBulk :many
INSERT INTO monster_formations (data_hash, version, monster_selection_id, formation_data_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('monster_selection_id')::int[]),
    unnest(sqlc.arg('formation_data_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateMonsterSelectionBulk :many
INSERT INTO monster_selections (data_hash)
SELECT
    unnest(sqlc.arg('data_hash')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateEncounterAreaBulk :many
INSERT INTO encounter_areas (data_hash, area_id, specification)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('specification')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateFormationBossSongBulk :many
INSERT INTO formation_boss_songs (data_hash, song_id, celebrate_victory)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('song_id')::int[]),
    unnest(sqlc.arg('celebrate_victory')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


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
RETURNING id;


-- name: CreateFormationTriggerCommandBulk :many
INSERT INTO formation_trigger_commands (data_hash, trigger_command_id, condition, use_amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('trigger_command_id')::int[]),
    unnest(sqlc.arg('condition')::null_string[]),
    unnest(sqlc.arg('use_amount')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


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
RETURNING id;


-- name: CreateMonsterEquipmentBulk :many
INSERT INTO monster_equipment (data_hash, monster_id, drop_chance, power, critical_plus)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('drop_chance')::int[]),
    unnest(sqlc.arg('power')::int[]),
    unnest(sqlc.arg('critical_plus')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateMonsterEquipmentSlotsBulk :many
INSERT INTO monster_equipment_slots (data_hash, monster_equipment_id, min_amount, max_amount, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_equipment_id')::int[]),
    unnest(sqlc.arg('min_amount')::int[]),
    unnest(sqlc.arg('max_amount')::int[]),
    unnest(sqlc.arg('type')::equipment_slots_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateEquipmentSlotsChanceBulk :many
INSERT INTO equipment_slots_chances (data_hash, amount, chance)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('amount')::int[]),
    unnest(sqlc.arg('chance')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateEquipmentDropBulk :many
INSERT INTO equipment_drops (data_hash, auto_ability_id, is_forced, probability, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('auto_ability_id')::int[]),
    unnest(sqlc.arg('is_forced')::boolean[]),
    unnest(sqlc.arg('probability')::null_auto_ability_probability[]),
    unnest(sqlc.arg('type')::equip_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAlteredStateBulk :many
INSERT INTO altered_states (data_hash, monster_id, condition, is_temporary)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('monster_id')::int[]),
    unnest(sqlc.arg('condition')::text[]),
    unnest(sqlc.arg('is_temporary')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateAltStateChangeBulk :many
INSERT INTO alt_state_changes (data_hash, altered_state_id, alteration_type, distance, added_status_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('altered_state_id')::int[]),
    unnest(sqlc.arg('alteration_type')::alteration_type[]),
    unnest(sqlc.arg('distance')::null_int[]),
    unnest(sqlc.arg('added_status_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateMonsterAbilityBulk :many
INSERT INTO monster_abilities (data_hash, ability_id, is_forced, is_unused)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('is_forced')::boolean[]),
    unnest(sqlc.arg('is_unused')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;
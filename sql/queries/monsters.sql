-- name: CreateMonster :one
INSERT INTO monsters (data_hash, name, version, specification, notes, species, is_story_based, is_repeatable, can_be_captured, area_conquest_location, ctb_icon_type, has_overdrive, is_underwater, is_zombie, distance, ap, ap_overkill, overkill_damage, gil, steal_gil, doom_countdown, poison_rate, threaten_chance, zanmato_level, monster_arena_price, sensor_text, scan_text)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monsters.data_hash
RETURNING *;


-- name: GetMonster :one
SELECT * FROM monsters WHERE id = $1;


-- name: GetMonstersByName :many
SELECT * FROM monsters WHERE name = $1;


-- name: GetMonsterItems :one
SELECT
    mi.*,

    i1.id AS steal_common_item_id,
    mi1.name AS steal_common_item,
    mi1.type AS steal_common_item_type,
    ia1.amount AS steal_common_amount,

    i2.id AS steal_rare_item_id,
    mi2.name AS steal_rare_item,
    mi2.type AS steal_rare_item_type,
    ia2.amount AS steal_rare_amount,

    i3.id AS drop_common_item_id,
    mi3.name AS drop_common_item,
    mi3.type AS drop_common_item_type,
    ia3.amount AS drop_common_amount,

    i4.id AS drop_rare_item_id,
    mi4.name AS drop_rare_item,
    mi4.type AS drop_rare_item_type,
    ia4.amount AS drop_rare_amount,

    i5.id AS sec_drop_common_item_id,
    mi5.name AS sec_drop_common_item,
    mi5.type AS sec_drop_common_item_type,
    ia5.amount AS sec_drop_common_amount,

    i6.id AS sec_drop_rare_item_id,
    mi6.name AS sec_drop_rare_item,
    mi6.type AS sec_drop_rare_item_type,
    ia6.amount AS sec_drop_rare_amount,

    i7.id AS bribe_item_id,
    mi7.name AS bribe_item,
    mi7.type AS bribe_item_type,
    ia7.amount AS bribe_amount

FROM monster_items mi

LEFT JOIN item_amounts ia1 ON mi.steal_common_id = ia1.id
LEFT JOIN master_items mi1 ON ia1.master_item_id = mi1.id
LEFT JOIN items i1 ON i1.master_item_id = mi1.id

LEFT JOIN item_amounts ia2 ON mi.steal_rare_id = ia2.id
LEFT JOIN master_items mi2 ON ia2.master_item_id = mi2.id
LEFT JOIN items i2 ON i2.master_item_id = mi2.id

LEFT JOIN item_amounts ia3 ON mi.drop_common_id = ia3.id
LEFT JOIN master_items mi3 ON ia3.master_item_id = mi3.id
LEFT JOIN items i3 ON i3.master_item_id = mi3.id

LEFT JOIN item_amounts ia4 ON mi.drop_rare_id = ia4.id
LEFT JOIN master_items mi4 ON ia4.master_item_id = mi4.id
LEFT JOIN items i4 ON i4.master_item_id = mi4.id

LEFT JOIN item_amounts ia5 ON mi.secondary_drop_common_id = ia5.id
LEFT JOIN master_items mi5 ON ia5.master_item_id = mi5.id
LEFT JOIN items i5 ON i5.master_item_id = mi5.id

LEFT JOIN item_amounts ia6 ON mi.secondary_drop_rare_id = ia6.id
LEFT JOIN master_items mi6 ON ia6.master_item_id = mi6.id
LEFT JOIN items i6 ON i6.master_item_id = mi6.id

LEFT JOIN item_amounts ia7 ON mi.bribe_id = ia7.id
LEFT JOIN master_items mi7 ON ia7.master_item_id = mi7.id
LEFT JOIN items i7 ON i7.master_item_id = mi7.id

WHERE mi.monster_id = $1;


-- name: GetMonsterOtherItems :many
SELECT
    i.id AS item_id,
    mi.name AS item,
    mi.type AS item_type,
    ia.amount AS amount,
    pi.chance AS chance
FROM j_monster_items_other_items j
LEFT JOIN possible_items pi ON j.possible_item_id = pi.id
LEFT JOIN item_amounts ia ON pi.item_amount_id = ia.id
LEFT JOIN master_items mi ON ia.master_item_id = mi.id
LEFT JOIN items i ON i.master_item_id = mi.id
WHERE j.monster_items_id = $1
ORDER BY chance DESC;


-- name: GetMonsterEquipment :one
SELECT * FROM monster_equipment WHERE monster_id = $1;


-- name: GetMonsterEquipmentSlots :many
SELECT * FROM monster_equipment_slots
WHERE monster_equipment_id = $1
ORDER BY id;


-- name: GetMonsterEquipmentSlotsChances :many
SELECT
    esc.amount AS amount,
    esc.chance AS chance
FROM equipment_slots_chances esc
LEFT JOIN j_monster_equipment_slots_chances j ON j.slots_chance_id = esc.id
WHERE j.monster_equipment_id = $1
AND j.equipment_slots_id = $2;


-- name: GetMonsterEquipmentAbilities :many
SELECT
    ed.id AS id,
    aa.name AS auto_ability,
    aa.id AS auto_ability_id,
    ed.is_forced AS is_forced,
    ed.probability AS probability
FROM j_monster_equipment_abilities j
LEFT JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
LEFT JOIN auto_abilities aa ON ed.auto_ability_id = aa.id
WHERE j.monster_equipment_id = $1
AND ed.type = $2;


-- name: GetEquipmentDropCharacters :many
SELECT
    c.id AS character_id,
    pu.name AS character_name
FROM j_equipment_drops_characters j
LEFT JOIN characters c ON j.character_id = c.id
LEFT JOIN player_units pu ON c.unit_id = pu.id
WHERE j.monster_equipment_id = $1
AND j.equipment_drop_id = $2;


-- name: GetMonsterProperties :many
SELECT
    p.id AS property_id,
    p.name AS property
FROM j_monsters_properties j
LEFT JOIN properties p ON j.property_id = p.id
WHERE j.monster_id = $1
ORDER BY p.id;


-- name: GetMonsterAutoAbilities :many
SELECT
    a.id AS auto_ability_id,
    a.name AS auto_ability
FROM j_monsters_auto_abilities j
LEFT JOIN auto_abilities a ON j.auto_ability_id = a.id
WHERE j.monster_id = $1
ORDER BY a.id;


-- name: GetMonsterRonsoRages :many
SELECT
    o.id AS ronso_rage_id,
    o.name AS ronso_rage
FROM j_monsters_ronso_rages j
LEFT JOIN overdrives o ON j.overdrive_id = o.id
WHERE j.monster_id = $1
ORDER BY o.id;


-- name: GetMonsterBaseStats :many
SELECT
    s.id AS stat_id,
    s.name AS stat,
    bs.value AS value
FROM j_monsters_base_stats j
LEFT JOIN base_stats bs ON j.base_stat_id = bs.id
LEFT JOIN stats s ON bs.stat_id = s.id
WHERE j.monster_id = $1
ORDER BY s.id;


-- name: GetMonsterElemResists :many
SELECT
    e.id AS element_id,
    e.name AS element,
    a.id AS affinity_id,
    a.name AS affinity
FROM j_monsters_elem_resists j
LEFT JOIN elemental_resists er ON j.elem_resist_id = er.id
LEFT JOIN elements e ON er.element_id = e.id
LEFT JOIN affinities a ON er.affinity_id = a.id
WHERE j.monster_id = $1
ORDER BY e.id;


-- name: GetMonsterStatusResists :many
SELECT
    sc.id AS status_id,
    sc.name AS status,
    sr.resistance AS resistance
FROM j_monsters_status_resists j
LEFT JOIN status_resists sr ON j.status_resist_id = sr.id
LEFT JOIN status_conditions sc ON sr.status_condition_id = sc.id
WHERE j.monster_id = $1
ORDER BY sc.id;


-- name: GetMonsterImmunities :many
SELECT
    sc.id AS status_id,
    sc.name AS status
FROM j_monsters_immunities j
LEFT JOIN status_conditions sc ON j.status_condition_id = sc.id
WHERE j.monster_id = $1
ORDER BY sc.id;


-- name: GetMonsterAbilities :many
SELECT 
    a.id AS ability_id,
    a.name AS ability,
    a.version AS version,
    a.specification AS specification,
    a.type AS ability_type,
    ma.is_forced AS is_forced,
    ma.is_unused AS is_unused
FROM j_monsters_abilities j
LEFT JOIN monster_abilities ma ON j.monster_ability_id = ma.id
LEFT JOIN abilities a ON ma.ability_id = a.id
WHERE j.monster_id = $1
ORDER BY a.id;


-- name: GetMonsterLocations :many
SELECT DISTINCT
    l.id AS location_id,
	l.name AS location,
    s.id AS sublocation_id,
	s.name AS sublocation,
    a.id AS area_id,
	a.name AS area,
	a.version
FROM locations l
LEFT JOIN sub_locations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sub_location_id = s.id
LEFT JOIN encounter_locations el ON el.area_id = a.id
LEFT JOIN j_encounter_location_formations jelf ON jelf.encounter_location_id = el.id
LEFT JOIN monster_formations mf ON jelf.monster_formation_id = mf.id
LEFT JOIN j_monster_formations_monsters jmfm ON jmfm.monster_formation_id = mf.id
LEFT JOIN monster_amounts ma ON jmfm.monster_amount_id = ma.id
WHERE ma.monster_id = $1;


-- name: GetMonsterMonsterFormations :many
SELECT
    mf.id,
    mf.category,
    mf.is_forced_ambush,
    mf.can_escape,
    mf.notes
FROM monster_formations mf
LEFT JOIN j_monster_formations_monsters j ON j.monster_formation_id = mf.id 
LEFT JOIN monster_amounts ma ON j.monster_amount_id = ma.id
WHERE ma.monster_id = $1
ORDER BY mf.id;


-- name: GetMonsterAlteredStates :many
SELECT * FROM altered_states WHERE monster_id = $1;


-- name: GetAltStateChanges :many
SELECT * FROM alt_state_changes WHERE altered_state_id = $1;


-- name: GetAltStateProperties :many
SELECT
    p.name AS property,
    p.id AS property_id
FROM j_alt_state_changes_properties j
LEFT JOIN properties p ON j.property_id = p.id
WHERE j.alt_state_change_id = $1
ORDER BY p.id;


-- name: GetAltStateAutoAbilities :many
SELECT
    a.name AS auto_ability,
    a.id AS auto_ability_id
FROM j_alt_state_changes_auto_abilities j
LEFT JOIN auto_abilities a ON j.auto_ability_id = a.id
WHERE j.alt_state_change_id = $1
ORDER BY a.id;


-- name: GetAltStateBaseStats :many
SELECT
    s.name AS stat,
    s.id AS stat_id,
    bs.value AS value
FROM j_alt_state_changes_base_stats j
LEFT JOIN base_stats bs ON j.base_stat_id = bs.id
LEFT JOIN stats s ON bs.stat_id = s.id
WHERE j.alt_state_change_id = $1
ORDER BY s.id;


-- name: GetAltStateElemResists :many
SELECT
    e.id AS element_id,
    e.name AS element,
    a.id AS affinity_id,
    a.name AS affinity
FROM j_alt_state_changes_elem_resists j
LEFT JOIN elemental_resists er ON j.elem_resist_id = er.id
LEFT JOIN elements e ON er.element_id = e.id
LEFT JOIN affinities a ON er.affinity_id = a.id
WHERE j.alt_state_change_id = $1
ORDER BY e.id;


-- name: GetAltStateImmunities :many
SELECT
    sc.id AS status_id,
    sc.name AS status
FROM j_alt_state_changes_status_immunities j
LEFT JOIN status_conditions sc ON j.status_condition_id = sc.id
WHERE j.alt_state_change_id = $1
ORDER BY sc.id;


-- name: GetAltStateStatusses :one
SELECT
    sc.id AS status_id,
    sc.name AS status,
    isc.probability AS probability,
    isc.duration_type AS duration_type,
    isc.amount AS amount
FROM alt_state_changes astc
LEFT JOIN inflicted_statusses isc ON astc.added_status_id = isc.id
LEFT JOIN status_conditions sc ON isc.status_condition_id = sc.id
WHERE astc.id = $1
ORDER BY sc.id;


-- name: GetMonsters :many
SELECT * FROM monsters ORDER BY id;


-- name: GetMonstersByElemResistIDs :many
SELECT m.*
FROM monsters m
JOIN j_monsters_elem_resists jmer ON jmer.monster_id = m.id
WHERE jmer.elem_resist_id = ANY(sqlc.arg(elem_resist_ids)::int[])
GROUP BY m.id
HAVING COUNT(DISTINCT jmer.elem_resist_id)
       = array_length(sqlc.arg(elem_resist_ids)::int[], 1)
ORDER BY m.id;


-- name: GetMonstersByStatusResists :many
WITH wanted_statuses AS (
    SELECT unnest(sqlc.arg(status_condition_ids)::int[]) AS status_condition_id
),
monster_status_match AS (
    SELECT
        m.id                           AS monster_id,
        ws.status_condition_id         AS status_condition_id
    FROM monsters m
    JOIN wanted_statuses ws ON TRUE
    LEFT JOIN j_monsters_immunities jmi
        ON jmi.monster_id = m.id
       AND jmi.status_condition_id = ws.status_condition_id
    LEFT JOIN j_monsters_status_resists jmsr
        ON jmsr.monster_id = m.id
    LEFT JOIN status_resists sr
        ON sr.id = jmsr.status_resist_id
       AND sr.status_condition_id = ws.status_condition_id
    WHERE
        jmi.status_condition_id IS NOT NULL
        OR (sr.status_condition_id IS NOT NULL AND sr.resistance >= sqlc.arg(min_resistance))
)
SELECT m.*
FROM monsters m
JOIN monster_status_match msm ON msm.monster_id = m.id
GROUP BY m.id
HAVING COUNT(DISTINCT msm.status_condition_id)
       = array_length(sqlc.arg(status_condition_ids)::int[], 1)
ORDER BY m.id;


-- name: GetMonstersByItem :many
SELECT DISTINCT m.*
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
LEFT JOIN j_monster_items_other_items jmio
  ON jmio.monster_items_id = mi.id
LEFT JOIN possible_items pi
  ON pi.id = jmio.possible_item_id
JOIN item_amounts ia
  ON ia.id IN (
      mi.steal_common_id,
      mi.steal_rare_id,
      mi.drop_common_id,
      mi.drop_rare_id,
      mi.secondary_drop_common_id,
      mi.secondary_drop_rare_id,
      mi.bribe_id,
      pi.item_amount_id
  )
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonstersByItemSteal :many
SELECT DISTINCT m.*
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia
  ON ia.id = mi.steal_common_id
  OR ia.id = mi.steal_rare_id
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonstersByItemDrop :many
SELECT DISTINCT m.*
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia
  ON ia.id IN (
      mi.drop_common_id,
      mi.drop_rare_id,
      mi.secondary_drop_common_id,
      mi.secondary_drop_rare_id
  )
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonstersByItemBribe :many
SELECT DISTINCT m.*
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON ia.id = mi.bribe_id
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonstersByItemOther :many
SELECT DISTINCT m.*
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN j_monster_items_other_items jmio
  ON jmio.monster_items_id = mi.id
JOIN possible_items pi
  ON pi.id = jmio.possible_item_id
JOIN item_amounts ia
  ON ia.id = pi.item_amount_id
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonstersByAutoAbilityIDs :many
SELECT m.*
FROM monsters m
LEFT JOIN monster_equipment me ON me.monster_id = m.id
LEFT JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
LEFT JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
WHERE ed.auto_ability_id = ANY(sqlc.arg(auto_ability_ids)::int[])
GROUP BY m.id
HAVING COUNT(DISTINCT ed.auto_ability_id) >= 1
ORDER BY m.id;


-- name: GetMonstersByRonsoRageID :many
SELECT
    m.*
FROM monsters m
LEFT JOIN j_monsters_ronso_rages j ON j.monster_id = m.id
WHERE j.overdrive_id = $1
ORDER BY m.id;


-- name: GetMonstersByDistance :many
SELECT * FROM monsters WHERE distance = $1;


-- name: GetMonstersByIsStoryBased :many
SELECT * FROM monsters WHERE is_story_based = $1;


-- name: GetMonstersByIsRepeatable :many
SELECT * FROM monsters WHERE is_repeatable = $1;


-- name: GetMonstersByCanBeCaptured :many
SELECT * FROM monsters WHERE can_be_captured = $1;


-- name: GetMonstersByHasOverdrive :many
SELECT * FROM monsters WHERE has_overdrive = $1;


-- name: GetMonstersByIsUnderwater :many
SELECT * FROM monsters WHERE is_underwater = $1;


-- name: GetMonstersByIsZombie :many
SELECT * FROM monsters WHERE is_zombie = $1;


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
INSERT INTO j_monsters_ronso_rages (data_hash, monster_id, overdrive_id)
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
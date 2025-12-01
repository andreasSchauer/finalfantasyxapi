-- name: CreateMonster :one
INSERT INTO monsters (data_hash, name, version, specification, notes, species, is_story_based, can_be_captured, area_conquest_location, ctb_icon_type, has_overdrive, is_underwater, is_zombie, distance, ap, ap_overkill, overkill_damage, gil, steal_gil, doom_countdown, poison_rate, threaten_chance, zanmato_level, monster_arena_price, sensor_text, scan_text)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)
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
FROM j_monster_items_other_items jmoi
LEFT JOIN possible_items pi ON jmoi.possible_item_id = pi.id
LEFT JOIN item_amounts ia ON pi.item_amount_id = ia.id
LEFT JOIN master_items mi ON ia.master_item_id = mi.id
LEFT JOIN items i ON i.master_item_id = mi.id
WHERE jmoi.monster_items_id = $1
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
LEFT JOIN j_monster_equipment_slots_chances jmesc ON jmesc.slots_chance_id = esc.id
WHERE jmesc.monster_equipment_id = $1
AND jmesc.equipment_slots_id = $2;


-- name: GetMonsterEquipmentAbilities :many
SELECT
    ed.id AS id,
    aa.name AS auto_ability,
    aa.id AS auto_ability_id,
    ed.is_forced AS is_forced,
    ed.probability AS probability
FROM j_monster_equipment_abilities jmea
LEFT JOIN equipment_drops ed ON jmea.equipment_drop_id = ed.id
LEFT JOIN auto_abilities aa ON ed.auto_ability_id = aa.id
WHERE jmea.monster_equipment_id = $1
AND ed.type = $2;


-- name: GetEquipmentDropCharacters :many
SELECT
    c.id AS character_id,
    pu.name AS character_name
FROM j_equipment_drops_characters jedc
LEFT JOIN characters c ON jedc.character_id = c.id
LEFT JOIN player_units pu ON c.unit_id = pu.id
WHERE jedc.monster_equipment_id = $1
AND jedc.equipment_drop_id = $2;


-- name: GetMonsterProperties :many
SELECT
    p.id AS property_id,
    p.name AS property
FROM j_monsters_properties jmp
LEFT JOIN properties p ON jmp.property_id = p.id
WHERE jmp.monster_id = $1
ORDER BY p.id;


-- name: GetMonsterAutoAbilities :many
SELECT
    a.id AS auto_ability_id,
    a.name AS auto_ability
FROM j_monsters_auto_abilities jma
LEFT JOIN auto_abilities a ON jma.auto_ability_id = a.id
WHERE jma.monster_id = $1
ORDER BY a.id;


-- name: GetMonsterRonsoRages :many
SELECT
    o.id AS ronso_rage_id,
    o.name AS ronso_rage
FROM j_monsters_ronso_rages jmr
LEFT JOIN overdrives o ON jmr.overdrive_id = o.id
WHERE jmr.monster_id = $1
ORDER BY o.id;


-- name: GetMonsterBaseStats :many
SELECT
    s.id AS stat_id,
    s.name AS stat,
    bs.value AS value
FROM j_monsters_base_stats jmbs
LEFT JOIN base_stats bs ON jmbs.base_stat_id = bs.id
LEFT JOIN stats s ON bs.stat_id = s.id
WHERE jmbs.monster_id = $1
ORDER BY s.id;


-- name: GetMonsterElemResists :many
SELECT
    e.id AS element_id,
    e.name AS element,
    a.id AS affinity_id,
    a.name AS affinity
FROM j_monsters_elem_resists jmer
LEFT JOIN elemental_resists er ON jmer.elem_resist_id = er.id
LEFT JOIN elements e ON er.element_id = e.id
LEFT JOIN affinities a ON er.affinity_id = a.id
WHERE jmer.monster_id = $1
ORDER BY e.id;


-- name: GetMonsterStatusResists :many
SELECT
    sc.id AS status_id,
    sc.name AS status,
    sr.resistance AS resistance
FROM j_monsters_status_resists jmsr
LEFT JOIN status_resists sr ON jmsr.status_resist_id = sr.id
LEFT JOIN status_conditions sc ON sr.status_condition_id = sc.id
WHERE jmsr.monster_id = $1
ORDER BY sc.id;


-- name: GetMonsterImmunities :many
SELECT
    sc.id AS status_id,
    sc.name AS status
FROM j_monsters_immunities jmi
LEFT JOIN status_conditions sc ON jmi.status_condition_id = sc.id
WHERE jmi.monster_id = $1
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
FROM j_monsters_abilities jma
LEFT JOIN monster_abilities ma ON jma.monster_ability_id = ma.id
LEFT JOIN abilities a ON ma.ability_id = a.id
WHERE jma.monster_id = $1
ORDER BY a.id;


-- name: GetMonsters :many
SELECT * FROM monsters ORDER BY id;


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
INSERT INTO alt_state_changes (data_hash, altered_state_id, alteration_type, distance)
VALUES ($1, $2, $3, $4)
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


-- name: CreateAltStateChangesAddedStatussesJunction :exec
INSERT INTO j_alt_state_changes_added_statusses (data_hash, alt_state_change_id, inflicted_status_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
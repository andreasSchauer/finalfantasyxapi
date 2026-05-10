-- name: GetMonsterIDsByName :many
SELECT id FROM monsters WHERE name = $1;


-- name: GetMonsterAreaIDs :many
SELECT DISTINCT ea.area_id
FROM encounter_areas ea
JOIN j_monster_formations_encounter_areas j1 ON j1.encounter_area_id = ea.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_selections_monsters j2 ON mf.monster_selection_id = j2.monster_selection_id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
WHERE ma.monster_id = $1
ORDER BY ea.area_id;


-- name: GetMonsterAreaIdPairs :many
SELECT DISTINCT
  ma.monster_id,
  a.id AS area_id
FROM areas a
JOIN encounter_areas ea ON ea.area_id = a.id
JOIN j_monster_formations_encounter_areas j1 ON j1.encounter_area_id = ea.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
WHERE ma.monster_id = ANY(sqlc.arg('monster_ids')::int[])
ORDER BY ma.monster_id, a.id;


-- name: GetMonsterMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_selections_monsters j ON mf.monster_selection_id = j.monster_selection_id
JOIN monster_amounts ma ON j.monster_amount_id = ma.id
WHERE ma.monster_id = $1
ORDER BY mf.id;


-- name: GetMonsterAbilityIDs :many
SELECT DISTINCT ma.ability_id
FROM monster_abilities ma
JOIN j_monsters_abilities j ON j.monster_ability_id = ma.id
WHERE j.monster_id = $1
ORDER BY ma.ability_id;


-- name: GetMonsterIDs :many
SELECT id FROM monsters ORDER BY id;


-- name: GetMonsterIDsByElemResistIDs :many
SELECT j.monster_id
FROM j_monsters_elem_resists j
WHERE j.elem_resist_id = ANY(sqlc.arg(elem_resist_ids)::int[])
GROUP BY j.monster_id
HAVING COUNT(DISTINCT j.elem_resist_id) = cardinality(sqlc.arg(elem_resist_ids)::int[])
ORDER BY j.monster_id;


-- name: GetMonsterIDsByStatusResists :many
WITH wanted AS (
    SELECT sqlc.arg('status_condition_ids')::int[] AS ids
),
all_matches AS (
  SELECT monster_id, status_condition_id
  FROM j_monsters_immunities

  UNION

  SELECT jmsr.monster_id, sr.status_condition_id
  FROM j_monsters_status_resists jmsr
  JOIN status_resists sr ON sr.id = jmsr.status_resist_id
  WHERE sr.resistance >= sqlc.arg('min_resistance')::int
)
SELECT m.monster_id
FROM all_matches m
JOIN wanted w ON m.status_condition_id = ANY(w.ids)
GROUP BY m.monster_id, w.ids
HAVING COUNT(DISTINCT m.status_condition_id) = cardinality(w.ids)
ORDER BY m.monster_id;


-- name: GetMonsterIDsByItem :many
WITH monster_item_amounts AS (
    SELECT mi.monster_id, mi.steal_common_id AS item_amount_id FROM monster_items mi
    UNION ALL SELECT mi.monster_id, mi.steal_rare_id AS item_amount_id FROM monster_items mi
    UNION ALL SELECT mi.monster_id, mi.drop_common_id AS item_amount_id FROM monster_items mi
    UNION ALL SELECT mi.monster_id, mi.drop_rare_id AS item_amount_id FROM monster_items mi
    UNION ALL SELECT mi.monster_id, mi.secondary_drop_common_id AS item_amount_id FROM monster_items mi
    UNION ALL SELECT mi.monster_id, mi.secondary_drop_rare_id AS item_amount_id FROM monster_items mi
    UNION ALL SELECT mi.monster_id, mi.bribe_id AS item_amount_id FROM monster_items mi
    UNION ALL
    SELECT mi.monster_id, pi.item_amount_id
    FROM possible_items pi
    JOIN j_monster_items_other_items jmio ON jmio.possible_item_id = pi.id
    JOIN monster_items mi ON jmio.monster_items_id = mi.id
)
SELECT DISTINCT mia.monster_id
FROM monster_item_amounts mia
JOIN item_amounts ia ON mia.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY mia.monster_id;


-- name: GetMonsterIDsByItemSteal :many
WITH monster_item_amounts_steal AS (
    SELECT mi.monster_id, mi.steal_common_id AS item_amount_id FROM monster_items mi
    UNION ALL
    SELECT mi.monster_id, mi.steal_rare_id AS item_amount_id FROM monster_items mi
)
SELECT DISTINCT mia.monster_id
FROM monster_item_amounts_steal mia
JOIN item_amounts ia ON mia.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY mia.monster_id;


-- name: GetMonsterIDsByItemDrop :many
WITH monster_item_amounts_drop AS (
    SELECT mi.monster_id, mi.drop_common_id AS item_amount_id FROM monster_items mi
    UNION ALL
    SELECT mi.monster_id, mi.drop_rare_id AS item_amount_id FROM monster_items mi
    UNION ALL
    SELECT mi.monster_id, mi.secondary_drop_common_id AS item_amount_id FROM monster_items mi
    UNION ALL
    SELECT mi.monster_id, mi.secondary_drop_rare_id AS item_amount_id FROM monster_items mi
)
SELECT DISTINCT mia.monster_id
FROM monster_item_amounts_drop mia
JOIN item_amounts ia ON mia.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY mia.monster_id;


-- name: GetMonsterIDsByItemBribe :many
SELECT DISTINCT mi.monster_id
FROM monster_items mi
JOIN item_amounts ia ON ia.id = mi.bribe_id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY mi.monster_id;


-- name: GetMonsterIDsByItemOther :many
SELECT DISTINCT mi.monster_id
FROM monster_items mi
JOIN j_monster_items_other_items j ON j.monster_items_id = mi.id
JOIN possible_items pi ON pi.id = j.possible_item_id
JOIN item_amounts ia ON ia.id = pi.item_amount_id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY mi.monster_id;


-- name: GetMonsterIDsByAutoAbility :many
SELECT DISTINCT me.monster_id
FROM monster_equipment me
JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
WHERE ed.auto_ability_id = $1
ORDER BY me.monster_id;


-- name: GetMonsterIDsByEmptySlots :many
SELECT DISTINCT me.monster_id
FROM monster_equipment me
JOIN monster_equipment_slots asl ON asl.monster_equipment_id = me.id AND asl.type = 'ability-slots'
JOIN monster_equipment_slots aa ON aa.monster_equipment_id = me.id AND aa.type = 'attached-abilities'
JOIN j_monster_equipment_slots_chances j ON j.monster_equipment_id = me.id AND j.equipment_slots_id = aa.id
JOIN equipment_slots_chances esc ON j.slots_chance_id = esc.id
WHERE esc.amount = 0 AND (asl.min_amount = ANY(sqlc.arg(slots)::int[]) OR asl.max_amount = ANY(sqlc.arg(slots)::int[]))
ORDER BY me.monster_id;


-- name: GetMonsterIDsByAutoAbilityIsForced :many
SELECT DISTINCT me.monster_id
FROM monster_equipment me
JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
WHERE ed.auto_ability_id = $1 AND ed.is_forced = $2
ORDER BY me.monster_id;


-- name: GetMonsterIDsByRonsoRage :many
SELECT monster_id FROM j_monsters_ronso_rages WHERE ronso_rage_id = $1 ORDER BY monster_id;


-- name: GetMonsterIDsByDistance :many
SELECT id FROM monsters WHERE distance = ANY(sqlc.arg(distances)::int[]);


-- name: GetMonsterIDsBySpecies :many
SELECT id FROM monsters WHERE species = $1;


-- name: GetCaptureMonsterIDsBySpecies :many
SELECT id FROM monsters WHERE species = $1 AND can_be_captured = true;


-- name: GetMonsterIDsByMaCreationArea :many
SELECT id FROM monsters WHERE area_conquest_location = $1;


-- name: GetCaptureMonsterIDsByMaCreationArea :many
SELECT id FROM monsters WHERE area_conquest_location = $1 AND can_be_captured = true;


-- name: GetMonsterIDsByCategory :many
SELECT id FROM monsters WHERE category = ANY(sqlc.narg('category')::monster_category[]);


-- name: GetMonsterIDsByAvailability :many
SELECT id FROM monsters WHERE availability = ANY(sqlc.narg('availability')::availability_type[]) ORDER BY id;


-- name: GetMonsterIDsByIsRepeatable :many
SELECT id FROM monsters WHERE is_repeatable = $1;


-- name: GetMonsterIDsByCanBeCaptured :many
SELECT id FROM monsters WHERE can_be_captured = $1;


-- name: GetMonsterIDsByHasOverdrive :many
SELECT id FROM monsters WHERE has_overdrive = $1;


-- name: GetMonsterIDsByIsUnderwater :many
SELECT id FROM monsters WHERE is_underwater = $1;


-- name: GetCaptureMonsterIDsByIsUnderwater :many
SELECT id FROM monsters WHERE is_underwater = true AND can_be_captured = true;


-- name: GetMonsterIDsByIsZombie :many
SELECT id FROM monsters WHERE is_zombie = $1;








-- name: GetMonsterFormationIDs :many
SELECT id FROM monster_formations ORDER BY id;


-- name: GetMonsterFormationIDsByCategory :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN formation_data fd ON mf.formation_data_id = fd.id
WHERE fd.category = ANY(sqlc.narg('category')::monster_formation_category[])
ORDER BY mf.id;


-- name: GetMonsterFormationIDsByAvailability :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN formation_data fd ON mf.formation_data_id = fd.id
WHERE fd.availability = ANY(sqlc.narg('availability')::availability_type[])
ORDER BY mf.id;


-- name: GetMonsterFormationIDsByRepeatable :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN formation_data fd ON mf.formation_data_id = fd.id
WHERE fd.category = 'random-encounter' OR fd.category = 'on-demand-fight'
ORDER BY mf.id;


-- name: GetMonsterFormationIDsByForcedAmbush :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN formation_data fd ON mf.formation_data_id = fd.id
WHERE fd.is_forced_ambush = $1
ORDER BY mf.id;


-- name: GetMonsterFormationIDsByMonster :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_selections_monsters j ON mf.monster_selection_id = j.monster_selection_id
JOIN monster_amounts ma ON j.monster_amount_id = ma.id
WHERE ma.monster_id = $1
ORDER BY mf.id;


-- name: GetMonsterFormationIDsByArea :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
WHERE ea.area_id = $1
ORDER BY mf.id;


-- name: GetMonsterFormationIDsBySublocation :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY mf.id;


-- name: GetMonsterFormationIDsByLocation :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY mf.id;


-- name: GetMonsterFormationMonsterIDs :many
SELECT DISTINCT ma.monster_id
FROM monster_amounts ma
JOIN j_monster_selections_monsters j ON j.monster_amount_id = ma.id
JOIN monster_formations mf ON j.monster_selection_id = mf.monster_selection_id
WHERE mf.id = $1
ORDER BY ma.monster_id;
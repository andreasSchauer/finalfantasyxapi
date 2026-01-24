
-- name: GetMonsterIDsByName :many
SELECT id FROM monsters WHERE name = $1;


-- name: GetMonsterAreaIDs :many
SELECT DISTINCT a.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations jelf ON jelf.encounter_location_id = el.id
JOIN monster_formations mf ON jelf.monster_formation_id = mf.id
JOIN j_monster_formations_monsters jmfm ON jmfm.monster_formation_id = mf.id
JOIN monster_amounts ma ON jmfm.monster_amount_id = ma.id
WHERE ma.monster_id = $1
ORDER BY a.id;


-- name: GetMonsterMonsterFormationIDs :many
SELECT mf.id
FROM monster_formations mf
JOIN j_monster_formations_monsters j ON j.monster_formation_id = mf.id 
JOIN monster_amounts ma ON j.monster_amount_id = ma.id
WHERE ma.monster_id = $1
ORDER BY mf.id;


-- name: GetMonsterIDs :many
SELECT id FROM monsters ORDER BY id;


-- name: GetMonsterIDsByElemResistIDs :many
SELECT m.id
FROM monsters m
JOIN j_monsters_elem_resists jmer ON jmer.monster_id = m.id
WHERE jmer.elem_resist_id = ANY(sqlc.arg(elem_resist_ids)::int[])
GROUP BY m.id
HAVING COUNT(DISTINCT jmer.elem_resist_id)
       = array_length(sqlc.arg(elem_resist_ids)::int[], 1)
ORDER BY m.id;


-- name: GetMonsterIDsByStatusResists :many
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
SELECT m.id
FROM monsters m
JOIN monster_status_match msm ON msm.monster_id = m.id
GROUP BY m.id
HAVING COUNT(DISTINCT msm.status_condition_id)
       = array_length(sqlc.arg(status_condition_ids)::int[], 1)
ORDER BY m.id;


-- name: GetMonsterIDsByItem :many
SELECT DISTINCT m.id
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


-- name: GetMonsterIDsByItemSteal :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia
  ON ia.id = mi.steal_common_id
  OR ia.id = mi.steal_rare_id
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonsterIDsByItemDrop :many
SELECT DISTINCT m.id
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


-- name: GetMonsterIDsByItemBribe :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON ia.id = mi.bribe_id
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonsterIDsByItemOther :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
JOIN possible_items pi ON pi.id = jmio.possible_item_id
JOIN item_amounts ia ON ia.id = pi.item_amount_id
WHERE ia.master_item_id = sqlc.arg(item_id)
ORDER BY m.id;


-- name: GetMonsterIDsByAutoAbilityIDs :many
SELECT m.id
FROM monsters m
JOIN monster_equipment me ON me.monster_id = m.id
JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
WHERE ed.auto_ability_id = ANY(sqlc.arg(auto_ability_ids)::int[])
GROUP BY m.id
HAVING COUNT(DISTINCT ed.auto_ability_id) >= 1
ORDER BY m.id;


-- name: GetMonsterIDsByRonsoRage :many
SELECT m.id
FROM monsters m
JOIN j_monsters_ronso_rages j ON j.monster_id = m.id
WHERE j.ronso_rage_id = $1
ORDER BY m.id;


-- name: GetMonsterIDsByDistance :many
SELECT id FROM monsters WHERE distance = $1;


-- name: GetMonsterIDsBySpecies :many
SELECT id FROM monsters WHERE species = $1;


-- name: GetMonsterIDsByMaCreationArea :many
SELECT id FROM monsters WHERE area_conquest_location = $1;


-- name: GetMonsterIDsByCTBIconType :many
SELECT id FROM monsters WHERE ctb_icon_type = $1;


-- name: GetMonsterIDsByCTBIconTypeBoss :many
SELECT id FROM monsters WHERE ctb_icon_type = 'boss' OR ctb_icon_type = 'boss-numbered';


-- name: GetMonsterIDsByIsStoryBased :many
SELECT id FROM monsters WHERE is_story_based = $1;


-- name: GetMonsterIDsByIsRepeatable :many
SELECT id FROM monsters WHERE is_repeatable = $1;


-- name: GetMonsterIDsByCanBeCaptured :many
SELECT id FROM monsters WHERE can_be_captured = $1;


-- name: GetMonsterIDsByHasOverdrive :many
SELECT id FROM monsters WHERE has_overdrive = $1;


-- name: GetMonsterIDsByIsUnderwater :many
SELECT id FROM monsters WHERE is_underwater = $1;


-- name: GetMonsterIDsByIsZombie :many
SELECT id FROM monsters WHERE is_zombie = $1;
-- name: GetMonsterIDsByName :many
SELECT id FROM monsters WHERE name = $1;


-- name: GetMonsterAreaIDs :many
SELECT DISTINCT area_id FROM mv_monster_encounters WHERE monster_id = $1 ORDER BY area_id;


-- name: GetMonsterAreaIDsRel :many
WITH w AS (
    SELECT
      sqlc.arg('monster_id')::int AS monster_id,
      sqlc.narg('availability')::availability_type[] AS availability,
      sqlc.narg('repeatable')::boolean AS repeatable
)
SELECT DISTINCT me.area_id
FROM mv_monster_encounters me
CROSS JOIN w
WHERE me.monster_id = w.monster_id
  AND (w.availability IS NULL OR me.avl_area = ANY(w.availability))
  AND (w.repeatable IS NULL OR me.is_repeatable_loc = w.repeatable)
ORDER BY me.area_id;


-- name: GetMonsterAreaIdPairs :many
SELECT DISTINCT
  monster_id,
  area_id
FROM mv_monster_encounters
WHERE monster_id = ANY(sqlc.arg('monster_ids')::int[])
ORDER BY monster_id, area_id;


-- name: GetMonsterMonsterFormationIDs :many
SELECT DISTINCT formation_id FROM mv_monster_encounters WHERE monster_id = $1 ORDER BY formation_id;


-- name: GetMonsterMonsterFormationIDsRel :many
WITH w AS (
    SELECT
      sqlc.arg('monster_id')::int AS monster_id,
      sqlc.narg('availability')::availability_type[] AS availability,
      sqlc.narg('repeatable')::boolean AS repeatable
)
SELECT DISTINCT me.formation_id
FROM mv_monster_encounters me
CROSS JOIN w
WHERE me.monster_id = w.monster_id
  AND (w.availability IS NULL OR me.avl_area = ANY(w.availability))
  AND (w.repeatable IS NULL OR me.is_repeatable_loc = w.repeatable)
ORDER BY me.formation_id;


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
SELECT DISTINCT monster_id FROM mv_monster_item_drops WHERE item_id = $1 ORDER BY monster_id;


-- name: GetMonsterIDsByItemSteal :many
SELECT DISTINCT monster_id
FROM mv_monster_item_drops
WHERE item_id = $1 AND source_type LIKE 'steal%'
ORDER BY monster_id;


-- name: GetMonsterIDsByItemDrop :many
SELECT DISTINCT monster_id
FROM mv_monster_item_drops
WHERE item_id = $1 AND source_type LIKE 'drop%'
ORDER BY monster_id;


-- name: GetMonsterIDsByItemBribe :many
SELECT DISTINCT monster_id
FROM mv_monster_item_drops
WHERE item_id = $1 AND source_type = 'bribe'
ORDER BY monster_id;


-- name: GetMonsterIDsByItemOther :many
SELECT DISTINCT monster_id
FROM mv_monster_item_drops
WHERE item_id = $1 AND source_type = 'other'
ORDER BY monster_id;


-- name: GetMonsterIDsByAutoAbility :many
SELECT DISTINCT monster_id FROM mv_monster_equipment_drops WHERE auto_ability_id = $1 ORDER BY monster_id;


-- name: GetMonsterIDsByEmptySlots :many
SELECT DISTINCT me.monster_id
FROM mv_monster_equipment_drops me
JOIN monster_equipment_slots asl ON me.ability_slots_id = asl.id
JOIN j_monster_equipment_slots_chances j ON j.equipment_slots_id = me.attached_abilities_id
JOIN equipment_slots_chances esc ON j.slots_chance_id = esc.id
WHERE esc.amount = 0
  AND (
        asl.min_amount = ANY(sqlc.arg(slots)::int[])
        OR
        asl.max_amount = ANY(sqlc.arg(slots)::int[])
)
ORDER BY me.monster_id;


-- name: GetMonsterIDsByAutoAbilityIsForced :many
SELECT DISTINCT monster_id
FROM mv_monster_equipment_drops
WHERE auto_ability_id = $1 AND is_forced = $2
ORDER BY monster_id;


-- name: GetMonsterIDsByArea :many
SELECT DISTINCT monster_id FROM mv_monster_encounters WHERE area_id = $1 ORDER BY monster_id;


-- name: GetMonsterIDsBySublocation :many
SELECT DISTINCT me.monster_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY me.monster_id;


-- name: GetMonsterIDsByLocation :many
SELECT DISTINCT me.monster_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.location_id = $1
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








-- name: GetMonsterFormationMonsterIDs :many
SELECT DISTINCT monster_id FROM mv_monster_encounters WHERE formation_id = $1 ORDER BY monster_id;


-- name: GetMonsterFormationIDs :many
SELECT DISTINCT id FROM monster_formations ORDER BY id;


-- name: GetMonsterFormationIDsByCategory :many
SELECT DISTINCT formation_id
FROM mv_monster_encounters
WHERE category = ANY(sqlc.narg('category')::monster_formation_category[])
ORDER BY formation_id;


-- name: GetMonsterFormationIDsByRepeatable :many
SELECT DISTINCT formation_id
FROM mv_monster_encounters
WHERE category = 'random-encounter' OR category = 'on-demand-fight'
ORDER BY formation_id;


-- name: GetMonsterFormationIDsByForcedAmbush :many
SELECT DISTINCT formation_id FROM mv_monster_encounters WHERE is_forced_ambush = $1 ORDER BY formation_id;


-- name: GetMonsterFormationIDsByMonster :many
SELECT DISTINCT formation_id FROM mv_monster_encounters WHERE monster_id = $1 ORDER BY formation_id;


-- name: GetMonsterFormationIDsByArea :many
SELECT DISTINCT formation_id FROM mv_monster_encounters WHERE area_id = $1 ORDER BY formation_id;


-- name: GetMonsterFormationIDsBySublocation :many
SELECT DISTINCT me.formation_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY me.formation_id;


-- name: GetMonsterFormationIDsByLocation :many
SELECT DISTINCT me.formation_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY me.formation_id;
-- name: GetMasterItemObtainableBools :one
WITH w AS (
    SELECT
      sqlc.arg('master_item_id')::int AS master_item_id,
      sqlc.narg('availability')::availability_type[] AS availability,
      sqlc.narg('repeatable')::boolean AS repeatable
)
SELECT 
  EXISTS(
    SELECT 1 FROM mv_item_sources mis
    WHERE mis.master_item_id = w.master_item_id
      AND mis.source_type = 'monster'
      AND (w.availability IS NULL OR mis.avl_self = ANY(w.availability))
      AND (w.repeatable IS NULL OR mis.is_repeatable = w.repeatable)
  ) as monsters,
  EXISTS(
    SELECT 1 FROM mv_item_sources mis
    WHERE mis.master_item_id = w.master_item_id
      AND mis.source_type = 'treasure'
      AND (w.availability IS NULL OR mis.avl_self = ANY(w.availability))
      AND (w.repeatable IS NULL OR mis.is_repeatable = w.repeatable)
  ) as treasures,
  EXISTS(
    SELECT 1 FROM mv_item_sources mis
    WHERE mis.master_item_id = w.master_item_id
      AND mis.source_type = 'shop'
      AND (w.availability IS NULL OR mis.avl_context = ANY(w.availability))
      AND (w.repeatable IS NULL OR mis.is_repeatable = w.repeatable)
  ) as shops,
  EXISTS(
    SELECT 1 FROM mv_item_sources mis
    WHERE mis.master_item_id = w.master_item_id
      AND mis.source_type = 'quest'
      AND (w.availability IS NULL OR mis.avl_self = ANY(w.availability))
      AND (w.repeatable IS NULL OR mis.is_repeatable = w.repeatable)
  ) as quests,
  EXISTS(
    SELECT 1 FROM mv_item_sources mis
    WHERE mis.master_item_id = w.master_item_id
      AND mis.source_type = 'blitzball'
      AND (w.availability IS NULL OR mis.avl_self = ANY(w.availability))
      AND (w.repeatable IS NULL OR mis.is_repeatable = w.repeatable)
  ) as blitzball
FROM w;


-- name: GetMasterItemIDs :many
SELECT id FROM master_items ORDER BY id;


-- name: GetMasterItemIDsByType :many
SELECT id FROM master_items WHERE type = ANY(sqlc.narg('item_type')::item_type[]) ORDER BY id;


-- name: GetMasterItemIDsByMethod :many
SELECT master_item_id
FROM mv_item_sources
WHERE source_type = sqlc.arg('method')::text
ORDER BY master_item_id;


-- name: GetMasterItemIDsByLocation :many
WITH w AS (
  SELECT sqlc.narg('method')::text AS method
)
SELECT DISTINCT mis.master_item_id
FROM mv_item_sources mis
JOIN mv_geography g ON mis.area_id = g.area_id
CROSS JOIN w
WHERE g.location_id = $1
  AND (w.method IS NULL OR mis.source_type = w.method)
ORDER BY mis.master_item_id;


-- name: GetMasterItemIDsBySublocation :many
WITH w AS (
  SELECT sqlc.narg('method')::text AS method
)
SELECT DISTINCT mis.master_item_id
FROM mv_item_sources mis
JOIN mv_geography g ON mis.area_id = g.area_id
CROSS JOIN w
WHERE g.sublocation_id = $1
  AND (w.method IS NULL OR mis.source_type = w.method)
ORDER BY mis.master_item_id;


-- name: GetMasterItemIDsByArea :many
WITH w AS (
  SELECT sqlc.narg('method')::text AS method
)
SELECT DISTINCT mis.master_item_id
FROM mv_item_sources mis
CROSS JOIN w
WHERE mis.area_id = $1
  AND (w.method IS NULL OR mis.source_type = w.method)
ORDER BY master_item_id;





-- name: GetItemSourceIDs :many
WITH w AS (
    SELECT
      sqlc.arg('item_id')::int AS item_id,
      sqlc.arg('source_type')::text AS source_type,
      sqlc.narg('availability')::availability_type[] AS availability,
      sqlc.narg('repeatable')::boolean AS repeatable
)
SELECT DISTINCT mis.source_id
FROM mv_item_sources mis
JOIN items i ON mis.master_item_id = i.master_item_id
CROSS JOIN w
WHERE i.id = w.item_id
  AND mis.source_type = w.source_type
  AND (w.availability IS NULL OR CASE
    WHEN w.source_type = 'shop' THEN mis.avl_context
    ELSE mis.avl_self
  END = ANY(w.availability))
  AND (w.repeatable IS NULL OR mis.is_repeatable = w.repeatable)
ORDER BY mis.source_id;


-- name: GetItemPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN item_amounts ia ON pa.aeon_learn_item_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY pa.id;


-- name: GetItemAutoAbilityIDs :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN item_amounts ia ON aa.required_item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY aa.id;


-- name: GetItemMixIDs :many
SELECT DISTINCT mix_id
FROM mix_combinations
WHERE first_item_id = $1 OR second_item_id = $1
ORDER BY mix_id;


-- name: GetItemIDs :many
SELECT id FROM items ORDER BY id;


-- name: GetItemIDsCategory :many
SELECT id FROM items WHERE category = ANY(sqlc.narg('category')::item_category[]) ORDER BY id;


-- name: GetItemIDsWithAbility :many
SELECT item_id FROM item_abilities ORDER BY item_id;


-- name: GetItemIDsByRelatedStat :many
SELECT item_id FROM j_items_related_stats WHERE stat_id = $1 ORDER BY item_id;


-- name: GetItemIDsByMethod :many
SELECT DISTINCT i.id
FROM mv_item_sources mis
JOIN items i ON i.master_item_id = mis.master_item_id
WHERE mis.source_type = sqlc.arg('method')::text
ORDER BY i.id;


-- name: GetItemIDsByLocation :many
WITH w AS (
  SELECT sqlc.narg('method')::text AS method
)
SELECT DISTINCT i.id
FROM items i
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
CROSS JOIN w
WHERE g.location_id = $1
  AND (w.method IS NULL OR mis.source_type = w.method)
ORDER BY i.id;


-- name: GetItemIDsBySublocation :many
WITH w AS (
  SELECT sqlc.narg('method')::text AS method
)
SELECT DISTINCT i.id
FROM items i
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
CROSS JOIN w
WHERE g.sublocation_id = $1
  AND (w.method IS NULL OR mis.source_type = w.method)
ORDER BY i.id;


-- name: GetItemIDsByArea :many
WITH w AS (
  SELECT sqlc.narg('method')::text AS method
)
SELECT DISTINCT i.id
FROM items i
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
CROSS JOIN w
WHERE mis.area_id = $1
  AND (w.method IS NULL OR mis.source_type = w.method)
ORDER BY i.id;




-- name: GetKeyItemSourceIDs :many
WITH w AS (
    SELECT
      sqlc.arg('key_item_id')::int AS key_item_id,
      sqlc.arg('source_type')::text AS source_type,
      sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT mis.source_id
FROM mv_item_sources mis
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
CROSS JOIN w
WHERE ki.id = w.key_item_id
  AND mis.source_type = w.source_type
  AND (w.availability IS NULL OR mis.avl_self = ANY(w.availability))
ORDER BY mis.source_id;


-- name: GetKeyItemTreasureIDs :many
SELECT DISTINCT mis.source_id
FROM mv_item_sources mis
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'treasure'
ORDER BY mis.source_id;


-- name: GetKeyItemQuestIDs :many
SELECT DISTINCT mis.source_id
FROM mv_item_sources mis
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'quest'
ORDER BY mis.source_id;


-- name: GetKeyItemAreaIDs :many
SELECT ca.area_id
FROM completion_areas ca
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
JOIN mv_item_sources mis ON mis.source_id = q.id
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'quest'

UNION

SELECT t.area_id
FROM treasures t
JOIN mv_item_sources mis ON mis.source_id = t.id
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'treasure'
ORDER BY area_id;


-- name: GetKeyItemCelestialWeapon :one
SELECT id FROM celestial_weapons WHERE key_item_base = $1;


-- name: GetKeyItemIDs :many
SELECT id FROM key_items ORDER BY id;


-- name: GetKeyItemIDsCategory :many
SELECT id FROM key_items WHERE category = ANY(sqlc.narg('category')::key_item_category[]) ORDER BY id;


-- name: GetKeyItemIDsByMethod :many
SELECT ki.id
FROM mv_item_sources mis
JOIN key_items ki ON ki.master_item_id = mis.master_item_id
WHERE mis.source_type = sqlc.arg('method')::text
ORDER BY ki.id;


-- name: GetKeyItemIDsByLocation :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY ki.id;


-- name: GetKeyItemIDsBySublocation :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY ki.id;


-- name: GetKeyItemIDsByArea :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
WHERE mis.area_id = $1
ORDER BY ki.id;






-- name: GetSphereIDs :many
SELECT id FROM spheres ORDER BY id;


-- name: GetSphereIDsByColor :many
SELECT id FROM spheres WHERE sphere_color = ANY(sqlc.arg('color')::sphere_color[]) ORDER BY id;


-- name: GetSphereIDsByLocation :many
SELECT DISTINCT s.id
FROM spheres s
JOIN items i ON s.item_id = i.id
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY s.id;


-- name: GetSphereIDsBySublocation :many
SELECT DISTINCT s.id
FROM spheres s
JOIN items i ON s.item_id = i.id
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY s.id;


-- name: GetSphereIDsByArea :many
SELECT DISTINCT s.id
FROM spheres s
JOIN items i ON s.item_id = i.id
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
WHERE mis.area_id = $1
ORDER BY s.id;


-- name: GetSphereIDsByMethod :many
SELECT DISTINCT s.id
FROM spheres s
JOIN items i ON s.item_id = i.id
JOIN mv_item_sources mis ON i.master_item_id = mis.master_item_id
WHERE mis.source_type = sqlc.arg('method')::text
ORDER BY s.id;





-- name: GetPrimerTreasureIDs :many
SELECT DISTINCT j.treasure_id
FROM j_treasures_items j
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ia.master_item_id = ki.master_item_id
JOIN primers p ON p.key_item_id = ki.id
WHERE p.id = $1
ORDER BY j.treasure_id;


-- name: GetPrimerAreaIDs :many
SELECT DISTINCT mis.area_id
FROM mv_item_sources mis
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
JOIN primers p ON p.key_item_id = ki.id
WHERE p.id = $1
  AND mis.source_type = 'treasure'
ORDER BY mis.area_id;


-- name: GetPrimerIDs :many
SELECT id FROM primers ORDER BY id;






-- name: GetMixIDs :many
SELECT id FROM mixes ORDER BY id;


-- name: GetMixIDsByCategory :many
SELECT id FROM mixes WHERE category = ANY(sqlc.narg('category')::mix_category[]) ORDER BY id;


-- name: GetMixIDsByItems :many
WITH mi AS (
  SELECT
      sqlc.arg('first_item_id')::int AS first,
      sqlc.narg('second_item_id')::int AS second
)
SELECT DISTINCT mc.mix_id
FROM mix_combinations mc
CROSS JOIN mi
WHERE
  (mi.second IS NULL AND (mc.first_item_id = mi.first OR mc.second_item_id = mi.first))
  OR 
  (mi.second IS NOT NULL AND (
      (mc.first_item_id = mi.first AND mc.second_item_id = mi.second)
      OR
      (mc.first_item_id = mi.second AND mc.second_item_id = mi.first)
  ))
ORDER BY mc.mix_id;
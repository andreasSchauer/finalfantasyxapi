-- name: GetMasterItemMonstersBool :one
SELECT EXISTS (
  SELECT 1
  FROM mv_monster_item_drops
  WHERE master_item_id = $1
) AS obtainable_from_monsters;


-- name: GetMasterItemTreasuresBool :one
SELECT EXISTS (
  SELECT 1
  FROM j_treasures_items j
  JOIN item_amounts ia ON j.item_amount_id = ia.id
  WHERE ia.master_item_id = $1
) AS obtainable_from_treasures;


-- name: GetMasterItemShopsBool :one
SELECT EXISTS (
  SELECT 1
  FROM j_shops_items j
  JOIN shop_items si ON j.shop_item_id = si.id
  JOIN items i ON si.item_id = i.id
  WHERE i.master_item_id = $1
) AS obtainable_from_shops;


-- name: GetMasterItemQuestsBool :one
SELECT EXISTS (
  SELECT 1
  FROM quests q
  JOIN quest_completions qc ON q.completion_id = qc.id
  JOIN item_amounts ia ON qc.item_amount_id = ia.id
  WHERE ia.master_item_id = $1
) AS obtainable_from_quests;


-- name: GetMasterItemIDs :many
SELECT id FROM master_items ORDER BY id;


-- name: GetMasterItemIDsByType :many
SELECT id FROM master_items WHERE type = ANY(sqlc.narg('item_type')::item_type[]) ORDER BY id;


-- name: GetMasterItemIDsMonster :many
SELECT DISTINCT master_item_id FROM mv_monster_item_drops ORDER BY master_item_id;


-- name: GetMasterItemIDsTreasure :many
SELECT DISTINCT ia.master_item_id
FROM item_amounts ia
JOIN j_treasures_items j ON j.item_amount_id = ia.id
ORDER BY ia.master_item_id;


-- name: GetMasterItemIDsShop :many
SELECT DISTINCT i.master_item_id
FROM items i
JOIN shop_items si ON si.item_id = i.id
JOIN j_shops_items j ON j.shop_item_id = si.id
ORDER BY i.master_item_id;


-- name: GetMasterItemIDsQuest :many
SELECT DISTINCT ia.master_item_id
FROM item_amounts ia
JOIN quest_completions qc ON qc.item_amount_id = ia.id
ORDER BY ia.master_item_id;








-- name: GetItemMonsterIDs :many
WITH w AS (
    SELECT
      sqlc.narg('repeatable')::BOOLEAN AS repeatable,
      sqlc.narg('availability')::availability_type[] AS availability,
      (SELECT i.master_item_id FROM items i WHERE i.id = sqlc.arg('item_id'))::int AS target_master_id
)
SELECT DISTINCT m.id
FROM monsters m
CROSS JOIN w
JOIN mv_monster_item_drops md ON md.monster_id = m.id
WHERE md.master_item_id = w.target_master_id
  AND (w.repeatable IS NULL OR m.is_repeatable = w.repeatable)
  AND (w.availability IS NULL OR m.availability = ANY(w.availability))
ORDER BY m.id;


-- name: GetItemTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
WHERE
  i.id = sqlc.arg(item_id)
  AND (w.availability IS NULL OR t.availability = ANY(w.availability))
ORDER BY t.id;


-- name: GetItemShopIDs :many
SELECT DISTINCT j.shop_id
FROM j_shops_items j
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
CROSS JOIN LATERAL (
  SELECT CASE j.shop_type
      WHEN 'pre-airship' THEN 'pre-story'::availability_type
      WHEN 'post-airship' THEN 'post'::availability_type
  END AS shop_availability
) calc
WHERE i.id = sqlc.arg(item_id)::int
  AND (w.availability IS NULL OR calc.shop_availability = ANY(w.availability))
ORDER BY j.shop_id;


-- name: GetItemQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
CROSS JOIN (
    SELECT
        sqlc.narg('repeatable')::BOOLEAN AS repeatable,
        sqlc.narg('availability')::availability_type[] AS availability
) w
WHERE
  i.id = sqlc.arg(item_id)
  AND (w.repeatable IS NULL OR q.is_repeatable = w.repeatable)
  AND (w.availability IS NULL OR q.availability = ANY(w.availability))
ORDER BY q.id;


-- name: GetItemBlitzballPrizeIDs :many
SELECT DISTINCT bi.position_id
FROM blitzball_items bi
JOIN possible_items pi ON bi.possible_item_id = pi.id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY bi.position_id;


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


-- name: GetItemIDsMonster :many
SELECT DISTINCT item_id FROM mv_monster_item_drops ORDER BY master_item_id;


-- name: GetItemIDsTreasure :many
SELECT DISTINCT i.id
FROM items i
JOIN item_amounts ia ON i.master_item_id = ia.master_item_id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
ORDER BY i.id;


-- name: GetItemIDsShop :many
SELECT DISTINCT i.id
FROM items i
JOIN shop_items si ON si.item_id = i.id
JOIN j_shops_items j ON j.shop_item_id = si.id
ORDER BY i.id;


-- name: GetItemIDsQuest :many
SELECT DISTINCT i.id
FROM items i
JOIN item_amounts ia ON i.master_item_id = ia.master_item_id
JOIN quest_completions qc ON qc.item_amount_id = ia.id
ORDER BY i.id;






-- name: GetKeyItemTreasureIDs :many
SELECT DISTINCT j.treasure_id
FROM j_treasures_items j
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ia.master_item_id = ki.master_item_id
WHERE ki.id = $1
ORDER BY j.treasure_id;


-- name: GetKeyItemQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN key_items ki ON ia.master_item_id = ki.master_item_id
WHERE ki.id = $1
ORDER BY q.id;


-- name: GetKeyItemAreaIDs :many
SELECT ca.area_id
FROM completion_areas ca
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1

UNION

SELECT t.area_id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY area_id;


-- name: GetKeyItemCelestialWeapon :one
SELECT id FROM celestial_weapons WHERE key_item_base = $1;


-- name: GetKeyItemIDs :many
SELECT id FROM key_items ORDER BY id;


-- name: GetKeyItemIDsCategory :many
SELECT id FROM key_items WHERE category = ANY(sqlc.narg('category')::key_item_category[]) ORDER BY id;


-- name: GetKeyItemIDsTreasure :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN item_amounts ia ON ki.master_item_id = ia.master_item_id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
ORDER BY ki.id;


-- name: GetKeyItemIDsQuest :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN item_amounts ia ON ki.master_item_id = ia.master_item_id
JOIN quest_completions qc ON qc.item_amount_id = ia.id
ORDER BY ki.id;


-- name: GetKeyItemIDsByAvailability :many
WITH w AS (
  SELECT sqlc.narg('availability')::availability_type[] AS availability
)
SELECT ki.id
FROM key_items ki
JOIN item_amounts ia ON ki.master_item_id = ia.master_item_id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
JOIN treasures t ON j.treasure_id = t.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
WHERE w.availability IS NULL OR t.availability = ANY(w.availability)

UNION

SELECT ki.id
FROM key_items ki
JOIN item_amounts ia ON ki.master_item_id = ia.master_item_id
JOIN quest_completions qc ON qc.item_amount_id = ia.id
JOIN quests q ON q.completion_id = qc.id
CROSS JOIN w
WHERE w.availability IS NULL OR q.availability = ANY(w.availability)
ORDER BY id;







-- name: GetSphereIDs :many
SELECT id FROM spheres ORDER BY id;


-- name: GetSphereIDsByColor :many
SELECT id FROM spheres WHERE sphere_color = ANY(sqlc.arg('color')::sphere_color[]) ORDER BY id;





-- name: GetPrimerTreasureIDs :many
SELECT DISTINCT j.treasure_id
FROM j_treasures_items j
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ia.master_item_id = ki.master_item_id
JOIN primers p ON p.key_item_id = ki.id
WHERE p.id = $1
ORDER BY j.treasure_id;


-- name: GetPrimerAreaIDs :many
SELECT DISTINCT t.area_id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ia.master_item_id = ki.master_item_id
JOIN primers p ON p.key_item_id = ki.id
WHERE p.id = $1
ORDER BY t.area_id;


-- name: GetPrimerIDs :many
SELECT id FROM primers ORDER BY id;


-- name: GetPrimerIDsByAvailability :many
SELECT DISTINCT p.id
FROM primers p
JOIN key_items ki ON p.key_item_id = ki.id
JOIN item_amounts ia ON ki.master_item_id = ia.master_item_id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
JOIN treasures t ON j.treasure_id = t.id
WHERE t.availability = ANY(sqlc.arg('availability')::availability_type[])
ORDER BY p.id;






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
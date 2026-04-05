-- name: GetItemMonsterIDs :many
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
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON i.master_item_id = mit.id
WHERE
    i.id = sqlc.arg(item_id)
    AND (sqlc.narg('repeatable')::BOOLEAN IS NULL OR m.is_repeatable = sqlc.narg('repeatable')::BOOLEAN)
    AND (sqlc.narg('availability')::availability_type[] IS NULL OR m.availability = ANY(sqlc.narg('availability')::availability_type[]))
ORDER BY m.id;


-- name: GetItemTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE
  i.id = sqlc.arg(item_id)
  AND (sqlc.narg('availability')::availability_type[] IS NULL OR t.availability = ANY(sqlc.narg('availability')::availability_type[]))
ORDER BY t.id;


-- name: GetItemShopIDs :many
SELECT DISTINCT s.id
FROM shops s
JOIN j_shops_items j ON j.shop_id = s.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
WHERE
  i.id = sqlc.arg(item_id)
  AND (sqlc.narg('availability')::availability_type[] IS NULL OR s.availability = ANY(sqlc.narg('availability')::availability_type[]))
ORDER BY s.id;


-- name: GetItemQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE
  i.id = sqlc.arg(item_id)
  AND (sqlc.narg('repeatable')::BOOLEAN IS NULL OR q.is_repeatable = sqlc.narg('repeatable')::BOOLEAN)
  AND (sqlc.narg('availability')::availability_type[] IS NULL OR q.availability = ANY(sqlc.narg('availability')::availability_type[]))
ORDER BY q.id;


-- name: GetItemBlitzballPrizeIDs :many
SELECT DISTINCT bp.id
FROM blitzball_positions bp
JOIN blitzball_items bi ON bi.position_id = bp.id
JOIN possible_items pi ON bi.possible_item_id = pi.id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY bp.id;


-- name: GetItemPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN item_amounts ia ON pa.aeon_learn_item_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY pa.id;


-- name: GetItemAutoAbilityIDs :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN item_amounts ia ON aa.required_item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY aa.id;


-- name: GetItemMixIDs :many
SELECT DISTINCT m.id
FROM mixes m
JOIN mix_combinations mc ON mc.mix_id = m.id
JOIN items i ON mc.first_item_id = i.id OR mc.second_item_id = i.id
WHERE i.id = $1
ORDER BY m.id;


-- name: GetItemIDs :many
SELECT id FROM items ORDER BY id;


-- name: GetItemIDsCategory :many
SELECT id FROM items WHERE category = ANY(sqlc.narg('category')::item_category[]) ORDER BY id;


-- name: GetItemIDsWithAbility :many
SELECT i.id
FROM items i
JOIN item_abilities ia ON ia.item_id = i.id
ORDER BY i.id;


-- name: GetItemIDsByRelatedStat :many
SELECT i.id
FROM items i
JOIN j_items_related_stats j ON j.item_id = i.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY i.id;


-- name: GetItemIDsMonster :many
SELECT DISTINCT i.id
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
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON i.master_item_id = mit.id
ORDER BY i.id;


-- name: GetItemIDsTreasure :many
SELECT DISTINCT i.id
FROM items i
JOIN master_items mi ON i.master_item_id = mi.id
JOIN item_amounts ia ON ia.master_item_id = mi.id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
JOIN treasures t ON j.treasure_id = t.id
ORDER BY i.id;


-- name: GetItemIDsShop :many
SELECT DISTINCT i.id
FROM items i
JOIN shop_items si ON si.item_id = i.id
JOIN j_shops_items j ON j.shop_item_id = si.id
JOIN shops s ON j.shop_id = s.id
ORDER BY i.id;


-- name: GetItemIDsQuest :many
SELECT DISTINCT i.id
FROM items i
JOIN master_items mi ON i.master_item_id = mi.id
JOIN item_amounts ia ON ia.master_item_id = mi.id
JOIN quest_completions qc ON qc.item_amount_id = ia.id
JOIN quests q ON q.completion_id = qc.id
ORDER BY i.id;




-- name: GetKeyItemTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY t.id;


-- name: GetKeyItemQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY q.id;


-- name: GetKeyItemAreaIDs :many
SELECT DISTINCT a.id
FROM areas a
WHERE
  EXISTS (
    SELECT 1
    FROM completion_areas ca
    JOIN quest_completions qc ON ca.completion_id = qc.id
    JOIN item_amounts ia ON qc.item_amount_id = ia.id
    JOIN master_items mi ON ia.master_item_id = mi.id
    JOIN key_items ki ON ki.master_item_id = mi.id
    WHERE ca.area_id = a.id
      AND ki.id = $1
  )
  OR
  EXISTS (
    SELECT 1
    FROM treasures t
    JOIN j_treasures_items j ON j.treasure_id = t.id
    JOIN item_amounts ia ON j.item_amount_id = ia.id
    JOIN master_items mi ON ia.master_item_id = mi.id
    JOIN key_items ki ON ki.master_item_id = mi.id
    WHERE t.area_id = a.id
      AND ki.id = $1
  )
ORDER BY a.id;


-- name: GetKeyItemCelestialWeapon :one
SELECT id FROM celestial_weapons WHERE key_item_base = $1;


-- name: GetKeyItemIDs :many
SELECT id FROM key_items ORDER BY id;


-- name: GetKeyItemIDsCategory :many
SELECT id FROM key_items WHERE category = ANY(sqlc.narg('category')::key_item_category[]) ORDER BY id;


-- name: GetKeyItemIDsTreasure :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN master_items mi ON ki.master_item_id = mi.id
JOIN item_amounts ia ON ia.master_item_id = mi.id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
JOIN treasures t ON j.treasure_id = t.id
ORDER BY ki.id;


-- name: GetKeyItemIDsQuest :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN master_items mi ON ki.master_item_id = mi.id
JOIN item_amounts ia ON ia.master_item_id = mi.id
JOIN quest_completions qc ON qc.item_amount_id = ia.id
JOIN quests q ON q.completion_id = qc.id
ORDER BY ki.id;


-- name: GetKeyItemIDsByAvailability :many
SELECT DISTINCT ki.id
FROM key_items ki
JOIN master_items mi ON ki.master_item_id = mi.id
JOIN item_amounts ia ON ia.master_item_id = mi.id
WHERE
  EXISTS (
    SELECT 1
    FROM j_treasures_items j
    JOIN treasures t ON j.treasure_id = t.id
    WHERE j.item_amount_id = ia.id
      AND (sqlc.narg('availability')::availability_type[] IS NULL OR t.availability = ANY(sqlc.narg('availability')::availability_type[]))
  )
  OR
  EXISTS (
    SELECT 1
    FROM quest_completions qc
    JOIN quests q ON q.completion_id = qc.id
    WHERE qc.item_amount_id = ia.id
      AND (sqlc.narg('availability')::availability_type[] IS NULL OR q.availability = ANY(sqlc.narg('availability')::availability_type[]))
  )
ORDER BY ki.id;







-- name: GetPrimerTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
JOIN primers p ON p.key_item_id = ki.id
WHERE p.id = $1
ORDER BY t.id;


-- name: GetPrimerAreaIDs :many
SELECT DISTINCT a.id
FROM areas a
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
JOIN primers p ON p.key_item_id = ki.id
WHERE p.id = $1
ORDER BY a.id;


-- name: GetPrimerIDs :many
SELECT id FROM primers ORDER BY id;


-- name: GetPrimerIDsByAvailability :many
SELECT DISTINCT p.id
FROM primers p
JOIN key_items ki ON p.key_item_id = ki.id
JOIN master_items mi ON ki.master_item_id = mi.id
JOIN item_amounts ia ON ia.master_item_id = mi.id
JOIN j_treasures_items j ON j.item_amount_id = ia.id
JOIN treasures t ON j.treasure_id = t.id
WHERE t.availability = ANY(sqlc.narg('availability')::availability_type[])
ORDER BY p.id;
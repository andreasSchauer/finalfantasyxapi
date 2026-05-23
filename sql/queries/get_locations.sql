-- name: GetAreaIDs :many
SELECT id FROM areas ORDER BY id;


-- name: GetAreaConnectionIDs :many
SELECT DISTINCT ca_id FROM mv_geography_graph WHERE a_id = $1 ORDER BY ca_id;


-- name: GetAreaCharacterIDs :many
SELECT DISTINCT id FROM characters WHERE area_id = $1 ORDER BY id;


-- name: GetAreaAeonIDs :many
SELECT DISTINCT id FROM aeons WHERE area_id = $1 ORDER BY id;


-- name: GetAreaShopIDs :many
SELECT id FROM shops WHERE area_id = $1 ORDER BY id;


-- name: GetAreaShopIdPairs :many
SELECT DISTINCT
  area_id,
  id AS shop_id
FROM shops
WHERE area_id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY area_id, shop_id;


-- name: GetAreaTreasureIDs :many
SELECT id FROM treasures WHERE area_id = $1 ORDER BY id;


-- name: GetAreaTreasureIdPairs :many
SELECT DISTINCT
  area_id,
  id AS treasure_id
FROM treasures 
WHERE area_id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY area_id, treasure_id;


-- name: GetAreaMonsterIDs :many
SELECT DISTINCT monster_id FROM mv_monster_encounters WHERE area_id = $1 ORDER BY monster_id;


-- name: GetAreaMonsterIdPairs :many
SELECT DISTINCT
  area_id,
  monster_id
FROM mv_monster_encounters
WHERE area_id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY area_id, monster_id;


-- name: GetAreaMonsterFormationIDs :many
SELECT DISTINCT formation_id FROM mv_monster_encounters WHERE area_id = $1 ORDER BY formation_id;


-- name: GetAreaBossSongIDs :many
SELECT DISTINCT song_id::int FROM mv_monster_encounters WHERE area_id = $1 AND song_id IS NOT NULL ORDER BY song_id;


-- name: GetAreaCueSongIDs :many
SELECT c.song_id
FROM cues c
WHERE c.trigger_area_id = $1

UNION

SELECT c.song_id
FROM cues c
JOIN j_cues_areas j ON j.cue_id = c.id
WHERE j.included_area_id = $1
ORDER BY song_id;


-- name: GetAreaBackgroundMusicSongIDs :many
SELECT DISTINCT song_id FROM j_songs_background_music WHERE area_id = $1 ORDER BY song_id;


-- name: GetAreaFMVSongIDs :many
SELECT DISTINCT song_id::int FROM fmvs WHERE song_id IS NOT NULL AND area_id = $1 ORDER BY song_id;


-- name: GetAreaFmvIDs :many
SELECT id FROM fmvs WHERE area_id = $1 ORDER BY id;


-- name: GetAreaQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_areas ca ON ca.completion_id = qc.id
WHERE ca.area_id = $1
ORDER BY q.id;


-- name: GetAreaIDsWithSaveSphere :many
SELECT id FROM areas WHERE has_save_sphere = $1 ORDER BY id;


-- name: GetAreaIDsWithCompSphere :many
SELECT id FROM areas WHERE has_compilation_sphere = $1 ORDER BY id;


-- name: GetAreaIDsWithDropOff :many
SELECT id FROM areas WHERE airship_drop_off = $1 ORDER BY id;


-- name: GetAreaIDsChocobo :many
SELECT id FROM areas WHERE can_ride_chocobo = $1 ORDER BY id;


-- name: GetAreaIDsWithCharacters :many
SELECT DISTINCT area_id::int FROM characters ORDER BY area_id;


-- name: GetAreaIDsWithAeons :many
SELECT DISTINCT area_id::int FROM aeons ORDER BY area_id;


-- name: GetAreaIDsWithMonsters :many
SELECT DISTINCT area_id FROM mv_monster_encounters ORDER BY area_id;


-- name: GetAreaIDsWithBosses :many
SELECT DISTINCT me.area_id
FROM mv_monster_encounters me
JOIN monsters m ON me.monster_id = m.id
WHERE m.category = 'boss'
ORDER BY me.area_id;


-- name: GetAreaIDsWithShops :many
SELECT DISTINCT area_id FROM shops ORDER BY area_id;


-- name: GetAreaIDsWithTreasures :many
SELECT DISTINCT area_id FROM treasures ORDER BY area_id;


-- name: GetAreasByMonster :many
SELECT DISTINCT area_id FROM mv_monster_encounters WHERE monster_id = $1 ORDER BY area_id;


-- name: GetAreaIDsWithItemFromMethod :many
WITH w AS (
    SELECT sqlc.arg('method')::text[] AS method
)
SELECT DISTINCT mis.area_id
FROM mv_item_sources mis
JOIN items i ON mis.master_item_id = i.master_item_id
CROSS JOIN w
WHERE i.id = $1
  AND (w.method IS NULL OR mis.source_type = ANY(w.method))
ORDER BY mis.area_id;


-- name: GetAreaIDsWithKeyItem :many
SELECT DISTINCT mis.area_id
FROM mv_item_sources mis
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
ORDER BY mis.area_id;


-- name: GetAreaIDsWithSidequests :many
SELECT DISTINCT area_id FROM completion_areas ORDER BY area_id;


-- name: GetAreaIDsWithFMVs :many
SELECT DISTINCT area_id FROM fmvs ORDER BY area_id;







-- name: GetSublocationIDs :many
SELECT id FROM sublocations ORDER BY id;


-- name: GetSublocationAreaIDs :many
SELECT DISTINCT id FROM areas WHERE sublocation_id = $1 ORDER BY id;


-- name: GetConnectedSublocationIDs :many
SELECT DISTINCT cs_id FROM mv_geography_graph WHERE s_id = $1 AND cs_id != $1 ORDER BY cs_id;


-- name: GetSublocationCharacterIDs :many
SELECT c.id
FROM characters c
JOIN mv_geography g ON c.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY c.id;


-- name: GetSublocationAeonIDs :many
SELECT a.id
FROM aeons a
JOIN mv_geography g ON a.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY a.id;


-- name: GetSublocationShopIDs :many
SELECT sh.id
FROM shops sh
JOIN mv_geography g ON sh.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY sh.id;


-- name: GetSublocationShopIdPairs :many
SELECT DISTINCT
  g.sublocation_id,
  sh.id AS shop_id
FROM shops sh
JOIN mv_geography g ON sh.area_id = g.area_id
WHERE g.sublocation_id = ANY(sqlc.arg('sublocation_ids')::int[])
ORDER BY g.sublocation_id, shop_id;


-- name: GetSublocationTreasureIDs :many
SELECT t.id
FROM treasures t
JOIN mv_geography g ON t.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY t.id;


-- name: GetSublocationTreasureIdPairs :many
SELECT DISTINCT
  g.sublocation_id,
  t.id AS treasure_id
FROM treasures t
JOIN mv_geography g ON t.area_id = g.area_id
WHERE g.sublocation_id = ANY(sqlc.arg('location_ids')::int[])
ORDER BY g.sublocation_id, treasure_id;


-- name: GetSublocationMonsterIDs :many
SELECT DISTINCT me.monster_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY me.monster_id;


-- name: GetSublocationMonsterIdPairs :many
SELECT DISTINCT
  g.sublocation_id,
  me.monster_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.sublocation_id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY g.sublocation_id, me.monster_id;


-- name: GetSublocationMonsterFormationIDs :many
SELECT DISTINCT me.formation_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY me.formation_id;


-- name: GetSublocationBossSongIDs :many
SELECT DISTINCT me.song_id::int
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.sublocation_id = $1 AND me.song_id IS NOT NULL
ORDER BY me.song_id;


-- name: GetSublocationCueSongIDs :many
SELECT c.song_id
FROM cues c
JOIN mv_geography g ON c.trigger_area_id = g.area_id
WHERE g.sublocation_id = $1

UNION

SELECT c.song_id
FROM cues c
JOIN j_cues_areas j ON j.cue_id = c.id
JOIN mv_geography g ON j.included_area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY song_id;



-- name: GetSublocationBackgroundMusicSongIDs :many
SELECT DISTINCT j.song_id
FROM j_songs_background_music j
JOIN mv_geography g ON j.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY j.song_id;


-- name: GetSublocationFMVSongIDs :many
SELECT DISTINCT f.song_id::int
FROM fmvs f
JOIN mv_geography g ON f.area_id = g.area_id
WHERE f.song_id IS NOT NULL
  AND g.sublocation_id = $1
ORDER BY f.song_id;


-- name: GetSublocationFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN mv_geography g ON f.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY f.id;


-- name: GetSublocationQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_areas ca ON ca.completion_id = qc.id
JOIN mv_geography g ON ca.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY q.id;


-- name: GetSublocationIDsWithCharacters :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN characters c ON c.area_id = g.area_id
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithAeons :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN aeons a ON a.area_id = g.area_id
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithMonsters :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_monster_encounters me ON me.area_id = g.area_id
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithBosses :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_monster_encounters me ON me.area_id = g.area_id
WHERE me.song_id IS NOT NULL
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithShops :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN shops sh ON sh.area_id = g.area_id
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithTreasures :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN treasures t ON t.area_id = g.area_id
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithItemFromMonster :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'monster'
ORDER BY g.sublocation_id;



-- name: GetSublocationIDsWithItemFromShop :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'shop'
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithItemFromTreasure :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'treasure'
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithItemFromQuest :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'quest'
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithKeyItemFromTreasure :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'treasure'
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithKeyItemFromQuest :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'quest'
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithSidequests :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN completion_areas ca ON ca.area_id = g.area_id
ORDER BY g.sublocation_id;


-- name: GetSublocationIDsWithFMVs :many
SELECT DISTINCT g.sublocation_id
FROM mv_geography g
JOIN fmvs f ON f.area_id = g.area_id
ORDER BY g.sublocation_id;






-- name: GetLocationIDs :many
SELECT id FROM locations ORDER BY id;


-- name: GetLocationSublocationIDs :many
SELECT DISTINCT sublocation_id FROM mv_geography WHERE location_id = $1 ORDER BY sublocation_id;


-- name: GetLocationAreaIDs :many
SELECT DISTINCT area_id FROM mv_geography WHERE location_id = $1 ORDER BY area_id;


-- name: GetConnectedLocationIDs :many
SELECT DISTINCT cl_id FROM mv_geography_graph WHERE l_id = $1 AND cl_id != $1 ORDER BY cl_id;


-- name: GetLocationCharacterIDs :many
SELECT c.id
FROM characters c
JOIN mv_geography g ON c.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY c.id;


-- name: GetLocationAeonIDs :many
SELECT a.id
FROM aeons a
JOIN mv_geography g ON a.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY a.id;


-- name: GetLocationShopIDs :many
SELECT sh.id
FROM shops sh
JOIN mv_geography g ON sh.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY sh.id;


-- name: GetLocationShopIdPairs :many
SELECT DISTINCT
  g.location_id,
  sh.id AS shop_id
FROM shops sh
JOIN mv_geography g ON sh.area_id = g.area_id
WHERE g.location_id = ANY(sqlc.arg('location_ids')::int[])
ORDER BY g.location_id, shop_id;


-- name: GetLocationTreasureIDs :many
SELECT t.id
FROM treasures t
JOIN mv_geography g ON t.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY t.id;


-- name: GetLocationTreasureIdPairs :many
SELECT DISTINCT
  g.location_id,
  t.id AS treasure_id
FROM treasures t
JOIN mv_geography g ON t.area_id = g.area_id
WHERE g.location_id = ANY(sqlc.arg('location_ids')::int[])
ORDER BY g.location_id, treasure_id;


-- name: GetLocationMonsterIDs :many
SELECT DISTINCT me.monster_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY me.monster_id;


-- name: GetLocationMonsterIdPairs :many
SELECT DISTINCT
  g.location_id,
  me.monster_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.location_id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY g.location_id, me.monster_id;


-- name: GetLocationMonsterFormationIDs :many
SELECT DISTINCT me.formation_id
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY me.formation_id;


-- name: GetLocationBossSongIDs :many
SELECT DISTINCT me.song_id::int
FROM mv_monster_encounters me
JOIN mv_geography g ON me.area_id = g.area_id
WHERE g.location_id = $1 AND me.song_id IS NOT NULL
ORDER BY me.song_id;


-- name: GetLocationCueSongIDs :many
SELECT c.song_id
FROM cues c
JOIN mv_geography g ON c.trigger_area_id = g.area_id
WHERE g.location_id = $1

UNION

SELECT c.song_id
FROM cues c
JOIN j_cues_areas j ON j.cue_id = c.id
JOIN mv_geography g ON j.included_area_id = g.area_id
WHERE g.location_id = $1
ORDER BY song_id;


-- name: GetLocationBackgroundMusicSongIDs :many
SELECT DISTINCT j.song_id
FROM j_songs_background_music j
JOIN mv_geography g ON j.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY j.song_id;


-- name: GetLocationFMVSongIDs :many
SELECT DISTINCT f.song_id::int
FROM fmvs f
JOIN mv_geography g ON f.area_id = g.area_id
WHERE f.song_id IS NOT NULL
  AND g.location_id = $1
ORDER BY f.song_id;


-- name: GetLocationFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN mv_geography g ON f.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY f.id;


-- name: GetLocationQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_areas ca ON ca.completion_id = qc.id
JOIN mv_geography g ON ca.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY q.id;


-- name: GetLocationIDsWithCharacters :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN characters c ON c.area_id = g.area_id
ORDER BY g.location_id;


-- name: GetLocationIDsWithAeons :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN aeons a ON a.area_id = g.area_id
ORDER BY g.location_id;


-- name: GetLocationIDsWithMonsters :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_monster_encounters me ON me.area_id = g.area_id
ORDER BY g.location_id;


-- name: GetLocationIDsWithBosses :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_monster_encounters me ON me.area_id = g.area_id
WHERE me.song_id IS NOT NULL
ORDER BY g.location_id;


-- name: GetLocationIDsWithShops :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN shops sh ON sh.area_id = g.area_id
ORDER BY g.location_id;


-- name: GetLocationIDsWithTreasures :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN treasures t ON t.area_id = g.area_id
ORDER BY g.location_id;


-- name: GetLocationIDsWithItemFromMonster :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'monster'
ORDER BY g.location_id;


-- name: GetLocationIDsWithItemFromShop :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'shop'
ORDER BY g.location_id;


-- name: GetLocationIDsWithItemFromTreasure :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'treasure'
ORDER BY g.location_id;


-- name: GetLocationIDsWithItemFromQuest :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = $1
  AND mis.source_type = 'quest'
ORDER BY g.location_id;


-- name: GetLocationIDsWithKeyItemFromTreasure :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'treasure'
ORDER BY g.location_id;


-- name: GetLocationIDsWithKeyItemFromQuest :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN mv_item_sources mis ON mis.area_id = g.area_id
JOIN key_items ki ON mis.master_item_id = ki.master_item_id
WHERE ki.id = $1
  AND mis.source_type = 'quest'
ORDER BY g.location_id;


-- name: GetLocationIDsWithSidequests :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN completion_areas ca ON ca.area_id = g.area_id
ORDER BY g.location_id;


-- name: GetLocationIDsWithFMVs :many
SELECT DISTINCT g.location_id
FROM mv_geography g
JOIN fmvs f ON f.area_id = g.area_id
ORDER BY g.location_id;






-- name: GetTreasureIDs :many
SELECT id FROM treasures ORDER BY id;


-- name: GetTreasureIDsByLootType :many
SELECT id FROM treasures WHERE loot_type = $1 ORDER BY id;


-- name: GetTreasureIDsByTreasureType :many
SELECT id FROM treasures WHERE treasure_type = $1 ORDER BY id;


-- name: GetTreasureIDsByItem :many
SELECT DISTINCT mis.source_id
FROM mv_item_sources mis
JOIN items i ON mis.master_item_id = i.master_item_id
WHERE i.id = 1
  AND mis.source_type = 'treasure'
ORDER BY mis.source_id;


-- name: GetTreasureIDsByIsAnimaTreasure :many
SELECT id FROM treasures WHERE is_anima_treasure = $1 ORDER BY id;







-- name: GetShopIDs :many
SELECT id FROM shops ORDER BY id;


-- name: GetShopIDsByCategory :many
SELECT id FROM shops WHERE category = ANY(sqlc.narg('category')::shop_category[]) ORDER BY id;


-- name: GetShopIDsEquipmentFilter :many
SELECT DISTINCT es.source_id
FROM mv_equipment_sources es
LEFT JOIN equipment_names en ON es.name_id = en.id
CROSS JOIN (
    SELECT
        sqlc.narg('empty_slots')::int[] AS empty_slots,
        sqlc.narg('character_id')::int AS character_id,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id
) w
WHERE 
    (w.empty_slots IS NULL OR es.empty_slots_amount::int = ANY(w.empty_slots))
    AND (w.character_id IS NULL OR en.character_id = w.character_id)
    AND (w.auto_ability_id IS NULL OR es.auto_ability_id = w.auto_ability_id)
    AND es.source_type = 'shop'
ORDER BY es.source_id;


-- name: GetShopIDsWithItems :many
SELECT DISTINCT s_id 
FROM mv_availabilities 
WHERE source_type = 'shop' AND sub_type = 'item'
ORDER BY s_id;


-- name: GetShopIDsWithEquipment :many
SELECT DISTINCT s_id 
FROM mv_availabilities 
WHERE source_type = 'shop' AND sub_type = 'equip'
ORDER BY s_id;


-- name: GetShopIDsByLocation :many
WITH w AS (
  SELECT sqlc.arg('location_id')::int AS location_id
)
SELECT sh.id
FROM shops sh
CROSS JOIN w
JOIN mv_geography g ON sh.area_id = g.area_id
JOIN mv_item_sources mis ON mis.source_id = sh.id AND source_type = 'shop'
WHERE g.location_id = w.location_id

UNION

SELECT sh.id
FROM shops sh
CROSS JOIN w
JOIN mv_geography g ON sh.area_id = g.area_id
JOIN mv_equipment_sources mes ON mes.source_id = sh.id AND source_type = 'shop'
WHERE g.location_id = w.location_id
ORDER BY id;


-- name: GetShopIDsBySublocation :many
WITH w AS (
  SELECT sqlc.arg('sublocation_id')::int AS sublocation_id
)
SELECT sh.id
FROM shops sh
CROSS JOIN w
JOIN mv_geography g ON sh.area_id = g.area_id
JOIN mv_item_sources mis ON mis.source_id = sh.id AND source_type = 'shop'
WHERE g.sublocation_id = w.sublocation_id

UNION

SELECT sh.id
FROM shops sh
CROSS JOIN w
JOIN mv_geography g ON sh.area_id = g.area_id
JOIN mv_equipment_sources mes ON mes.source_id = sh.id AND source_type = 'shop'
WHERE g.sublocation_id = w.sublocation_id
ORDER BY id;
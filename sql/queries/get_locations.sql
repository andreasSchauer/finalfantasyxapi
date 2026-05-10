-- name: GetAreaIDs :many
SELECT id FROM areas ORDER BY id;


-- name: GetAreaConnectionIDs :many
SELECT DISTINCT ca.id
FROM j_area_connected_areas j
JOIN area_connections ac ON j.connection_id = ac.id
JOIN areas ca ON ac.area_id = ca.id
WHERE j.area_id = $1
ORDER BY ca.id;


-- name: GetAreaCharacterIDs :many
SELECT DISTINCT id FROM characters WHERE area_id = $1 ORDER BY id;


-- name: GetAreaAeonIDs :many
SELECT DISTINCT id FROM aeons WHERE area_id = $1 ORDER BY id;


-- name: GetAreaShopIDs :many
SELECT id FROM shops WHERE area_id = $1 ORDER BY id;


-- name: GetAreaShopIdPairs :many
SELECT DISTINCT
  a.id AS area_id,
  sh.id AS shop_id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
WHERE a.id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY a.id, sh.id;


-- name: GetAreaTreasureIDs :many
SELECT id FROM treasures WHERE area_id = $1 ORDER BY id;


-- name: GetAreaTreasureIdPairs :many
SELECT DISTINCT
  a.id AS area_id,
  t.id AS treasure_id
FROM treasures t
JOIN areas a ON t.area_id = a.id
WHERE a.id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY a.id, t.id;


-- name: GetAreaMonsterIDs :many
SELECT DISTINCT ma.monster_id
FROM monster_amounts ma
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_formations mf ON mf.monster_selection_id = j1.monster_selection_id
JOIN j_monster_formations_encounter_areas j2 ON j2.monster_formation_id = mf.id
JOIN encounter_areas ea ON j2.encounter_area_id = ea.id
WHERE ea.area_id = $1
ORDER BY ma.monster_id;


-- name: GetAreaMonsterIdPairs :many
SELECT DISTINCT
  a.id AS area_id,
  m.id AS monster_id
FROM monsters m
JOIN monster_amounts ma ON ma.monster_id = m.id
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_selections ms ON j1.monster_selection_id = ms.id
JOIN monster_formations mf ON mf.monster_selection_id = ms.id
JOIN j_monster_formations_encounter_areas j2 ON j2.monster_formation_id = mf.id
JOIN encounter_areas ea ON j2.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
WHERE a.id = ANY(sqlc.arg('area_ids')::int[])
ORDER BY a.id, m.id;


-- name: GetAreaMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
WHERE ea.area_id = $1
ORDER BY mf.id;


-- name: GetAreaBossSongIDs :many
SELECT DISTINCT bs.song_id
FROM formation_boss_songs bs
JOIN formation_data fd ON fd.boss_song_id = bs.id
JOIN monster_formations mf ON mf.formation_data_id = fd.id
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
WHERE ea.area_id = $1
ORDER BY bs.song_id;


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
JOIN completion_areas ca ON ca.completion_id = q.completion_id
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


-- name: GetAreaIDsByAvailability :many
SELECT id FROM areas WHERE availability = ANY(sqlc.narg('availability')::availability_type[]) ORDER BY id;


-- name: GetAreaIDsWithCharacters :many
SELECT DISTINCT area_id::int FROM characters ORDER BY area_id;


-- name: GetAreaIDsWithAeons :many
SELECT DISTINCT area_id::int FROM aeons ORDER BY area_id;


-- name: GetAreaIDsWithMonsters :many
SELECT DISTINCT ea.area_id
FROM encounter_areas ea
JOIN j_monster_formations_encounter_areas j1 ON j1.encounter_area_id = ea.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_selections_monsters j2 ON mf.monster_selection_id = j2.monster_selection_id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
ORDER BY ea.area_id;


-- name: GetAreaIDsWithBosses :many
SELECT DISTINCT ea.area_id
FROM encounter_areas ea
JOIN j_monster_formations_encounter_areas j ON j.encounter_area_id = ea.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs bs ON fd.boss_song_id = bs.id
ORDER BY ea.area_id;


-- name: GetAreaIDsWithShops :many
SELECT DISTINCT area_id FROM shops ORDER BY area_id;


-- name: GetAreaIDsWithTreasures :many
SELECT DISTINCT area_id FROM treasures ORDER BY area_id;


-- name: GetAreaIDsWithItemFromMonster :many
WITH target_monster_ids AS (
    SELECT mi.monster_id
    FROM monster_items mi
    JOIN item_amounts ia ON ia.id IN (
      mi.steal_common_id, mi.steal_rare_id, mi.drop_common_id, mi.drop_rare_id, mi.secondary_drop_common_id, mi.secondary_drop_rare_id, mi.bribe_id
    )
    JOIN items i ON i.master_item_id = ia.master_item_id
    WHERE i.id = sqlc.arg('item_id')::int

    UNION

    SELECT mi.monster_id
    FROM monster_items mi
    JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
    JOIN possible_items pi ON jmio.possible_item_id = pi.id
    JOIN item_amounts ia ON pi.item_amount_id = ia.id
    JOIN items i ON i.master_item_id = ia.master_item_id
    WHERE i.id = sqlc.arg('item_id')::int
)
SELECT DISTINCT ea.area_id
FROM target_monster_ids tmi
JOIN monster_amounts ma ON ma.monster_id = tmi.monster_id
JOIN j_monster_selections_monsters jmsm ON jmsm.monster_amount_id = ma.id
JOIN monster_formations mf ON mf.monster_selection_id = jmsm.monster_selection_id
JOIN j_monster_formations_encounter_areas jme ON jme.monster_formation_id = mf.id
JOIN encounter_areas ea ON jme.encounter_area_id = ea.id
ORDER BY ea.area_id;


-- name: GetAreaIDsWithItemFromShop :many
SELECT DISTINCT sh.area_id
FROM shops sh
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
WHERE si.item_id = $1
ORDER BY sh.area_id;


-- name: GetAreaIDsWithItemFromTreasure :many
SELECT DISTINCT t.area_id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY t.area_id;


-- name: GetAreaIDsWithItemFromQuest :many
SELECT DISTINCT ca.area_id
FROM completion_areas ca
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY ca.area_id;


-- name: GetAreaIDsWithKeyItemFromTreasure :many
SELECT DISTINCT t.area_id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY t.area_id;


-- name: GetAreaIDsWithKeyItemFromQuest :many
SELECT DISTINCT ca.area_id
FROM completion_areas ca
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY ca.area_id;


-- name: GetAreaIDsWithSidequests :many
SELECT DISTINCT area_id FROM completion_areas ORDER BY area_id;


-- name: GetAreaIDsWithFMVs :many
SELECT DISTINCT area_id FROM fmvs ORDER BY area_id;







-- name: GetSublocationIDs :many
SELECT id FROM sublocations ORDER BY id;


-- name: GetSublocationAreaIDs :many
SELECT DISTINCT id FROM areas WHERE sublocation_id = $1 ORDER BY id;


-- name: GetConnectedSublocationIDs :many
SELECT DISTINCT ca.sublocation_id
FROM areas a
JOIN j_area_connected_areas j ON j.area_id = a.id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN areas ca ON ac.area_id = ca.id
WHERE a.sublocation_id = $1 AND ca.sublocation_id != $1
ORDER BY ca.sublocation_id;


-- name: GetSublocationCharacterIDs :many
SELECT c.id
FROM characters c
JOIN areas a ON c.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY c.id;


-- name: GetSublocationAeonIDs :many
SELECT ae.id
FROM aeons ae
JOIN areas a ON ae.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY ae.id;


-- name: GetSublocationShopIDs :many
SELECT sh.id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY sh.id;


-- name: GetSublocationShopIdPairs :many
SELECT DISTINCT
  s.id AS sublocation_id,
  sh.id AS shop_id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = ANY(sqlc.arg('sublocation_ids')::int[])
ORDER BY s.id, sh.id;


-- name: GetSublocationTreasureIDs :many
SELECT t.id
FROM treasures t
JOIN areas a ON t.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY t.id;


-- name: GetSublocationTreasureIdPairs :many
SELECT DISTINCT
  s.id AS sublocation_id,
  t.id AS treasure_id
FROM treasures t
JOIN areas a ON t.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = ANY(sqlc.arg('sublocation_ids')::int[])
ORDER BY s.id, t.id;


-- name: GetSublocationMonsterIDs :many
SELECT DISTINCT ma.monster_id
FROM monster_amounts ma
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_formations mf ON mf.monster_selection_id = j1.monster_selection_id
JOIN j_monster_formations_encounter_areas j2 ON j2.monster_formation_id = mf.id
JOIN encounter_areas ea ON j2.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY ma.monster_id;


-- name: GetSublocationMonsterIdPairs :many
SELECT DISTINCT
  s.id AS sublocation_id,
  m.id AS monster_id
FROM monsters m
JOIN monster_amounts ma ON ma.monster_id = m.id
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_selections ms ON j1.monster_selection_id = ms.id
JOIN monster_formations mf ON mf.monster_selection_id = ms.id
JOIN j_monster_formations_encounter_areas j2 ON j2.monster_formation_id = mf.id
JOIN encounter_areas ea ON j2.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = ANY(sqlc.arg('sublocation_ids')::int[])
ORDER BY s.id, m.id;


-- name: GetSublocationMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY mf.id;


-- name: GetSublocationBossSongIDs :many
SELECT DISTINCT bs.song_id
FROM formation_boss_songs bs
JOIN formation_data fd ON fd.boss_song_id = bs.id
JOIN monster_formations mf ON mf.formation_data_id = fd.id
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY bs.song_id;


-- name: GetSublocationCueSongIDs :many
SELECT c.song_id
FROM cues c
JOIN areas a ON c.trigger_area_id = a.id
WHERE a.sublocation_id = $1

UNION

SELECT c.song_id
FROM cues c
JOIN j_cues_areas j ON j.cue_id = c.id
JOIN areas a ON j.included_area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY song_id;


-- name: GetSublocationBackgroundMusicSongIDs :many
SELECT DISTINCT j.song_id
FROM j_songs_background_music j
JOIN areas a ON j.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY j.song_id;


-- name: GetSublocationFMVSongIDs :many
SELECT DISTINCT f.song_id::int
FROM fmvs f
JOIN areas a ON f.area_id = a.id
WHERE f.song_id IS NOT NULL
  AND a.sublocation_id = $1
ORDER BY f.song_id;


-- name: GetSublocationFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN areas a ON f.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY f.id;


-- name: GetSublocationQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN completion_areas ca ON ca.completion_id = q.completion_id
JOIN areas a ON ca.area_id = a.id
WHERE a.sublocation_id = $1
ORDER BY q.id;


-- name: GetSublocationIDsWithCharacters :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN characters c ON c.area_id = a.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithAeons :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN aeons ae ON ae.area_id = a.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithMonsters :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN encounter_areas ea ON ea.area_id = a.id
JOIN j_monster_formations_encounter_areas j1 ON j1.encounter_area_id = ea.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_selections_monsters j2 ON mf.monster_selection_id = j2.monster_selection_id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithBosses :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN encounter_areas ea ON ea.area_id = a.id
JOIN j_monster_formations_encounter_areas j ON j.encounter_area_id = ea.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs bs ON fd.boss_song_id = bs.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithShops :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN shops sh ON sh.area_id = a.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithTreasures :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN treasures t ON t.area_id = a.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithItemFromMonster :many
WITH target_monster_ids AS (
    SELECT mi.monster_id
    FROM monster_items mi
    JOIN item_amounts ia ON ia.id IN (
      mi.steal_common_id, mi.steal_rare_id, mi.drop_common_id, mi.drop_rare_id, mi.secondary_drop_common_id, mi.secondary_drop_rare_id, mi.bribe_id
    )
    JOIN items i ON i.master_item_id = ia.master_item_id
    WHERE i.id = sqlc.arg('item_id')::int

    UNION

    SELECT mi.monster_id
    FROM monster_items mi
    JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
    JOIN possible_items pi ON jmio.possible_item_id = pi.id
    JOIN item_amounts ia ON pi.item_amount_id = ia.id
    JOIN items i ON i.master_item_id = ia.master_item_id
    WHERE i.id = sqlc.arg('item_id')::int
)
SELECT DISTINCT a.sublocation_id
FROM target_monster_ids tmi
JOIN monster_amounts ma ON ma.monster_id = tmi.monster_id
JOIN j_monster_selections_monsters jmsm ON jmsm.monster_amount_id = ma.id
JOIN monster_formations mf ON mf.monster_selection_id = jmsm.monster_selection_id
JOIN j_monster_formations_encounter_areas jme ON jme.monster_formation_id = mf.id
JOIN encounter_areas ea ON jme.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithItemFromShop :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN shops sh ON sh.area_id = a.id
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
WHERE si.item_id = $1
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithItemFromTreasure :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithItemFromQuest :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN completion_areas ca ON ca.area_id = a.id
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithKeyItemFromTreasure :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithKeyItemFromQuest :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN completion_areas ca ON ca.area_id = a.id
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithSidequests :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN completion_areas ca ON ca.area_id = a.id
ORDER BY a.sublocation_id;


-- name: GetSublocationIDsWithFMVs :many
SELECT DISTINCT a.sublocation_id
FROM areas a
JOIN fmvs f ON f.area_id = a.id
ORDER BY a.sublocation_id;






-- name: GetLocationIDs :many
SELECT id FROM locations ORDER BY id;


-- name: GetLocationSublocationIDs :many
SELECT DISTINCT id FROM sublocations WHERE location_id = $1 ORDER BY id;


-- name: GetLocationAreaIDs :many
SELECT DISTINCT a.id
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY a.id;


-- name: GetConnectedLocationIDs :many
SELECT DISTINCT cs.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN j_area_connected_areas j ON j.area_id = a.id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN areas ca ON ac.area_id = ca.id
JOIN sublocations cs ON ca.sublocation_id = cs.id
WHERE s.location_id = $1 AND cs.location_id != $1
ORDER BY cs.location_id;


-- name: GetLocationCharacterIDs :many
SELECT c.id
FROM characters c
JOIN areas a ON c.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY c.id;


-- name: GetLocationAeonIDs :many
SELECT ae.id
FROM aeons ae
JOIN areas a ON ae.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY ae.id;


-- name: GetLocationShopIDs :many
SELECT sh.id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY sh.id;


-- name: GetLocationShopIdPairs :many
SELECT DISTINCT
  l.id AS location_id,
  sh.id AS shop_id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = ANY(sqlc.arg('location_ids')::int[])
ORDER BY l.id, sh.id;


-- name: GetLocationTreasureIDs :many
SELECT t.id
FROM treasures t
JOIN areas a ON t.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY t.id;


-- name: GetLocationTreasureIdPairs :many
SELECT DISTINCT
  l.id AS location_id,
  t.id AS treasure_id
FROM treasures t
JOIN areas a ON t.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = ANY(sqlc.arg('location_ids')::int[])
ORDER BY l.id, t.id;


-- name: GetLocationMonsterIDs :many
SELECT DISTINCT ma.monster_id
FROM monster_amounts ma
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_formations mf ON mf.monster_selection_id = j1.monster_selection_id
JOIN j_monster_formations_encounter_areas j2 ON j2.monster_formation_id = mf.id
JOIN encounter_areas ea ON j2.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY ma.monster_id;


-- name: GetLocationMonsterIdPairs :many
SELECT DISTINCT
  l.id AS location_id,
  m.id AS monster_id
FROM monsters m
JOIN monster_amounts ma ON ma.monster_id = m.id
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_selections ms ON j1.monster_selection_id = ms.id
JOIN monster_formations mf ON mf.monster_selection_id = ms.id
JOIN j_monster_formations_encounter_areas j2 ON j2.monster_formation_id = mf.id
JOIN encounter_areas ea ON j2.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = ANY(sqlc.arg('location_ids')::int[])
ORDER BY l.id, m.id;


-- name: GetLocationMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY mf.id;


-- name: GetLocationBossSongIDs :many
SELECT DISTINCT bs.song_id
FROM formation_boss_songs bs
JOIN formation_data fd ON fd.boss_song_id = bs.id
JOIN monster_formations mf ON mf.formation_data_id = fd.id
JOIN j_monster_formations_encounter_areas j ON j.monster_formation_id = mf.id
JOIN encounter_areas ea ON j.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY bs.song_id;


-- name: GetLocationCueSongIDs :many
SELECT c.song_id
FROM cues c
JOIN areas a ON c.trigger_area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1

UNION

SELECT c.song_id
FROM cues c
JOIN j_cues_areas j ON j.cue_id = c.id
JOIN areas a ON j.included_area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY song_id;


-- name: GetLocationBackgroundMusicSongIDs :many
SELECT DISTINCT j.song_id
FROM j_songs_background_music j
JOIN areas a ON j.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY j.song_id;


-- name: GetLocationFMVSongIDs :many
SELECT DISTINCT f.song_id::int
FROM fmvs f
JOIN areas a ON f.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE f.song_id IS NOT NULL
  AND s.location_id = $1
ORDER BY f.song_id;


-- name: GetLocationFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN areas a ON f.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY f.id;


-- name: GetLocationQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN completion_areas ca ON ca.completion_id = q.completion_id
JOIN areas a ON ca.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.location_id = $1
ORDER BY q.id;


-- name: GetLocationIDsWithCharacters :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN characters c ON c.area_id = a.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithAeons :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN aeons ae ON ae.area_id = a.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithMonsters :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_areas ea ON ea.area_id = a.id
JOIN j_monster_formations_encounter_areas j1 ON j1.encounter_area_id = ea.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_selections_monsters j2 ON mf.monster_selection_id = j2.monster_selection_id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithBosses :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_areas ea ON ea.area_id = a.id
JOIN j_monster_formations_encounter_areas j ON j.encounter_area_id = ea.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs bs ON fd.boss_song_id = bs.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithShops :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN shops sh ON sh.area_id = a.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithTreasures :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithItemFromMonster :many
WITH target_monster_ids AS (
    SELECT mi.monster_id
    FROM monster_items mi
    JOIN item_amounts ia ON ia.id IN (
      mi.steal_common_id, mi.steal_rare_id, mi.drop_common_id, mi.drop_rare_id, mi.secondary_drop_common_id, mi.secondary_drop_rare_id, mi.bribe_id
    )
    JOIN items i ON i.master_item_id = ia.master_item_id
    WHERE i.id = sqlc.arg('item_id')::int

    UNION

    SELECT mi.monster_id
    FROM monster_items mi
    JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
    JOIN possible_items pi ON jmio.possible_item_id = pi.id
    JOIN item_amounts ia ON pi.item_amount_id = ia.id
    JOIN items i ON i.master_item_id = ia.master_item_id
    WHERE i.id = sqlc.arg('item_id')::int
)
SELECT DISTINCT s.location_id
FROM target_monster_ids tmi
JOIN monster_amounts ma ON ma.monster_id = tmi.monster_id
JOIN j_monster_selections_monsters jmsm ON jmsm.monster_amount_id = ma.id
JOIN monster_formations mf ON mf.monster_selection_id = jmsm.monster_selection_id
JOIN j_monster_formations_encounter_areas jme ON jme.monster_formation_id = mf.id
JOIN encounter_areas ea ON jme.encounter_area_id = ea.id
JOIN areas a ON ea.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithItemFromShop :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN shops sh ON sh.area_id = a.id
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
WHERE si.item_id = $1
ORDER BY s.location_id;


-- name: GetLocationIDsWithItemFromTreasure :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY s.location_id;


-- name: GetLocationIDsWithItemFromQuest :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_areas ca ON ca.area_id = a.id
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN items i ON i.master_item_id = ia.master_item_id
WHERE i.id = $1
ORDER BY s.location_id;


-- name: GetLocationIDsWithKeyItemFromTreasure :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY s.location_id;


-- name: GetLocationIDsWithKeyItemFromQuest :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_areas ca ON ca.area_id = a.id
JOIN quest_completions qc ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN key_items ki ON ki.master_item_id = ia.master_item_id
WHERE ki.id = $1
ORDER BY s.location_id;


-- name: GetLocationIDsWithSidequests :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_areas ca ON ca.area_id = a.id
ORDER BY s.location_id;


-- name: GetLocationIDsWithFMVs :many
SELECT DISTINCT s.location_id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN fmvs f ON f.area_id = a.id
ORDER BY s.location_id;






-- name: GetTreasureIDs :many
SELECT id FROM treasures ORDER BY id;


-- name: GetTreasureIDsByLootType :many
SELECT id FROM treasures WHERE loot_type = $1 ORDER BY id;


-- name: GetTreasureIDsByTreasureType :many
SELECT id FROM treasures WHERE treasure_type = $1 ORDER BY id;


-- name: GetTreasureIDsByItem :many
SELECT DISTINCT j.treasure_id
FROM j_treasures_items j
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id
WHERE i.id = $1
ORDER BY j.treasure_id;


-- name: GetTreasureIDsByIsAnimaTreasure :many
SELECT id FROM treasures WHERE is_anima_treasure = $1 ORDER BY id;


-- name: GetTreasureIDsByAvailability :many
SELECT id FROM treasures WHERE availability = ANY(sqlc.narg('availability')::availability_type[]) ORDER BY id;






-- name: GetShopIDs :many
SELECT id FROM shops ORDER BY id;


-- name: GetShopIDsByCategory :many
SELECT id FROM shops WHERE category = ANY(sqlc.narg('category')::shop_category[]) ORDER BY id;


-- name: GetShopIDsEquipmentFilter :many
SELECT DISTINCT se.shop_id
FROM shop_equipment_pieces se
LEFT JOIN equipment_names en ON se.equipment_name_id = en.id
LEFT JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
CROSS JOIN (
    SELECT
        sqlc.narg('shop_type')::shop_type AS shop_type,
        sqlc.narg('empty_slots')::int[] AS empty_slots,
        sqlc.narg('character_id')::int AS character_id,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id
) w
WHERE 
    (w.shop_type IS NULL OR se.shop_type = w.shop_type)
    AND (w.empty_slots IS NULL OR se.empty_slots_amount::int = ANY(w.empty_slots))
    AND (w.character_id IS NULL OR en.character_id = w.character_id)
    AND (w.auto_ability_id IS NULL OR j.auto_ability_id = w.auto_ability_id)
ORDER BY se.shop_id;


-- name: GetShopIDsWithItems :many
SELECT DISTINCT shop_id FROM j_shops_items ORDER BY shop_id;


-- name: GetShopIDsWithEquipment :many
SELECT DISTINCT shop_id FROM shop_equipment_pieces ORDER BY shop_id;


-- name: GetShopIDsByAvailability :many
SELECT id FROM shops WHERE availability = ANY(sqlc.narg('availability')::availability_type[]) ORDER BY id;
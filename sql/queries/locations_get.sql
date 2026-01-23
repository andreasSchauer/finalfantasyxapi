-- name: GetAreaIDs :many
SELECT id FROM areas ORDER BY id;


-- name: GetAreaConnectionIDs :many
SELECT a.id
FROM area_connections ac
JOIN j_area_connected_areas j ON j.connection_id = ac.id
JOIN areas a ON ac.area_id = a.id
JOIN areas a2 ON j.area_id = a2.id
WHERE a2.id = $1
ORDER BY a.id;


-- name: GetAreaCharacterIDs :many
SELECT c.id
FROM areas a
JOIN characters c ON c.area_id = a.id
JOIN player_units pu ON c.unit_id = pu.id
WHERE a.id = $1
ORDER BY c.id;


-- name: GetAreaAeonIDs :many
SELECT ae.id
FROM areas a
JOIN aeons ae ON ae.area_id = a.id
JOIN player_units pu ON ae.unit_id = pu.id
WHERE a.id = $1
ORDER BY ae.id;


-- name: GetAreaShopIDs :many
SELECT id FROM shops WHERE area_id = $1 ORDER BY id;


-- name: GetAreaTreasureIDs :many
SELECT id FROM treasures WHERE area_id = $1 ORDER BY id;


-- name: GetAreaMonsterIDs :many
SELECT DISTINCT m.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
WHERE a.id = $1
ORDER BY m.id;


-- name: GetAreaMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
WHERE a.id = $1
ORDER BY mf.id;


-- name: GetAreaBossSongIDs :many
SELECT DISTINCT so.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_boss_songs bs ON mf.boss_song_id = bs.id
JOIN songs so ON bs.song_id = so.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaCues :many
SELECT DISTINCT so.id, c.replaces_encounter_music
FROM cues c
JOIN songs so ON c.song_id = so.id
JOIN j_songs_cues j ON j.cue_id = c.id
JOIN areas a ON COALESCE(c.trigger_area_id, j.included_area_id) = a.id
WHERE j.included_area_id = $1 OR c.trigger_area_id = $1
ORDER BY so.id;


-- name: GetAreaBackgroundMusic :many
SELECT so.id, bm.replaces_encounter_music
FROM areas a
JOIN j_songs_background_music j ON j.area_id = a.id 
JOIN background_music bm ON j.bm_id = bm.id
JOIN songs so ON j.song_id = so.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFMVSongIDs :many
SELECT DISTINCT so.id
FROM areas a
JOIN fmvs f ON f.area_id = a.id
JOIN songs so ON f.song_id = so.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFmvIDs :many
SELECT f.id
FROM areas a
JOIN fmvs f ON f.area_id = a.id
WHERE a.id = $1
ORDER BY f.id;


-- name: GetAreaQuestIDs :many
SELECT DISTINCT q.id
FROM areas a
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
WHERE a.id = $1
ORDER BY q.id;


-- name: GetAreaIDsWithSaveSphere :many
SELECT id FROM areas WHERE has_save_sphere = $1 ORDER BY id;


-- name: GetAreaIDsWithCompSphere :many
SELECT id FROM areas WHERE has_compilation_sphere = $1 ORDER BY id;


-- name: GetAreaIDsWithDropOff :many
SELECT id FROM areas WHERE airship_drop_off = $1 ORDER BY id;


-- name: GetAreaIDsChocobo :many
SELECT id FROM areas WHERE can_ride_chocobo = $1 ORDER BY id;


-- name: GetAreaIDsStoryOnly :many
SELECT id FROM areas WHERE story_only = $1 ORDER BY id;


-- name: GetAreaIDsWithCharacters :many
SELECT a.id
FROM areas a
JOIN characters c ON c.area_id = a.id
JOIN player_units pu ON c.unit_id = pu.id
ORDER BY a.id;


-- name: GetAreaIDsWithAeons :many
SELECT a.id
FROM areas a
JOIN aeons ae ON ae.area_id = a.id
JOIN player_units pu ON ae.unit_id = pu.id
ORDER BY a.id;


-- name: GetAreaIDsWithMonsters :many
SELECT DISTINCT a.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
ORDER BY a.id;


-- name: GetAreaIDsWithBosses :many
SELECT DISTINCT a.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_boss_songs bs ON mf.boss_song_id = bs.id
JOIN songs so ON bs.song_id = so.id
ORDER BY a.id;


-- name: GetAreaIDsWithShops :many
SELECT DISTINCT a.id
FROM areas a
JOIN shops sh ON sh.area_id = a.id
ORDER BY a.id;


-- name: GetAreaIDsWithTreasures :many
SELECT DISTINCT a.id
FROM areas a
JOIN treasures t ON t.area_id = a.id
ORDER BY a.id;


-- name: GetAreaIDsWithItemFromMonster :many
SELECT DISTINCT a.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
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
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreaIDsWithItemFromShop :many
SELECT DISTINCT a.id
FROM areas a
JOIN shops sh ON sh.area_id = a.id
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreaIDsWithItemFromTreasure :many
SELECT DISTINCT a.id
FROM areas a
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreaIDsWithItemFromQuest :many
SELECT DISTINCT a.id
FROM areas a
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreaIDsWithKeyItemFromTreasure :many
SELECT DISTINCT a.id
FROM areas a
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY a.id;


-- name: GetAreaIDsWithKeyItemFromQuest :many
SELECT DISTINCT a.id
FROM areas a
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY a.id;


-- name: GetAreaIDsWithSidequests :many
SELECT DISTINCT a.id
FROM areas a
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
ORDER BY a.id;


-- name: GetAreaIDsWithFMVs :many
SELECT DISTINCT a.id
FROM areas a
JOIN fmvs f ON f.area_id = a.id
ORDER BY a.id;










-- name: GetSublocationAreaIDs :many
SELECT a.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY a.id;


-- name: GetSublocationConnectionIDs :many
SELECT DISTINCT s.id, s.name
FROM area_connections ac
JOIN j_area_connected_areas j ON j.connection_id = ac.id
JOIN areas a ON ac.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN areas a2 ON j.area_id = a2.id
JOIN sublocations s2 ON a2.sublocation_id = s2.id
WHERE s2.id = $1 AND s2.id != s.id
ORDER BY s.id;


-- name: GetSublocationMonsterIDs :many
SELECT DISTINCT m.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
WHERE s.id = $1
ORDER BY m.id;











-- name: GetLocationAreaIDs :many
SELECT a.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id 
WHERE l.id = $1
ORDER BY a.id;


-- name: GetLocationConnectionIDs :many
SELECT DISTINCT l.id
FROM area_connections ac
JOIN j_area_connected_areas j ON j.connection_id = ac.id
JOIN areas a ON ac.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN areas a2 ON j.area_id = a2.id
JOIN sublocations s2 ON a2.sublocation_id = s2.id
JOIN locations l2 ON s2.location_id = l2.id
WHERE l2.id = $1 AND l2.id != l.id
ORDER BY l.id;


-- name: GetLocationMonsterIDs :many
SELECT DISTINCT m.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
WHERE l.id = $1
ORDER BY m.id;
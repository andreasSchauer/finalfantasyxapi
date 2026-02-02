-- name: GetAreaIDs :many
SELECT id FROM areas ORDER BY id;


-- name: GetAreaConnectionIDs :many
SELECT ca.id
FROM areas a
JOIN j_area_connected_areas j ON j.area_id = a.id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN areas ca ON ac.area_id = ca.id
WHERE a.id = $1
ORDER BY ca.id;


-- name: GetAreaCharacterIDs :many
SELECT c.id
FROM characters c
JOIN areas a ON c.area_id = a.id
WHERE a.id = $1
ORDER BY c.id;


-- name: GetAreaAeonIDs :many
SELECT ae.id
FROM aeons ae
JOIN areas a ON ae.area_id = a.id
WHERE a.id = $1
ORDER BY ae.id;


-- name: GetAreaShopIDs :many
SELECT id FROM shops WHERE area_id = $1 ORDER BY id;


-- name: GetAreaTreasureIDs :many
SELECT id FROM treasures WHERE area_id = $1 ORDER BY id;


-- name: GetAreaMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_amounts ma ON ma.monster_id = m.id
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_selections ms ON j1.monster_selection_id = ms.id
JOIN monster_formations mf ON mf.monster_selection_id = ms.id
JOIN j_monster_formations_encounter_locations j2 ON j2.monster_formation_id = mf.id
JOIN encounter_locations el ON j2.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
WHERE a.id = $1
ORDER BY m.id;


-- name: GetAreaMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_locations j ON j.monster_formation_id = mf.id
JOIN encounter_locations el ON j.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
WHERE a.id = $1
ORDER BY mf.id;


-- name: GetAreaBossSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN formation_boss_songs bs ON bs.song_id = so.id
JOIN formation_data fd ON fd.boss_song_id = bs.id
JOIN monster_formations mf ON mf.formation_data_id = fd.id
JOIN j_monster_formations_encounter_locations j ON j.monster_formation_id = mf.id
JOIN encounter_locations el ON j.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaCueSongIDs :many
SELECT DISTINCT so.id
FROM cues c
JOIN songs so ON c.song_id = so.id
JOIN j_songs_cues j ON j.cue_id = c.id
JOIN areas a ON COALESCE(c.trigger_area_id, j.included_area_id) = a.id
WHERE j.included_area_id = $1 OR c.trigger_area_id = $1
ORDER BY so.id;


-- name: GetAreaBackgroundMusicSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN j_songs_background_music j ON j.song_id = so.id
JOIN background_music bm ON j.bm_id = bm.id
JOIN areas a ON j.area_id = a.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFMVSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN fmvs f ON f.song_id = so.id
JOIN areas a ON f.area_id = a.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN areas a ON f.area_id = a.id
WHERE a.id = $1
ORDER BY f.id;


-- name: GetAreaQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_locations cl ON cl.completion_id = qc.id
JOIN areas a ON cl.area_id = a.id
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
SELECT DISTINCT a.id
FROM areas a
JOIN characters c ON c.area_id = a.id
ORDER BY a.id;


-- name: GetAreaIDsWithAeons :many
SELECT DISTINCT a.id
FROM areas a
JOIN aeons ae ON ae.area_id = a.id
ORDER BY a.id;


-- name: GetAreaIDsWithMonsters :many
SELECT DISTINCT a.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
ORDER BY a.id;


-- name: GetAreaIDsWithBosses :many
SELECT DISTINCT a.id
FROM areas a
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs bs ON fd.boss_song_id = bs.id
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
JOIN j_monster_formations_encounter_locations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
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









-- name: GetSublocationIDs :many
SELECT id FROM sublocations ORDER BY id;


-- name: GetSublocationAreaIDs :many
SELECT DISTINCT a.id
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY a.id;


-- name: GetConnectedSublocationIDs :many
SELECT DISTINCT cs.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN j_area_connected_areas j ON j.area_id = a.id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN areas ca ON ac.area_id = ca.id
JOIN sublocations cs ON ca.sublocation_id = cs.id
WHERE s.id = $1 AND s.id != cs.id
ORDER BY cs.id;


-- name: GetSublocationCharacterIDs :many
SELECT c.id
FROM characters c
JOIN areas a ON c.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY c.id;


-- name: GetSublocationAeonIDs :many
SELECT ae.id
FROM aeons ae
JOIN areas a ON ae.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY ae.id;


-- name: GetSublocationShopIDs :many
SELECT sh.id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY sh.id;


-- name: GetSublocationTreasureIDs :many
SELECT t.id
FROM treasures t
JOIN areas a ON t.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY t.id;


-- name: GetSublocationMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_amounts ma ON ma.monster_id = m.id
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_selections ms ON j1.monster_selection_id = ms.id
JOIN monster_formations mf ON mf.monster_selection_id = ms.id
JOIN j_monster_formations_encounter_locations j2 ON j2.monster_formation_id = mf.id
JOIN encounter_locations el ON j2.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY m.id;


-- name: GetSublocationMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_locations j ON j.monster_formation_id = mf.id
JOIN encounter_locations el ON j.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY mf.id;


-- name: GetSublocationBossSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN formation_boss_songs bs ON bs.song_id = so.id
JOIN formation_data fd ON fd.boss_song_id = bs.id
JOIN monster_formations mf ON mf.formation_data_id = fd.id
JOIN j_monster_formations_encounter_locations j ON j.monster_formation_id = mf.id
JOIN encounter_locations el ON j.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY so.id;


-- name: GetSublocationCueSongIDs :many
SELECT DISTINCT so.id
FROM cues c
JOIN songs so ON c.song_id = so.id
JOIN j_songs_cues j ON j.cue_id = c.id
JOIN areas a ON COALESCE(c.trigger_area_id, j.included_area_id) = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY so.id;


-- name: GetSublocationBackgroundMusicSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN j_songs_background_music j ON j.song_id = so.id
JOIN background_music bm ON j.bm_id = bm.id
JOIN areas a ON j.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY so.id;


-- name: GetSublocationFMVSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN fmvs f ON f.song_id = so.id
JOIN areas a ON f.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY so.id;


-- name: GetSublocationFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN areas a ON f.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY f.id;


-- name: GetSublocationQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_locations cl ON cl.completion_id = qc.id
JOIN areas a ON cl.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY q.id;


-- name: GetSublocationIDsWithCharacters :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN characters c ON c.area_id = a.id
ORDER BY s.id;


-- name: GetSublocationIDsWithAeons :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN aeons ae ON ae.area_id = a.id
ORDER BY s.id;


-- name: GetSublocationIDsWithMonsters :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
ORDER BY s.id;


-- name: GetSublocationIDsWithBosses :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs bs ON fd.boss_song_id = bs.id
JOIN songs so ON bs.song_id = so.id
ORDER BY s.id;


-- name: GetSublocationIDsWithShops :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN shops sh ON sh.area_id = a.id
ORDER BY s.id;


-- name: GetSublocationIDsWithTreasures :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
ORDER BY s.id;


-- name: GetSublocationIDsWithItemFromMonster :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
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
ORDER BY s.id;


-- name: GetSublocationIDsWithItemFromShop :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN shops sh ON sh.area_id = a.id
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
WHERE i.id = $1
ORDER BY s.id;


-- name: GetSublocationIDsWithItemFromTreasure :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY s.id;


-- name: GetSublocationIDsWithItemFromQuest :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY s.id;


-- name: GetSublocationIDsWithKeyItemFromTreasure :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY s.id;


-- name: GetSublocationIDsWithKeyItemFromQuest :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY s.id;


-- name: GetSublocationIDsWithSidequests :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
ORDER BY s.id;


-- name: GetSublocationIDsWithFMVs :many
SELECT DISTINCT s.id
FROM sublocations s
JOIN areas a ON a.sublocation_id = s.id
JOIN fmvs f ON f.area_id = a.id
ORDER BY s.id;






-- name: GetLocationIDs :many
SELECT id FROM locations ORDER BY id;


-- name: GetLocationSublocationIDs :many
SELECT s.id
FROM sublocations s
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY s.id;


-- name: GetLocationAreaIDs :many
SELECT DISTINCT a.id
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY a.id;


-- name: GetConnectedLocationIDs :many
SELECT DISTINCT cl.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN j_area_connected_areas j ON j.area_id = a.id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN areas ca ON ac.area_id = ca.id
JOIN sublocations cs ON ca.sublocation_id = cs.id
JOIN locations cl ON cs.location_id = cl.id
WHERE l.id = $1 AND l.id != cl.id
ORDER BY cl.id;


-- name: GetLocationCharacterIDs :many
SELECT c.id
FROM characters c
JOIN areas a ON c.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY c.id;


-- name: GetLocationAeonIDs :many
SELECT ae.id
FROM aeons ae
JOIN areas a ON ae.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY ae.id;


-- name: GetLocationShopIDs :many
SELECT sh.id
FROM shops sh
JOIN areas a ON sh.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY sh.id;


-- name: GetLocationTreasureIDs :many
SELECT t.id
FROM treasures t
JOIN areas a ON t.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY t.id;


-- name: GetLocationMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_amounts ma ON ma.monster_id = m.id
JOIN j_monster_selections_monsters j1 ON j1.monster_amount_id = ma.id
JOIN monster_selections ms ON j1.monster_selection_id = ms.id
JOIN monster_formations mf ON mf.monster_selection_id = ms.id
JOIN j_monster_formations_encounter_locations j2 ON j2.monster_formation_id = mf.id
JOIN encounter_locations el ON j2.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE location_id = $1
ORDER BY m.id;


-- name: GetLocationMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN j_monster_formations_encounter_locations j ON j.monster_formation_id = mf.id
JOIN encounter_locations el ON j.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY mf.id;


-- name: GetLocationBossSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN formation_boss_songs bs ON bs.song_id = so.id
JOIN formation_data fd ON fd.boss_song_id = bs.id
JOIN monster_formations mf ON mf.formation_data_id = fd.id
JOIN j_monster_formations_encounter_locations j ON j.monster_formation_id = mf.id
JOIN encounter_locations el ON j.encounter_location_id = el.id
JOIN areas a ON el.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY so.id;


-- name: GetLocationCueSongIDs :many
SELECT DISTINCT so.id
FROM cues c
JOIN songs so ON c.song_id = so.id
JOIN j_songs_cues j ON j.cue_id = c.id
JOIN areas a ON COALESCE(c.trigger_area_id, j.included_area_id) = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY so.id;


-- name: GetLocationBackgroundMusicSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN j_songs_background_music j ON j.song_id = so.id
JOIN background_music bm ON j.bm_id = bm.id
JOIN areas a ON j.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY so.id;


-- name: GetLocationFMVSongIDs :many
SELECT DISTINCT so.id
FROM songs so
JOIN fmvs f ON f.song_id = so.id
JOIN areas a ON f.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY so.id;


-- name: GetLocationFmvIDs :many
SELECT f.id
FROM fmvs f
JOIN areas a ON f.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY f.id;


-- name: GetLocationQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_locations cl ON cl.completion_id = qc.id
JOIN areas a ON cl.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY q.id;


-- name: GetLocationIDsWithCharacters :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN characters c ON c.area_id = a.id
ORDER BY l.id;


-- name: GetLocationIDsWithAeons :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN aeons ae ON ae.area_id = a.id
ORDER BY l.id;


-- name: GetLocationIDsWithMonsters :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
ORDER BY l.id;


-- name: GetLocationIDsWithBosses :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs bs ON fd.boss_song_id = bs.id
JOIN songs so ON bs.song_id = so.id
ORDER BY l.id;


-- name: GetLocationIDsWithShops :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN shops sh ON sh.area_id = a.id
ORDER BY l.id;


-- name: GetLocationIDsWithTreasures :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
ORDER BY l.id;


-- name: GetLocationIDsWithItemFromMonster :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_monster_formations_encounter_locations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN monster_selections ms ON mf.monster_selection_id = ms.id
JOIN j_monster_selections_monsters j2 ON j2.monster_selection_id = ms.id
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
ORDER BY l.id;


-- name: GetLocationIDsWithItemFromShop :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN shops sh ON sh.area_id = a.id
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
WHERE i.id = $1
ORDER BY l.id;


-- name: GetLocationIDsWithItemFromTreasure :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY l.id;


-- name: GetLocationIDsWithItemFromQuest :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY l.id;


-- name: GetLocationIDsWithKeyItemFromTreasure :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY l.id;


-- name: GetLocationIDsWithKeyItemFromQuest :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY l.id;


-- name: GetLocationIDsWithSidequests :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
ORDER BY l.id;


-- name: GetLocationIDsWithFMVs :many
SELECT DISTINCT l.id
FROM locations l
JOIN sublocations s ON s.location_id = l.id
JOIN areas a ON a.sublocation_id = s.id
JOIN fmvs f ON f.area_id = a.id
ORDER BY l.id;


-- name: GetTreasureIDs :many
SELECT id FROM treasures ORDER BY id;


-- name: GetTreasureIDsByLootType :many
SELECT id FROM treasures WHERE loot_type = $1 ORDER BY id;


-- name: GetTreasureIDsByTreasureType :many
SELECT id FROM treasures WHERE treasure_type = $1 ORDER BY id;


-- name: GetTreasureIDsByIsAnimaTreasure :many
SELECT id FROM treasures WHERE is_anima_treasure = $1 ORDER BY id;


-- name: GetTreasureIDsByIsPostAirship :many
SELECT id FROM treasures WHERE is_post_airship = $1 ORDER BY id;
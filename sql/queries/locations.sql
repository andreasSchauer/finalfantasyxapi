-- name: GetLocationAreaByAreaName :one
SELECT l.name, s.name, a.name, a.version FROM locations l JOIN sublocations s ON s.location_id = l.id JOIN areas a ON a.sublocation_id = a.id
WHERE l.id = $1 AND s.id = $2 AND a.name = $3 AND a.version = $4;

-- name: CreateLocation :one
INSERT INTO locations (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = locations.data_hash
RETURNING *;


-- name: CreateSubLocation :one
INSERT INTO sublocations (data_hash, location_id, name, specification)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = sublocations.data_hash
RETURNING *;


-- name: CreateArea :one
INSERT INTO areas (data_hash, sublocation_id, name, version, specification, story_only, has_save_sphere, airship_drop_off, has_compilation_sphere, can_ride_chocobo)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = areas.data_hash
RETURNING *;


-- name: CreateAreaConnection :one
INSERT INTO area_connections (data_hash, area_id, connection_type, story_only, notes)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = area_connections.data_hash
RETURNING *;


-- name: CreateAreaConnectedAreasJunction :exec
INSERT INTO j_area_connected_areas (data_hash, area_id, connection_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: GetAreas :many
SELECT
    l.id AS location_id,
    l.name AS location,
    s.id AS sublocation_id,
    s.name AS sublocation,
    a.*
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id 
JOIN locations l ON s.location_id = l.id;


-- name: GetArea :one
SELECT
    l.id AS location_id,
    l.name AS location,
    s.name AS sublocation,
    a.*
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id 
JOIN locations l ON s.location_id = l.id
WHERE a.id = $1;


-- name: GetAreaConnections :many
SELECT
    ac.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS version,
    a.specification AS specification
FROM area_connections ac
JOIN j_area_connected_areas j ON j.connection_id = ac.id
JOIN areas a ON ac.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN areas a2 ON j.area_id = a2.id
WHERE a2.id = $1
ORDER BY ac.id;


-- name: GetAreaCharacters :many
SELECT
    c.*,
    pu.name,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area, a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN characters c ON c.area_id = a.id
JOIN player_units pu ON c.unit_id = pu.id
WHERE a.id = $1
ORDER BY c.id;


-- name: GetAreaAeons :many
SELECT
    ae.*,
    pu.name,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area, a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN aeons ae ON ae.area_id = a.id
JOIN player_units pu ON ae.unit_id = pu.id
WHERE a.id = $1
ORDER BY ae.id;


-- name: GetAreaShops :many
SELECT * FROM shops WHERE area_id = $1 ORDER BY id;


-- name: GetAreaTreasures :many
SELECT * FROM treasures WHERE area_id = $1 ORDER BY id;


-- name: GetAreaMonsters :many
SELECT DISTINCT
    m.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
WHERE a.id = $1
ORDER BY m.id;


-- name: GetAreaMonsterFormations :many
SELECT DISTINCT
    mf.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
WHERE a.id = $1
ORDER BY mf.id;


-- name: GetAreaBossSongs :many
SELECT DISTINCT
    so.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_boss_songs bs ON mf.boss_song_id = bs.id
JOIN songs so ON bs.song_id = so.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaCues :many
SELECT DISTINCT
    so.*,
    c.replaces_encounter_music,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM cues c
JOIN songs so ON c.song_id = so.id
JOIN j_songs_cues j ON j.cue_id = c.id
JOIN areas a ON COALESCE(c.trigger_area_id, j.included_area_id) = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE j.included_area_id = $1 OR c.trigger_area_id = $1
ORDER BY so.id;


-- name: GetAreaBackgroundMusic :many
SELECT
    so.*,
    bm.replaces_encounter_music,
    l.name AS location,
    s.name As sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN j_songs_background_music j ON j.area_id = a.id 
JOIN background_music bm ON j.bm_id = bm.id
JOIN songs so ON j.song_id = so.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFMVSongs :many
SELECT DISTINCT
    so.*,
    l.name AS location,
    s.name As sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN fmvs f ON f.area_id = a.id
JOIN songs so ON f.song_id = so.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFMVs :many
SELECT
    f.*,
    l.name AS location,
    s.name As sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN fmvs f ON f.area_id = a.id
WHERE a.id = $1
ORDER BY f.id;


-- name: GetAreaQuests :many
SELECT DISTINCT
    q.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
WHERE a.id = $1
ORDER BY q.id;


-- name: GetAreasWithSaveSphere :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE a.has_save_sphere = $1
ORDER BY a.id;


-- name: GetAreasWithCompSphere :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE a.has_compilation_sphere = $1
ORDER BY a.id;


-- name: GetAreasWithDropOff :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE a.airship_drop_off = $1
ORDER BY a.id;


-- name: GetAreasWithChocobo :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE a.can_ride_chocobo = $1
ORDER BY a.id;


-- name: GetAreasStoryOnly :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
WHERE a.story_only = $1
ORDER BY a.id;


-- name: GetAreasWithCharacters :many
SELECT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN characters c ON c.area_id = a.id
JOIN player_units pu ON c.unit_id = pu.id
ORDER BY a.id;


-- name: GetAreasWithAeons :many
SELECT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN aeons ae ON ae.area_id = a.id
JOIN player_units pu ON ae.unit_id = pu.id
ORDER BY a.id;


-- name: GetAreasWithMonsters :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j1 ON j1.encounter_location_id = el.id
JOIN monster_formations mf ON j1.monster_formation_id = mf.id
JOIN j_monster_formations_monsters j2 ON j2.monster_formation_id = mf.id
JOIN monster_amounts ma ON j2.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
ORDER BY a.id;


-- name: GetAreasWithItemMonster :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
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


-- name: GetAreasWithBosses :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN encounter_locations el ON el.area_id = a.id
JOIN j_encounter_location_formations j ON j.encounter_location_id = el.id
JOIN monster_formations mf ON j.monster_formation_id = mf.id
JOIN formation_boss_songs bs ON mf.boss_song_id = bs.id
JOIN songs so ON bs.song_id = so.id
ORDER BY a.id;


-- name: GetAreasWithShops :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN shops sh ON sh.area_id = a.id
ORDER BY a.id;


-- name: GetAreasWithItemShop :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN shops sh ON sh.area_id = a.id
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreasWithTreasures :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN treasures t ON t.area_id = a.id
ORDER BY a.id;


-- name: GetAreasWithItemTreasure :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN treasures t ON t.area_id = a.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreasWithKeyItemTreasure :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN treasures t ON t.area_id = a.id
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY a.id;


-- name: GetAreasWithSidequests :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN quests q ON qc.quest_id = q.id
ORDER BY a.id;


-- name: GetAreasWithItemQuest :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY a.id;


-- name: GetAreasWithKeyItemQuest :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN completion_locations cl ON cl.area_id = a.id
JOIN quest_completions qc ON cl.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN key_items ki ON ki.master_item_id = mi.id
WHERE ki.id = $1
ORDER BY a.id;


-- name: GetAreasWithFMVs :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name As sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id
JOIN fmvs f ON f.area_id = a.id
ORDER BY a.id;




-- name: GetLocationAreas :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id 
JOIN locations l ON s.location_id = l.id
WHERE l.id = $1
ORDER BY a.id;


-- name: GetLocationMonsters :many
SELECT DISTINCT
    m.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM monsters m
LEFT JOIN monster_amounts ma ON ma.monster_id = m.id
LEFT JOIN j_monster_formations_monsters j1 ON j1.monster_amount_id = ma.id
LEFT JOIN monster_formations mf ON j1.monster_formation_id = mf.id
LEFT JOIN j_encounter_location_formations j2 ON j2.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j2.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE location_id = $1
ORDER BY m.id;


-- name: GetSublocationAreas :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id 
JOIN locations l ON s.location_id = l.id
WHERE s.id = $1
ORDER BY a.id;


-- name: GetSublocationMonsters :many
SELECT DISTINCT
    m.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM monsters m
LEFT JOIN monster_amounts ma ON ma.monster_id = m.id
LEFT JOIN j_monster_formations_monsters j1 ON j1.monster_amount_id = ma.id
LEFT JOIN monster_formations mf ON j1.monster_formation_id = mf.id
LEFT JOIN j_encounter_location_formations j2 ON j2.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j2.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE sublocation_id = $1
ORDER BY m.id;




-- name: CreateTreasure :one
INSERT INTO treasures (data_hash, area_id, version, treasure_type, loot_type, is_post_airship, is_anima_treasure, notes, gil_amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = treasures.data_hash
RETURNING *;


-- name: UpdateTreasure :exec
UPDATE treasures
SET data_hash = $1,
    found_equipment_id = $2
WHERE id = $3;


-- name: CreateShop :one
INSERT INTO shops (data_hash, version, area_id, notes, category)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = shops.data_hash
RETURNING *;


-- name: CreateEncounterLocation :one
INSERT INTO encounter_locations (data_hash, version, area_id, notes)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = encounter_locations.data_hash
RETURNING *;


-- name: CreateFormationBossSong :one
INSERT INTO formation_boss_songs (data_hash, song_id, celebrate_victory)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = formation_boss_songs.data_hash
RETURNING *;


-- name: CreateMonsterFormation :one
INSERT INTO monster_formations (data_hash, category, is_forced_ambush, can_escape, boss_song_id, notes)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_formations.data_hash
RETURNING *;


-- name: CreateEncounterLocationFormationsJunction :exec
INSERT INTO j_encounter_location_formations (data_hash, encounter_location_id, monster_formation_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationsMonstersJunction :exec
INSERT INTO j_monster_formations_monsters (data_hash, monster_formation_id, monster_amount_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationsTriggerCommandsJunction :exec
INSERT INTO j_monster_formations_trigger_commands (data_hash, monster_formation_id, trigger_command_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateFoundEquipmentPiece :one
INSERT INTO found_equipment_pieces (data_hash, equipment_name_id, empty_slots_amount)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = found_equipment_pieces.data_hash
RETURNING *;


-- name: CreateFoundEquipmentAbilitiesJunction :exec
INSERT INTO j_found_equipment_abilities (data_hash, found_equipment_id, auto_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTreasuresItemsJunction :exec
INSERT INTO j_treasures_items (data_hash, treasure_id, item_amount_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShopItem :one
INSERT INTO shop_items (data_hash, item_id, price)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = shop_items.data_hash
RETURNING *;


-- name: CreateShopEquipmentPiece :one
INSERT INTO shop_equipment_pieces (data_hash, found_equipment_id, price)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = shop_equipment_pieces.data_hash
RETURNING *;


-- name: CreateShopsItemsJunction :exec
INSERT INTO j_shops_items (data_hash, shop_id, shop_item_id, shop_type)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShopsEquipmentJunction :exec
INSERT INTO j_shops_equipment (data_hash, shop_id, shop_equipment_id, shop_type)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
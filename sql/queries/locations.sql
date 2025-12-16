-- name: GetLocationAreaByAreaName :one
SELECT l.name, s.name, a.name, a.version FROM locations l LEFT JOIN sublocations s ON s.location_id = l.id LEFT JOIN areas a ON a.sublocation_id = a.id
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

-- name: GetAreaCount :one
SELECT COUNT(id) FROM areas;


-- name: GetAreas :many
SELECT
    l.id AS location_id,
    l.name AS location,
    s.id AS sublocation_id,
    s.name AS sublocation,
    a.* FROM areas a
LEFT JOIN sublocations s ON a.sublocation_id = s.id 
LEFT JOIN locations l ON s.location_id = l.id;


-- name: GetArea :one
SELECT
    l.id AS location_id,
    l.name AS location,
    s.name AS sublocation,
    a.* FROM areas a
LEFT JOIN sublocations s ON a.sublocation_id = s.id 
LEFT JOIN locations l ON s.location_id = l.id
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
LEFT JOIN j_area_connected_areas j ON j.connection_id = ac.id
LEFT JOIN areas a ON ac.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
LEFT JOIN areas a2 ON j.area_id = a2.id
WHERE a2.id = $1
ORDER BY ac.id;


-- name: GetAreaCharacters :many
SELECT
    c.*,
    pu.name,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area, a.version AS area_version
FROM player_units pu
LEFT JOIN characters c ON c.unit_id = pu.id
LEFT JOIN areas a ON c.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY c.id;


-- name: GetAreaAeons :many
SELECT
    ae.*,
    pu.name,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area, a.version AS area_version
FROM player_units pu
LEFT JOIN aeons ae ON ae.unit_id = pu.id
LEFT JOIN areas a ON ae.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
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
FROM monsters m
LEFT JOIN monster_amounts ma ON ma.monster_id = m.id
LEFT JOIN j_monster_formations_monsters j1 ON j1.monster_amount_id = ma.id
LEFT JOIN monster_formations mf ON j1.monster_formation_id = mf.id
LEFT JOIN j_encounter_location_formations j2 ON j2.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j2.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY m.id;


-- name: GetAreaMonsterFormations :many
SELECT DISTINCT
    mf.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM monster_formations mf
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY mf.id;


-- name: GetAreaBossSongs :many
SELECT DISTINCT
    so.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM songs so
LEFT JOIN formation_boss_songs bs ON bs.song_id = so.id
LEFT JOIN monster_formations mf ON mf.boss_song_id = bs.id
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
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
LEFT JOIN songs so ON c.song_id = so.id
LEFT JOIN j_songs_cues j ON j.cue_id = c.id
LEFT JOIN areas a ON COALESCE(c.trigger_area_id, j.included_area_id) = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
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
FROM background_music bm
LEFT JOIN j_songs_background_music j ON j.bm_id = bm.id
LEFT JOIN songs so ON j.song_id = so.id
LEFT JOIN areas a ON j.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFMVSongs :many
SELECT DISTINCT
    so.*,
    l.name AS location,
    s.name As sublocation,
    a.name AS area,
    a.version AS area_version
FROM songs so
LEFT JOIN fmvs f ON f.song_id = so.id
LEFT JOIN areas a ON f.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY so.id;


-- name: GetAreaFMVs :many
SELECT
    f.*,
    l.name AS location,
    s.name As sublocation,
    a.name AS area,
    a.version AS area_version
FROM fmvs f
LEFT JOIN areas a ON f.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY f.id;


-- name: GetAreaQuests :many
SELECT DISTINCT
    q.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM completion_locations cl
LEFT JOIN quest_completions qc ON cl.completion_id = qc.id
LEFT JOIN quests q ON qc.quest_id = q.id
LEFT JOIN areas a ON cl.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.id = $1
ORDER BY q.id;


-- name: GetAreasWithSaveSphere :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE a.has_save_sphere = $1
ORDER BY a.id;


-- name: GetAreasWithCompSphere :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE a.has_compilation_sphere = $1
ORDER BY a.id;


-- name: GetAreasWithDropOff :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE a.airship_drop_off = $1
ORDER BY a.id;


-- name: GetAreasWithChocobo :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE a.can_ride_chocobo = $1
ORDER BY a.id;


-- name: GetAreasStoryOnly :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE a.story_only = $1
ORDER BY a.id;


-- name: GetSublocationAreas :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE s.id = $1
ORDER BY a.id;


-- name: GetAreasWithBosses :many
SELECT DISTINCT
    a.id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM songs so
LEFT JOIN formation_boss_songs bs ON bs.song_id = so.id
LEFT JOIN monster_formations mf ON mf.boss_song_id = bs.id
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
ORDER BY a.id;


-- name: GetSublocationShops :many
SELECT
    sh.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version area_version
FROM shops sh
LEFT JOIN areas a ON sh.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.sublocation_id = $1
ORDER BY sh.id;


-- name: GetSublocationTreasures :many
SELECT
    t.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM treasures t
LEFT JOIN areas a ON t.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE a.sublocation_id = $1
ORDER BY t.id;


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


-- name: GetSublocationMonsterFormations :many
SELECT DISTINCT
    mf.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM monster_formations mf
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE sublocation_id = $1
ORDER BY mf.id;


-- name: GetSublocationsWithBosses :many
SELECT DISTINCT
    s.id,
    l.name AS location,
    s.name AS sublocation
FROM songs so
INNER JOIN formation_boss_songs bs ON bs.song_id = so.id
LEFT JOIN monster_formations mf ON mf.boss_song_id = bs.id
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
ORDER BY s.id;


-- name: GetLocationAreas :many
SELECT
    l.name AS location, 
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM locations l
LEFT JOIN sublocations s ON s.location_id = l.id
LEFT JOIN areas a ON a.sublocation_id = s.id
WHERE l.id = $1
ORDER BY a.id;


-- name: GetLocationShops :many
SELECT
    sh.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM shops sh
LEFT JOIN areas a ON sh.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE s.location_id = $1
ORDER BY sh.id;


-- name: GetLocationTreasures :many
SELECT
    t.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM treasures t
LEFT JOIN areas a ON t.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE s.location_id = $1
ORDER BY t.id;


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


-- name: GetLocationMonsterFormations :many
SELECT DISTINCT
    mf.*,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version AS area_version
FROM monster_formations mf
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
WHERE location_id = $1
ORDER BY mf.id;


-- name: GetLocationsWithBosses :many
SELECT DISTINCT
    l.id,
    l.name AS location
FROM songs so
INNER JOIN formation_boss_songs bs ON bs.song_id = so.id
LEFT JOIN monster_formations mf ON mf.boss_song_id = bs.id
LEFT JOIN j_encounter_location_formations j ON j.monster_formation_id = mf.id
LEFT JOIN encounter_locations el ON j.encounter_location_id = el.id
LEFT JOIN areas a ON el.area_id = a.id
LEFT JOIN sublocations s ON a.sublocation_id = s.id
LEFT JOIN locations l ON s.location_id = l.id
ORDER BY l.id;


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
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
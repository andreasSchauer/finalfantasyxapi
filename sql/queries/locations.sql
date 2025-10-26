-- name: CreateLocation :one
INSERT INTO locations (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = locations.data_hash
RETURNING *;


-- name: CreateSubLocation :one
INSERT INTO sub_locations (data_hash, location_id, name, specification)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = sub_locations.data_hash
RETURNING *;


-- name: CreateArea :one
INSERT INTO areas (data_hash, sub_location_id, name, version, specification, story_only, has_save_sphere, airship_drop_off, has_compilation_sphere, can_ride_chocobo)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = areas.data_hash
RETURNING *;


-- name: CreateAreaConnection :one
INSERT INTO area_connections (data_hash, area_id, connection_type, story_only, notes)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = area_connections.data_hash
RETURNING *;


-- name: CreateAreaConnectionJunction :exec
INSERT INTO j_area_connection (data_hash, area_id, connection_id)
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


-- name: CreateFormationLocation :one
INSERT INTO formation_locations (data_hash, version, area_id, notes)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = formation_locations.data_hash
RETURNING *;


-- name: CreateFormationBossSong :one
INSERT INTO formation_boss_songs (data_hash, song_id, celebrate_victory)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = formation_boss_songs.data_hash
RETURNING *;


-- name: CreateMonsterFormation :one
INSERT INTO monster_formations (data_hash, formation_location_id, category, is_forced_ambush, can_escape, boss_song_id, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_formations.data_hash
RETURNING *;


-- name: CreateLocationMonsterFormationJunction :exec
INSERT INTO j_location_monster_formation (data_hash, formation_location_id, monster_formation_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationMonsterAmountJunction :exec
INSERT INTO j_monster_formation_monster_amount (data_hash, monster_formation_id, monster_amount_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationTriggerCommandJunction :exec
INSERT INTO j_monster_formation_trigger_command (data_hash, monster_formation_id, trigger_command_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateFoundEquipmentPiece :one
INSERT INTO found_equipment_pieces (data_hash, equipment_name_id, empty_slots_amount)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = found_equipment_pieces.data_hash
RETURNING *;


-- name: CreateFoundEquipmentAutoAbilityJunction :exec
INSERT INTO j_found_equipment_auto_ability (data_hash, found_equipment_id, auto_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTreasureItemAmountJunction :exec
INSERT INTO j_treasure_item_amount (data_hash, treasure_id, item_amount_id)
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


-- name: CreateShopShopItemJunction :exec
INSERT INTO j_shop_shop_item (data_hash, shop_id, shop_item_id, shop_type)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShopShopEquipmentJunction :exec
INSERT INTO j_shop_shop_equipment (data_hash, shop_id, shop_equipment_id, shop_type)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
-- name: CreateLocation :one
INSERT INTO locations (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = locations.data_hash
RETURNING *;


-- name: CreateSublocation :one
INSERT INTO sublocations (data_hash, location_id, name, specification)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = sublocations.data_hash
RETURNING *;


-- name: CreateArea :one
INSERT INTO areas (data_hash, sublocation_id, name, version, specification, availability, has_save_sphere, airship_drop_off, has_compilation_sphere, can_ride_chocobo)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = areas.data_hash
RETURNING *;


-- name: CreateAreaConnection :one
INSERT INTO area_connections (data_hash, area_id, connection_type, is_story_based, notes)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = area_connections.data_hash
RETURNING *;


-- name: CreateAreaConnectedAreasJunction :exec
INSERT INTO j_area_connected_areas (data_hash, area_id, connection_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTreasure :one
INSERT INTO treasures (data_hash, area_id, version, treasure_type, loot_type, availability, is_anima_treasure, notes, gil_amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = treasures.data_hash
RETURNING *;


-- name: CreateShop :one
INSERT INTO shops (data_hash, version, area_id, notes, category, availability)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = shops.data_hash
RETURNING *;


-- name: CreateTreasureEquipmentPiece :one
INSERT INTO treasure_equipment_pieces (data_hash, treasure_id, equipment_name_id, empty_slots_amount)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = treasure_equipment_pieces.data_hash
RETURNING *;


-- name: CreateTreasureEquipmentAbilitiesJunction :exec
INSERT INTO j_treasure_equipment_abilities (data_hash, treasure_equipment_id, auto_ability_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShopEquipmentAbilitiesJunction :exec
INSERT INTO j_shop_equipment_abilities (data_hash, shop_equipment_id, auto_ability_id)
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
INSERT INTO shop_equipment_pieces (data_hash, shop_id, equipment_name_id, shop_type, empty_slots_amount, price)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = shop_equipment_pieces.data_hash
RETURNING *;


-- name: CreateShopsItemsJunction :exec
INSERT INTO j_shops_items (data_hash, shop_id, shop_item_id, shop_type)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
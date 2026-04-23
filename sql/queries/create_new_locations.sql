-- name: CreateLocationBulk :many
INSERT INTO locations (data_hash, name)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateSublocationBulk :many
INSERT INTO sublocations (data_hash, location_id, name, specification)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('location_id')::int[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('specification')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAreaBulk :many
INSERT INTO areas (data_hash, sublocation_id, name, version, specification, availability, has_save_sphere, airship_drop_off, has_compilation_sphere, can_ride_chocobo)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('sublocation_id')::int[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('specification')::null_string[]),
    unnest(sqlc.arg('availability')::availability_type[]),
    unnest(sqlc.arg('has_save_sphere')::boolean[]),
    unnest(sqlc.arg('airship_drop_off')::boolean[]),
    unnest(sqlc.arg('has_compilation_sphere')::boolean[]),
    unnest(sqlc.arg('can_ride_chocobo')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateAreaConnectionBulk :many
INSERT INTO area_connections (data_hash, area_id, connection_type, is_story_based, notes)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('connection_type')::area_connection_type[]),
    unnest(sqlc.arg('is_story_based')::boolean[]),
    unnest(sqlc.arg('notes')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateTreasureBulk :many
INSERT INTO treasures (data_hash, area_id, version, treasure_type, loot_type, availability, is_anima_treasure, notes, gil_amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('version')::int[]),
    unnest(sqlc.arg('treasure_type')::treasure_type[]),
    unnest(sqlc.arg('loot_type')::loot_type[]),
    unnest(sqlc.arg('availability')::availability_type[]),
    unnest(sqlc.arg('is_anima_treasure')::boolean[]),
    unnest(sqlc.arg('notes')::null_string[]),
    unnest(sqlc.arg('gil_amount')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateShopBulk :many
INSERT INTO shops (data_hash, version, area_id, notes, category, availability)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('version')::null_int[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('notes')::null_string[]),
    unnest(sqlc.arg('category')::shop_category[]),
    unnest(sqlc.arg('availability')::availability_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateTreasureEquipmentPieceBulk :many
INSERT INTO treasure_equipment_pieces (data_hash, treasure_id, equipment_name_id, empty_slots_amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('treasure_id')::int[]),
    unnest(sqlc.arg('equipment_name_id')::int[]),
    unnest(sqlc.arg('empty_slots_amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateShopItemBulk :many
INSERT INTO shop_items (data_hash, item_id, price)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('item_id')::int[]),
    unnest(sqlc.arg('price')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateShopEquipmentPieceBulk :many
INSERT INTO shop_equipment_pieces (data_hash, shop_id, equipment_name_id, shop_type, empty_slots_amount, price)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('shop_id')::int[]),
    unnest(sqlc.arg('equipment_name_id')::int[]),
    unnest(sqlc.arg('shop_type')::shop_type[]),
    unnest(sqlc.arg('empty_slots_amount')::int[]),
    unnest(sqlc.arg('price')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;





-- name: CreateAreaConnectedAreasJunctionBulk :exec
INSERT INTO j_area_connected_areas (data_hash, area_id, connection_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('connection_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTreasureEquipmentAbilitiesJunctionBulk :exec
INSERT INTO j_treasure_equipment_abilities (data_hash, treasure_equipment_id, auto_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('treasure_equipment_id')::int[]),
    unnest(sqlc.arg('auto_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShopEquipmentAbilitiesJunctionBulk :exec
INSERT INTO j_shop_equipment_abilities (data_hash, shop_equipment_id, auto_ability_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('shop_equipment_id')::int[]),
    unnest(sqlc.arg('auto_ability_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTreasuresItemsJunctionBulk :exec
INSERT INTO j_treasures_items (data_hash, treasure_id, item_amount_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('treasure_id')::int[]),
    unnest(sqlc.arg('item_amount_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShopsItemsJunctionBulk :exec
INSERT INTO j_shops_items (data_hash, shop_id, shop_item_id, shop_type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('shop_id')::int[]),
    unnest(sqlc.arg('shop_item_id')::int[]),
    unnest(sqlc.arg('shop_type')::int[])
ON CONFLICT(data_hash) DO NOTHING;
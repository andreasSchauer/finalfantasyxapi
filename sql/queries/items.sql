-- name: CreateMasterItem :one
INSERT INTO master_items (data_hash, name, type)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = master_items.data_hash
RETURNING *;


-- name: CreateItem :one
INSERT INTO items (data_hash, master_item_id, description, effect, sphere_grid_description, category, usability, base_price, sell_value)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = items.data_hash
RETURNING *;


-- name: CreateItemAbility :exec
INSERT INTO item_abilities (data_hash, item_id, ability_id, cursor)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateKeyItem :exec
INSERT INTO key_items (data_hash, master_item_id, category, description, effect)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: GetKeyItemByName :one
SELECT
    ki.id as key_item_id,
    ki.data_hash as key_item_data_hash,
    mi.name,
    ki.category,
    ki.description,
    ki.effect,
    mi.id as master_item_id,
    mi.data_hash as master_item_data_hash,
    mi.type
FROM key_items ki
LEFT JOIN master_items mi
ON mi.id = ki.master_item_id
WHERE mi.name = $1;


-- name: CreatePrimer :exec
INSERT INTO primers (data_hash, key_item_id, al_bhed_letter, english_letter)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
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


-- name: GetItemByName :one
SELECT
    i.id as item_id,
    i.data_hash as item_data_hash,
    mi.id as master_item_id,
    mi.name,
    i.description,
    i.effect,
    i.sphere_grid_description,
    i.category,
    i.usability,
    i.base_price,
    i.sell_value,
    mi.data_hash as master_item_data_hash,
    mi.type
FROM items i
LEFT JOIN master_items mi
ON mi.id = i.master_item_id
WHERE mi.name = $1;


-- name: GetItems :many
SELECT
    i.id as item_id,
    i.data_hash as item_data_hash,
    mi.id as master_item_id,
    mi.name,
    i.description,
    i.effect,
    i.sphere_grid_description,
    i.category,
    i.usability,
    i.base_price,
    i.sell_value,
    mi.data_hash as master_item_data_hash,
    mi.type
FROM items i
LEFT JOIN master_items mi
ON mi.id = i.master_item_id;


-- name: CreateKeyItem :one
INSERT INTO key_items (data_hash, master_item_id, category, description, effect)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = key_items.data_hash
RETURNING *;


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


-- name: GetKeyItems :many
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
ON mi.id = ki.master_item_id;



-- name: CreatePrimer :exec
INSERT INTO primers (data_hash, key_item_id, al_bhed_letter, english_letter)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;



-- name: CreateMix :one
INSERT INTO mixes (data_hash, overdrive_id, category)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = mixes.data_hash
RETURNING *;


-- name: CreateMixCombination :one
INSERT INTO mix_combinations (data_hash, first_item_id, second_item_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = mix_combinations.data_hash
RETURNING *;


-- name: CreateMixComboJunction :exec
INSERT INTO mix_combo_junctions (data_hash, mix_id, combo_id, is_best_combo)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
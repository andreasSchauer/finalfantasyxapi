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


-- name: CreateItemsRelatedStatsJunction :exec
INSERT INTO j_items_related_stats (data_hash, item_id, stat_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateItemsAvailableMenusJunction :exec
INSERT INTO j_items_available_menus (data_hash, item_id, submenu_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateItemAbility :one
INSERT INTO item_abilities (data_hash, item_id, ability_id, cursor)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = item_abilities.data_hash
RETURNING *;



-- name: CreateKeyItem :one
INSERT INTO key_items (data_hash, master_item_id, category, description, effect)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = key_items.data_hash
RETURNING *;



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


-- name: CreateMixesCombinationsJunction :exec
INSERT INTO j_mixes_combinations (data_hash, mix_id, combo_id, is_best_combo)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateItemAmount :one
INSERT INTO item_amounts (data_hash, master_item_id, amount)
VALUES ( $1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = item_amounts.data_hash
RETURNING *;
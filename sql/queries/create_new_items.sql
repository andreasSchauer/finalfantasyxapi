-- name: CreateMasterItemBulk :many
INSERT INTO master_items (data_hash, name, type)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('type')::item_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateItemBulk :many
INSERT INTO items (data_hash, master_item_id, description, effect, category, usability, base_price, sell_value)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('master_item_id')::int[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('category')::item_category[]),
    unnest(sqlc.arg('usability')::item_usability[]),
    unnest(sqlc.arg('base_price')::null_int[]),
    unnest(sqlc.arg('sell_value')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateItemAbilityBulk :many
INSERT INTO item_abilities (data_hash, item_id, ability_id, cursor)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('item_id')::int[]),
    unnest(sqlc.arg('ability_id')::int[]),
    unnest(sqlc.arg('cursor')::target_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;



-- name: CreateSphereBulk :many
INSERT INTO spheres (data_hash, item_id, sphere_grid_description, sphere_color, sphere_effect, target_node_position, target_node_state)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('item_id')::int[]),
    unnest(sqlc.arg('sphere_grid_description')::text[]),
    unnest(sqlc.arg('sphere_color')::sphere_color[]),
    unnest(sqlc.arg('sphere_effect')::sphere_effect[]),
    unnest(sqlc.arg('target_node_position')::node_position[]),
    unnest(sqlc.arg('target_node_state')::null_node_state[]),
    unnest(sqlc.arg('created_node_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateSphereTargetableNodeBulk :many
INSERT INTO spheres_targetable_nodes (data_hash, sphere_id, node)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('sphere_id')::int[]),
    unnest(sqlc.arg('node')::node_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateCreatedNodeBulk :many
INSERT INTO created_nodes(data_hash, node, value)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('node')::node_type[]),
    unnest(sqlc.arg('value')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateKeyItemBulk :many
INSERT INTO key_items (data_hash, master_item_id, category, description, effect)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('master_item_id')::int[]),
    unnest(sqlc.arg('category')::key_item_category[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;



-- name: CreatePrimerBulk :many
INSERT INTO primers (data_hash, key_item_id, al_bhed_letter, english_letter)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('key_item_id')::int[]),
    unnest(sqlc.arg('al_bhed_letter')::text[]),
    unnest(sqlc.arg('english_letter')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;



-- name: CreateMixBulk :many
INSERT INTO mixes (data_hash, overdrive_id, category)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('overdrive_id')::int[]),
    unnest(sqlc.arg('category')::mix_category[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateMixCombinationBulk :many
INSERT INTO mix_combinations (data_hash, mix_id, first_item_id, second_item_id, is_best_combo)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('mix_id')::int[]),
    unnest(sqlc.arg('first_item_id')::int[]),
    unnest(sqlc.arg('second_item_id')::int[]),
    unnest(sqlc.arg('is_best_combo')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreateItemAmountBulk :many
INSERT INTO item_amounts (data_hash, master_item_id, amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('master_item_id')::int[]),
    unnest(sqlc.arg('amount')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;


-- name: CreatePossibleItemBulk :many
INSERT INTO possible_items (data_hash, item_amount_id, chance)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('item_amount_id')::int[]),
    unnest(sqlc.arg('chance')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;





-- name: CreateItemsRelatedStatsJunctionBulk :exec
INSERT INTO j_items_related_stats (data_hash, item_id, stat_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('item_id')::int[]),
    unnest(sqlc.arg('stat_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateItemsAvailableMenusJunctionBulk :exec
INSERT INTO j_items_available_menus (data_hash, item_id, submenu_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('item_id')::int[]),
    unnest(sqlc.arg('submenu_id')::int[])
ON CONFLICT(data_hash) DO NOTHING;
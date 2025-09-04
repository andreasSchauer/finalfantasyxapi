-- name: CreateTreasureList :one
INSERT INTO treasure_lists DEFAULT VALUES
RETURNING *;


-- name: CreateTreasure :exec
INSERT INTO treasures (data_hash, treasure_list_id, treasure_type, loot_type, is_post_airship, is_anima_treasure, notes, gil_amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShop :exec
INSERT INTO shops (data_hash, version, notes, category)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBlitzballItemList :exec
INSERT INTO blitzball_items_lists (data_hash, category, slot)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateQuest :one
INSERT INTO quests (data_hash, name, type)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = quests.data_hash
RETURNING *;


-- name: CreateSidequest :one
INSERT INTO sidequests (quest_id)
VALUES ($1)
RETURNING *;


-- name: CreateSubquest :exec
INSERT INTO subquests (quest_id, parent_sidequest_id)
VALUES ($1, $2);


-- name: CreateMonsterArenaCreation :exec
INSERT INTO monster_arena_creations (data_hash, name, category, required_area, required_species, underwater_only, creations_unlocked_category, amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO NOTHING;
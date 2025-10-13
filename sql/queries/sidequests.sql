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
INSERT INTO sidequests (data_hash, quest_id)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = sidequests.data_hash
RETURNING *;


-- name: CreateSubquest :one
INSERT INTO subquests (data_hash, quest_id, parent_sidequest_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = subquests.data_hash
RETURNING *;


-- name: CreateMonsterArenaCreation :exec
INSERT INTO monster_arena_creations (data_hash, subquest_id, category, required_area, required_species, underwater_only, creations_unlocked_category, amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO NOTHING;
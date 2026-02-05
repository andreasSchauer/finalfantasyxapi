-- name: CreateBlitzballPosition :one
INSERT INTO blitzball_positions (data_hash, category, slot)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = blitzball_positions.data_hash
RETURNING *;


-- name: CreateBlitzballItem :exec
INSERT INTO blitzball_items (data_hash, position_id, possible_item_id)
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
INSERT INTO subquests (data_hash, quest_id, sidequest_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = subquests.data_hash
RETURNING *;


-- name: CreateQuestCompletion :one
INSERT INTO quest_completions (data_hash, quest_id, condition, item_amount_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = quest_completions.data_hash
RETURNING *;


-- name: CreateCompletionLocation :exec
INSERT INTO completion_locations (data_hash, completion_id, area_id, notes)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterArenaCreation :one
INSERT INTO monster_arena_creations (data_hash, subquest_id, category, required_area, required_species, underwater_only, creations_unlocked_category, amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO Update SET data_hash = monster_arena_creations.data_hash
RETURNING *;


-- name: UpdateMonsterArenaCreation :exec
UPDATE monster_arena_creations
SET data_hash = $1,
    monster_id = $2
WHERE id = $3;
-- name: CreateBlitzballPositionBulk :many
INSERT INTO blitzball_positions (data_hash, category, slot)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('category')::blitzball_tournament_category[]),
    unnest(sqlc.arg('slot')::blitzball_position_slot[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateBlitzballItemBulk :many
INSERT INTO blitzball_items (data_hash, position_id, possible_item_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('position_id')::int[]),
    unnest(sqlc.arg('possible_item_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateQuestBulk :many
INSERT INTO quests (data_hash, name, type, availability, is_repeatable)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('type')::quest_type[]),
    unnest(sqlc.arg('availability')::availability_type[]),
    unnest(sqlc.arg('is_repeatable')::boolean[]),
    unnest(sqlc.arg('completion_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateSidequestBulk :many
INSERT INTO sidequests (data_hash, quest_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('quest_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateSubquestBulk :many
INSERT INTO subquests (data_hash, quest_id, sidequest_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('quest_id')::int[]),
    unnest(sqlc.arg('sidequest_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateQuestCompletionBulk :many
INSERT INTO quest_completions (data_hash, condition, item_amount_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('condition')::null_string[]),
    unnest(sqlc.arg('item_amount_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateCompletionAreaBulk :many
INSERT INTO completion_areas (data_hash, completion_id, area_id, notes)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('completion_id')::int[]),
    unnest(sqlc.arg('area_id')::int[]),
    unnest(sqlc.arg('notes')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateMonsterArenaCreationBulk :many
INSERT INTO monster_arena_creations (data_hash, subquest_id, category, required_area, required_species, underwater_only, creations_unlocked_category, amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('subquest_id')::int[]),
    unnest(sqlc.arg('category')::ma_creation_category[]),
    unnest(sqlc.arg('required_area')::null_ma_creation_area[]),
    unnest(sqlc.arg('required_species')::null_ma_creation_species[]),
    unnest(sqlc.arg('underwater_only')::boolean[]),
    unnest(sqlc.arg('creations_unlocked_category')::null_creations_unlocked_category[]),
    unnest(sqlc.arg('amount')::int[]),
    unnest(sqlc.arg('monster_id')::null_int[])
ON CONFLICT(data_hash) DO Update SET data_hash = EXCLUDED.data_hash
RETURNING id;
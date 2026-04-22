-- name: CreateAeonCommandBulk :many
INSERT INTO aeon_commands (data_hash, name, description, effect, cursor)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('effect')::text[]),
    unnest(sqlc.arg('cursor')::null_target_type[]),
    unnest(sqlc.arg('topmenu_id')::null_int[]),
    unnest(sqlc.arg('submenu_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateTopmenuBulk :many
INSERT INTO topmenus (data_hash, name)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateSubmenuBulk :many
INSERT INTO submenus (data_hash, name, description, effect)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('description')::null_string[]),
    unnest(sqlc.arg('effect')::text[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateOverdriveCommandBulk :many
INSERT INTO overdrive_commands (data_hash, name, description, rank)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('description')::text[]),
    unnest(sqlc.arg('rank')::int[]),
    unnest(sqlc.arg('character_class_id')::null_int[]),
    unnest(sqlc.arg('topmenu_id')::null_int[]),
    unnest(sqlc.arg('submenu_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;
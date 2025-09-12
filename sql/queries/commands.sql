-- name: CreateCommand :exec
INSERT INTO commands (data_hash, name, description, effect, category, cursor)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveCommand :one
INSERT INTO overdrive_commands (data_hash, name, description, rank, open_menu)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = overdrive_commands.data_hash
RETURNING *;


-- name: CreateOverdrive :exec
INSERT INTO overdrives (data_hash, od_command_id, name, version, description, effect, rank, appears_in_help_bar, can_copycat, unlock_condition, countdown_in_sec, cursor)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT(data_hash) DO NOTHING;
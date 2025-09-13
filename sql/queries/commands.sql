-- name: CreateAeonCommand :exec
INSERT INTO aeon_commands (data_hash, name, description, effect, cursor)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMenuCommand :exec
INSERT INTO menu_commands (data_hash, name, description, effect)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
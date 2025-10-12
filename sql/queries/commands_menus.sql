-- name: CreateAeonCommand :exec
INSERT INTO aeon_commands (data_hash, name, description, effect, topmenu, cursor)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateSubmenu :exec
INSERT INTO submenus (data_hash, name, description, effect, topmenu)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;
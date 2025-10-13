-- name: CreateAeonCommand :exec
INSERT INTO aeon_commands (data_hash, name, description, effect, topmenu, cursor)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateSubmenu :one
INSERT INTO submenus (data_hash, name, description, effect, topmenu)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = submenus.data_hash
RETURNING *;


-- name: CreateSubmenuCharacterClassJunction :exec
INSERT INTO j_submenu_character_class (data_hash, submenu_id, character_class_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
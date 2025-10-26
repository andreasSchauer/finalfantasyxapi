-- name: CreateAeonCommand :one
INSERT INTO aeon_commands (data_hash, name, description, effect, topmenu, cursor)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = aeon_commands.data_hash
RETURNING *;


-- name: UpdateAeonCommand :exec
UPDATE aeon_commands
SET data_hash = $1,
    submenu_id = $2
WHERE id = $3;


-- name: CreateAeonCommandAbilityJunction :exec
INSERT INTO j_aeon_command_ability (data_hash, aeon_command_id, ability_id, character_class_id)
VALUES ($1, $2, $3, $4)
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


-- name: CreateOverdriveCommand :one
INSERT INTO overdrive_commands (data_hash, name, description, rank, topmenu)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = overdrive_commands.data_hash
RETURNING *;


-- name: UpdateOverdriveCommand :exec
UPDATE overdrive_commands
SET data_hash = $1,
    character_class_id = $2,
    submenu_id = $3
WHERE id = $4;
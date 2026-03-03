-- name: CreateAeonCommand :one
INSERT INTO aeon_commands (data_hash, name, description, effect, cursor)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = aeon_commands.data_hash
RETURNING *;


-- name: UpdateAeonCommand :exec
UPDATE aeon_commands
SET data_hash = $1,
    topmenu_id = $2,
    submenu_id = $3
WHERE id = $4;


-- name: CreateAeonCommandsPossibleAbilitiesJunction :exec
INSERT INTO j_aeon_commands_possible_abilities (data_hash, aeon_command_id, character_class_id, ability_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTopmenu :one
INSERT INTO topmenus (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = topmenus.data_hash
RETURNING *;


-- name: CreateSubmenu :one
INSERT INTO submenus (data_hash, name, description, effect)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = submenus.data_hash
RETURNING *;


-- name: UpdateSubmenu :exec
UPDATE submenus
SET data_hash = $1,
    topmenu_id = $2
WHERE id = $3;


-- name: CreateSubmenusUsersJunction :exec
INSERT INTO j_submenus_users (data_hash, submenu_id, character_class_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveCommand :one
INSERT INTO overdrive_commands (data_hash, name, description, rank)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = overdrive_commands.data_hash
RETURNING *;


-- name: UpdateOverdriveCommand :exec
UPDATE overdrive_commands
SET data_hash = $1,
    character_class_id = $2,
    topmenu_id = $3,
    submenu_id = $4
WHERE id = $5;
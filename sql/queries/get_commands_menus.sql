-- name: GetOverdriveCommandOverdriveAbilityIDs :many
SELECT DISTINCT j.overdrive_ability_id
FROM j_overdrives_overdrive_abilities j
JOIN overdrives o ON j.overdrive_id = o.id
WHERE o.od_command_id = $1
ORDER BY j.overdrive_ability_id;


-- name: GetOverdriveCommandIDs :many
SELECT id FROM overdrive_commands ORDER BY id;


-- name: GetOverdriveCommandOverdriveIDs :many
SELECT id FROM overdrives WHERE od_command_id = $1 ORDER BY id;




-- name: GetSubmenuIDs :many
SELECT id FROM submenus ORDER BY id;


-- name: GetSubmenuAbilityIDs :many
SELECT pa.ability_id FROM player_abilities pa WHERE pa.submenu_id = $1
UNION
SELECT oa.ability_id FROM misc_abilities oa WHERE oa.submenu_id = $1
ORDER BY ability_id;


-- name: GetSubmenuOpenedByAeonCommandID :one
SELECT id FROM aeon_commands WHERE submenu_id = $1;


-- name: GetSubmenuOpenedByAbilityID :one
SELECT pa.ability_id FROM player_abilities pa WHERE pa.open_submenu_id = $1
UNION
SELECT oa.ability_id FROM misc_abilities oa WHERE oa.open_submenu_id = $1
ORDER BY ability_id;


-- name: GetSubmenuOpenedByOverdriveCommandIDs :many
SELECT id FROM overdrive_commands WHERE submenu_id = $1 ORDER BY id;




-- name: GetTopmenuSubmenuIDs :many
SELECT id FROM submenus s WHERE topmenu_id = $1 ORDER BY id;


-- name: GetTopmenuAeonCommandIDs :many
SELECT id FROM aeon_commands WHERE topmenu_id = $1 ORDER BY id;


-- name: GetTopmenuOverdriveCommandIDs :many
SELECT id FROM overdrive_commands WHERE topmenu_id = $1 ORDER BY id;


-- name: GetTopmenuOverdriveIDs :many
SELECT id FROM overdrives WHERE topmenu_id = $1 ORDER BY id;


-- name: GetTopmenuAbilityIDs :many
SELECT pa.ability_id FROM player_abilities pa WHERE pa.topmenu_id = $1
UNION
SELECT ma.ability_id FROM misc_abilities ma WHERE ma.topmenu_id = $1
UNION
SELECT tc.ability_id FROM trigger_commands tc WHERE tc.topmenu_id = $1
ORDER BY ability_id;


-- name: GetTopmenuIDs :many
SELECT id FROM topmenus ORDER BY id;


-- name: GetAeonCommandIDs :many
SELECT id FROM aeon_commands ORDER BY id;
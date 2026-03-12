-- name: GetOverdriveCommandOverdriveAbilityIDs :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN j_overdrives_overdrive_abilities j ON j.overdrive_ability_id = oa.id
JOIN overdrives o ON j.overdrive_id = o.id
JOIN overdrive_commands oc ON o.od_command_id = oc.id
WHERE oc.id = $1
ORDER BY oa.id;


-- name: GetOverdriveCommandIDs :many
SELECT id FROM overdrive_commands ORDER BY id;


-- name: GetOverdriveCommandOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN overdrive_commands oc ON o.od_command_id = oc.id
WHERE oc.id = $1
ORDER BY o.id;




-- name: GetSubmenuIDs :many
SELECT id FROM submenus ORDER BY id;


-- name: GetSubmenuAbilityIDs :many
SELECT a.id
FROM abilities a
LEFT JOIN player_abilities pa ON pa.ability_id = a.id
LEFT JOIN unspecified_abilities ua ON ua.ability_id = a.id
WHERE pa.submenu_id = $1 OR ua.submenu_id = $1
ORDER BY a.id;


-- name: GetSubmenuOpenedByAeonCommandID :one
SELECT id FROM aeon_commands WHERE submenu_id = $1;


-- name: GetSubmenuOpenedByAbilityID :one
SELECT a.id
FROM abilities a
LEFT JOIN player_abilities pa ON pa.ability_id = a.id
LEFT JOIN unspecified_abilities ua ON ua.ability_id = a.id
WHERE pa.open_submenu_id = $1 OR ua.open_submenu_id = $1;


-- name: GetSubmenuOpenedByOverdriveCommandIDs :many
SELECT oc.id
FROM overdrive_commands oc
JOIN submenus s ON oc.submenu_id = s.id
WHERE s.id = $1
ORDER BY oc.id;







-- name: GetTopmenuSubmenuIDs :many
SELECT s.id
FROM submenus s
JOIN topmenus t ON s.topmenu_id = t.id
WHERE t.id = $1
ORDER BY s.id;


-- name: GetTopmenuAeonCommandIDs :many
SELECT ac.id
FROM aeon_commands ac
JOIN topmenus t ON ac.topmenu_id = t.id
WHERE t.id = $1
ORDER BY ac.id;


-- name: GetTopmenuOverdriveCommandIDs :many
SELECT oc.id
FROM overdrive_commands oc
JOIN topmenus t ON oc.topmenu_id = t.id
WHERE t.id = $1
ORDER BY oc.id;


-- name: GetTopmenuOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN topmenus t ON o.topmenu_id = t.id
WHERE t.id = $1
ORDER BY o.id;


-- name: GetTopmenuAbilityIDs :many
SELECT a.id
FROM abilities a
LEFT JOIN player_abilities pa ON pa.ability_id = a.id
LEFT JOIN unspecified_abilities ua ON ua.ability_id = a.id
LEFT JOIN trigger_commands tc ON tc.ability_id = a.id
WHERE pa.topmenu_id = $1 OR ua.topmenu_id = $1 OR tc.topmenu_id = $1
ORDER BY a.id;


-- name: GetTopmenuIDs :many
SELECT id FROM topmenus ORDER BY id;



-- name: GetAeonCommandIDs :many
SELECT id FROM aeon_commands ORDER BY id;
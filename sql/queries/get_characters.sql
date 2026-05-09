-- name: GetPlayerUnitIDs :many
SELECT id FROM player_units ORDER BY id;


-- name: GetPlayerUnitIDsByType :many
SELECT id FROM player_units WHERE type = $1 ORDER BY id;


-- name: GetCharacterCharClassIDs :many
SELECT j.class_id
FROM j_character_class_player_units j
JOIN characters c ON j.unit_id = c.unit_id
WHERE c.id = $1
ORDER BY j.class_id;


-- name: GetCharacterOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN j_character_class_player_units j ON j.class_id = o.character_class_id
JOIN characters c ON j.unit_id = c.unit_id
WHERE c.id = $1
ORDER BY o.id;


-- name: GetCharacterDefaultAbilityIDs :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_default_abilities da ON da.ability_id = a.id
JOIN j_character_class_player_units j ON j.class_id = da.class_id
JOIN characters c ON j.unit_id = c.unit_id
WHERE a.type = 'player-ability' AND c.id = $1
ORDER BY pa.id;


-- name: GetCharacterSgAbilityIDs :many
SELECT id FROM player_abilities WHERE standard_grid_char_id = sqlc.arg('character_id')::int ORDER BY id;


-- name: GetCharacterEgAbilityIDs :many
SELECT id FROM player_abilities WHERE expert_grid_char_id = sqlc.arg('character_id')::int ORDER BY id;


-- name: GetCharacterOverdriveCommandID :one
SELECT oc.id
FROM overdrive_commands oc
JOIN j_character_class_player_units j ON j.class_id = oc.character_class_id
JOIN characters c ON j.unit_id = c.unit_id
WHERE c.id = $1;


-- name: GetCharacterOverdriveAbilityIDs :many
SELECT jooa.overdrive_ability_id
FROM j_overdrives_overdrive_abilities jooa
JOIN overdrives o ON jooa.overdrive_id = o.id
JOIN j_character_class_player_units jcp ON jcp.class_id = o.character_class_id
JOIN characters c ON jcp.unit_id = c.unit_id
WHERE c.id = $1
ORDER BY jooa.overdrive_ability_id;


-- name: GetCharacterCelestialWeaponID :one
SELECT id FROM celestial_weapons WHERE character_id = sqlc.arg('character_id')::int ORDER BY id;


-- name: GetCharacterIDs :many
SELECT id FROM characters ORDER BY id;


-- name: GetCharacterIDsStoryBased :many
SELECT id FROM characters WHERE is_story_based = $1 ORDER BY id;


-- name: GetCharacterIDsCanFightUnderwater :many
SELECT id FROM characters WHERE can_fight_underwater = $1 ORDER BY id;





-- name: GetAeonCharClassIDs :many
SELECT j.class_id
FROM j_character_class_player_units j
JOIN aeons a ON a.unit_id = j.unit_id
WHERE a.id = $1
ORDER BY j.class_id;


-- name: GetAeonDefaultAbilityIDs :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_default_abilities da ON da.ability_id = a.id
JOIN j_character_class_player_units j ON j.class_id = da.class_id
JOIN aeons ae ON ae.unit_id = j.unit_id
WHERE a.type = 'player-ability' AND ae.id = $1
ORDER BY pa.id;


-- name: GetAeonCelestialWeaponID :one
SELECT id FROM celestial_weapons WHERE aeon_id = sqlc.arg('aeon_id')::int;


-- name: GetAeonAeonCommandIDs :many
SELECT DISTINCT j1.aeon_command_id
FROM j_aeon_commands_possible_abilities j1
JOIN j_character_class_player_units j2 ON j1.character_class_id = j2.class_id
JOIN aeons a ON a.unit_id = j2.unit_id
WHERE a.id = $1
ORDER BY j1.aeon_command_id;


-- name: GetAeonOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN j_character_class_player_units j ON j.class_id = o.character_class_id
JOIN aeons a ON a.unit_id = j.unit_id
WHERE a.id = $1
ORDER BY o.id;


-- name: GetAeonOverdriveAbilityIDs :many
SELECT jooa.overdrive_ability_id
FROM j_overdrives_overdrive_abilities jooa
JOIN overdrives o ON jooa.overdrive_id = o.id
JOIN j_character_class_player_units jcp ON jcp.class_id = o.character_class_id
JOIN aeons a ON a.unit_id = jcp.unit_id
WHERE a.id = $1
ORDER BY jooa.overdrive_ability_id;


-- name: GetAeonIDs :many
SELECT id FROM aeons ORDER BY id;


-- name: GetAeonIDsOptional :many
SELECT id FROM aeons WHERE is_optional = $1 ORDER BY id;





-- name: GetCharacterClassUnitIDs :many
SELECT unit_id FROM j_character_class_player_units WHERE class_id = $1 ORDER BY unit_id;


-- name: GetCharacterClassDefaultAbilityIDs :many
SELECT ability_id FROM j_default_abilities WHERE class_id = $1 ORDER BY ability_id;


-- name: GetCharacterClassLearnableAbilityIDs :many
SELECT pa.ability_id
FROM player_abilities pa
JOIN j_player_abilities_learned_by j ON j.player_ability_id = pa.id
WHERE j.character_class_id = $1

UNION

SELECT ma.ability_id
FROM misc_abilities ma
JOIN j_misc_abilities_learned_by j ON j.misc_ability_id = ma.id
WHERE j.character_class_id = $1
ORDER BY ability_id;


-- name: GetCharacterClassDefaultOverdriveIDs :many
SELECT id
FROM overdrives
WHERE character_class_id = sqlc.arg('class_id')::int
  AND unlock_condition IS NULL
ORDER BY id;


-- name: GetCharacterClassLearnableOverdriveIDs :many
SELECT id
FROM overdrives
WHERE character_class_id = sqlc.arg('class_id')::int
  AND unlock_condition IS NOT NULL
ORDER BY id;


-- name: GetCharacterClassSubmenuIDs :many
SELECT submenu_id FROM j_submenus_users WHERE character_class_id = $1 ORDER BY submenu_id;


-- name: GetCharacterClassesIDs :many
SELECT id FROM character_classes ORDER BY id;


-- name: GetCharacterClassesIDsByCategory :many
SELECT id FROM character_classes WHERE category = ANY(sqlc.narg('category')::character_class_category[]) ORDER BY id;
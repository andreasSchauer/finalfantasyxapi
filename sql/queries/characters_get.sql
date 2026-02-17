-- name: GetCharacterCharClassIDs :many
SELECT cc.id
FROM character_classes cc
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN characters c ON c.unit_id = pu.id
WHERE c.id = $1
ORDER BY cc.id;


-- name: GetCharacterDefaultAbilityIDs :many
SELECT pl.id
FROM player_abilities pl
JOIN abilities a ON pl.ability_id = a.id
JOIN default_abilities da ON da.ability_id = a.id
JOIN character_classes cc ON da.class_id = cc.id
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN characters c ON c.unit_id = pu.id
WHERE a.type = 'player-ability' AND c.id = $1
ORDER BY pl.id;


-- name: GetCharacterSgAbilityIDs :many
SELECT pl.id
FROM player_abilities pl
JOIN characters c ON pl.standard_grid_char_id = c.id
WHERE c.id = $1
ORDER BY pl.id;


-- name: GetCharacterEgAbilityIDs :many
SELECT pl.id
FROM player_abilities pl
JOIN characters c ON pl.expert_grid_char_id = c.id
WHERE c.id = $1
ORDER BY pl.id;


-- name: GetCharacterOverdriveCommandID :one
SELECT oc.id
FROM overdrive_commands oc
JOIN character_classes cc ON oc.character_class_id = cc.id
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN characters c ON c.unit_id = pu.id
WHERE c.id = $1;


-- name: GetCharacterCelestialWeaponID :one
SELECT cw.id
FROM celestial_weapons cw
JOIN characters c ON cw.character_id = c.id
WHERE c.id = $1;


-- name: GetCharacterIDs :many
SELECT id FROM characters ORDER BY id;


-- name: GetCharacterIDsStoryOnly :many
SELECT id FROM characters WHERE story_only = $1 ORDER BY id;


-- name: GetAeonCharClassIDs :many
SELECT cc.id
FROM character_classes cc
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN aeons a ON a.unit_id = pu.id
WHERE a.id = $1
ORDER BY cc.id;


-- name: GetAeonDefaultAbilityIDs :many
SELECT pl.id
FROM player_abilities pl
JOIN abilities a ON pl.ability_id = a.id
JOIN default_abilities da ON da.ability_id = a.id
JOIN character_classes cc ON da.class_id = cc.id
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN aeons ae ON ae.unit_id = pu.id
WHERE a.type = 'player-ability' AND ae.id = $1
ORDER BY pl.id;


-- name: GetAeonCelestialWeaponID :one
SELECT cw.id
FROM celestial_weapons cw
JOIN aeons a ON cw.aeon_id = a.id
WHERE a.id = $1;


-- name: GetAeonAeonCommandIDs :many
SELECT DISTINCT ac.id
FROM aeon_commands ac
JOIN j_aeon_commands_possible_abilities j1 ON j1.aeon_command_id = ac.id
JOIN character_classes cc ON j1.character_class_id = cc.id
JOIN j_character_class_player_units j2 ON j2.class_id = cc.id
JOIN player_units pu ON j2.unit_id = pu.id
JOIN aeons a ON a.unit_id = pu.id
WHERE a.id = $1
ORDER BY ac.id;


-- name: GetAeonOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN character_classes cc ON o.character_class_id = cc.id
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN aeons a ON a.unit_id = pu.id
WHERE a.id = $1
ORDER BY o.id;


-- name: GetAeonIDs :many
SELECT id FROM aeons ORDER BY id;


-- name: GetAeonIDsOptional :many
SELECT id FROM aeons WHERE is_optional = $1 ORDER BY id;
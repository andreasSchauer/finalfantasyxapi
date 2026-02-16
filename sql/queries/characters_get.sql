-- name: GetCharacterCharClassIDs :many
SELECT cc.id
FROM character_classes cc
JOIN j_player_units_character_class j ON j.class_id = cc.id
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
JOIN j_player_units_character_class j ON j.class_id = cc.id
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
JOIN j_player_units_character_class j ON j.class_id = cc.id
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



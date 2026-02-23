-- name: GetCharacterCharClassIDs :many
SELECT cc.id
FROM character_classes cc
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN characters c ON c.unit_id = pu.id
WHERE c.id = $1
ORDER BY cc.id;


-- name: GetCharacterDefaultAbilityIDs :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN default_abilities da ON da.ability_id = a.id
JOIN character_classes cc ON da.class_id = cc.id
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN characters c ON c.unit_id = pu.id
WHERE a.type = 'player-ability' AND c.id = $1
ORDER BY pa.id;


-- name: GetCharacterSgAbilityIDs :many
SELECT pa.id
FROM player_abilities pa
JOIN characters c ON pa.standard_grid_char_id = c.id
WHERE c.id = $1
ORDER BY pa.id;


-- name: GetCharacterEgAbilityIDs :many
SELECT pa.id
FROM player_abilities pa
JOIN characters c ON pa.expert_grid_char_id = c.id
WHERE c.id = $1
ORDER BY pa.id;


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


-- name: GetCharacterIDsCanFightUnderwater :many
SELECT id FROM characters WHERE can_fight_underwater = $1 ORDER BY id;




-- name: GetAeonCharClassIDs :many
SELECT cc.id
FROM character_classes cc
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN aeons a ON a.unit_id = pu.id
WHERE a.id = $1
ORDER BY cc.id;


-- name: GetAeonDefaultAbilityIDs :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN default_abilities da ON da.ability_id = a.id
JOIN character_classes cc ON da.class_id = cc.id
JOIN j_character_class_player_units j ON j.class_id = cc.id
JOIN player_units pu ON j.unit_id = pu.id
JOIN aeons ae ON ae.unit_id = pu.id
WHERE a.type = 'player-ability' AND ae.id = $1
ORDER BY pa.id;


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





-- name: GetCharacterClassUnitIDs :many
SELECT pu.id
FROM player_units pu
JOIN j_character_class_player_units j ON j.unit_id = pu.id
JOIN character_classes cc ON j.class_id = cc.id
WHERE cc.id = $1
ORDER BY pu.id;


-- name: GetCharacterClassDefaultAbilityIDs :many
SELECT a.id
FROM abilities a
JOIN default_abilities da ON da.ability_id = a.id
JOIN character_classes cc ON da.class_id = cc.id
WHERE cc.id = $1
ORDER BY a.id;


-- name: GetCharacterClassLearnableAbilityIDs :many
SELECT a.id
FROM abilities a
LEFT JOIN player_abilities pa ON pa.ability_id = a.id
LEFT JOIN j_player_abilities_learned_by j1 ON j1.player_ability_id = pa.id
LEFT JOIN other_abilities ga ON ga.ability_id = a.id
LEFT JOIN j_other_abilities_learned_by j2 ON j2.other_ability_id = ga.id
LEFT JOIN character_classes cc ON j1.character_class_id = cc.id OR j2.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY a.id;


-- name: GetCharacterClassDefaultOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN character_classes cc ON o.character_class_id = cc.id
WHERE cc.id = $1 AND o.unlock_condition IS NULL
ORDER BY o.id;


-- name: GetCharacterClassLearnableOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN character_classes cc ON o.character_class_id = cc.id
WHERE cc.id = $1 AND o.unlock_condition IS NOT NULL
ORDER BY o.id;


-- name: GetCharacterClassSubmenuIDs :many
SELECT s.id
FROM submenus s
JOIN j_submenus_users j ON j.submenu_id = s.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY s.id;


-- name: GetCharacterClassesIDs :many
SELECT id FROM character_classes ORDER BY id;


-- name: GetCharacterClassesIDsByCategory :many
SELECT id FROM character_classes WHERE category = $1 ORDER BY id;
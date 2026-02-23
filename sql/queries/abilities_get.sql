-- name: GetPlayerAbilityMonsterIDs :many
SELECT m.id
FROM monsters m
JOIN j_monsters_abilities j ON j.monster_id = m.id
JOIN monster_abilities ma ON j.monster_ability_id = ma.id
JOIN abilities a ON ma.ability_id = a.id
JOIN player_abilities pa ON pa.ability_id = a.id
WHERE pa.id = $1
ORDER BY m.id;


-- name: GetPlayerAbilityIDs :many
SELECT id FROM player_abilities ORDER BY id;


-- name: GetPlayerAbilityIDsByCategory :many
SELECT id FROM player_abilities WHERE category = $1 ORDER BY id;
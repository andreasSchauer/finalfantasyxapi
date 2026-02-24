-- name: GetPlayerAbilityIDsByName :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
WHERE a.name = $1;


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


-- name: GetPlayerAbilityIDsBasedOnPhysAttack :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_phys_attack = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsDarkable :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.battle_interaction_id = bi.id
WHERE j2.status_condition_id = 4
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByElement :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN elements e ON d.element_id = e.id
WHERE e.id = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByDamageType :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_type = $1
ORDER BY pa.id;
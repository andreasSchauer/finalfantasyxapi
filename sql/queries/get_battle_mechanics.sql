-- name: GetOverdriveModeIDs :many
SELECT id FROM overdrive_modes ORDER BY id;


-- name: GetOverdriveModeIDsByType :many
SELECT id FROM overdrive_modes WHERE type = $1 ORDER BY id;


-- name: GetAgilityTierByAgility :one
SELECT * FROM agility_tiers
WHERE (sqlc.arg(agility)::int) >= min_agility
AND (sqlc.arg(agility)::int) <= max_agility;







-- name: GetElementStatusConditionID :one
SELECT sc.id
FROM status_conditions sc
JOIN elemental_resists er ON sc.added_elem_resist_id = er.id
JOIN elements e ON er.element_id = e.id
WHERE e.id = $1;


-- name: GetElementAutoAbilityIDs :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
WHERE EXISTS (
    SELECT 1
    FROM elemental_resists er
    JOIN elements e ON er.element_id = e.id
    WHERE aa.added_elem_resist_id = er.id
      AND e.id = $1

    UNION ALL

    SELECT 1
    FROM elements e
    WHERE aa.on_hit_element_id = e.id
      AND e.id = $1
)
ORDER BY aa.id;


-- name: GetElementPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN elements e ON d.element_id = e.id
WHERE e.id = $1
ORDER BY pa.id;


-- name: GetElementOverdriveAbilityIDs :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN elements e ON d.element_id = e.id
WHERE e.id = $1
ORDER BY oa.id;


-- name: GetElementItemAbilityIDs :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN elements e ON d.element_id = e.id
WHERE e.id = $1
ORDER BY ia.id;


-- name: GetElementEnemyAbilityIDs :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN elements e ON d.element_id = e.id
WHERE e.id = $1
ORDER BY ea.id;


-- name: GetElementMonsterIDsWeak :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_elem_resists j ON j.monster_id = m.id
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'weak'
ORDER BY m.id;


-- name: GetElementMonsterIDsHalved :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_elem_resists j ON j.monster_id = m.id
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'halved'
ORDER BY m.id;


-- name: GetElementMonsterIDsImmune :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_elem_resists j ON j.monster_id = m.id
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'immune'
ORDER BY m.id;


-- name: GetElementMonsterIDsAbsorb :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_elem_resists j ON j.monster_id = m.id
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'absorb'
ORDER BY m.id;


-- name: GetElementIDs :many
SELECT id FROM elements ORDER BY id;
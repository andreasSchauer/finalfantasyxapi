-- name: GetPlayerAbilityIDsByName :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
WHERE a.name = $1;


-- name: GetPlayerAbilityMonsterIDs :many
SELECT DISTINCT m.id
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


-- name: GetPlayerAbilityIDsByMpCost :many
SELECT id FROM player_abilities WHERE mp_cost = $1 ORDER BY id;


-- name: GetPlayerAbilityIDsByMpCostMin :many
SELECT id FROM player_abilities WHERE mp_cost >= $1 ORDER BY id;


-- name: GetPlayerAbilityIDsByMpCostMax :many
SELECT id FROM player_abilities WHERE mp_cost <= $1 ORDER BY id;


-- name: GetPlayerAbilityIDsCanUseOutsideBattle :many
SELECT id FROM player_abilities WHERE can_use_outside_battle = $1 ORDER BY id;


-- name: GetPlayerAbilityIDsStdSgChar :many
SELECT id FROM player_abilities WHERE standard_grid_char_id = $1 ORDER BY id;


-- name: GetPlayerAbilityIDsExpSgChar :many
SELECT id FROM player_abilities WHERE expert_grid_char_id = $1 ORDER BY id;


-- name: GetPlayerAbilityIDsByCharClass :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_player_abilities_learned_by j ON j.player_ability_id = pa.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByLearnItem :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN item_amounts ia ON pa.aeon_learn_item_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByRelatedStat :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_player_abilities_related_stats j ON j.player_ability_id = pa.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByRank :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByCanCopycat :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.can_copycat = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsBasedOnPhysAttack :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_phys_attack = true
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsDarkable :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 4
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsSilenceable :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 13
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsReflectable :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 28
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsDealsDelay :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_delay j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON j2.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByInflictedStatus :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByRemovedStatus :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_removed_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsWithStatChanges :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsWithModifierChanges :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByElement :many
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


-- name: GetPlayerAbilityIDsByDamageType :many
SELECT DISTINCT pa.id
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


-- name: GetPlayerAbilityIDsByAttackType :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = $1
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByDamageFormula :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY pa.id;
-- name: GetEnemyAbilityIDsByName :many
SELECT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
WHERE a.name = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_abilities j ON j.monster_id = m.id
JOIN monster_abilities ma ON j.monster_ability_id = ma.id
JOIN abilities a ON ma.ability_id = a.id
JOIN enemy_abilities ea ON ea.ability_id = a.id
WHERE ea.id = $1
ORDER BY m.id;


-- name: GetEnemyAbilityIDs :many
SELECT id FROM enemy_abilities ORDER BY id;


-- name: GetEnemyAbilityIDsByMonster :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN monster_abilities ma ON ma.ability_id = a.id
JOIN j_monsters_abilities j ON j.monster_ability_id = ma.id
JOIN monsters m ON j.monster_id = m.id
WHERE m.id = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByRank :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsDarkable :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 4
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsSilenceable :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 13
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsReflectable :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 28
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsDealsDelay :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_delay j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON j2.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByInflictedStatus :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByRemovedStatus :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_removed_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByElement :many
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


-- name: GetEnemyAbilityIDsByDamageType :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_type = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByAttackType :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByDamageFormula :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY ea.id;







-- name: GetItemAbilityIDs :many
SELECT id FROM item_abilities ORDER BY id;


-- name: GetItemAbilityIDsCanUseOutsideBattle :many
SELECT ia.id
FROM item_abilities ia
JOIN items i ON ia.item_id = i.id
WHERE i.usability = 'always' OR i.usability = 'outside-battle'
ORDER BY ia.id;


-- name: GetItemAbilityIDsByCategory :many
SELECT ia.id
FROM item_abilities ia
JOIN items i ON ia.item_id = i.id
WHERE i.category = $1
ORDER BY ia.id;


-- name: GetItemAbilityIDsByRelatedStat :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_items_related_stats j ON j.player_ability_id = pa.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY ia.id;


-- name: GetItemAbilityIDsDealsDelay :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_delay j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON j2.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY ia.id;


-- name: GetItemAbilityIDsByInflictedStatus :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY ia.id;


-- name: GetItemAbilityIDsByRemovedStatus :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_removed_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY ia.id;


-- name: GetItemAbilityIDsWithStatChanges :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
ORDER BY ia.id;


-- name: GetItemAbilityIDsWithModifierChanges :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
ORDER BY ia.id;


-- name: GetItemAbilityIDsByElement :many
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


-- name: GetItemAbilityIDsByAttackType :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = $1
ORDER BY ia.id;


-- name: GetItemAbilityIDsByDamageFormula :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY ia.id;







-- name: GetOtherAbilityIDsByName :many
SELECT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
WHERE a.name = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDs :many
SELECT id FROM other_abilities ORDER BY id;


-- name: GetOtherAbilityIDsByCharClass :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN j_other_abilities_learned_by j ON j.other_ability_id = oa.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByRank :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByCanCopycat :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.can_copycat = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsBasedOnPhysAttack :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_phys_attack = true
ORDER BY oa.id;


-- name: GetOtherAbilityIDsDarkable :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 4
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByInflictedStatus :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsDealsDelay :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_delay j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON j2.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByDamageType :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_type = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByAttackType :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = $1
ORDER BY oa.id;


-- name: GetOtherAbilityIDsByDamageFormula :many
SELECT DISTINCT oa.id
FROM other_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY oa.id;


-- name: GetAbilityAttributes :many
SELECT aa.rank, aa.can_copycat, aa.appears_in_help_bar
FROM ability_attributes aa
JOIN abilities a ON a.attributes_id = aa.id
WHERE a.id = $1;


-- name: GetOverdriveAbilityAttributes :many
SELECT aa.rank, aa.can_copycat, aa.appears_in_help_bar
FROM ability_attributes aa
JOIN overdrives o ON o.attributes_id = aa.id
JOIn j_overdrives_overdrive_abilities j ON j.overdrive_id = o.id
JOIN overdrive_abilities oa ON j.overdrive_ability_id = oa.id
JOIN abilities a ON oa.ability_id = a.id
WHERE a.id = $1;


-- name: GetOverdriveAbilityOverdriveIDs :many
SELECT o.id
FROM overdrives o
JOIN j_overdrives_overdrive_abilities j ON j.overdrive_id = o.id
JOIN overdrive_abilities oa ON j.overdrive_ability_id = oa.id
WHERE oa.id = $1;


-- name: GetOverdriveAbilityIDsByName :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
WHERE a.name = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDs :many
SELECT id FROM overdrive_abilities ORDER BY id;


-- name: GetOverdriveAbilityIDsByCharClass :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_overdrives_overdrive_abilities j ON j.overdrive_ability_id = oa.id
JOIN overdrives o ON j.overdrive_id = o.id
JOIN character_classes cc ON o.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByRelatedStat :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_overdrive_abilities_related_stats j ON j.overdrive_ability_id = oa.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByRank :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsDealsDelay :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_delay j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON j2.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByInflictedStatus :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByRemovedStatus :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_removed_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsWithStatChanges :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsWithModifierChanges :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByElement :many
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


-- name: GetOverdriveAbilityIDsByAttackType :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByDamageFormula :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY oa.id;





-- name: GetPlayerAbilityIDsByName :many
SELECT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
WHERE a.name = $1
ORDER BY pa.id;


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




-- name: GetTriggerCommandIDsByName :many
SELECT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
WHERE a.name = $1
ORDER BY tc.id;


-- name: GetTriggerCommandMonsterFormationIDs :many
SELECT mf.id
FROM monster_formations mf
JOIN j_monster_formations_trigger_commands j ON j.monster_formation_id = mf.id
JOIN formation_trigger_commands ftc ON j.trigger_command_id = ftc.id
JOIN trigger_commands tc ON ftc.trigger_command_id = tc.id
WHERE tc.id = $1
ORDER BY mf.id;


-- name: GetTriggerCommandCharClassIDs :many
SELECT cc.id
FROM character_classes cc
JOIN j_formation_trigger_commands_users j ON j.character_class_id = cc.id
JOIN formation_trigger_commands ftc ON j.trigger_command_id = ftc.id
JOIN trigger_commands tc ON ftc.trigger_command_id = tc.id
WHERE tc.id = $1
ORDER BY cc.id;


-- name: GetTriggerCommandIDs :many
SELECT id FROM trigger_commands ORDER BY id;


-- name: GetTriggerCommandIDsByCharClass :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN formation_trigger_commands ftc ON ftc.trigger_command_id = tc.id
Join j_formation_trigger_commands_users j ON j.trigger_command_id = ftc.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY tc.id;


-- name: GetTriggerCommandIDsByRelatedStat :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN j_trigger_commands_related_stats j ON j.trigger_command = tc.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY tc.id;


-- name: GetTriggerCommandIDsByRank :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = $1
ORDER BY tc.id;


-- name: GetTriggerCommandIDsWithStatChanges :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
ORDER BY tc.id;


-- name: GetTriggerCommandIDsWithModifierChanges :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
ORDER BY tc.id;
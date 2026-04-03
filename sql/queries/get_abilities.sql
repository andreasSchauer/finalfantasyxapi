-- name: GetAbilityMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_abilities j ON j.monster_id = m.id
JOIN monster_abilities ma ON j.monster_ability_id = ma.id
JOIN abilities a ON ma.ability_id = a.id
WHERE a.id = $1
ORDER BY m.id;


-- name: GetAbilityIDs :many
SELECT id FROM abilities ORDER BY id;


-- name: GetAbilityIDsByMonster :many
SELECT DISTINCT a.id
FROM abilities a
JOIN monster_abilities ma ON ma.ability_id = a.id
JOIN j_monsters_abilities j ON j.monster_ability_id = ma.id
JOIN monsters m ON j.monster_id = m.id
WHERE m.id = $1
ORDER BY a.id;


-- name: GetAbilityIDsByType :many
SELECT id FROM abilities WHERE type = ANY(sqlc.narg('ability_type')::ability_type[]) ORDER BY id;


-- name: GetAbilityIDsByRank :many
SELECT DISTINCT a.id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY a.id;


-- name: GetAbilityIDsByCanCopycat :many
SELECT DISTINCT a.id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.can_copycat = $1
ORDER BY a.id;


-- name: GetAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT a.id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY a.id;


-- name: GetAbilityIDsBasedOnUserAttack :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_user_attack = true
ORDER BY a.id;


-- name: GetAbilityIDsByTargetType :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY a.id;


-- name: GetAbilityIDsDarkable :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 4
ORDER BY a.id;


-- name: GetAbilityIDsSilenceable :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 13
ORDER BY a.id;


-- name: GetAbilityIDsReflectable :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 28
ORDER BY a.id;


-- name: GetAbilityIDsDealsDelay :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY a.id;


-- name: GetAbilityIDsByInflictedStatus :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY a.id;


-- name: GetAbilityIDsByRemovedStatus :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_removed_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY a.id;


-- name: GetAbilityIDsWithStatChanges :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
ORDER BY a.id;


-- name: GetAbilityIDsWithModifierChanges :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
ORDER BY a.id;


-- name: GetAbilityIDsCanCrit :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
WHERE d.critical IS NOT NULL
ORDER BY a.id;


-- name: GetAbilityIDsBreakDmgLimit :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
WHERE d.break_dmg_limit IS NOT NULL
ORDER BY a.id;


-- name: GetAbilityIDsByElement :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN elements e ON d.element_id = e.id
WHERE e.id = $1
ORDER BY a.id;


-- name: GetAbilityIDsByDamageType :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
ORDER BY a.id;


-- name: GetAbilityIDsByAttackType :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY a.id;


-- name: GetAbilityIDsByDamageFormula :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY a.id;






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
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByTargetType :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
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
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
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


-- name: GetEnemyAbilityIDsCanCrit :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
WHERE d.critical IS NOT NULL
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsBreakDmgLimit :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
WHERE d.break_dmg_limit IS NOT NULL
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
WHERE ad.damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
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
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
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
WHERE i.category = ANY(sqlc.narg('category')::item_category[])
ORDER BY ia.id;


-- name: GetItemAbilityIDsByTargetType :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
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
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
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
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
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







-- name: GetUnspecifiedAbilityIDsByName :many
SELECT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
WHERE a.name = $1
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDs :many
SELECT id FROM unspecified_abilities ORDER BY id;


-- name: GetUnspecifiedAbilityIDsByCharClass :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN j_unspecified_abilities_learned_by j ON j.unspecified_ability_id = ua.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByRank :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByCanCopycat :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.can_copycat = $1
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsBasedOnUserAttack :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_user_attack = true
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByTargetType :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsDarkable :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_affected_by j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = 4
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByInflictedStatus :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsDealsDelay :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByAttackType :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY ua.id;


-- name: GetUnspecifiedAbilityIDsByDamageFormula :many
SELECT DISTINCT ua.id
FROM unspecified_abilities ua
JOIN abilities a ON ua.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
JOIN j_damages_damage_calc j3 ON j3.ability_id = a.id AND j3.battle_interaction_id = bi.id AND j3.damage_id = d.id
JOIN ability_damages ad ON j3.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY ua.id;





-- name: GetOverdriveAbilityAttributes :one
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
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByTargetType :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsDealsDelay :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
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


-- name: GetOverdriveAbilityIDsCanCrit :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_damage j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN damages d ON j2.damage_id = d.id
WHERE d.critical IS NOT NULL
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
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
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
SELECT id FROM player_abilities WHERE category = ANY(sqlc.narg('category')::player_ability_category[]) ORDER BY id;


-- name: GetPlayerAbilityIDsByMpCost :many
SELECT id FROM player_abilities WHERE mp_cost = ANY(sqlc.arg(mp_cost)::int[]) ORDER BY id;


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
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
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


-- name: GetPlayerAbilityIDsBasedOnUserAttack :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_user_attack = true
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByTargetType :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
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
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
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
WHERE ad.damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
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
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
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
JOIN j_trigger_commands_related_stats j ON j.trigger_command_id = tc.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY tc.id;


-- name: GetTriggerCommandIDsByRank :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY tc.id;


-- name: GetTriggerCommandIDsByTargetType :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
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






-- name: GetOverdriveOverdriveAbilityIDs :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN j_overdrives_overdrive_abilities j ON j.overdrive_ability_id = oa.id
JOIN overdrives o ON j.overdrive_id = o.id
WHERE o.id = $1
ORDER BY oa.id;


-- name: GetOverdriveIDs :many
SELECT id FROM overdrives ORDER BY id;


-- name: GetOverdriveIDsByRank :many
SELECT o.id
FROM overdrives o
JOIN ability_attributes aa ON o.attributes_id = aa.id
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY o.id;


-- name: GetOverdriveIDsByUser :many
SELECT o.id
FROM overdrives o
JOIN character_classes cc ON o.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY o.id;






-- name: GetRonsoRageIDs :many
SELECT id FROM ronso_rages ORDER BY id;


-- name: GetRonsoRageMonsterIDs :many
SELECT m.id
FROM monsters m
JOIN j_monsters_ronso_rages j ON j.monster_id = m.id
JOIN ronso_rages r ON j.ronso_rage_id = r.id
WHERE r.id = $1
ORDER BY m.id;
-- name: GetAbilityIdRankPairs :many
SELECT DISTINCT
    a.id AS ability_id,
    aa.rank AS rank
FROM abilities a
JOIN overdrive_abilities oa ON oa.ability_id = a.id
JOIN j_overdrives_overdrive_abilities j ON j.overdrive_ability_id = oa.id
JOIN overdrives o ON j.overdrive_id = o.id
JOIN ability_attributes aa ON o.attributes_id = aa.id
WHERE a.id = ANY(sqlc.arg('ability_ids')::int[])
ORDER BY a.id, aa.rank;


-- name: GetAbilityMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN j_monsters_abilities j ON j.monster_id = m.id
JOIN monster_abilities ma ON j.monster_ability_id = ma.id
WHERE ma.ability_id = $1
ORDER BY m.id;


-- name: GetAbilityIDs :many
SELECT id FROM abilities ORDER BY id;


-- name: GetAbilityIDsByMonster :many
SELECT DISTINCT a.id
FROM abilities a
JOIN monster_abilities ma ON ma.ability_id = a.id
JOIN j_monsters_abilities j ON j.monster_ability_id = ma.id
WHERE j.monster_id = $1
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
SELECT DISTINCT j.ability_id
FROM j_abilities_battle_interactions j
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_user_attack = true
ORDER BY j.ability_id;


-- name: GetAbilityIDsByTargetType :many
SELECT DISTINCT j.ability_id
FROM j_abilities_battle_interactions j
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY j.ability_id;


-- name: GetAbilityIDsDarkable :many
SELECT DISTINCT ability_id
FROM j_battle_interactions_affected_by
WHERE status_condition_id = 4
ORDER BY ability_id;


-- name: GetAbilityIDsSilenceable :many
SELECT DISTINCT ability_id
FROM j_battle_interactions_affected_by
WHERE status_condition_id = 13
ORDER BY ability_id;


-- name: GetAbilityIDsReflectable :many
SELECT DISTINCT ability_id
FROM j_battle_interactions_affected_by
WHERE status_condition_id = 28
ORDER BY ability_id;


-- name: GetAbilityIDsDealsDelay :many
SELECT DISTINCT j.ability_id
FROM j_abilities_battle_interactions j
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY j.ability_id;


-- name: GetAbilityIDsByInflictedStatus :many
SELECT j.ability_id
FROM j_battle_interactions_inflicted_status_conditions j
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.narg('status_id')
  AND sqlc.narg('status_id')::int IS NOT NULL 
  AND sqlc.narg('status_id')::int != 6

UNION

SELECT j.ability_id
FROM j_abilities_battle_interactions j
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.inflicted_delay_id IS NOT NULL
  AND sqlc.narg('status_id')::int = 6

UNION

SELECT a.id AS ability_id
FROM abilities a
WHERE sqlc.narg('status_id')::int IS NULL
  AND a.id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT j.ability_id FROM j_abilities_battle_interactions j 
    JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
    WHERE bi.inflicted_delay_id IS NOT NULL
)
ORDER BY ability_id;


-- name: GetAbilityIDsByRemovedStatus :many
SELECT ability_id
FROM j_battle_interactions_removed_status_conditions
WHERE sqlc.narg('status_id')::int IS NOT NULL
  AND status_condition_id = sqlc.narg('status_id')::int

UNION

SELECT a.id AS ability_id
FROM abilities a
WHERE sqlc.narg('status_id')::int IS NULL
  AND a.id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
ORDER BY ability_id;


-- name: GetAbilityIDsWithStatChanges :many
SELECT DISTINCT ability_id
FROM j_battle_interactions_stat_changes
ORDER BY ability_id;


-- name: GetAbilityIDsWithModifierChanges :many
SELECT DISTINCT ability_id
FROM j_battle_interactions_modifier_changes
ORDER BY ability_id;


-- name: GetAbilityIDsCanCrit :many
SELECT DISTINCT j.ability_id
FROM j_battle_interactions_damage j
JOIN damages d ON j.damage_id = d.id
WHERE d.critical IS NOT NULL
ORDER BY j.ability_id;


-- name: GetAbilityIDsBreakDmgLimit :many
SELECT DISTINCT j.ability_id
FROM j_battle_interactions_damage j
JOIN damages d ON j.damage_id = d.id
WHERE d.break_dmg_limit IS NOT NULL
ORDER BY j.ability_id;


-- name: GetAbilityIDsByElement :many
SELECT j.ability_id
FROM j_battle_interactions_damage j
JOIN damages d ON j.damage_id = d.id
WHERE sqlc.narg('element_id')::int[] IS NOT NULL
  AND d.element_id = ANY(sqlc.narg('element_id')::int[])

UNION

SELECT a.id AS ability_id
FROM abilities a
WHERE sqlc.narg('element_id')::int[] IS NULL
  AND a.id NOT IN (
    SELECT j.ability_id
    FROM j_battle_interactions_damage j
    JOIN damages d ON j.damage_id = d.id
    AND d.element_id IS NOT NULL
  )
ORDER BY ability_id;


-- name: GetAbilityIDsByDamageType :many
SELECT DISTINCT j.ability_id
FROM j_damages_damage_calc j
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
ORDER BY j.ability_id;


-- name: GetAbilityIDsByAttackType :many
SELECT DISTINCT j.ability_id
FROM j_damages_damage_calc j
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY j.ability_id;


-- name: GetAbilityIDsByDamageFormula :many
SELECT DISTINCT j.ability_id
FROM j_damages_damage_calc j
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY j.ability_id;






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
JOIN enemy_abilities ea ON ma.ability_id = ea.ability_id
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
WHERE j.monster_id = $1
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
JOIN j_abilities_battle_interactions j ON j.ability_id = ea.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsDarkable :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_affected_by j ON j.ability_id = ea.ability_id
WHERE j.status_condition_id = 4
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsSilenceable :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_affected_by j ON j.ability_id = ea.ability_id
WHERE j.status_condition_id = 13
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsReflectable :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_affected_by j ON j.ability_id = ea.ability_id
WHERE j.status_condition_id = 28
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsDealsDelay :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_abilities_battle_interactions j ON j.ability_id = ea.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByInflictedStatus :many
SELECT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_inflicted_status_conditions j ON j.ability_id = ea.ability_id
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.narg('status_id')
  AND sqlc.narg('status_id')::int IS NOT NULL 
  AND sqlc.narg('status_id')::int != 6

UNION

SELECT ea.id
FROM enemy_abilities ea
JOIN j_abilities_battle_interactions j ON j.ability_id = ea.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.inflicted_delay_id IS NOT NULL
  AND sqlc.narg('status_id')::int = 6

UNION

SELECT ea.id
FROM enemy_abilities ea
WHERE sqlc.narg('status_id')::int IS NULL
  AND ea.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT j.ability_id FROM j_abilities_battle_interactions j 
    JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
    WHERE bi.inflicted_delay_id IS NOT NULL
)
ORDER BY id;


-- name: GetEnemyAbilityIDsByRemovedStatus :many
SELECT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_removed_status_conditions j ON j.ability_id = ea.ability_id
WHERE sqlc.narg('status_id')::int IS NOT NULL
  AND status_condition_id = sqlc.narg('status_id')::int

UNION

SELECT ea.id
FROM enemy_abilities ea
WHERE sqlc.narg('status_id')::int IS NULL
  AND ea.ability_id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
ORDER BY id;


-- name: GetEnemyAbilityIDsCanCrit :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_damage j ON j.ability_id = ea.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.critical IS NOT NULL
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsBreakDmgLimit :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_damage j ON j.ability_id = ea.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.break_dmg_limit IS NOT NULL
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByElement :many
SELECT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_damage j ON j.ability_id = ea.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE sqlc.narg('element_id')::int[] IS NOT NULL
  AND d.element_id = ANY(sqlc.narg('element_id')::int[])

UNION

SELECT ea.id
FROM enemy_abilities ea
WHERE sqlc.narg('element_id')::int[] IS NULL
  AND ea.ability_id NOT IN (
    SELECT j.ability_id
    FROM j_battle_interactions_damage j
    JOIN damages d ON j.damage_id = d.id
    AND d.element_id IS NOT NULL
  )
ORDER BY id;


-- name: GetEnemyAbilityIDsByDamageType :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_damages_damage_calc j ON j.ability_id = ea.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByAttackType :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_damages_damage_calc j ON j.ability_id = ea.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY ea.id;


-- name: GetEnemyAbilityIDsByDamageFormula :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_damages_damage_calc j ON j.ability_id = ea.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
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
JOIN j_abilities_battle_interactions j ON j.ability_id = ia.ability_id
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
JOIN j_abilities_battle_interactions j ON j.ability_id = ia.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY ia.id;


-- name: GetItemAbilityIDsByInflictedStatus :many
SELECT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_inflicted_status_conditions j ON j.ability_id = ia.ability_id
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.narg('status_id')
  AND sqlc.narg('status_id')::int IS NOT NULL 
  AND sqlc.narg('status_id')::int != 6

UNION

SELECT ia.id
FROM item_abilities ia
JOIN j_abilities_battle_interactions j ON j.ability_id = ia.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.inflicted_delay_id IS NOT NULL
  AND sqlc.narg('status_id')::int = 6

UNION

SELECT ia.id
FROM item_abilities ia
WHERE sqlc.narg('status_id')::int IS NULL
  AND ia.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT j.ability_id FROM j_abilities_battle_interactions j 
    JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
    WHERE bi.inflicted_delay_id IS NOT NULL
)
ORDER BY id;


-- name: GetItemAbilityIDsByRemovedStatus :many
SELECT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_removed_status_conditions j ON j.ability_id = ia.ability_id
WHERE sqlc.narg('status_id')::int IS NOT NULL
  AND status_condition_id = sqlc.narg('status_id')::int

UNION

SELECT ia.id
FROM item_abilities ia
WHERE sqlc.narg('status_id')::int IS NULL
  AND ia.ability_id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
ORDER BY id;


-- name: GetItemAbilityIDsWithStatChanges :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_stat_changes j ON j.ability_id = ia.ability_id
ORDER BY ia.id;


-- name: GetItemAbilityIDsWithModifierChanges :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = ia.ability_id
ORDER BY ia.id;


-- name: GetItemAbilityIDsByElement :many
SELECT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_damage j ON j.ability_id = ia.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE sqlc.narg('element_id')::int[] IS NOT NULL
  AND d.element_id = ANY(sqlc.narg('element_id')::int[])

UNION

SELECT ia.id
FROM item_abilities ia
WHERE sqlc.narg('element_id')::int[] IS NULL
  AND ia.ability_id NOT IN (
    SELECT j.ability_id
    FROM j_battle_interactions_damage j
    JOIN damages d ON j.damage_id = d.id
    AND d.element_id IS NOT NULL
  )
ORDER BY id;


-- name: GetItemAbilityIDsByAttackType :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_damages_damage_calc j ON j.ability_id = ia.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY ia.id;


-- name: GetItemAbilityIDsByDamageFormula :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_damages_damage_calc j ON j.ability_id = ia.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY ia.id;







-- name: GetMiscAbilityIDsByName :many
SELECT ma.id
FROM misc_abilities ma
JOIN abilities a ON ma.ability_id = a.id
WHERE a.name = $1
ORDER BY ma.id;


-- name: GetMiscAbilityIDs :many
SELECT id FROM misc_abilities ORDER BY id;


-- name: GetMiscAbilityIDsByCharClass :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_misc_abilities_learned_by j ON j.misc_ability_id = ma.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByRank :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN abilities a ON ma.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByCanCopycat :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN abilities a ON ma.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.can_copycat = $1
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN abilities a ON ma.ability_id = a.id
JOIN ability_attributes aa ON a.attributes_id = aa.id
WHERE aa.appears_in_help_bar = $1
ORDER BY ma.id;


-- name: GetMiscAbilityIDsBasedOnUserAttack :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_abilities_battle_interactions j ON j.ability_id = ma.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_user_attack = true
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByTargetType :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_abilities_battle_interactions j ON j.ability_id = ma.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY ma.id;


-- name: GetMiscAbilityIDsDarkable :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_battle_interactions_affected_by j ON j.ability_id = ma.ability_id
WHERE j.status_condition_id = 4
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByInflictedStatus :many
SELECT ma.id
FROM misc_abilities ma
JOIN j_battle_interactions_inflicted_status_conditions j ON j.ability_id = ma.ability_id
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.narg('status_id')
  AND sqlc.narg('status_id')::int IS NOT NULL 
  AND sqlc.narg('status_id')::int != 6

UNION

SELECT ma.id
FROM misc_abilities ma
JOIN j_abilities_battle_interactions j ON j.ability_id = ma.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.inflicted_delay_id IS NOT NULL
  AND sqlc.narg('status_id')::int = 6

UNION

SELECT ma.id
FROM misc_abilities ma
WHERE sqlc.narg('status_id')::int IS NULL
  AND ma.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT j.ability_id FROM j_abilities_battle_interactions j 
    JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
    WHERE bi.inflicted_delay_id IS NOT NULL
)
ORDER BY id;


-- name: GetMiscAbilityIDsDealsDelay :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_abilities_battle_interactions j ON j.ability_id = ma.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByAttackType :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_damages_damage_calc j ON j.ability_id = ma.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY ma.id;


-- name: GetMiscAbilityIDsByDamageFormula :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_damages_damage_calc j ON j.ability_id = ma.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY ma.id;






-- name: GetOverdriveAbilityIdRankPairs :many
SELECT DISTINCT
    oa.id AS overdrive_ability_id,
    aa.rank AS rank
FROM overdrive_abilities oa
JOIN j_overdrives_overdrive_abilities j ON j.overdrive_ability_id = oa.id
JOIN overdrives o ON j.overdrive_id = o.id
JOIN ability_attributes aa ON o.attributes_id = aa.id
WHERE oa.id = ANY(sqlc.arg('overdrive_ability_ids')::int[])
ORDER BY oa.id, aa.rank;


-- name: GetOverdriveAbilityIDsByName :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
WHERE a.name = $1
ORDER BY oa.id;


-- name: GetOverdriveAbilityOverdriveIDs :many
SELECT overdrive_id
FROM j_overdrives_overdrive_abilities
WHERE overdrive_ability_id = $1
ORDER BY overdrive_id;


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
JOIN j_abilities_battle_interactions j ON j.ability_id = oa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsDealsDelay :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_abilities_battle_interactions j ON j.ability_id = oa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByInflictedStatus :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_inflicted_status_conditions j ON j.ability_id = oa.ability_id
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.narg('status_id')
  AND sqlc.narg('status_id')::int IS NOT NULL 
  AND sqlc.narg('status_id')::int != 6

UNION

SELECT oa.id
FROM overdrive_abilities oa
JOIN j_abilities_battle_interactions j ON j.ability_id = oa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.inflicted_delay_id IS NOT NULL
  AND sqlc.narg('status_id')::int = 6

UNION

SELECT oa.id
FROM overdrive_abilities oa
WHERE sqlc.narg('status_id')::int IS NULL
  AND oa.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT j.ability_id FROM j_abilities_battle_interactions j 
    JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
    WHERE bi.inflicted_delay_id IS NOT NULL
)
ORDER BY id;


-- name: GetOverdriveAbilityIDsByRemovedStatus :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_removed_status_conditions j ON j.ability_id = oa.ability_id
WHERE sqlc.narg('status_id')::int IS NOT NULL
  AND status_condition_id = sqlc.narg('status_id')::int

UNION

SELECT oa.id
FROM overdrive_abilities oa
WHERE sqlc.narg('status_id')::int IS NULL
  AND oa.ability_id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
ORDER BY id;


-- name: GetOverdriveAbilityIDsCanCrit :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_damage j ON j.ability_id = oa.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.critical IS NOT NULL
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsWithStatChanges :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_stat_changes j ON j.ability_id = oa.ability_id
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsWithModifierChanges :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = oa.ability_id
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByElement :many
SELECT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_damage j ON j.ability_id = oa.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE sqlc.narg('element_id')::int[] IS NOT NULL
  AND d.element_id = ANY(sqlc.narg('element_id')::int[])

UNION

SELECT oa.id
FROM overdrive_abilities oa
WHERE sqlc.narg('element_id')::int[] IS NULL
  AND oa.ability_id NOT IN (
    SELECT j.ability_id
    FROM j_battle_interactions_damage j
    JOIN damages d ON j.damage_id = d.id
    AND d.element_id IS NOT NULL
  )
ORDER BY id;


-- name: GetOverdriveAbilityIDsByAttackType :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_damages_damage_calc j ON j.ability_id = oa.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY oa.id;


-- name: GetOverdriveAbilityIDsByDamageFormula :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_damages_damage_calc j ON j.ability_id = oa.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
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
JOIN player_abilities pa ON ma.ability_id = pa.ability_id
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
JOIN j_abilities_battle_interactions j ON j.ability_id = pa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.based_on_user_attack = true
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByTargetType :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_abilities_battle_interactions j ON j.ability_id = pa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsDarkable :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_affected_by j ON j.ability_id = pa.ability_id
WHERE j.status_condition_id = 4
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsSilenceable :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_affected_by j ON j.ability_id = pa.ability_id
WHERE j.status_condition_id = 13
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsReflectable :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_affected_by j ON j.ability_id = pa.ability_id
WHERE j.status_condition_id = 28
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsDealsDelay :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_abilities_battle_interactions j ON j.ability_id = pa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
WHERE idl.ctb_attack_type = 'attack'
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByInflictedStatus :many
SELECT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_inflicted_status_conditions j ON j.ability_id = pa.ability_id
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.narg('status_id')
  AND sqlc.narg('status_id')::int IS NOT NULL 
  AND sqlc.narg('status_id')::int != 6

UNION

SELECT pa.id
FROM player_abilities pa
JOIN j_abilities_battle_interactions j ON j.ability_id = pa.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.inflicted_delay_id IS NOT NULL
  AND sqlc.narg('status_id')::int = 6

UNION

SELECT pa.id
FROM player_abilities pa
WHERE sqlc.narg('status_id')::int IS NULL
  AND pa.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT j.ability_id FROM j_abilities_battle_interactions j 
    JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
    WHERE bi.inflicted_delay_id IS NOT NULL
)
ORDER BY id;


-- name: GetPlayerAbilityIDsByRemovedStatus :many
SELECT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_removed_status_conditions j ON j.ability_id = pa.ability_id
WHERE sqlc.narg('status_id')::int IS NOT NULL
  AND status_condition_id = sqlc.narg('status_id')::int

UNION

SELECT pa.id
FROM player_abilities pa
WHERE sqlc.narg('status_id')::int IS NULL
  AND pa.ability_id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
ORDER BY id;


-- name: GetPlayerAbilityIDsWithStatChanges :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_stat_changes j ON j.ability_id = pa.ability_id
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsWithModifierChanges :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = pa.ability_id
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByElement :many
SELECT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_damage j ON j.ability_id = pa.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE sqlc.narg('element_id')::int[] IS NOT NULL
  AND d.element_id = ANY(sqlc.narg('element_id')::int[])

UNION

SELECT pa.id
FROM player_abilities pa
WHERE sqlc.narg('element_id')::int[] IS NULL
  AND pa.ability_id NOT IN (
    SELECT j.ability_id
    FROM j_battle_interactions_damage j
    JOIN damages d ON j.damage_id = d.id
    AND d.element_id IS NOT NULL
  )
ORDER BY id;


-- name: GetPlayerAbilityIDsByDamageType :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_damages_damage_calc j ON j.ability_id = pa.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByAttackType :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_damages_damage_calc j ON j.ability_id = pa.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
ORDER BY pa.id;


-- name: GetPlayerAbilityIDsByDamageFormula :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_damages_damage_calc j ON j.ability_id = pa.ability_id
JOIN ability_damages ad ON j.ability_damage_id = ad.id
WHERE ad.damage_formula = $1
ORDER BY pa.id;




-- name: GetTriggerCommandIDsByName :many
SELECT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
WHERE a.name = $1
ORDER BY tc.id;


-- name: GetTriggerCommandMonsterFormationIDs :many
SELECT j.monster_formation_id
FROM j_monster_formations_trigger_commands j
JOIN formation_trigger_commands ftc ON j.trigger_command_id = ftc.id
WHERE ftc.trigger_command_id = $1
ORDER BY j.monster_formation_id;


-- name: GetTriggerCommandCharClassIDs :many
SELECT j.character_class_id
FROM j_formation_trigger_commands_users j
JOIN formation_trigger_commands ftc ON j.trigger_command_id = ftc.id
WHERE ftc.trigger_command_id = $1
ORDER BY j.character_class_id;


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
JOIN j_abilities_battle_interactions j ON j.ability_id = tc.ability_id
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE bi.target = ANY(sqlc.narg('target_type')::target_type[])
ORDER BY tc.id;


-- name: GetTriggerCommandIDsWithStatChanges :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN j_battle_interactions_stat_changes j ON j.ability_id = tc.ability_id
ORDER BY tc.id;


-- name: GetTriggerCommandIDsWithModifierChanges :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = tc.ability_id
ORDER BY tc.id;






-- name: GetOverdriveOverdriveAbilityIDs :many
SELECT overdrive_ability_id
FROM j_overdrives_overdrive_abilities
WHERE overdrive_id = $1
ORDER BY overdrive_ability_id;


-- name: GetOverdriveIDs :many
SELECT id FROM overdrives ORDER BY id;


-- name: GetOverdriveIDsByRank :many
SELECT o.id
FROM overdrives o
JOIN ability_attributes aa ON o.attributes_id = aa.id
WHERE aa.rank = ANY(sqlc.arg(rank)::int[])
ORDER BY o.id;


-- name: GetOverdriveIDsByUser :many
SELECT id FROM overdrives WHERE character_class_id = $1 ORDER BY id;




-- name: GetRonsoRageIDs :many
SELECT id FROM ronso_rages ORDER BY id;


-- name: GetRonsoRageMonsterIDs :many
SELECT m.id
FROM monsters m
JOIN j_monsters_ronso_rages j ON j.monster_id = m.id
WHERE j.ronso_rage_id = $1
ORDER BY m.id;
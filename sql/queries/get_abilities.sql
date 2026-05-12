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
SELECT DISTINCT ability_id FROM mv_abilities WHERE rank = ANY(sqlc.arg(rank)::int[]) ORDER BY ability_id;


-- name: GetAbilityIDsByCanCopycat :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE can_copycat = $1 ORDER BY ability_id;


-- name: GetAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE appears_in_help_bar = $1 ORDER BY ability_id;


-- name: GetAbilityIDsBasedOnUserAttack :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE based_on_user_attack = true ORDER BY ability_id;


-- name: GetAbilityIDsByTargetType :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE target = ANY(sqlc.narg('target_type')::target_type[]) ORDER BY ability_id;


-- name: GetAbilityIDsDarkable :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE affected_status_id = 4 ORDER BY ability_id;


-- name: GetAbilityIDsSilenceable :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE affected_status_id = 13 ORDER BY ability_id;


-- name: GetAbilityIDsReflectable :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE affected_status_id = 28 ORDER BY ability_id;


-- name: GetAbilityIDsDealsDelay :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE ctb_attack_type = 'attack' ORDER BY ability_id;


-- name: GetAbilityIDsByInflictedStatus :many
WITH w AS (
    SELECT sqlc.narg('status_id')::int AS status_id
)
SELECT j.ability_id
FROM j_battle_interactions_inflicted_status_conditions j
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
CROSS JOIN w
WHERE ist.status_condition_id = w.status_id
  AND w.status_id IS NOT NULL 
  AND w.status_id != 6

UNION

SELECT a.ability_id
FROM mv_abilities a
CROSS JOIN w
WHERE a.inflicted_delay_id IS NOT NULL
  AND w.status_id = 6

UNION

SELECT a.ability_id
FROM mv_abilities a
CROSS JOIN w
WHERE w.status_id IS NULL
  AND a.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT ability_id FROM mv_abilities
    WHERE inflicted_delay_id IS NOT NULL
)
ORDER BY ability_id;


-- name: GetAbilityIDsByRemovedStatus :many
WITH w AS (
    SELECT sqlc.narg('status_id')::int AS status_id
)
SELECT ability_id
FROM j_battle_interactions_removed_status_conditions
CROSS JOIN w
WHERE w.status_id IS NOT NULL
  AND status_condition_id = w.status_id

UNION

SELECT a.id AS ability_id
FROM abilities a
CROSS JOIN w
WHERE w.status_id IS NULL
  AND a.id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
ORDER BY ability_id;


-- name: GetAbilityIDsWithStatChanges :many
SELECT DISTINCT ability_id FROM j_battle_interactions_stat_changes ORDER BY ability_id;


-- name: GetAbilityIDsWithModifierChanges :many
SELECT DISTINCT ability_id FROM j_battle_interactions_modifier_changes ORDER BY ability_id;


-- name: GetAbilityIDsCanCrit :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE critical IS NOT NULL ORDER BY ability_id;


-- name: GetAbilityIDsBreakDmgLimit :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE break_dmg_limit IS NOT NULL ORDER BY ability_id;


-- name: GetAbilityIDsByElement :many
WITH w AS (
    SELECT sqlc.narg('element_id')::int[] AS ids
)
SELECT a.ability_id
FROM mv_abilities a
CROSS JOIN w
WHERE w.ids IS NOT NULL
  AND a.element_id = ANY(w.ids)

UNION

SELECT a.id AS ability_id
FROM abilities a
CROSS JOIN w
WHERE w.ids IS NULL
  AND a.id NOT IN (
    SELECT ability_id FROM mv_abilities WHERE element_id IS NOT NULL
  )
ORDER BY ability_id;


-- name: GetAbilityIDsByDamageType :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE damage_type = ANY(sqlc.narg('damage_type')::damage_type[]) ORDER BY ability_id;


-- name: GetAbilityIDsByAttackType :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE attack_type = ANY(sqlc.narg('attack_type')::attack_type[]) ORDER BY ability_id;


-- name: GetAbilityIDsByDamageFormula :many
SELECT DISTINCT ability_id FROM mv_abilities WHERE damage_formula = sqlc.arg('damage_formula')::damage_formula ORDER BY ability_id;







-- name: GetTypedAbilityIDsByName :many
SELECT DISTINCT typed_id FROM mv_abilities WHERE name = $1 AND type = sqlc.arg('type')::ability_type ORDER BY typed_id;


-- name: GetTypedAbilityIDsByRank :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE rank = ANY(sqlc.arg(rank)::int[])
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByCanCopycat :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE can_copycat = $1
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByAppearsInHelpBar :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE appears_in_help_bar = $1
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsBasedOnUserAttack :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE based_on_user_attack = true
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByTargetType :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE target = ANY(sqlc.narg('target_type')::target_type[])
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsDarkable :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE affected_status_id = 4
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsSilenceable :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE affected_status_id = 13
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsReflectable :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE affected_status_id = 28
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsDealsDelay :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE ctb_attack_type = 'attack'
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByInflictedStatus :many
WITH w AS (
    SELECT
      sqlc.narg('status_id')::int AS status_id,
      sqlc.arg('type')::ability_type AS ability_type
)
SELECT a.typed_id
FROM mv_abilities a
JOIN j_battle_interactions_inflicted_status_conditions j ON j.ability_id = a.ability_id
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
CROSS JOIN w
WHERE ist.status_condition_id = w.status_id
  AND w.status_id IS NOT NULL 
  AND w.status_id != 6
  AND a.type = w.ability_type

UNION

SELECT a.typed_id
FROM mv_abilities a
CROSS JOIN w
WHERE a.inflicted_delay_id IS NOT NULL
  AND w.status_id = 6
  AND a.type = w.ability_type

UNION

SELECT a.typed_id
FROM mv_abilities a
CROSS JOIN w
WHERE w.status_id IS NULL
  AND a.ability_id NOT IN (
    SELECT ability_id FROM j_battle_interactions_inflicted_status_conditions
    UNION
    SELECT ability_id FROM mv_abilities
    WHERE inflicted_delay_id IS NOT NULL
  )
  AND a.type = w.ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByRemovedStatus :many
WITH w AS (
    SELECT
      sqlc.narg('status_id')::int AS status_id,
      sqlc.arg('type')::ability_type AS ability_type
)
SELECT a.typed_id
FROM mv_abilities a
JOIN j_battle_interactions_removed_status_conditions j ON j.ability_id = a.ability_id
CROSS JOIN w
WHERE w.status_id IS NOT NULL
  AND j.status_condition_id = w.status_id
  AND a.type = w.ability_type

UNION

SELECT a.typed_id
FROM mv_abilities a
CROSS JOIN w
WHERE w.status_id IS NULL
  AND a.ability_id NOT IN (
    SELECT ability_id
    FROM j_battle_interactions_removed_status_conditions
  )
  AND a.type = w.ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByElement :many
WITH w AS (
    SELECT sqlc.narg('element_id')::int[] AS ids,
    sqlc.arg('type')::ability_type AS ability_type
)
SELECT a.typed_id
FROM mv_abilities a
CROSS JOIN w
WHERE w.ids IS NOT NULL
  AND a.element_id = ANY(w.ids)
  AND a.type = w.ability_type

UNION

SELECT a.typed_id
FROM mv_abilities a
CROSS JOIN w
WHERE w.ids IS NULL
  AND a.ability_id NOT IN (
    SELECT ability_id FROM mv_abilities WHERE element_id IS NOT NULL
  )
  AND a.type = w.ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsWithStatChanges :many
SELECT DISTINCT a.typed_id
FROM mv_abilities a
JOIN j_battle_interactions_stat_changes j ON j.ability_id = a.ability_id
AND type = sqlc.arg('type')::ability_type
ORDER BY a.typed_id;


-- name: GetTypedAbilityIDsWithModifierChanges :many
SELECT DISTINCT a.typed_id
FROM mv_abilities a
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = a.ability_id
AND type = sqlc.arg('type')::ability_type
ORDER BY a.typed_id;


-- name: GetTypedAbilityIDsCanCrit :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE critical IS NOT NULL
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsBreakDmgLimit :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE break_dmg_limit IS NOT NULL
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByDamageType :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE damage_type = ANY(sqlc.narg('damage_type')::damage_type[])
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByAttackType :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE attack_type = ANY(sqlc.narg('attack_type')::attack_type[])
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;


-- name: GetTypedAbilityIDsByDamageFormula :many
SELECT DISTINCT typed_id
FROM mv_abilities
WHERE damage_formula = sqlc.arg('damage_formula')::damage_formula
AND type = sqlc.arg('type')::ability_type
ORDER BY typed_id;







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


-- name: GetItemAbilityIDsByRelatedStat :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_items_related_stats j ON j.player_ability_id = pa.id
JOIN stats s ON j.stat_id = s.id
WHERE s.id = $1
ORDER BY ia.id;







-- name: GetMiscAbilityIDs :many
SELECT id FROM misc_abilities ORDER BY id;


-- name: GetMiscAbilityIDsByCharClass :many
SELECT DISTINCT ma.id
FROM misc_abilities ma
JOIN j_misc_abilities_learned_by j ON j.misc_ability_id = ma.id
JOIN character_classes cc ON j.character_class_id = cc.id
WHERE cc.id = $1
ORDER BY ma.id;







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
SELECT DISTINCT ftc.trigger_command_id
FROM formation_trigger_commands ftc
JOIN j_formation_trigger_commands_users j ON j.trigger_command_id = ftc.id
WHERE j.character_class_id = $1
ORDER BY ftc.trigger_command_id;


-- name: GetTriggerCommandIDsByRelatedStat :many
SELECT DISTINCT trigger_command_id FROM j_trigger_commands_related_stats WHERE stat_id = $1 ORDER BY trigger_command_id;







-- name: GetOverdriveOverdriveAbilityIDs :many
SELECT overdrive_ability_id FROM j_overdrives_overdrive_abilities WHERE overdrive_id = $1 ORDER BY overdrive_ability_id;


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
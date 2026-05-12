-- name: GetOverdriveModeIDs :many
SELECT id FROM overdrive_modes ORDER BY id;


-- name: GetOverdriveModeIDsByType :many
SELECT id FROM overdrive_modes WHERE type = $1 ORDER BY id;





-- name: GetAgilityTierIDs :many
SELECT id FROM agility_tiers ORDER BY id;


-- name: GetAgilityTierIDsByAgility :many
SELECT id FROM agility_tiers
WHERE (sqlc.arg(agility)::int) >= min_agility
  AND (sqlc.arg(agility)::int) <= max_agility
ORDER BY id;


-- name: GetAgilityTierByAgility :one
SELECT * FROM agility_tiers
WHERE (sqlc.arg(agility)::int) >= min_agility
  AND (sqlc.arg(agility)::int) <= max_agility;







-- name: GetElementStatusConditionID :one
SELECT DISTINCT sc.id
FROM status_conditions sc
JOIN elemental_resists er ON sc.added_elem_resist_id = er.id
WHERE er.element_id = $1;


-- name: GetElementAutoAbilityIDs :many
SELECT aa.id
FROM auto_abilities aa
JOIN elemental_resists er ON aa.added_elem_resist_id = er.id
WHERE er.element_id = $1

UNION

SELECT id
FROM auto_abilities
WHERE on_hit_element_id = $1

ORDER BY id;


-- name: GetElementPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_damage j ON j.ability_id = pa.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.element_id = $1
ORDER BY pa.id;


-- name: GetElementOverdriveAbilityIDs :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_damage j ON j.ability_id = oa.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.element_id = $1
ORDER BY oa.id;


-- name: GetElementItemAbilityIDs :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_damage j ON j.ability_id = ia.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.element_id = $1
ORDER BY ia.id;


-- name: GetElementEnemyAbilityIDs :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_damage j ON j.ability_id = ea.ability_id
JOIN damages d ON j.damage_id = d.id
WHERE d.element_id = $1
ORDER BY ea.id;


-- name: GetElementMonsterIDsWeak :many
SELECT DISTINCT j.monster_id
FROM j_monsters_elem_resists j
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'weak'
ORDER BY j.monster_id;


-- name: GetElementMonsterIDsHalved :many
SELECT DISTINCT j.monster_id
FROM j_monsters_elem_resists j
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'halved'
ORDER BY j.monster_id;


-- name: GetElementMonsterIDsImmune :many
SELECT DISTINCT j.monster_id
FROM j_monsters_elem_resists j
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'immune'
ORDER BY j.monster_id;


-- name: GetElementMonsterIDsAbsorb :many
SELECT DISTINCT j.monster_id
FROM j_monsters_elem_resists j
JOIN elemental_resists er ON j.elem_resist_id = er.id
WHERE er.element_id = $1 AND er.affinity = 'absorb'
ORDER BY j.monster_id;


-- name: GetElementIDs :many
SELECT id FROM elements ORDER BY id;







-- name: GetStatusConditionAutoAbilityIDs :many
SELECT jast.auto_ability_id
FROM j_auto_abilities_added_statusses jast
WHERE jast.status_condition_id = $1

UNION

SELECT jasr.auto_ability_id
FROM j_auto_abilities_added_status_resists jasr
JOIN status_resists sr ON jasr.status_resist_id = sr.id
WHERE sr.status_condition_id = $1

UNION

SELECT aa.id AS auto_ability_id
FROM auto_abilities aa
JOIN inflicted_statusses ist ON aa.on_hit_status_id = ist.id
WHERE ist.status_condition_id = $1

ORDER BY auto_ability_id;


-- name: GetStatusConditionAbilityIDsInflicted :many
SELECT j.ability_id
FROM j_battle_interactions_inflicted_status_conditions j
JOIN inflicted_statusses ist ON j.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.arg(status_condition_id)::int
  AND ist.probability BETWEEN
        sqlc.arg('min_rate')::int
    AND sqlc.arg('max_rate')::int

UNION

SELECT j.ability_id
FROM j_abilities_battle_interactions j
JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
WHERE sqlc.arg(status_condition_id)::int = 6 AND bi.inflicted_delay_id IS NOT NULL
ORDER BY ability_id;


-- name: GetStatusConditionAbilityIDsRemoved :many
SELECT ability_id
FROM j_battle_interactions_removed_status_conditions
WHERE status_condition_id = $1
ORDER BY ability_id;


-- name: GetStatusConditionInflictedDelayConditionIDs :many
SELECT DISTINCT sc.id
FROM status_conditions sc
JOIN inflicted_delays idl ON sc.inflicted_delay_id = idl.id
WHERE sqlc.arg(status_id)::int = 6
  AND sc.inflicted_delay_id IS NOT NULL
  AND idl.ctb_attack_type = 'attack'
ORDER BY sc.id;


-- name: GetStatusConditionRemovedConditionIDs :many
SELECT parent_condition_id
FROM j_status_conditions_removed_status_conditions
WHERE child_condition_id = $1
ORDER BY parent_condition_id;


-- name: GetStatusConditionResistingMonsterIDs :many
SELECT jmi.monster_id
FROM j_monsters_immunities jmi
WHERE jmi.status_condition_id = $1

UNION

SELECT jsr.monster_id
FROM j_monsters_status_resists jsr
JOIN status_resists sr ON sr.id = jsr.status_resist_id
WHERE sr.status_condition_id = $1
  AND sr.resistance >= sqlc.arg('min_resistance')::int
ORDER BY monster_id;


-- name: GetStatusConditionIDs :many
SELECT id FROM status_conditions ORDER BY id;


-- name: GetStatusConditionIDsByCategory :many
SELECT id FROM status_conditions WHERE category = ANY(sqlc.arg('category')::status_condition_category[]) ORDER BY id;






-- name: GetStatSphereIDs :many
SELECT DISTINCT sph.id
FROM spheres sph
JOIN j_items_related_stats j ON j.item_id = sph.item_id
WHERE j.stat_id = $1
ORDER BY sph.id;


-- name: GetStatAutoAbilityIDs :many
SELECT auto_ability_id
FROM j_auto_abilities_related_stats
WHERE stat_id = $1
ORDER BY auto_ability_id;


-- name: GetStatStatusConditionIDs :many
SELECT status_condition_id
FROM j_status_conditions_related_stats
WHERE stat_id = $1
ORDER BY status_condition_id;


-- name: GetStatPropertyIDs :many
SELECT property_id
FROM j_properties_related_stats
WHERE stat_id = $1
ORDER BY property_id;


-- name: GetStatPlayerAbilityIDs :many
SELECT player_ability_id
FROM j_player_abilities_related_stats
WHERE stat_id = $1
ORDER BY player_ability_id;


-- name: GetStatOverdriveAbilityIDs :many
SELECT overdrive_ability_id
FROM j_overdrive_abilities_related_stats
WHERE stat_id = $1
ORDER BY overdrive_ability_id;


-- name: GetStatItemAbilityIDs :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_items_related_stats j ON j.item_id = ia.item_id
WHERE j.stat_id = $1
ORDER BY ia.id;


-- name: GetStatTriggerCommandIDs :many
SELECT trigger_command_id
FROM j_trigger_commands_related_stats
WHERE stat_id = $1
ORDER BY trigger_command_id;


-- name: GetStatAutoAbilityIDsStatChange :many
SELECT j.auto_ability_id
FROM j_auto_abilities_stat_changes j
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY j.auto_ability_id;


-- name: GetStatStatusConditionIDsStatChange :many
SELECT j.status_condition_id
FROM j_status_conditions_stat_changes j
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY j.status_condition_id;


-- name: GetStatPlayerAbilityIDsStatChange :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_stat_changes j ON j.ability_id = pa.ability_id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY pa.id;


-- name: GetStatOverdriveAbilityIDsStatChange :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_stat_changes j ON j.ability_id = oa.ability_id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY oa.id;


-- name: GetStatItemAbilityIDsStatChange :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_stat_changes j ON j.ability_id = ia.ability_id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY ia.id;


-- name: GetStatTriggerCommandIDsStatChange :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN j_battle_interactions_stat_changes j ON j.ability_id = tc.ability_id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY tc.id;


-- name: GetStatIDs :many
SELECT id FROM stats ORDER BY id;








-- name: GetModifierAutoAbilityIDs :many
SELECT DISTINCT j.auto_ability_id
FROM j_auto_abilities_modifier_changes j
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY j.auto_ability_id;


-- name: GetModifierStatusConditionIDs :many
SELECT DISTINCT j.status_condition_id
FROM j_status_conditions_modifier_changes j
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY j.status_condition_id;


-- name: GetModifierPropertyIDs :many
SELECT DISTINCT p.id
FROM properties p
JOIN modifier_changes mc ON p.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY p.id;


-- name: GetModifierPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = pa.ability_id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY pa.id;


-- name: GetModifierOverdriveAbilityIDs :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = oa.ability_id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY oa.id;


-- name: GetModifierItemAbilityIDs :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = ia.ability_id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY ia.id;


-- name: GetModifierTriggerCommandIDs :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = tc.ability_id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY tc.id;


-- name: GetModifierEnemyAbilityIDs :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN j_battle_interactions_modifier_changes j ON j.ability_id = ea.ability_id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY ea.id;


-- name: GetModifierIDs :many
SELECT id FROM modifiers ORDER BY id;


-- name: GetModifierIDsByCategory :many
SELECT id FROM modifiers WHERE category = ANY(sqlc.arg('category')::modifier_category[]) ORDER BY id;




-- name: GetPropertyAutoAbilityIDs :many
SELECT id
FROM auto_abilities
WHERE added_property_id = $1
ORDER BY id;


-- name: GetPropertyMonsterIDs :many
SELECT monster_id
FROM j_monsters_properties
WHERE property_id = $1
ORDER BY monster_id;


-- name: GetPropertyIDs :many
SELECT id FROM properties ORDER BY id;
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
JOIN elements e ON er.element_id = e.id
WHERE e.id = $1;


-- name: GetElementAutoAbilityIDs :many
SELECT aa.id
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







-- name: GetStatusConditionAutoAbilityIDs :many
SELECT aa.id
FROM auto_abilities aa
JOIN j_auto_abilities_added_statusses j ON j.auto_ability_id = aa.id
JOIN status_conditions sc ON j.status_condition_id = sc.id
WHERE sc.id = $1

UNION

SELECT aa.id
FROM auto_abilities aa
JOIN j_auto_abilities_added_status_resists j ON j.auto_ability_id = aa.id
JOIN status_resists sr ON j.status_resist_id = sr.id
JOIN status_conditions sc ON sr.status_condition_id = sc.id
WHERE sc.id = $1

UNION

SELECT aa.id
FROM auto_abilities aa
JOIN inflicted_statusses ist ON aa.on_hit_status_id = ist.id
JOIN status_conditions sc ON ist.status_condition_id = sc.id
WHERE sc.id = $1

ORDER BY id;



-- name: GetStatusConditionAbilityIDsInflicted :many
SELECT a.id FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN j_battle_interactions_inflicted_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = j1.battle_interaction_id
JOIN inflicted_statusses ist ON j2.inflicted_status_id = ist.id
WHERE ist.status_condition_id = sqlc.arg(status_condition_id)::int AND ist.probability BETWEEN sqlc.arg('min_rate')::int AND sqlc.arg('max_rate')::int

UNION

SELECT a.id FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
WHERE sqlc.arg(status_condition_id)::int = 6 AND bi.inflicted_delay_id IS NOT NULL

ORDER BY id;


-- name: GetStatusConditionAbilityIDsRemoved :many
SELECT DISTINCT a.id
FROM abilities a
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_removed_status_conditions j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN status_conditions sc ON j2.status_condition_id = sc.id
WHERE sc.id = $1
ORDER BY a.id;


-- name: GetStatusConditionInflictedDelayConditionIDs :many
SELECT DISTINCT sc.id
FROM status_conditions sc
JOIN inflicted_delays idl ON sc.inflicted_delay_id = idl.id
WHERE sqlc.arg(status_id)::int = 6
  AND sc.inflicted_delay_id IS NOT NULL
  AND idl.ctb_attack_type = 'attack'
ORDER BY sc.id;


-- name: GetStatusConditionRemovedConditionIDs :many
SELECT DISTINCT scp.id
FROM status_conditions scp
JOIN j_status_conditions_removed_status_conditions j ON j.parent_condition_id = scp.id
JOIN status_conditions scc ON j.child_condition_id = scc.id
WHERE scc.id = $1
ORDER BY scp.id;


-- name: GetStatusConditionResistingMonsterIDs :many
SELECT m.id
FROM monsters m
JOIN j_monsters_immunities jmi ON jmi.monster_id = m.id
WHERE jmi.status_condition_id = sqlc.arg(status_condition_id)::int

UNION

SELECT m.id
FROM monsters m
JOIN j_monsters_status_resists jmsr ON jmsr.monster_id = m.id
JOIN status_resists sr ON sr.id = jmsr.status_resist_id
WHERE sr.status_condition_id = sqlc.arg(status_condition_id)::int
  AND sr.resistance >= sqlc.arg('min_resistance')::int

ORDER BY id;


-- name: GetStatusConditionIDs :many
SELECT id FROM status_conditions ORDER BY id;


-- name: GetStatusConditionIDsByCategory :many
SELECT id FROM status_conditions WHERE category = ANY(sqlc.arg('category')::status_condition_category[]) ORDER BY id;






-- name: GetStatSphereIDs :many
SELECT DISTINCT sph.id
FROM spheres sph
JOIN items i ON sph.item_id = i.id
JOIN j_items_related_stats j ON j.item_id = i.id
WHERE j.stat_id = $1
ORDER BY sph.id;


-- name: GetStatAutoAbilityIDs :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN j_auto_abilities_related_stats j ON j.auto_ability_id = aa.id
WHERE j.stat_id = $1
ORDER BY aa.id;


-- name: GetStatStatusConditionIDs :many
SELECT DISTINCT sc.id
FROM status_conditions sc
JOIN j_status_conditions_related_stats j ON j.status_condition_id = sc.id
WHERE j.stat_id = $1
ORDER BY sc.id;


-- name: GetStatPropertyIDs :many
SELECT DISTINCT p.id
FROM properties p
JOIN j_properties_related_stats j ON j.property_id = p.id
WHERE j.stat_id = $1
ORDER BY p.id;


-- name: GetStatPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN j_player_abilities_related_stats j ON j.player_ability_id = pa.id
WHERE j.stat_id = $1
ORDER BY pa.id;


-- name: GetStatOverdriveAbilityIDs :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN j_overdrive_abilities_related_stats j ON j.overdrive_ability_id = oa.id
WHERE j.stat_id = $1
ORDER BY oa.id;


-- name: GetStatItemAbilityIDs :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN items i ON ia.item_id = i.id
JOIN j_items_related_stats j ON j.item_id = i.id
WHERE j.stat_id = $1
ORDER BY ia.id;


-- name: GetStatTriggerCommandIDs :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN j_trigger_commands_related_stats j ON j.trigger_command_id = tc.id
WHERE j.stat_id = $1
ORDER BY tc.id;


-- name: GetStatAutoAbilityIDsStatChange :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN j_auto_abilities_stat_changes j ON j.auto_ability_id = aa.id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY aa.id;


-- name: GetStatStatusConditionIDsStatChange :many
SELECT DISTINCT scon.id
FROM status_conditions scon
JOIN j_status_conditions_stat_changes j ON j.status_condition_id = scon.id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY scon.id;


-- name: GetStatPropertyIDsStatChange :many
SELECT DISTINCT p.id
FROM properties p
JOIN j_properties_stat_changes j ON j.property_id = p.id
JOIN stat_changes sc ON j.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY p.id;


-- name: GetStatPlayerAbilityIDsStatChange :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY pa.id;


-- name: GetStatOverdriveAbilityIDsStatChange :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY oa.id;


-- name: GetStatItemAbilityIDsStatChange :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY ia.id;


-- name: GetStatTriggerCommandIDsStatChange :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_stat_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN stat_changes sc ON j2.stat_change_id = sc.id
WHERE sc.stat_id = $1
ORDER BY tc.id;


-- name: GetStatIDs :many
SELECT id FROM stats ORDER BY id;








-- name: GetModifierAutoAbilityIDs :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN j_auto_abilities_modifier_changes j ON j.auto_ability_id = aa.id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY aa.id;


-- name: GetModifierStatusConditionIDs :many
SELECT DISTINCT sc.id
FROM status_conditions sc
JOIN j_status_conditions_modifier_changes j ON j.status_condition_id = sc.id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY sc.id;


-- name: GetModifierPropertyIDs :many
SELECT DISTINCT p.id
FROM properties p
JOIN j_properties_modifier_changes j ON j.property_id = p.id
JOIN modifier_changes mc ON j.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY p.id;


-- name: GetModifierPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN abilities a ON pa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY pa.id;


-- name: GetModifierOverdriveAbilityIDs :many
SELECT DISTINCT oa.id
FROM overdrive_abilities oa
JOIN abilities a ON oa.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY oa.id;


-- name: GetModifierItemAbilityIDs :many
SELECT DISTINCT ia.id
FROM item_abilities ia
JOIN abilities a ON ia.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY ia.id;


-- name: GetModifierTriggerCommandIDs :many
SELECT DISTINCT tc.id
FROM trigger_commands tc
JOIN abilities a ON tc.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY tc.id;


-- name: GetModifierEnemyAbilityIDs :many
SELECT DISTINCT ea.id
FROM enemy_abilities ea
JOIN abilities a ON ea.ability_id = a.id
JOIN j_abilities_battle_interactions j1 ON j1.ability_id = a.id
JOIN battle_interactions bi ON j1.battle_interaction_id = bi.id
JOIN j_battle_interactions_modifier_changes j2 ON j2.ability_id = a.id AND j2.battle_interaction_id = bi.id
JOIN modifier_changes mc ON j2.modifier_change_id = mc.id
WHERE mc.modifier_id = $1
ORDER BY ea.id;


-- name: GetModifierIDs :many
SELECT id FROM modifiers ORDER BY id;


-- name: GetModifierIDsByCategory :many
SELECT id FROM modifiers WHERE category = ANY(sqlc.arg('category')::modifier_category[]) ORDER BY id;




-- name: GetPropertyAutoAbilityIDs :many
SELECT aa.id
FROM auto_abilities aa
WHERE aa.added_property_id = sqlc.arg('property_id')::int
ORDER BY aa.id;


-- name: GetPropertyMonsterIDs :many
SELECT m.id
FROM monsters m
JOIN j_monsters_properties j ON j.monster_id = m.id
WHERE j.property_id = $1
ORDER BY m.id;


-- name: GetPropertyIDs :many
SELECT id FROM properties ORDER BY id;
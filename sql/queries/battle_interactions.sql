-- name: CreateDamage :one
INSERT INTO damages (data_hash, critical, critical_plus_val, is_piercing, break_dmg_limit, element_id)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = damages.data_hash
RETURNING *;


-- name: CreateAbilityDamage :one
INSERT INTO ability_damages (data_hash, condition, attack_type, stat_id, damage_type, damage_formula, damage_constant)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = ability_damages.data_hash
RETURNING *;


-- name: CreateDamagesDamageCalcJunction :exec
INSERT INTO j_damages_damage_calc (data_hash, ability_id, battle_interaction_id, damage_id, ability_damage_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAbilityAccuracy :one
INSERT INTO ability_accuracies (data_hash, acc_source, hit_chance, acc_modifier)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = ability_accuracies.data_hash
RETURNING *;


-- name: CreateInflictedStatus :one
INSERT INTO inflicted_statusses (data_hash, status_condition_id, probability, duration_type, amount)
VALUES ( $1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = inflicted_statusses.data_hash
RETURNING *;


-- name: CreateInflictedDelay :one
INSERT INTO inflicted_delays (data_hash, condition, ctb_attack_type, delay_type, damage_constant)
VALUES ( $1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = inflicted_delays.data_hash
RETURNING *;


-- name: CreateStatChange :one
INSERT INTO stat_changes (data_hash, stat_id, calculation_type, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = stat_changes.data_hash
RETURNING *;


-- name: CreateModifierChange :one
INSERT INTO modifier_changes (data_hash, modifier_id, calculation_type, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = modifier_changes.data_hash
RETURNING *;


-- name: CreateBattleInteraction :one
INSERT INTO battle_interactions (data_hash, target, based_on_phys_attack, range, shatter_rate, accuracy_id, hit_amount, special_action)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = battle_interactions.data_hash
RETURNING *;


-- name: CreateBattleIntDamageJunction :exec
INSERT INTO j_battle_interaction_damage (data_hash, ability_id, battle_interaction_id, damage_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntAffectedByJunction :exec
INSERT INTO j_battle_interactions_affected_by (data_hash, ability_id, battle_interaction_id, status_condition_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntInflictedDelayJunction :exec
INSERT INTO j_battle_interactions_inflicted_delay (data_hash, ability_id, battle_interaction_id, inflicted_delay_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntInflictedConditionsJunction :exec
INSERT INTO j_battle_interactions_inflicted_status_conditions (data_hash, ability_id, battle_interaction_id, inflicted_status_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntRemovedConditionsJunction :exec
INSERT INTO j_battle_interactions_removed_status_conditions (data_hash, ability_id, battle_interaction_id, status_condition_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntCopiedConditionsJunction :exec
INSERT INTO j_battle_interactions_copied_status_conditions (data_hash, ability_id, battle_interaction_id, inflicted_status_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntStatChangesJunction :exec
INSERT INTO j_battle_interactions_stat_changes (data_hash, ability_id, battle_interaction_id, stat_change_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateBattleIntModifierChangesJunction :exec
INSERT INTO j_battle_interactions_modifier_changes (data_hash, ability_id, battle_interaction_id, modifier_change_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
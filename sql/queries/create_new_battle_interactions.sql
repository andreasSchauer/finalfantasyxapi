-- name: CreateDamageBulk :many
INSERT INTO damages (data_hash, critical, critical_plus_val, is_piercing, break_dmg_limit, element_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('critical')::null_critical_type[]),
    unnest(sqlc.arg('critical_plus_val')::int[]),
    unnest(sqlc.arg('is_piercing')::boolean[]),
    unnest(sqlc.arg('break_dmg_limit')::null_break_dmg_lmt_type[]),
    unnest(sqlc.arg('element_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = damages.data_hash
RETURNING id;


-- name: CreateAbilityDamageBulk :many
INSERT INTO ability_damages (data_hash, condition, attack_type, stat_id, damage_type, damage_formula, damage_constant)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('condition')::null_string[]),
    unnest(sqlc.arg('attack_type')::attack_type[]),
    unnest(sqlc.arg('stat_id')::int[]),
    unnest(sqlc.arg('damage_type')::damage_type[]),
    unnest(sqlc.arg('damage_formula')::damage_formula[]),
    unnest(sqlc.arg('damage_constant')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = ability_damages.data_hash
RETURNING id;


-- name: CreateAbilityAccuracyBulk :many
INSERT INTO ability_accuracies (data_hash, acc_source, hit_chance, acc_modifier)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('acc_source')::acc_source_type[]),
    unnest(sqlc.arg('hit_chance')::null_int[]),
    unnest(sqlc.arg('acc_modifier')::null_float[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = ability_accuracies.data_hash
RETURNING id;


-- name: CreateInflictedStatusBulk :many
INSERT INTO inflicted_statusses (data_hash, status_condition_id, probability, duration_type, amount)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('status_condition_id')::int[]),
    unnest(sqlc.arg('probability')::int[]),
    unnest(sqlc.arg('duration_type')::duration_type[]),
    unnest(sqlc.arg('amount')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = inflicted_statusses.data_hash
RETURNING id;


-- name: CreateInflictedDelayBulk :many
INSERT INTO inflicted_delays (data_hash, condition, ctb_attack_type, delay_type, damage_constant)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('condition')::null_string[]),
    unnest(sqlc.arg('ctb_attack_type')::ctb_attack_type[]),
    unnest(sqlc.arg('delay_type')::delay_type[]),
    unnest(sqlc.arg('damage_constant')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = inflicted_delays.data_hash
RETURNING id;


-- name: CreateStatChangeBulk :many
INSERT INTO stat_changes (data_hash, stat_id, calculation_type, value)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('stat_id')::int[]),
    unnest(sqlc.arg('calculation_type')::calculation_type[]),
    unnest(sqlc.arg('value')::real[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = stat_changes.data_hash
RETURNING id;


-- name: CreateModifierChangeBulk :many
INSERT INTO modifier_changes (data_hash, modifier_id, calculation_type, value)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('modifier_id')::int[]),
    unnest(sqlc.arg('calculation_type')::calculation_type[]),
    unnest(sqlc.arg('value')::real[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = modifier_changes.data_hash
RETURNING id;


-- name: CreateBattleInteractionBulk :many
INSERT INTO battle_interactions (data_hash, target, based_on_user_attack, range, shatter_rate, accuracy_id, inflicted_delay_id, hit_amount, special_action)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('target')::target_type[]),
    unnest(sqlc.arg('based_on_user_attack')::boolean[]),
    unnest(sqlc.arg('range')::null_int[]),
    unnest(sqlc.arg('shatter_rate')::int[]),
    unnest(sqlc.arg('accuracy_id')::int[]),
    unnest(sqlc.arg('inflicted_delay_id')::null_int[]),
    unnest(sqlc.arg('hit_amount')::int[]),
    unnest(sqlc.arg('special_action')::null_special_action_type[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = battle_interactions.data_hash
RETURNING id;
-- name: CreateStat :exec
INSERT INTO stats (data_hash, name, effect, min_val, max_val, max_val_2)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateElement :exec
INSERT INTO elements (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAffinity :exec
INSERT INTO affinities (data_hash, name, damage_factor)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAgilityTier :one
INSERT INTO agility_tiers (data_hash, min_agility, max_agility, tick_speed, monster_min_icv, monster_max_icv, character_max_icv)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (data_hash) DO UPDATE SET data_hash = agility_tiers.data_hash
RETURNING *;


-- name: CreateAgilitySubtier :exec
INSERT INTO agility_subtiers (data_hash, agility_tier_id, subtier_min_agility, subtier_max_agility, character_min_icv)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateOverdriveMode :exec
INSERT INTO overdrive_modes (data_hash, name, description, effect, type, fill_rate)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateStatusCondition :exec
INSERT INTO status_conditions (data_hash, name, effect)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateProperty :exec
INSERT INTO properties (data_hash, name, effect)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
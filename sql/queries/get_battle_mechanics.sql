-- name: GetOverdriveModeIDs :many
SELECT id FROM overdrive_modes ORDER BY id;


-- name: GetOverdriveModeIDsByType :many
SELECT id FROM overdrive_modes WHERE type = $1 ORDER BY id;


-- name: GetAgilityTierByAgility :one
SELECT * FROM agility_tiers
WHERE (sqlc.arg(agility)::int) >= min_agility
AND (sqlc.arg(agility)::int) <= max_agility;
-- name: CreateStat :one
INSERT INTO stats (data_hash, name, effect, min_val, max_val, max_val_2)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING
RETURNING *;
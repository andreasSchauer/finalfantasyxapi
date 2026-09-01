-- name: GetDbState :one
SELECT data_hash FROM db_state WHERE id = 1;


-- name: UpdateDbState :exec
UPDATE db_state
SET data_hash = $1,
    updated_at = NOW()
WHERE id = 1;

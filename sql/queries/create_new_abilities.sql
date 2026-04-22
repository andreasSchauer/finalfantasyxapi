-- name: CreateAbilitiesBulk :many
INSERT INTO abilities (data_hash, name, version, specification, attributes_id, type)
SELECT 
    unnest(sqlc.arg('data_hashes')::text[]), 
    unnest(sqlc.arg('names')::text[]), 
    unnest(sqlc.arg('versions')::null_int[]), 
    unnest(sqlc.arg('specifications')::null_string[]),
    unnest(sqlc.arg('attributes_ids')::null_int[]),
    unnest(sqlc.arg('types')::null_nullify_armored[])
ON CONFLICT (data_hash) DO UPDATE SET 
    data_hash = EXCLUDED.data_hash
RETURNING id, data_hash;
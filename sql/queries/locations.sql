-- name: CreateLocation :one
INSERT INTO locations (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = locations.data_hash
RETURNING *;


-- name: CreateSubLocation :one
INSERT INTO sub_locations (data_hash, location_id, name, specification)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = sub_locations.data_hash
RETURNING *;


-- name: CreateArea :exec
INSERT INTO areas (data_hash, sub_location_id, name, section, can_revisit, has_save_sphere, airship_drop_off, has_compilation_sphere)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(data_hash) DO NOTHING;
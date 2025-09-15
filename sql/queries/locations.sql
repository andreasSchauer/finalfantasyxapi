-- name: CreateLocation :one
INSERT INTO locations (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = locations.data_hash
RETURNING *;


-- name: CreateSubLocation :one
INSERT INTO sub_locations (data_hash, location_id, name, version, specification)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = sub_locations.data_hash
RETURNING *;


-- name: CreateArea :exec
INSERT INTO areas (data_hash, sub_location_id, name, version, specification, can_revisit, has_save_sphere, airship_drop_off, has_compilation_sphere)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO NOTHING;


-- name: GetLocationAreas :many
SELECT
    l.id as location_id,
    s.id as sub_location_id,
    a.id as area_id,
    l.name as location_name,
    s.name as sub_location_name,
    a.name as area_name,
    s.version as s_version,
    a.version as a_version
FROM areas a
LEFT JOIN sub_locations s
ON a.sub_location_id = s.id
LEFT JOIN locations l
ON s.location_id = l.id;



-- name: CreateTreasure :exec
INSERT INTO treasures (data_hash, area_id, version, treasure_type, loot_type, is_post_airship, is_anima_treasure, notes, gil_amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShop :exec
INSERT INTO shops (data_hash, version, area_id, notes, category)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationList :exec
INSERT INTO monster_formation_lists (data_hash, version, area_id, notes)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;
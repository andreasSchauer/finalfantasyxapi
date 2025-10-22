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


-- name: CreateArea :one
INSERT INTO areas (data_hash, sub_location_id, name, version, specification, story_only, has_save_sphere, airship_drop_off, has_compilation_sphere, can_ride_chocobo)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = areas.data_hash
RETURNING *;


-- name: CreateAreaConnection :one
INSERT INTO area_connections (data_hash, area_id, connection_type, story_only, notes)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = area_connections.data_hash
RETURNING *;


-- name: CreateAreaConnectionJunction :exec
INSERT INTO j_area_connection (data_hash, area_id, connection_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateTreasure :exec
INSERT INTO treasures (data_hash, area_id, version, treasure_type, loot_type, is_post_airship, is_anima_treasure, notes, gil_amount)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateShop :exec
INSERT INTO shops (data_hash, version, area_id, notes, category)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateFormationLocation :one
INSERT INTO formation_locations (data_hash, version, area_id, notes)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = formation_locations.data_hash
RETURNING *;


-- name: CreateFormationBossSong :one
INSERT INTO formation_boss_songs (data_hash, song_id, celebrate_victory)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = formation_boss_songs.data_hash
RETURNING *;


-- name: CreateMonsterFormation :one
INSERT INTO monster_formations (data_hash, formation_location_id, category, is_forced_ambush, can_escape, boss_song_id, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = monster_formations.data_hash
RETURNING *;


-- name: CreateLocationMonsterFormationJunction :exec
INSERT INTO j_location_monster_formation (data_hash, formation_location_id, monster_formation_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationMonsterAmountJunction :exec
INSERT INTO j_monster_formation_monster_amount (data_hash, monster_formation_id, monster_amount_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateMonsterFormationTriggerCommandJunction :exec
INSERT INTO j_monster_formation_trigger_command (data_hash, monster_formation_id, trigger_command_id)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO NOTHING;
-- name: CreateSongCredit :one
INSERT INTO song_credits (data_hash, composer, arranger, performer, lyricist)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = song_credits.data_hash
RETURNING *;



-- name: CreateSong :one
INSERT INTO songs (data_hash, name, streaming_name, in_game_name, ost_name, translation, streaming_track_number, music_sphere_id, ost_disc, ost_track_number, duration_in_seconds, can_loop, special_use_case, credits_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = songs.data_hash
RETURNING *;


-- name: CreateFMV :exec
INSERT INTO fmvs (data_hash, name, translation, cutscene_description, song_id, area_id)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;
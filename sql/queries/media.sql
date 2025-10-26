-- name: CreateSong :one
INSERT INTO songs (data_hash, name, streaming_name, in_game_name, ost_name, translation, streaming_track_number, music_sphere_id, ost_disc, ost_track_number, duration_in_seconds, can_loop, special_use_case)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = songs.data_hash
RETURNING *;


-- name: UpdateSong :exec
UPDATE songs
SET data_hash = $1,
    credits_id = $2
WHERE id = $3;


-- name: CreateSongCredit :one
INSERT INTO song_credits (data_hash, composer, arranger, performer, lyricist)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = song_credits.data_hash
RETURNING *;


-- name: CreateBackgroundMusic :one
INSERT INTO background_music (data_hash, condition, replaces_encounter_music)
VALUES ($1, $2, $3)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = background_music.data_hash
RETURNING *;


-- name: CreateSongsBackgroundMusicJunction :exec
INSERT INTO j_songs_background_music (data_hash, song_id, bm_id, area_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateCue :one
INSERT INTO cues (data_hash, scene_description, area_id, replaces_bg_music, end_trigger, replaces_encounter_music)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = cues.data_hash
RETURNING *;


-- name: CreateSongsCuesJunction :exec
INSERT INTO j_songs_cues (data_hash, song_id, cue_id, area_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateFMV :exec
INSERT INTO fmvs (data_hash, name, translation, cutscene_description, song_id, area_id)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;
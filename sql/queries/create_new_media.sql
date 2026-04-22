-- name: CreateSongBulk :many
INSERT INTO songs (data_hash, name, streaming_name, in_game_name, ost_name, translation, streaming_track_number, music_sphere_id, ost_disc, ost_track_number, duration_in_seconds, can_loop, special_use_case)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('streaming_name')::null_string[]),
    unnest(sqlc.arg('in_game_name')::null_string[]),
    unnest(sqlc.arg('ost_name')::null_string[]),
    unnest(sqlc.arg('translation')::null_string[]),
    unnest(sqlc.arg('streaming_track_number')::null_int[]),
    unnest(sqlc.arg('music_sphere_id')::null_int[]),
    unnest(sqlc.arg('ost_disc')::null_int[]),
    unnest(sqlc.arg('ost_track_number')::null_int[]),
    unnest(sqlc.arg('duration_in_seconds')::int[]),
    unnest(sqlc.arg('can_loop')::boolean[]),
    unnest(sqlc.arg('special_use_case')::null_music_use_case[]),
    unnest(sqlc.arg('credits_id')::null_int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateSongCreditBulk :many
INSERT INTO song_credits (data_hash, composer, arranger, performer, lyricist)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('composer')::null_composer[]),
    unnest(sqlc.arg('arranger')::null_arranger[]),
    unnest(sqlc.arg('performer')::null_string[]),
    unnest(sqlc.arg('lyricist')::null_string[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateBackgroundMusicBulk :many
INSERT INTO background_music (data_hash, condition, replaces_encounter_music)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('condition')::null_string[]),
    unnest(sqlc.arg('replaces_encounter_music')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateCueBulk :many
INSERT INTO cues (data_hash, song_id, scene_description, trigger_area_id, replaces_bg_music, end_trigger, replaces_encounter_music)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('song_id')::int[]),
    unnest(sqlc.arg('scene_description')::text[]),
    unnest(sqlc.arg('trigger_area_id')::null_int[]),
    unnest(sqlc.arg('replaces_bg_music')::null_bg_replacement_type[]),
    unnest(sqlc.arg('end_trigger')::null_string[]),
    unnest(sqlc.arg('replaces_encounter_music')::boolean[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;


-- name: CreateFMVBulk :many
INSERT INTO fmvs (data_hash, name, translation, cutscene_description, song_id, area_id)
SELECT
    unnest(sqlc.arg('data_hash')::text[]),
    unnest(sqlc.arg('name')::text[]),
    unnest(sqlc.arg('translation')::null_string[]),
    unnest(sqlc.arg('cutscene_description')::text[]),
    unnest(sqlc.arg('song_id')::null_int[]),
    unnest(sqlc.arg('area_id')::int[])
ON CONFLICT(data_hash) DO UPDATE SET data_hash = EXCLUDED.data_hash
RETURNING id;
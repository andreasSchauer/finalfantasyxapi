-- name: GetSongFmvIDs :many
SELECT id FROM fmvs WHERE song_id = $1 ORDER BY id;


-- name: GetSongMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs fbs ON fd.boss_song_id = fbs.id
WHERE fbs.song_id = $1
ORDER BY mf.id;


-- name: GetSongIDs :many
SELECT id FROM songs ORDER BY id;


-- name: GetSongIDsWithFMVs :many
SELECT DISTINCT song_id::int FROM fmvs WHERE song_id IS NOT NULL ORDER BY song_id;


-- name: GetSongIDsWithSpecialUseCase :many
SELECT id FROM songs WHERE special_use_case IS NOT NULL;


-- name: GetSongIDsByComposer :many
SELECT s.id
FROM songs s
JOIN song_credits c ON s.credits_id = c.id
WHERE c.composer = $1
ORDER BY s.id;


-- name: GetSongIDsByArranger :many
SELECT s.id
FROM songs s
JOIN song_credits c ON s.credits_id = c.id
WHERE c.arranger = $1
ORDER BY s.id;


-- name: GetFmvIDs :many
SELECT id FROM fmvs ORDER BY id;
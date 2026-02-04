-- name: GetSongFmvIDs :many
SELECT DISTINCT f.id
FROM fmvs f
JOIN songs s ON f.song_id = s.id
WHERE s.id = $1
ORDER BY f.id;


-- name: GetSongMonsterFormationIDs :many
SELECT DISTINCT mf.id
FROM monster_formations mf
JOIN formation_data fd ON mf.formation_data_id = fd.id
JOIN formation_boss_songs fbs ON fd.boss_song_id = fbs.id
JOIN songs s ON fbs.song_id = s.id
WHERE s.id = $1
ORDER BY mf.id;


-- name: GetSongIDs :many
SELECT id FROM songs ORDER BY id;


-- name: GetSongIDsWithFMVs :many
SELECT s.id
FROM songs s
JOIN fmvs f ON f.song_id = s.id
ORDER BY s.id;


-- name: GetSongIDsWithSpecialUseCase :many
SELECT id FROM songs WHERE special_use_case != NULL;


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
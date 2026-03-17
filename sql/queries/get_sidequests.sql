-- name: GetParentSidequest :one
SELECT q.*
FROM subquests su
LEFT JOIN sidequests si ON su.sidequest_id = si.id
LEFT JOIN quests q ON si.quest_id = q.id
WHERE su.id = $1;


-- name: GetSidequestSubquestIDs :many
SELECT id FROM subquests WHERE sidequest_id = $1 ORDER BY id;


-- name: GetSidequestIDs :many
SELECT id FROM sidequests ORDER BY id;


-- name: GetSidequestIDsByPostAirship :many
SELECT DISTINCT si.id
FROM sidequests si
JOIN quests qsi ON si.quest_id = qsi.id
JOIN quest_completions qcsi ON qcsi.quest_id = qsi.id
JOIN subquests su ON su.sidequest_id = si.id
JOIN quests qsu ON su.quest_id = qsu.id
JOIN quest_completions qcsu ON qcsu.quest_id = qsu.id
WHERE qcsi.is_post_airship = $1 OR qcsu.is_post_airship = $1
ORDER BY si.id;


-- name: GetSubquestIDs :many
SELECT id FROM subquests ORDER BY id;


-- name: GetSubquestIDsByPostAirship :many
SELECT DISTINCT s.id
FROM subquests s
JOIN quests q ON s.quest_id = q.id
JOIN quest_completions qc ON qc.quest_id = q.id
WHERE qc.is_post_airship = true
ORDER BY s.id;


-- name: GetArenaCreationIDs :many
SELECT id FROM monster_arena_creations ORDER BY id;


-- name: GetArenaCreationIDsByCategory :many
SELECT id FROM monster_arena_creations WHERE category = $1 ORDER BY id;


-- name: GetBlitzballPrizeIDs :many
SELECT id FROM blitzball_positions ORDER BY id;


-- name: GetBlitzballPrizeIDsByCategory :many
SELECT id FROM blitzball_positions WHERE category = $1 ORDER BY id;
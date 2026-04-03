-- name: GetQuestIDs :many
SELECT id FROM quests ORDER BY id;


-- name: GetQuestIDsByType :many
SELECT id FROM quests WHERE type = $1 ORDER BY id;


-- name: GetQuestIDsByAvailability :many
SELECT id FROM quests WHERE availability = ANY(sqlc.narg('availability')::availability_type[]) ORDER BY id;


-- name: GetQuestIDsByRepeatable :many
SELECT id FROM quests WHERE is_repeatable = $1 ORDER BY id;


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


-- name: GetSidequestIDsByAvailability :many
SELECT DISTINCT s.id
FROM sidequests s
JOIN quests q ON s.quest_id = q.id
WHERE q.availability = ANY(sqlc.narg('availability')::availability_type[])
ORDER BY s.id;


-- name: GetSubquestIDs :many
SELECT id FROM subquests ORDER BY id;


-- name: GetSubquestIDsByAvailability :many
SELECT DISTINCT s.id
FROM subquests s
JOIN quests q ON s.quest_id = q.id
WHERE q.availability = ANY(sqlc.narg('availability')::availability_type[])
ORDER BY s.id;


-- name: GetSubquestIDsByRepeatable :many
SELECT s.id
FROM subquests s
JOIN quests q ON s.quest_id = q.id
WHERE q.is_repeatable = $1
ORDER BY s.id;


-- name: GetArenaCreationIDs :many
SELECT id FROM monster_arena_creations ORDER BY id;


-- name: GetArenaCreationIDsByCategory :many
SELECT id FROM monster_arena_creations WHERE category = ANY(sqlc.narg('category')::ma_creation_category[]) ORDER BY id;


-- name: GetBlitzballPrizeIDs :many
SELECT id FROM blitzball_positions ORDER BY id;


-- name: GetBlitzballPrizeIDsByCategory :many
SELECT id FROM blitzball_positions WHERE category = $1 ORDER BY id;
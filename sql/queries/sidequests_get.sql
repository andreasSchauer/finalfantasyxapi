
-- name: GetParentSidequest :one
SELECT q.*
FROM subquests su
LEFT JOIN sidequests si ON su.sidequest_id = si.id
LEFT JOIN quests q ON si.quest_id = q.id
WHERE su.id = $1;
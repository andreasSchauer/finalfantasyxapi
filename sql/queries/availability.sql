-- name: FilterMonsterIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('context_id')::int AS context_id,
        sqlc.narg('context_type')::text AS context_type
),
available_monsters AS (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END as current_avl
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    CROSS JOIN w
    WHERE a.source_type = 'monster'
      AND a.s_id = ANY(w.ids)
      AND (w.context_id IS NULL OR CASE
           WHEN w.context_type = 'location' THEN g.location_id
           WHEN w.context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.context_type = 'area' THEN g.area_id
          END = w.context_id)
)
SELECT s_id
FROM available_monsters
GROUP BY s_id
HAVING ARRAY[MIN(current_avl)] && (SELECT availability FROM w);



-- might use a source type array, if the sources need to be specified
-- needs to change (see monster query)
-- name: FilterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability
)
SELECT DISTINCT i.id
FROM items i
JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
JOIN mv_availabilities a ON mis.source_id = a.s_id AND mis.source_type = a.source_type
CROSS JOIN w
WHERE i.id = ANY(w.ids)
  AND CASE
        WHEN w.avl_type = 'self' THEN a.avl_self
        WHEN w.avl_type = 'context' THEN a.avl_context
        WHEN w.avl_type = 'area' THEN a.avl_area
      END = ANY(w.availability)
ORDER BY i.id;

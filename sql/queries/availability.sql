-- name: FilterIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('source_type')::text AS source_type,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability
)
SELECT DISTINCT a.s_id
FROM mv_availabilities a
CROSS JOIN w
WHERE a.s_id = ANY(w.ids)
  AND a.source_type = w.source_type
  AND CASE
        WHEN w.avl_type = 'self' THEN a.avl_self
        WHEN w.avl_type = 'context' THEN a.avl_context
        WHEN w.avl_type = 'area' THEN a.avl_area
      END = ANY(w.availability)
ORDER BY a.s_id;



WITH w AS (
    SELECT
        ARRAY[1, 5, 8, 19, 30, 36]::int[] AS ids,
        'shop'::text AS source_type,
        'context'::text AS avl_type,
        ARRAY['post']::availability_type[] AS availability
)
SELECT DISTINCT a.*
FROM mv_availabilities a
CROSS JOIN w
WHERE a.s_id = ANY(w.ids)
  AND a.source_type = w.source_type
  AND CASE
        WHEN w.avl_type = 'self' THEN a.avl_self
        WHEN w.avl_type = 'context' THEN a.avl_context
        WHEN w.avl_type = 'area' THEN a.avl_area
      END = ANY(w.availability)
ORDER BY a.s_id;
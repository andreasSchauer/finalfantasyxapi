-- name: FilterMonsterIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('source_type')::text AS source_type,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('context_id')::int AS context_id,
        sqlc.narg('context_type')::text AS context_type
)
SELECT DISTINCT a.s_id
FROM mv_availabilities a
JOIN mv_geography g ON a.a_id = g.area_id
CROSS JOIN w
WHERE a.source_type = w.source_type
  AND a.s_id = ANY(w.ids)
  AND CASE
        WHEN w.avl_type = 'self' THEN a.avl_self
        WHEN w.avl_type = 'context' THEN a.avl_context
        WHEN w.avl_type = 'area' THEN a.avl_area
      END = ANY(w.availability)
  AND (w.context_id IS NULL OR CASE
       WHEN w.context_type = 'location' THEN g.location_id
       WHEN w.context_type = 'sublocation' THEN g.sublocation_id
       WHEN w.context_type = 'area' THEN g.area_id
      END = w.context_id)
ORDER BY a.s_id;


WITH w AS (
    SELECT
        ARRAY[155, 159, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 301]::int[] AS ids,
        'monster'::text AS source_type,
        'context'::text AS avl_type,
        ARRAY['post']::availability_type[] AS availability
)
SELECT DISTINCT a.s_id, a.*
FROM mv_availabilities a
JOIN mv_geography g ON a.a_id = g.area_id
CROSS JOIN w
WHERE a.s_id = ANY(w.ids)
  AND a.source_type = w.source_type
  AND g.location_id = 23
  AND CASE
        WHEN w.avl_type = 'self' THEN a.avl_self
        WHEN w.avl_type = 'context' THEN a.avl_context
        WHEN w.avl_type = 'area' THEN a.avl_area
      END = ANY(w.availability)
ORDER BY a.s_id;


-- might use a source type array, if the sources need to be specified
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

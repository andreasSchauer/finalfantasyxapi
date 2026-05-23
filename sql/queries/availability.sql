-- name: FilterMonsterIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
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
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT s_id
FROM available_monsters
GROUP BY s_id
HAVING ARRAY[MIN(current_avl)] && (SELECT availability FROM w);



-- name: FilterMonsterFormationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
),
available_formations AS (
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
    WHERE a.source_type = 'monster-formation'
      AND a.s_id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT s_id
FROM available_formations
GROUP BY s_id
HAVING ARRAY[MIN(current_avl)] && (SELECT availability FROM w);


-- name: FilterShopIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('sub_types')::text[] AS sub_types
)
SELECT s_id
FROM (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END as current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'shop'
      AND a.s_id = ANY(w.ids)
      AND (w.sub_types IS NULL OR a.sub_type = ANY(w.sub_types))
) available_shops
GROUP BY s_id
HAVING ARRAY[MIN(current_avl)] && (SELECT availability FROM w);


-- name: FilterTreasureIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability
)
SELECT DISTINCT a.s_id
FROM mv_availabilities a
CROSS JOIN w
WHERE a.source_type = 'treasure'
  AND a.s_id = ANY(w.ids)
  AND ARRAY[a.avl_self] &&  w.availability
ORDER BY a.s_id;



-- name: FilterAreaIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('monster_id')::int AS monster_id
),
area_availabilities AS (
    SELECT a.a_id, a.avl_self as current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'area'
      AND a.a_id = ANY(w.ids)
      AND w.monster_id IS NULL

    UNION ALL

    SELECT a.a_id, a.avl_area as current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'monster'
      AND a.a_id = ANY(w.ids)
      AND a.s_id = w.monster_id
)
SELECT a_id
FROM area_availabilities
GROUP BY a_id
HAVING ARRAY[MAX(current_avl)] && (SELECT availability FROM w)
ORDER BY a_id;





WITH w AS (
    SELECT
        ARRAY[172]::int[] AS ids,
        ARRAY['post']::availability_type[] AS availability,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('sub_sources')::text[] AS sub_sources
),
area_availabilities AS (
    SELECT a.a_id, a.avl_self as current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'area'
      AND a.a_id = ANY(w.ids)
      AND w.monster_id IS NULL

    UNION ALL

    SELECT a.a_id, a.avl_self as current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'monster'
      AND a.a_id = ANY(w.ids)
      AND a.s_id = w.monster_id

    UNION ALL

    SELECT a.a_id, a.avl_area as current_avl
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = ANY(w.sub_sources)
      AND ARRAY[a.avl_area] && w.availability
)
SELECT a_id
FROM area_availabilities
GROUP BY a_id
HAVING ARRAY[MIN(current_avl)] && (SELECT availability FROM w)
  AND (
        (SELECT sub_sources FROM w) IS NULL OR 
        COUNT(DISTINCT source_type) > 1
      )
ORDER BY a_id;




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

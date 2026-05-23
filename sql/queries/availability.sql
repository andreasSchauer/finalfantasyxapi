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
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods
),
available_areas AS (
    SELECT a.a_id, a.avl_self as current_avl, 'area'::text AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = 'area'

    UNION ALL

    SELECT a.a_id, a.avl_area as current_avl, 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id

    UNION ALL

    SELECT
        mis.area_id AS a_id,
        'always'::availability_type AS current_avl,
        'item'::text AS s_type
    FROM mv_item_sources mis
    JOIN items i ON mis.master_item_id = i.master_item_id
    JOIN w ON mis.area_id = ANY(w.ids)
    WHERE i.id = w.item_id
      AND ARRAY[mis.availability] && (SELECT availability FROM w)
      AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

    UNION ALL

    SELECT
        mis.area_id AS a_id,
        'always'::availability_type AS current_avl,
        'key-item'::text AS s_type
    FROM mv_item_sources mis
    JOIN key_items ki ON mis.master_item_id = ki.master_item_id
    JOIN w ON mis.area_id = ANY(w.ids)
    WHERE ki.id = w.key_item_id
      AND ARRAY[mis.availability] && (SELECT availability FROM w)

    UNION ALL

    SELECT
        a.a_id,
        CASE
          WHEN a.source_type = 'monster' THEN a.avl_area
          ELSE a.avl_self
        END AS current_avl,
        CASE
          WHEN a.sub_type = 'boss' THEN 'boss'::text
          ELSE a.source_type
        END AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type != 'area'
      AND ARRAY[
        CASE
          WHEN a.source_type = 'monster' THEN a.avl_area
          ELSE a.avl_self
      END] && (SELECT availability FROM w)
)
SELECT a_id
FROM available_areas
GROUP BY a_id
HAVING ARRAY[MAX(current_avl)] && (SELECT availability FROM w)
AND (
    (SELECT reqs FROM w) IS NULL OR
    ARRAY_AGG(DISTINCT s_type) @> (SELECT reqs FROM w)
)
AND (
    (SELECT excls FROM w) IS NULL OR
    NOT ARRAY_AGG(DISTINCT s_type) && (SELECT excls FROM w)
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

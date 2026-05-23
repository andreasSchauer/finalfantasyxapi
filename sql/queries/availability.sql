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












WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('empty_slots')::int[] AS empty_slots,
        sqlc.narg('character_id')::int AS character_id,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id
),
available_shops AS (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END as current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    LEFT JOIN mv_equipment_sources es ON a.s_id = es.source_id 
        AND es.source_type = 'shop' -- Move this here
        AND (w.empty_slots IS NULL OR es.empty_slots_amount::int = ANY(w.empty_slots)) -- And these
        AND (w.auto_ability_id IS NULL OR es.auto_ability_id = w.auto_ability_id)
    LEFT JOIN equipment_names en ON es.name_id = en.id
        AND (w.character_id IS NULL OR en.character_id = w.character_id)
    WHERE a.source_type = 'shop'
      AND a.s_id = ANY(w.ids)
)
SELECT s_id
FROM available_shops
GROUP BY s_id
HAVING ARRAY[MIN(current_avl)] && (SELECT availability FROM w);



WITH w AS (
    SELECT
        ARRAY[1, 2, 16, 27, 28, 30, 35, 36]::int[] AS ids,
        'context'::text AS avl_type,
        ARRAY['post']::availability_type[] AS availability,
        NULL::int AS loc_context_id,
        NULL::text AS loc_context_type,
        ARRAY[3, 4]::int[] AS empty_slots,
        NULL::int AS character_id,
        NULL::int AS auto_ability_id
),
available_shops AS (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END as current_avl
    FROM mv_availabilities a
    JOIN mv_equipment_sources es ON a.s_id = es.source_id
    LEFT JOIN equipment_names en ON es.name_id = en.id
    JOIN mv_geography g ON a.a_id = g.area_id
    CROSS JOIN w
    WHERE a.source_type = 'shop'
      AND es.source_type = 'shop'
      AND a.s_id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
      AND (w.empty_slots IS NULL OR es.empty_slots_amount::int = ANY(w.empty_slots))
      AND (w.character_id IS NULL OR en.character_id = w.character_id)
      AND (w.auto_ability_id IS NULL OR es.auto_ability_id = w.auto_ability_id)
)
SELECT s_id
FROM available_shops
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

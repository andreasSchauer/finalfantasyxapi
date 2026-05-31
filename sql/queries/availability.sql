-- name: FilterMonsterIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_monsters AS (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN me.is_repeatable_loc
            ELSE me.is_repeatable
        END AS is_rep
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN mv_monster_encounters me ON me.monster_id = a.source_id AND a.source_type = 'monster' AND me.area_id = a.a_id
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
CROSS JOIN w
GROUP BY s_id, w.availability, w.is_repeatable
HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable);


-- name: FilterMasterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_master_items AS (
    SELECT 
        mis.master_item_id,
        CASE
            WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN a.avl_context
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM mv_item_sources mis
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE mis.master_item_id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT master_item_id
FROM available_master_items
CROSS JOIN w
GROUP BY master_item_id, w.availability, w.is_repeatable
HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable);


-- name: FilterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_items AS (
    SELECT 
        i.id AS item_id,
        CASE
            WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN a.avl_context
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM items i
    JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE i.id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT item_id
FROM available_items
CROSS JOIN w
GROUP BY item_id, w.availability, w.is_repeatable
HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable);


-- name: FilterKeyItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
),
available_key_items AS (
    SELECT 
        ki.id AS key_item_id,
        a.avl_self AS current_avl
    FROM key_items ki
    JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE ki.id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT key_item_id
FROM available_key_items
CROSS JOIN w
GROUP BY key_item_id, w.availability
HAVING MIN(current_avl) = ANY(w.availability);


-- name: FilterSphereIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_spheres AS (
    SELECT 
        s.id AS sphere_id,
        CASE
            WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN a.avl_context
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM spheres s
    JOIN items i ON s.item_id = i.id
    JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE s.id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT sphere_id
FROM available_spheres
CROSS JOIN w
GROUP BY sphere_id, w.availability, w.is_repeatable
HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable);


-- name: FilterPrimerIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability
),
available_primers AS (
    SELECT 
        p.id AS primer_id,
        a.avl_self AS current_avl
    FROM primers p
    JOIN key_items ki ON p.key_item_id = ki.id
    JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE p.id = ANY(w.ids)
)
SELECT primer_id
FROM available_primers
CROSS JOIN w
GROUP BY primer_id, w.availability
HAVING MIN(current_avl) = ANY(w.availability);


-- name: FilterAutoAbilityIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('character_id')::int AS character_id,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.narg('req_item')::boolean AS req_item
),
available_auto_abilities AS (
    SELECT 
        aas.auto_ability_id,
        CASE
            WHEN aas.source_type = 'shop' AND w.avl_type = 'self' THEN a.avl_context
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN aas.is_repeatable_loc
            ELSE aas.is_repeatable
        END AS is_rep
    FROM mv_auto_ability_sources aas
    JOIN mv_availabilities a ON a.s_id = aas.source_id
     AND a.source_type = aas.source_type
     AND a.a_id = aas.area_id
    JOIN mv_geography g ON aas.area_id = g.area_id
    CROSS JOIN w
    WHERE (w.req_item IS NULL OR w.req_item = FALSE)
      AND aas.auto_ability_id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
      AND (w.character_id IS NULL OR aas.character_id = w.character_id OR aas.character_id IS NULL)

    UNION ALL

    SELECT
        aa.id AS auto_ability_id,
        CASE
            WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN a.avl_context
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM auto_abilities aa
    JOIN item_amounts ia_req ON aa.required_item_amount_id = ia_req.id
    JOIN mv_item_sources mis ON mis.master_item_id = ia_req.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE w.req_item = TRUE
      AND a.id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
)
SELECT auto_ability_id
FROM available_auto_abilities
CROSS JOIN w
GROUP BY auto_ability_id, w.availability, w.is_repeatable
HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable);



-- name: FilterMonsterFormationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_formations AS (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl,
        me.is_repeatable_loc AS is_rep
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN mv_monster_encounters me ON me.formation_id = a.source_id AND a.source_type = 'monster-formation' AND me.area_id = a.a_id
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
CROSS JOIN w
GROUP BY s_id, w.availability, w.is_repeatable
HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable);


-- name: FilterShopIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('sub_types')::text[] AS sub_types
),
available_shops AS (
    SELECT 
        a.s_id,
        CASE 
            WHEN w.avl_type = 'self' THEN a.avl_self
            WHEN w.avl_type = 'context' THEN a.avl_context
            WHEN w.avl_type = 'area' THEN a.avl_area
        END AS current_avl
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'shop'
        AND a.s_id = ANY(w.ids)
        AND (w.sub_types IS NULL OR a.sub_type = ANY(w.sub_types))
)
SELECT s_id
FROM available_shops
CROSS JOIN w
GROUP BY s_id, w.availability
HAVING MIN(current_avl) = ANY(w.availability);


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
  AND a.avl_self = ANY(w.availability)
ORDER BY a.s_id;


-- name: FilterQuestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
)
SELECT DISTINCT a.s_id
FROM quests q
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE a.s_id = ANY(w.ids)
  AND (w.availability IS NULL OR a.avl_self = ANY(w.availability))
  AND (w.is_repeatable IS NULL OR q.is_repeatable = w.is_repeatable)
ORDER BY a.s_id;


-- name: FilterSidequestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
)
SELECT DISTINCT s.id
FROM sidequests s
JOIN quests q ON s.quest_id = q.id
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND (w.availability IS NULL OR a.avl_self = ANY(w.availability))
  AND (w.is_repeatable IS NULL OR q.is_repeatable = w.is_repeatable)
ORDER BY s.id;


-- name: FilterSubquestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
)
SELECT DISTINCT s.id
FROM subquests s
JOIN quests q ON s.quest_id = q.id
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND (w.availability IS NULL OR a.avl_self = ANY(w.availability))
  AND (w.is_repeatable IS NULL OR q.is_repeatable = w.is_repeatable)
ORDER BY s.id;


-- name: FilterAreaIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_areas AS (
    SELECT a.a_id, a.avl_self AS current_avl, a.is_repeatable_loc AS is_rep, 'area'::text AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = 'area'

    UNION ALL

    SELECT a.a_id, a.avl_area AS current_avl, a.is_repeatable_loc AS is_rep, 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id
      AND (w.availability IS NULL OR a.avl_area = ANY(w.availability))

    UNION ALL

    SELECT
        mis.area_id AS a_id,
        -- assigning the avl of the item's resource to the location would alter the results
        -- so we assign the lowest rank to not raise MAX avl of GROUP BY bucket
        'always'::availability_type AS current_avl,
        mis.is_repeatable_loc AS is_rep,
        'item'::text AS s_type
    FROM mv_item_sources mis
    JOIN items i ON mis.master_item_id = i.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN w ON mis.area_id = ANY(w.ids)
    WHERE i.id = w.item_id
      AND (w.availability IS NULL OR a.avl_area = ANY(w.availability))
      AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

    UNION ALL

    SELECT
        mis.area_id AS a_id,
        'always'::availability_type AS current_avl,
        mis.is_repeatable_loc AS is_rep,
        'key-item'::text AS s_type
    FROM mv_item_sources mis
    JOIN key_items ki ON mis.master_item_id = ki.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN w ON mis.area_id = ANY(w.ids)
    WHERE ki.id = w.key_item_id
      AND (w.availability IS NULL OR a.avl_area = ANY(w.availability))

    UNION ALL

    SELECT
        a.a_id,
        -- can't use alias in where clause,
        -- so we assign the uninvasive option here and filter in the where clause
        'always'::availability_type AS current_avl,
        a.is_repeatable_loc AS is_rep,
        CASE
          WHEN a.sub_type = 'boss' THEN 'boss'::text
          ELSE a.source_type
        END AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type != 'area'
      AND (w.availability IS NULL OR CASE
        WHEN a.source_type = 'monster' THEN a.avl_area
        ELSE a.avl_self
      END = ANY(w.availability))
)
SELECT a_id
FROM available_areas
CROSS JOIN w
GROUP BY a_id, w.availability, w.item_id, w.monster_id, w.key_item_id, w.reqs, w.excls, w.is_repeatable
HAVING (w.availability IS NULL OR MAX(current_avl) = ANY(w.availability))
   AND (w.item_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['item'])
   AND (w.monster_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['monster'])
   AND (w.key_item_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['key-item'])
   AND (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable)
ORDER BY a_id;




-- name: FilterSublocationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_sublocations AS (
    SELECT s.id AS sublocation_id, s.availability AS current_avl, NULL::boolean AS is_rep, 'sublocation'::text AS s_type
    FROM sublocations s
    JOIN w ON s.id = ANY(w.ids)

    UNION ALL

    SELECT g.sublocation_id, a.avl_context AS current_avl, a.is_repeatable_loc AS is_rep, 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN w ON g.sublocation_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id
      AND (w.availability IS NULL OR a.avl_context = ANY(w.availability))

    UNION ALL

    SELECT g.sublocation_id, 'always'::availability_type AS current_avl, mis.is_repeatable_loc AS is_rep, 'item'::text AS s_type
    FROM mv_item_sources mis
    JOIN items i ON mis.master_item_id = i.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    JOIN w ON g.sublocation_id = ANY(w.ids)
    WHERE i.id = w.item_id
      AND a.avl_context = ANY(w.availability)
      AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

    UNION ALL

    SELECT g.sublocation_id, 'always'::availability_type AS current_avl, mis.is_repeatable_loc AS is_rep, 'key-item'::text AS s_type
    FROM mv_item_sources mis
    JOIN key_items ki ON mis.master_item_id = ki.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    JOIN w ON g.sublocation_id = ANY(w.ids)
    WHERE ki.id = w.key_item_id
      AND a.avl_context = ANY(w.availability)

    UNION ALL

    SELECT
        g.sublocation_id,
        'always'::availability_type AS current_avl,
        a.is_repeatable_loc AS is_rep,
        CASE
          WHEN a.sub_type = 'boss' THEN 'boss'::text
          ELSE a.source_type
        END AS s_type
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN w ON g.sublocation_id = ANY(w.ids)
    WHERE a.source_type != 'area'
      AND CASE
        WHEN a.source_type = 'shop' THEN a.avl_self
        ELSE a.avl_context
      END = ANY(w.availability)
)
SELECT sublocation_id
FROM available_sublocations
CROSS JOIN w
GROUP BY sublocation_id, w.availability, w.item_id, w.monster_id, w.key_item_id, w.reqs, w.excls, w.is_repeatable
HAVING (w.availability IS NULL OR MAX(current_avl) = ANY(w.availability))
   AND (w.item_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['item'])
   AND (w.monster_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['monster'])
   AND (w.key_item_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['key-item'])
   AND (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable)
ORDER BY sublocation_id;





-- name: FilterLocationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::availability_type[] AS availability,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable
),
available_locations AS (
    SELECT l.id AS location_id, l.availability AS current_avl, NULL::boolean AS is_rep, 'location'::text AS s_type
    FROM locations l
    JOIN w ON l.id = ANY(w.ids)

    UNION ALL

    SELECT g.location_id, a.avl_context AS current_avl, a.is_repeatable_loc AS is_rep, 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN w ON g.location_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id
      AND (w.availability IS NULL OR a.avl_context = ANY(w.availability))

    UNION ALL

    SELECT g.location_id, 'always'::availability_type AS current_avl, mis.is_repeatable_loc AS is_rep, 'item'::text AS s_type
    FROM mv_item_sources mis
    JOIN items i ON mis.master_item_id = i.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    JOIN w ON g.location_id = ANY(w.ids)
    WHERE i.id = w.item_id
      AND a.avl_context = ANY(w.availability)
      AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

    UNION ALL

    SELECT g.location_id, 'always'::availability_type AS current_avl, mis.is_repeatable_loc AS is_rep, 'key-item'::text AS s_type
    FROM mv_item_sources mis
    JOIN key_items ki ON mis.master_item_id = ki.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    JOIN w ON g.location_id = ANY(w.ids)
    WHERE ki.id = w.key_item_id
      AND a.avl_context = ANY(w.availability)

    UNION ALL

    SELECT
        g.location_id,
        'always'::availability_type AS current_avl,
        a.is_repeatable_loc AS is_rep,
        CASE
          WHEN a.sub_type = 'boss' THEN 'boss'::text
          ELSE a.source_type
        END AS s_type
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN w ON g.location_id = ANY(w.ids)
    WHERE a.source_type != 'area'
      AND CASE
        WHEN a.source_type = 'shop' THEN a.avl_self
        ELSE a.avl_context
      END = ANY(w.availability)
)
SELECT location_id
FROM available_locations
CROSS JOIN w
GROUP BY location_id, w.availability, w.item_id, w.monster_id, w.key_item_id, w.reqs, w.excls, w.is_repeatable
HAVING (w.availability IS NULL OR MAX(current_avl) = ANY(w.availability))
   AND (w.item_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['item'])
   AND (w.monster_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['monster'])
   AND (w.key_item_id IS NULL OR ARRAY_AGG(s_type) @> ARRAY['key-item'])
   AND (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
   AND (w.is_repeatable IS NULL OR BOOL_OR(is_rep) = w.is_repeatable)
ORDER BY location_id;
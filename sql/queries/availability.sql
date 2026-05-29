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
        END AS current_avl,
        w.availability
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
GROUP BY s_id, availability
HAVING MIN(current_avl) = ANY(availability);


-- name: FilterMasterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
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
        w.availability
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
GROUP BY master_item_id, availability
HAVING MIN(current_avl) = ANY(availability);


-- name: FilterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
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
        w.availability
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
GROUP BY item_id, availability
HAVING MIN(current_avl) = ANY(availability);


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
        a.avl_self AS current_avl,
        w.availability
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
GROUP BY key_item_id, availability
HAVING MIN(current_avl) = ANY(availability);


-- name: FilterSphereIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
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
        w.availability
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
GROUP BY sphere_id, availability
HAVING MIN(current_avl) = ANY(availability);


-- name: FilterPrimerIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability
),
available_primers AS (
    SELECT 
        p.id AS primer_id,
        a.avl_self AS current_avl,
        w.availability
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
GROUP BY primer_id, availability
HAVING MIN(current_avl) = ANY(availability);


-- name: FilterAutoAbilityIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::availability_type[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('character_id')::int AS character_id
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
        w.availability
    FROM mv_auto_ability_sources aas
    JOIN mv_availabilities a ON a.s_id = aas.source_id
     AND a.source_type = aas.source_type
     AND a.a_id = aas.area_id
    JOIN mv_geography g ON aas.area_id = g.area_id
    CROSS JOIN w
    WHERE aas.auto_ability_id = ANY(w.ids)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
          END = w.loc_context_id)
      AND (w.character_id IS NULL OR aas.character_id = w.character_id OR aas.character_id IS NULL)
)
SELECT auto_ability_id
FROM available_auto_abilities
GROUP BY auto_ability_id, availability
HAVING MIN(current_avl) = ANY(availability);



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
        END AS current_avl
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
HAVING MIN(current_avl) = ANY(SELECT availability FROM w);


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
        END AS current_avl,
        w.availability
    FROM mv_availabilities a
    CROSS JOIN w
    WHERE a.source_type = 'shop'
      AND a.s_id = ANY(w.ids)
      AND (w.sub_types IS NULL OR a.sub_type = ANY(w.sub_types))
) available_shops
GROUP BY s_id, availability
HAVING MIN(current_avl) = ANY(availability);


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
        sqlc.arg('availability')::availability_type[] AS availability
)
SELECT DISTINCT a.s_id
FROM mv_availabilities a
CROSS JOIN w
WHERE a.source_type = 'quest'
  AND a.s_id = ANY(w.ids)
  AND a.avl_self = ANY(w.availability)
ORDER BY a.s_id;


-- name: FilterSidequestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability
)
SELECT DISTINCT s.id
FROM sidequests s
JOIN mv_availabilities a ON s.quest_id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND a.avl_self = ANY(w.availability)
ORDER BY s.id;


-- name: FilterSubquestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::availability_type[] AS availability
)
SELECT DISTINCT s.id
FROM subquests s
JOIN mv_availabilities a ON s.quest_id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND a.avl_self = ANY(w.availability)
ORDER BY s.id;


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
    SELECT a.a_id, a.avl_self AS current_avl, w.availability, 'area'::text AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = 'area'

    UNION ALL

    SELECT a.a_id, a.avl_area AS current_avl, w.availability 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id

    UNION ALL

    SELECT
        mis.area_id AS a_id,
        'always'::availability_type AS current_avl, -- lowest rank to not raise MAX avl of GROUP BY bucket
        w.availability,
        'item'::text AS s_type
    FROM mv_item_sources mis
    JOIN items i ON mis.master_item_id = i.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN w ON mis.area_id = ANY(w.ids)
    WHERE i.id = w.item_id
      AND a.avl_area = ANY(w.availability)
      AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

    UNION ALL

    SELECT
        mis.area_id AS a_id,
        'always'::availability_type AS current_avl,
        w.availability
        'key-item'::text AS s_type
    FROM mv_item_sources mis
    JOIN key_items ki ON mis.master_item_id = ki.master_item_id
    JOIN mv_availabilities a ON a.s_id = mis.source_id
     AND a.source_type = mis.source_type
     AND a.a_id = mis.area_id
    JOIN w ON mis.area_id = ANY(w.ids)
    WHERE ki.id = w.key_item_id
      AND a.avl_area = ANY(w.availability)

    UNION ALL

    SELECT
        a.a_id,
        'always'::availability_type AS current_avl, -- avl logic already in where clause
        w.availability,
        CASE
          WHEN a.sub_type = 'boss' THEN 'boss'::text
          ELSE a.source_type
        END AS s_type
    FROM mv_availabilities a
    JOIN w ON a.a_id = ANY(w.ids)
    WHERE a.source_type != 'area'
      AND CASE
        WHEN a.source_type = 'monster' THEN a.avl_area
        ELSE a.avl_self
      END = ANY(w.availability)
)
SELECT a_id
FROM available_areas
GROUP BY a_id, availability
HAVING MAX(current_avl) = ANY(availability)
AND (
    (SELECT reqs FROM w) IS NULL OR
    ARRAY_AGG(DISTINCT s_type) @> (SELECT reqs FROM w)
)
AND (
    (SELECT excls FROM w) IS NULL OR
    NOT ARRAY_AGG(DISTINCT s_type) && (SELECT excls FROM w)
)
ORDER BY a_id;




-- name: FilterSublocationIDsByAvailability :many
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
available_sublocations AS (
    SELECT s.id AS sublocation_id, s.availability AS current_avl, w.availability, 'sublocation'::text AS s_type
    FROM sublocations s
    JOIN w ON s.id = ANY(w.ids)

    UNION ALL

    SELECT g.sublocation_id, a.avl_context AS current_avl, w.availability 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN w ON g.sublocation_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id

    UNION ALL

    SELECT g.sublocation_id, 'always'::availability_type AS current_avl, w.availability, 'item'::text AS s_type
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

    SELECT g.sublocation_id, 'always'::availability_type AS current_avl, w.availability, 'key-item'::text AS s_type
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
        w.availability,
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
GROUP BY sublocation_id, availability
HAVING MAX(current_avl) = ANY(availability)
AND (
    (SELECT reqs FROM w) IS NULL OR
    ARRAY_AGG(DISTINCT s_type) @> (SELECT reqs FROM w)
)
AND (
    (SELECT excls FROM w) IS NULL OR
    NOT ARRAY_AGG(DISTINCT s_type) && (SELECT excls FROM w)
)
ORDER BY sublocation_id;




-- name: FilterLocationIDsByAvailability :many
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
available_locations AS (
    SELECT l.id AS location_id, l.availability AS current_avl, w.availability, 'location'::text AS s_type
    FROM locations l
    JOIN w ON l.id = ANY(w.ids)

    UNION ALL

    SELECT g.location_id, a.avl_context AS current_avl, w.availability, 'monster'::text AS s_type
    FROM mv_availabilities a
    JOIN mv_geography g ON a.a_id = g.area_id
    JOIN w ON g.location_id = ANY(w.ids)
    WHERE a.source_type = 'monster'
      AND w.monster_id IS NOT NULL
      AND a.s_id = w.monster_id

    UNION ALL

    SELECT g.location_id, 'always'::availability_type AS current_avl, w.availability, 'item'::text AS s_type
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

    SELECT g.location_id, 'always'::availability_type AS current_avl, w.availability, 'key-item'::text AS s_type
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
        w.availability,
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
GROUP BY location_id, availability
HAVING MAX(current_avl) = ANY(availability)
AND (
    (SELECT reqs FROM w) IS NULL OR
    ARRAY_AGG(DISTINCT s_type) @> (SELECT reqs FROM w)
)
AND (
    (SELECT excls FROM w) IS NULL OR
    NOT ARRAY_AGG(DISTINCT s_type) && (SELECT excls FROM w)
)
ORDER BY location_id;
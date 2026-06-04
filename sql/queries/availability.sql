CREATE OR REPLACE FUNCTION get_avl_rank(avl availability_type, pre_airship BOOLEAN) 
RETURNS INT AS $$
    SELECT CASE 
        WHEN avl = 'always' THEN 1

        WHEN NOT pre_airship AND avl = 'post' THEN 2
        WHEN NOT pre_airship AND avl = 'pre-story' THEN 3

        WHEN pre_airship AND avl = 'pre-story' THEN 2
        WHEN pre_airship AND avl = 'post' THEN 3

        WHEN avl = 'post-story' THEN 4
    END;
$$ LANGUAGE sql IMMUTABLE;


-- name: FilterMonsterIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
),
raw_monsters AS (
    SELECT
        me.monster_id,
        get_avl_rank(
            CASE 
                WHEN w.avl_type = 'self' THEN me.avl_self
                WHEN w.avl_type = 'context' THEN me.avl_context
                WHEN w.avl_type = 'area' THEN me.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN me.is_repeatable_loc
            ELSE me.is_repeatable
        END AS is_rep
    FROM mv_monster_encounters me
    JOIN mv_geography g ON me.area_id = g.area_id
    CROSS JOIN w
    WHERE me.monster_id = ANY(w.ids)
    AND (w.loc_context_id IS NULL OR CASE
        WHEN w.loc_context_type = 'location' THEN g.location_id
        WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
        WHEN w.loc_context_type = 'area' THEN g.area_id
    END = w.loc_context_id)
)
SELECT monster_id
FROM raw_monsters
CROSS JOIN w
GROUP BY monster_id, w.availability, w.is_repeatable
HAVING 
    CASE
        WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
        WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
        ELSE MIN(current_avl) FILTER (WHERE is_rep = w.is_repeatable) = ANY(w.availability)
    END
ORDER BY monster_id;



-- name: FilterMonsterFormationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
),
raw_formations AS (
    SELECT
        me.formation_id,
        get_avl_rank(
            CASE 
                WHEN w.avl_type = 'self' THEN me.avl_context
                WHEN w.avl_type = 'context' THEN me.avl_context
                WHEN w.avl_type = 'area' THEN me.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        me.is_repeatable_loc AS is_rep
    FROM mv_monster_encounters me
    JOIN mv_geography g ON me.area_id = g.area_id
    CROSS JOIN w
    WHERE me.formation_id = ANY(w.ids)
    AND (w.loc_context_id IS NULL OR CASE
        WHEN w.loc_context_type = 'location' THEN g.location_id
        WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
        WHEN w.loc_context_type = 'area' THEN g.area_id
    END = w.loc_context_id)
)
SELECT formation_id
FROM raw_formations
CROSS JOIN w
GROUP BY formation_id, w.availability, w.is_repeatable
HAVING 
    CASE
        WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
        WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
        ELSE MIN(current_avl) FILTER (WHERE is_rep = w.is_repeatable) = ANY(w.availability)
    END
ORDER BY formation_id;



-- name: FilterMasterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('method')::text AS method
),
raw_master_items AS (
    SELECT
        mis.master_item_id,
        get_avl_rank (
            CASE
                WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN mis.avl_context
                WHEN w.avl_type = 'self' THEN mis.avl_self
                WHEN w.avl_type = 'context' THEN mis.avl_context
                WHEN w.avl_type = 'area' THEN mis.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM mv_item_sources mis
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE mis.master_item_id = ANY(w.ids)
    AND (w.method IS NULL OR mis.source_type = w.method)
    AND (w.loc_context_id IS NULL OR CASE
        WHEN w.loc_context_type = 'location' THEN g.location_id
        WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
        WHEN w.loc_context_type = 'area' THEN g.area_id
    END = w.loc_context_id)
)
SELECT master_item_id
FROM raw_master_items
CROSS JOIN w
GROUP BY master_item_id, w.availability, w.is_repeatable
HAVING 
    CASE
        WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
        WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
        ELSE MIN(current_avl) FILTER (WHERE is_rep = w.is_repeatable) = ANY(w.availability)
    END
ORDER BY master_item_id;


-- name: FilterItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('method')::text AS method
),
raw_items AS (
    SELECT
        i.id AS item_id,
        get_avl_rank (
            CASE
                WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN mis.avl_context
                WHEN w.avl_type = 'self' THEN mis.avl_self
                WHEN w.avl_type = 'context' THEN mis.avl_context
                WHEN w.avl_type = 'area' THEN mis.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM items i
    JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE i.id = ANY(w.ids)
    AND (w.method IS NULL OR mis.source_type = w.method)
    AND (w.loc_context_id IS NULL OR CASE
        WHEN w.loc_context_type = 'location' THEN g.location_id
        WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
        WHEN w.loc_context_type = 'area' THEN g.area_id
    END = w.loc_context_id)
)
SELECT item_id
FROM raw_items
CROSS JOIN w
GROUP BY item_id, w.availability, w.is_repeatable
HAVING 
    CASE
        WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
        WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
        ELSE MIN(current_avl) FILTER (WHERE is_rep = w.is_repeatable) = ANY(w.availability)
    END
ORDER BY item_id;


-- name: FilterKeyItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::int[] AS availability,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type
)
SELECT ki.id
FROM key_items ki
JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
CROSS JOIN w
WHERE ki.id = ANY(w.ids)
  AND (w.loc_context_id IS NULL OR CASE
       WHEN w.loc_context_type = 'location' THEN g.location_id
       WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
       WHEN w.loc_context_type = 'area' THEN g.area_id
  END = w.loc_context_id)
GROUP BY ki.id, w.availability, w.pre_airship
HAVING MIN(get_avl_rank(mis.avl_self, w.pre_airship)) = ANY(w.availability)
ORDER BY ki.id;


-- name: FilterSphereIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('method')::text AS method
),
raw_spheres AS (
    SELECT
        s.id AS sphere_id,
        get_avl_rank (
            CASE
                WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN mis.avl_context
                WHEN w.avl_type = 'self' THEN mis.avl_self
                WHEN w.avl_type = 'context' THEN mis.avl_context
                WHEN w.avl_type = 'area' THEN mis.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM spheres s
    JOIN items i ON s.item_id = i.id
    JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE s.id = ANY(w.ids)
    AND (w.method IS NULL OR mis.source_type = w.method)
    AND (w.loc_context_id IS NULL OR CASE
        WHEN w.loc_context_type = 'location' THEN g.location_id
        WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
        WHEN w.loc_context_type = 'area' THEN g.area_id
    END = w.loc_context_id)
)
SELECT sphere_id
FROM raw_spheres
CROSS JOIN w
GROUP BY sphere_id, w.availability, w.is_repeatable
HAVING
    CASE
        WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
        WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
        ELSE MIN(current_avl) FILTER (WHERE is_rep = w.is_repeatable) = ANY(w.availability)
    END
ORDER BY sphere_id;


-- name: FilterPrimerIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::int[] AS availability,
        sqlc.arg('pre_airship')::boolean AS pre_airship
)
SELECT p.id
FROM primers p
JOIN key_items ki ON p.key_item_id = ki.id
JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
CROSS JOIN w
WHERE p.id = ANY(w.ids)
GROUP BY p.id, w.availability, w.pre_airship
HAVING MIN(get_avl_rank(mis.avl_self, w.pre_airship)) = ANY(w.availability)
ORDER BY p.id;


-- name: FilterAutoAbilityIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('character_id')::int AS character_id,
        sqlc.narg('req_item')::boolean AS req_item
),
available_auto_abilities AS (
    SELECT 
        aas.auto_ability_id,
        get_avl_rank (
            CASE
                WHEN aas.source_type = 'shop' AND w.avl_type = 'self' THEN aas.avl_context
                WHEN w.avl_type = 'self' THEN aas.avl_self
                WHEN w.avl_type = 'context' THEN aas.avl_context
                WHEN w.avl_type = 'area' THEN aas.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN aas.is_repeatable_loc
            ELSE aas.is_repeatable
        END AS is_rep
    FROM mv_auto_ability_sources aas
    JOIN mv_geography g ON aas.area_id = g.area_id
    CROSS JOIN w
    WHERE aas.auto_ability_id = ANY(w.ids)
      AND (w.req_item IS NULL OR w.req_item = FALSE)
      AND (w.loc_context_id IS NULL OR CASE
           WHEN w.loc_context_type = 'location' THEN g.location_id
           WHEN w.loc_context_type = 'sublocation' THEN g.sublocation_id
           WHEN w.loc_context_type = 'area' THEN g.area_id
      END = w.loc_context_id)
      AND (w.character_id IS NULL OR aas.character_id = w.character_id OR aas.character_id IS NULL)

    UNION ALL

    SELECT
        aa.id AS auto_ability_id,
        get_avl_rank (
            CASE
                WHEN mis.source_type = 'shop' AND w.avl_type = 'self' THEN mis.avl_context
                WHEN w.avl_type = 'self' THEN mis.avl_self
                WHEN w.avl_type = 'context' THEN mis.avl_context
                WHEN w.avl_type = 'area' THEN mis.avl_area
            END,
            w.pre_airship
        ) AS current_avl,
        CASE
            WHEN w.loc_context_id IS NOT NULL THEN mis.is_repeatable_loc
            ELSE mis.is_repeatable
        END AS is_rep
    FROM auto_abilities aa
    JOIN item_amounts ia_req ON aa.required_item_amount_id = ia_req.id
    JOIN mv_item_sources mis ON mis.master_item_id = ia_req.master_item_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE aa.id = ANY(w.ids)
      AND w.req_item = TRUE
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
HAVING
    CASE
        WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
        WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
        ELSE MIN(current_avl) FILTER(WHERE is_rep = w.is_repeatable) = ANY(w.availability)
    END
ORDER BY auto_ability_id;


-- name: FilterShopIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::int[] AS availability,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id,
        sqlc.narg('empty_slots')::int[] AS empty_slots,
        sqlc.narg('character_id')::int AS character_id
),
available_shops AS (
    SELECT shop_id, s_type
    FROM (
        SELECT
            a.s_id AS shop_id,
            get_avl_rank(
                CASE 
                    WHEN w.avl_type = 'self' THEN a.avl_self
                    WHEN w.avl_type = 'context' THEN a.avl_context
                    WHEN w.avl_type = 'area' THEN a.avl_area
                END,
                w.pre_airship
            ) AS current_avl,
            a.sub_type AS s_type
        FROM mv_availabilities a
        CROSS JOIN w
        WHERE a.s_id = ANY(w.ids)
          AND a.source_type = 'shop'

        UNION ALL

        SELECT
            es.source_id AS shop_id,
            get_avl_rank (
                CASE 
                    WHEN w.avl_type = 'self' THEN es.avl_self  
                    WHEN w.avl_type = 'context' THEN es.avl_context
                    WHEN w.avl_type = 'area' THEN es.avl_area
                END,
                w.pre_airship
            ) AS current_avl,
            'equip_filter' AS s_type
        FROM mv_equipment_sources es
        CROSS JOIN w
        WHERE es.source_id = ANY(w.ids)
        AND es.source_type = 'shop'
        AND (w.auto_ability_id IS NULL OR es.auto_ability_id = w.auto_ability_id)
        AND (w.empty_slots IS NULL OR es.empty_slots_amount::int = ANY(w.empty_slots))
        AND (w.character_id IS NULL OR es.character_id = w.character_id)
    ) AS raw_sources
    CROSS JOIN w
    GROUP BY shop_id, s_type, w.availability
    HAVING (w.availability IS NULL OR MIN(current_avl) = ANY(w.availability))
)
SELECT shop_id
FROM available_shops
CROSS JOIN w
GROUP BY shop_id, w.reqs, w.excls
HAVING (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY shop_id;


-- name: FilterTreasureIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::int[] AS availability,
        sqlc.arg('pre_airship')::boolean AS pre_airship
)
SELECT DISTINCT a.s_id
FROM mv_availabilities a
CROSS JOIN w
WHERE a.source_type = 'treasure'
  AND a.s_id = ANY(w.ids)
  AND get_avl_rank(a.avl_self, w.pre_airship) = ANY(w.availability)
ORDER BY a.s_id;


-- name: FilterQuestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship
)
SELECT DISTINCT a.s_id
FROM quests q
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE a.s_id = ANY(w.ids)
  AND (w.availability IS NULL OR get_avl_rank(a.avl_self, w.pre_airship) = ANY(w.availability))
  AND (w.is_repeatable IS NULL OR q.is_repeatable = w.is_repeatable)
ORDER BY a.s_id;


-- name: FilterSidequestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship
)
SELECT DISTINCT s.id
FROM sidequests s
JOIN quests q ON s.quest_id = q.id
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND (w.availability IS NULL OR get_avl_rank(a.avl_self, w.pre_airship) = ANY(w.availability))
  AND (w.is_repeatable IS NULL OR q.is_repeatable = w.is_repeatable)
ORDER BY s.id;


-- name: FilterSubquestIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship
)
SELECT DISTINCT s.id
FROM subquests s
JOIN quests q ON s.quest_id = q.id
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND (w.availability IS NULL OR get_avl_rank(a.avl_self, w.pre_airship) = ANY(w.availability))
  AND (w.is_repeatable IS NULL OR q.is_repeatable = w.is_repeatable)
ORDER BY s.id;


-- name: FilterAreaIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods
),
resource_evaluation AS (
    SELECT area_id, s_type FROM (
        SELECT
            a.id AS area_id,
            'area'::text AS s_type,
            get_avl_rank(a.availability, w.pre_airship) AS current_avl,
            FALSE AS is_rep
        FROM areas a
        CROSS JOIN w
        WHERE a.a_id = ANY(w.ids)

        UNION ALL

        SELECT
            me.area_id,
            'monster-single'::text AS s_type,
            get_avl_rank(me.avl_area, w.pre_airship) AS current_avl,
            me.is_repeatable_loc AS is_rep
        FROM mv_monster_encounters me
        CROSS JOIN w
        WHERE me.area_id = ANY(w.ids)
          AND me.monster_id = w.monster_id

        UNION ALL

        SELECT
            mis.area_id,
            'item'::text AS s_type,
            get_avl_rank(mis.avl_area, w.pre_airship) AS current_avl,
            mis.is_repeatable_loc AS is_rep
        FROM mv_item_sources mis
        JOIN items i ON mis.master_item_id = i.master_item_id
        CROSS JOIN w
        WHERE mis.area_id = ANY(w.ids)
          AND i.id = w.item_id
          AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

        UNION ALL

        SELECT
            mis.area_id,
            'key-item'::text AS s_type,
            get_avl_rank(mis.avl_area, w.pre_airship) AS current_avl,
            mis.is_repeatable_loc AS is_rep
        FROM mv_item_sources mis
        JOIN key_items ki ON mis.master_item_id = ki.master_item_id
        CROSS JOIN w
        WHERE mis.area_id = ANY(w.ids)
          AND ki.id = w.key_item_id
    ) as r
    CROSS JOIN w
    GROUP BY area_id, s_type, w.availability, w.is_repeatable
    HAVING
        CASE
            WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
            WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
            ELSE MIN(current_avl) FILTER(WHERE is_rep = w.is_repeatable) = ANY(w.availability)
        END
),
content_evaluation AS (
    SELECT area_id, s_type FROM (
        SELECT
            a.a_id AS area_id,
            a.s_id AS source_id,
            a.source_type AS s_type,
            get_avl_rank(
                CASE
                    WHEN a.source_type = 'monster' THEN a.avl_area
                    ELSE a.avl_self
                END,
                w.pre_airship
            ) AS current_avl
        FROM mv_availabilities a
        CROSS JOIN w
        WHERE a.a_id = ANY(w.ids)
          AND a.source_type != 'area'

        UNION ALL

        SELECT
            a.a_id,
            a.s_id as source_id,
            a.sub_type AS s_type,
            get_avl_rank(a.avl_area, w.pre_airship) AS current_avl
        FROM mv_availabilities a
        CROSS JOIN w
        WHERE a.a_id = ANY(w.ids)
          AND a.sub_type = 'boss'
    ) as c
    CROSS JOIN w
    GROUP BY area_id, source_id, s_type, w.availability
    HAVING w.availability IS NULL OR MIN(current_avl) = ANY(w.availability)
),
final_combination AS (
    SELECT area_id, s_type FROM resource_evaluation
    UNION ALL
    SELECT area_id, s_type FROM content_evaluation
)
SELECT area_id
FROM final_combination
CROSS JOIN w
GROUP BY area_id, w.reqs, w.excls
HAVING (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY area_id;


-- name: FilterSublocationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods
),
resource_evaluation AS (
    SELECT sublocation_id, s_type FROM (
        SELECT
            s.id AS sublocation_id,
            'sublocation'::text AS s_type,
            get_avl_rank(s.availability, w.pre_airship) AS current_avl,
            FALSE AS is_rep
        FROM sublocations s
        CROSS JOIN w
        WHERE s.id = ANY(w.ids)

        UNION ALL

        SELECT
            g.sublocation_id,
            'monster-single'::text AS s_type,
            get_avl_rank(me.avl_context, w.pre_airship) AS current_avl,
            me.is_repeatable_loc AS is_rep
        FROM mv_monster_encounters me
        JOIN mv_geography g ON me.area_id = g.area_id
        CROSS JOIN w
        WHERE g.sublocation_id = ANY(w.ids)
          AND me.monster_id = w.monster_id

        UNION ALL

        SELECT
            g.sublocation_id,
            'item'::text AS s_type,
            get_avl_rank(mis.avl_context, w.pre_airship) AS current_avl,
            mis.is_repeatable_loc AS is_rep
        FROM mv_item_sources mis
        JOIN items i ON mis.master_item_id = i.master_item_id
        JOIN mv_geography g ON mis.area_id = g.area_id
        CROSS JOIN w
        WHERE g.sublocation_id = ANY(w.ids)
          AND i.id = w.item_id
          AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

        UNION ALL

        SELECT
            g.sublocation_id,
            'key-item'::text AS s_type,
            get_avl_rank(mis.avl_context, w.pre_airship) AS current_avl,
            mis.is_repeatable_loc AS is_rep
        FROM mv_item_sources mis
        JOIN key_items ki ON mis.master_item_id = ki.master_item_id
        JOIN mv_geography g ON mis.area_id = g.area_id
        CROSS JOIN w
        WHERE g.sublocation_id = ANY(w.ids)
          AND ki.id = w.key_item_id
    ) as r
    CROSS JOIN w
    GROUP BY sublocation_id, s_type, w.availability, w.is_repeatable
    HAVING
        CASE
            WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
            WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
            ELSE MIN(current_avl) FILTER(WHERE is_rep = w.is_repeatable) = ANY(w.availability)
        END
),
content_evaluation AS (
    SELECT sublocation_id, s_type FROM (
        SELECT
            g.sublocation_id,
            a.s_id as source_id, 
            a.source_type AS s_type,
            get_avl_rank(
                CASE
                    WHEN a.source_type = 'shop' THEN a.avl_self
                    ELSE a.avl_context
                END,
                w.pre_airship
            ) AS current_avl
        FROM mv_availabilities a
        JOIN mv_geography g ON a.a_id = g.area_id
        CROSS JOIN w
        WHERE g.sublocation_id = ANY(w.ids)
          AND a.source_type != 'area'

        UNION ALL

        SELECT
            g.sublocation_id,
            a.s_id as source_id,
            a.sub_type AS s_type,
            get_avl_rank(a.avl_context, w.pre_airship) AS current_avl
        FROM mv_availabilities a
        JOIN mv_geography g ON a.a_id = g.area_id
        CROSS JOIN w
        WHERE g.sublocation_id = ANY(w.ids)
          AND a.sub_type = 'boss'
    ) as c
    CROSS JOIN w
    GROUP BY sublocation_id, source_id, s_type, w.availability
    HAVING w.availability IS NULL OR MIN(current_avl) = ANY(w.availability)
),
final_combination AS (
    SELECT sublocation_id, s_type FROM resource_evaluation
    UNION ALL
    SELECT sublocation_id, s_type FROM content_evaluation
)
SELECT sublocation_id
FROM final_combination
CROSS JOIN w
GROUP BY sublocation_id, w.reqs, w.excls
HAVING (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY sublocation_id;



-- name: FilterLocationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('methods')::text[] AS methods
),
resource_evaluation AS (
    SELECT location_id, s_type FROM (
        SELECT
            l.id AS location_id,
            'location'::text AS s_type,
            get_avl_rank(l.availability, w.pre_airship) AS current_avl,
            FALSE AS is_rep
        FROM locations l
        CROSS JOIN w
        WHERE l.id = ANY(w.ids)

        UNION ALL

        SELECT
            g.location_id,
            'monster-single'::text AS s_type,
            get_avl_rank(me.avl_context, w.pre_airship) AS current_avl,
            me.is_repeatable_loc AS is_rep
        FROM mv_monster_encounters me
        JOIN mv_geography g ON me.area_id = g.area_id
        CROSS JOIN w
        WHERE g.location_id = ANY(w.ids)
          AND me.monster_id = w.monster_id

        UNION ALL

        SELECT
            g.location_id,
            'item'::text AS s_type,
            get_avl_rank(mis.avl_context, w.pre_airship) AS current_avl,
            mis.is_repeatable_loc AS is_rep
        FROM mv_item_sources mis
        JOIN items i ON mis.master_item_id = i.master_item_id
        JOIN mv_geography g ON mis.area_id = g.area_id
        CROSS JOIN w
        WHERE g.location_id = ANY(w.ids)
          AND i.id = w.item_id
          AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))

        UNION ALL

        SELECT
            g.location_id,
            'key-item'::text AS s_type,
            get_avl_rank(mis.avl_context, w.pre_airship) AS current_avl,
            mis.is_repeatable_loc AS is_rep
        FROM mv_item_sources mis
        JOIN key_items ki ON mis.master_item_id = ki.master_item_id
        JOIN mv_geography g ON mis.area_id = g.area_id
        CROSS JOIN w
        WHERE g.location_id = ANY(w.ids)
          AND ki.id = w.key_item_id
    ) as r
    CROSS JOIN w
    GROUP BY location_id, s_type, w.availability, w.is_repeatable
    HAVING
        CASE
            WHEN w.is_repeatable IS NULL THEN MIN(current_avl) = ANY(w.availability)
            WHEN w.availability IS NULL THEN BOOL_OR(is_rep = w.is_repeatable)
            ELSE MIN(current_avl) FILTER(WHERE is_rep = w.is_repeatable) = ANY(w.availability)
        END
),
content_evaluation AS (
    SELECT location_id, s_type FROM (
        SELECT
            g.location_id,
            a.s_id as source_id, 
            a.source_type AS s_type,
            get_avl_rank(
                CASE
                    WHEN a.source_type = 'shop' THEN a.avl_self
                    ELSE a.avl_context
                END,
                w.pre_airship
            ) AS current_avl
        FROM mv_availabilities a
        JOIN mv_geography g ON a.a_id = g.area_id
        CROSS JOIN w
        WHERE g.location_id = ANY(w.ids)
          AND a.source_type != 'area'

        UNION ALL

        SELECT
            g.location_id,
            a.s_id as source_id,
            a.sub_type AS s_type,
            get_avl_rank(a.avl_context, w.pre_airship) AS current_avl
        FROM mv_availabilities a
        JOIN mv_geography g ON a.a_id = g.area_id
        CROSS JOIN w
        WHERE g.location_id = ANY(w.ids)
          AND a.sub_type = 'boss'
    ) as c
    CROSS JOIN w
    GROUP BY location_id, source_id, s_type, w.availability
    HAVING w.availability IS NULL OR MIN(current_avl) = ANY(w.availability)
),
final_combination AS (
    SELECT location_id, s_type FROM resource_evaluation
    UNION ALL
    SELECT location_id, s_type FROM content_evaluation
)
SELECT location_id
FROM final_combination
CROSS JOIN w
GROUP BY location_id, w.reqs, w.excls
HAVING (w.reqs IS NULL OR ARRAY_AGG(s_type) @> w.reqs)
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY location_id;
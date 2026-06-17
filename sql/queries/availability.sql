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
        get_avl_rank(select_avl(w.avl_type, me.avl_self, me.avl_context, me.avl_context_2), w.pre_airship) AS current_avl,
        get_is_rep(w.loc_context_id, me.is_repeatable, me.is_repeatable_loc) AS is_rep
    FROM mv_monster_encounters me
    JOIN mv_geography g ON me.area_id = g.area_id
    CROSS JOIN w
    WHERE me.monster_id = ANY(w.ids)
    AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
),
monsters_rep AS (
    SELECT
        monster_id,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_monsters
    GROUP BY monster_id, current_avl
)
SELECT m.monster_id
FROM monsters_rep m
CROSS JOIN w
GROUP BY m.monster_id, w.availability, w.is_repeatable
HAVING filter_avl_rep(
    MIN(m.current_avl),
    BOOL_OR(m.is_rep = w.is_repeatable),
    MIN(m.current_avl) FILTER (WHERE m.is_rep = w.is_repeatable),
    BOOL_OR(m.is_rep) FILTER (WHERE m.current_avl = ANY(w.availability)),
    MIN(m.current_avl) FILTER (WHERE m.is_rep = FALSE),
    w.availability,
    w.is_repeatable
)
ORDER BY m.monster_id;



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
        get_avl_rank(select_avl(w.avl_type, me.avl_context, me.avl_context, me.avl_context_2), w.pre_airship) AS current_avl,
        me.is_repeatable_loc AS is_rep
    FROM mv_monster_encounters me
    JOIN mv_geography g ON me.area_id = g.area_id
    CROSS JOIN w
    WHERE me.formation_id = ANY(w.ids)
    AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
),
formations_rep AS (
    SELECT
        formation_id,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_formations
    GROUP BY formation_id, current_avl
)
SELECT mf.formation_id
FROM formations_rep mf
CROSS JOIN w
GROUP BY mf.formation_id, w.availability, w.is_repeatable
HAVING filter_avl_rep(
    MIN(mf.current_avl),
    BOOL_OR(mf.is_rep = w.is_repeatable),
    MIN(mf.current_avl) FILTER (WHERE mf.is_rep = w.is_repeatable),
    BOOL_OR(mf.is_rep) FILTER (WHERE mf.current_avl = ANY(w.availability)),
    MIN(mf.current_avl) FILTER (WHERE mf.is_rep = FALSE),
    w.availability,
    w.is_repeatable
)
ORDER BY mf.formation_id;



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
        sqlc.narg('methods')::text[] AS methods
),
raw_master_items AS (
    SELECT
        mis.master_item_id,
        get_avl_rank_item(mis.source_type, w.avl_type, w.loc_context_type, w.pre_airship, mis.avl_self, mis.avl_context, mis.avl_context_2, mis.avl_shop_loc, mis.avl_shop_subloc, mis.avl_shop_area) AS current_avl,
        get_is_rep(w.loc_context_id, mis.is_repeatable, mis.is_repeatable_loc) AS is_rep
    FROM mv_item_sources mis
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE mis.master_item_id = ANY(w.ids)
    AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))
    AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
),
master_items_rep AS (
    SELECT
        master_item_id,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_master_items
    GROUP BY master_item_id, current_avl
)
SELECT mi.master_item_id
FROM master_items_rep mi
CROSS JOIN w
GROUP BY mi.master_item_id, w.availability, w.is_repeatable
HAVING filter_avl_rep(
    MIN(mi.current_avl),
    BOOL_OR(mi.is_rep = w.is_repeatable),
    MIN(mi.current_avl) FILTER (WHERE mi.is_rep = w.is_repeatable),
    BOOL_OR(mi.is_rep) FILTER (WHERE mi.current_avl = ANY(w.availability)),
    MIN(mi.current_avl) FILTER (WHERE mi.is_rep = FALSE),
    w.availability,
    w.is_repeatable
)
ORDER BY mi.master_item_id;



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
        sqlc.narg('methods')::text[] AS methods
),
raw_items AS (
    SELECT
        i.id AS item_id,
        get_avl_rank_item(mis.source_type, w.avl_type, w.loc_context_type, w.pre_airship, mis.avl_self, mis.avl_context, mis.avl_context_2, mis.avl_shop_loc, mis.avl_shop_subloc, mis.avl_shop_area) AS current_avl,
        get_is_rep(w.loc_context_id, mis.is_repeatable, mis.is_repeatable_loc) AS is_rep
    FROM items i
    JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE i.id = ANY(w.ids)
    AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))
    AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
),
items_rep AS (
    SELECT
        item_id,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_items
    GROUP BY item_id, current_avl
)
SELECT i.item_id
FROM items_rep i
CROSS JOIN w
GROUP BY i.item_id, w.availability, w.is_repeatable
HAVING filter_avl_rep(
    MIN(i.current_avl),
    BOOL_OR(i.is_rep = w.is_repeatable),
    MIN(i.current_avl) FILTER (WHERE i.is_rep = w.is_repeatable),
    BOOL_OR(i.is_rep) FILTER (WHERE i.current_avl = ANY(w.availability)),
    MIN(i.current_avl) FILTER (WHERE i.is_rep = FALSE),
    w.availability,
    w.is_repeatable
)
ORDER BY i.item_id;


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
        sqlc.narg('methods')::text[] AS methods
),
raw_spheres AS (
    SELECT
        s.id AS sphere_id,
        get_avl_rank_item(mis.source_type, w.avl_type, w.loc_context_type, w.pre_airship, mis.avl_self, mis.avl_context, mis.avl_context_2, mis.avl_shop_loc, mis.avl_shop_subloc, mis.avl_shop_area) AS current_avl,
        get_is_rep(w.loc_context_id, mis.is_repeatable, mis.is_repeatable_loc) AS is_rep
    FROM spheres s
    JOIN items i ON s.item_id = i.id
    JOIN mv_item_sources mis ON mis.master_item_id = i.master_item_id
    JOIN mv_geography g ON mis.area_id = g.area_id
    CROSS JOIN w
    WHERE s.id = ANY(w.ids)
    AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))
    AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
),
spheres_rep AS (
    SELECT
        sphere_id,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_spheres
    GROUP BY sphere_id, current_avl
)
SELECT s.sphere_id
FROM spheres_rep s
CROSS JOIN w
GROUP BY s.sphere_id, w.availability, w.is_repeatable
HAVING filter_avl_rep(
    MIN(s.current_avl),
    BOOL_OR(s.is_rep = w.is_repeatable),
    MIN(s.current_avl) FILTER (WHERE s.is_rep = w.is_repeatable),
    BOOL_OR(s.is_rep) FILTER (WHERE s.current_avl = ANY(w.availability)),
    MIN(s.current_avl) FILTER (WHERE s.is_rep = FALSE),
    w.availability,
    w.is_repeatable
)
ORDER BY s.sphere_id;


-- name: FilterKeyItemIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('availability')::int[] AS availability,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
        sqlc.narg('methods')::text[] AS methods
)
SELECT ki.id
FROM key_items ki
JOIN mv_item_sources mis ON mis.master_item_id = ki.master_item_id
JOIN mv_geography g ON mis.area_id = g.area_id
CROSS JOIN w
WHERE ki.id = ANY(w.ids)
  AND (w.methods IS NULL OR mis.source_type = ANY(w.methods))
  AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
GROUP BY ki.id, w.availability, w.pre_airship
HAVING MIN(get_avl_rank(mis.avl_self, w.pre_airship)) = ANY(w.availability)
ORDER BY ki.id;



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
        sqlc.narg('methods')::text[] AS methods,
        sqlc.arg('req_item')::boolean AS req_item
),
raw_auto_abilities AS (
    SELECT
        aas.auto_ability_id,
        get_avl_rank_auto_ability(aas.source_type, w.avl_type, w.loc_context_type, aas.character_id, w.pre_airship, aas.avl_self, aas.avl_context, aas.avl_context_2, aas.avl_shop_loc, aas.avl_shop_subloc, aas.avl_shop_area, aas.avl_shop_loc_char, aas.avl_shop_subloc_char, aas.avl_shop_area_char) AS current_avl,
        get_is_rep(w.loc_context_id, aas.is_repeatable, aas.is_repeatable_loc) AS is_rep
    FROM mv_auto_ability_sources aas
    JOIN mv_geography g ON aas.area_id = g.area_id
    CROSS JOIN w
    WHERE aas.auto_ability_id = ANY(w.ids)
    AND w.req_item = FALSE
    AND (w.character_id IS NULL OR aas.character_id = w.character_id OR aas.character_id IS NULL)
    AND (w.methods IS NULL OR aas.source_type = ANY(w.methods))
    AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)

    UNION ALL

    SELECT sub.auto_ability_id, sub.current_avl, sub.is_rep FROM (
        SELECT
            calc.auto_ability_id,
            calc.current_avl,
            calc.is_rep,
            BOOL_OR(calc.is_rep) FILTER (WHERE w.availability IS NULL OR calc.current_avl = ANY(w.availability)) 
                OVER (PARTITION BY calc.auto_ability_id) AS group_is_rep,
            SUM(calc.amount) FILTER (WHERE NOT calc.is_rep AND (w.availability IS NULL OR calc.current_avl = ANY(w.availability)))
                OVER (PARTITION BY calc.auto_ability_id) AS group_total_amt,
            calc.req_amount
        FROM (
            SELECT
                aa.id AS auto_ability_id,
                mis.amount,
                ia_req.amount AS req_amount,
                get_avl_rank_item(mis.source_type, w.avl_type, w.loc_context_type, w.pre_airship, mis.avl_self, mis.avl_context, mis.avl_context_2, mis.avl_shop_loc, mis.avl_shop_subloc, mis.avl_shop_area) AS current_avl,
                get_is_rep(w.loc_context_id, mis.is_repeatable, mis.is_repeatable_loc) AS is_rep
            FROM auto_abilities aa
            JOIN item_amounts ia_req ON aa.required_item_amount_id = ia_req.id
            JOIN mv_item_sources mis ON mis.master_item_id = ia_req.master_item_id
            JOIN mv_geography g ON mis.area_id = g.area_id
            CROSS JOIN w
            WHERE aa.id = ANY(w.ids)
            AND w.req_item = TRUE
            AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
        ) AS calc
        CROSS JOIN w
    ) AS sub
    WHERE sub.group_is_rep OR (COALESCE(sub.group_total_amt, 0) >= sub.req_amount)
),
auto_abilities_rep AS (
    SELECT
        auto_ability_id,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_auto_abilities
    GROUP BY auto_ability_id, current_avl
)
SELECT a.auto_ability_id
FROM auto_abilities_rep a
CROSS JOIN w
GROUP BY a.auto_ability_id, w.availability, w.is_repeatable
HAVING filter_avl_rep(
    MIN(a.current_avl),
    BOOL_OR(a.is_rep = w.is_repeatable),
    MIN(a.current_avl) FILTER (WHERE a.is_rep = w.is_repeatable),
    BOOL_OR(a.is_rep) FILTER (WHERE a.current_avl = ANY(w.availability)),
    MIN(a.current_avl) FILTER (WHERE a.is_rep = FALSE),
    w.availability,
    w.is_repeatable
)
ORDER BY a.auto_ability_id;


-- name: FilterShopIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.arg('avl_type')::text AS avl_type,
        sqlc.arg('availability')::int[] AS availability,
        sqlc.narg('loc_context_id')::int AS loc_context_id,
        sqlc.narg('loc_context_type')::text AS loc_context_type,
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
            get_avl_rank(select_avl(w.avl_type, a.avl_self, a.avl_context, a.avl_context_2), FALSE) AS current_avl,
            a.sub_type AS s_type
        FROM mv_availabilities a
        JOIN mv_geography g ON a.a_id = g.area_id
        CROSS JOIN w
        WHERE a.s_id = ANY(w.ids)
          AND (w.auto_ability_id IS NULL AND w.empty_slots IS NULL AND w.character_id IS NULL)
          AND a.source_type = 'shop'
          AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)

        UNION ALL

        SELECT
            sa.shop_id,
            get_avl_rank(select_shop_equip_avl(w.auto_ability_id, w.character_id, w.empty_slots, sa.avl_acs, sa.avl_ac, sa.avl_as, sa.avl_cs, sa.avl_a, sa.avl_c, sa.avl_s, sa.avl_self), FALSE) AS current_avl,
            'equip-filter' AS s_type
        FROM mv_shop_equip_avls sa
        JOIN mv_geography g ON sa.area_id = g.area_id
        CROSS JOIN w
        WHERE sa.shop_id = ANY(w.ids)
        AND (w.auto_ability_id IS NOT NULL OR w.empty_slots IS NOT NULL OR w.character_id IS NOT NULL)
        AND (w.auto_ability_id IS NULL OR sa.auto_ability_id = w.auto_ability_id)
        AND (w.empty_slots IS NULL OR sa.empty_slots_amount::int = ANY(w.empty_slots))
        AND (w.character_id IS NULL OR sa.character_id = w.character_id)
        AND (w.loc_context_id IS NULL OR get_loc_ctx_id(w.loc_context_type, g.location_id, g.sublocation_id, g.area_id) = w.loc_context_id)
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
        sqlc.arg('pre_airship')::boolean AS pre_airship
)
SELECT DISTINCT s.id
FROM sidequests s
JOIN quests q ON s.quest_id = q.id
JOIN mv_availabilities a ON q.id = a.s_id AND a.source_type = 'quest'
CROSS JOIN w
WHERE s.id = ANY(w.ids)
  AND (w.availability IS NULL OR get_avl_rank(a.avl_self, w.pre_airship) = ANY(w.availability))
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


-- name: FilterAreaIDsByAvailabilitySoft :many
SELECT DISTINCT id FROM areas WHERE get_avl_rank(availability, false) = ANY(sqlc.narg('availability')::int[]) ORDER BY id;


-- name: FilterSublocationIDsByAvailabilitySoft :many
SELECT DISTINCT id FROM sublocations WHERE get_avl_rank(availability, false) = ANY(sqlc.narg('availability')::int[]) ORDER BY id;


-- name: FilterLocationIDsByAvailabilitySoft :many
SELECT DISTINCT id FROM locations WHERE get_avl_rank(availability, false) = ANY(sqlc.narg('availability')::int[]) ORDER BY id;


-- name: FilterAreaIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.arg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id,
        sqlc.narg('methods')::text[] AS methods
),
all_areas AS (
    SELECT id AS area_id, 'area' AS s_type 
    FROM areas
    CROSS JOIN w
    WHERE id = ANY(w.ids)
),
raw_res_areas AS (
    SELECT
        me.area_id,
        'monster-single'::text AS s_type,
        get_avl_rank(me.avl_context_2, w.pre_airship) AS current_avl,
        me.is_repeatable_loc AS is_rep
    FROM mv_monster_encounters me
    CROSS JOIN w
    WHERE me.area_id = ANY(w.ids)
        AND me.monster_id = w.monster_id

    UNION ALL

    SELECT
        mis.area_id,
        'item'::text AS s_type,
        get_avl_rank(mis.avl_context_2, w.pre_airship) AS current_avl,
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
        get_avl_rank(mis.avl_context_2, w.pre_airship) AS current_avl,
        mis.is_repeatable_loc AS is_rep
    FROM mv_item_sources mis
    JOIN key_items ki ON mis.master_item_id = ki.master_item_id
    CROSS JOIN w
    WHERE mis.area_id = ANY(w.ids)
        AND ki.id = w.key_item_id

    UNION ALL

    SELECT
        aas.area_id,
        'auto-ability'::text AS s_type,
        get_avl_rank(aas.avl_context_2, w.pre_airship) AS current_avl,
        aas.is_repeatable_loc AS is_rep
    FROM mv_auto_ability_sources aas
    CROSS JOIN w
    WHERE aas.area_id = ANY(w.ids)
        AND aas.auto_ability_id = w.auto_ability_id
),
raw_res_areas_rep AS (
    SELECT
        area_id,
        s_type,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_res_areas
    GROUP BY area_id, s_type, current_avl
),
filtered_res_areas AS (
    SELECT a.area_id, a.s_type
    FROM raw_res_areas_rep a
    CROSS JOIN w
    GROUP BY a.area_id, a.s_type, w.availability, w.is_repeatable
    HAVING filter_avl_rep(
        MIN(a.current_avl),
        BOOL_OR(a.is_rep = w.is_repeatable),
        MIN(a.current_avl) FILTER (WHERE a.is_rep = w.is_repeatable),
        BOOL_OR(a.is_rep) FILTER (WHERE a.current_avl = ANY(w.availability)),
        MIN(a.current_avl) FILTER (WHERE a.is_rep = FALSE),
        w.availability,
        w.is_repeatable
    )
),
content_areas AS (
    SELECT area_id, s_type FROM (
        SELECT
            a.a_id AS area_id,
            a.s_id AS source_id,
            a.source_type AS s_type,
            get_avl_rank_condition(a.avl_context_2, a.avl_self, w.pre_airship, (a.source_type = 'monster')) AS current_avl
        FROM mv_availabilities a
        CROSS JOIN w
        WHERE a.a_id = ANY(w.ids)
          AND a.source_type != 'area'

        UNION ALL

        SELECT
            a.a_id as area_id,
            a.s_id as source_id,
            a.sub_type AS s_type,
            get_avl_rank(a.avl_context_2, w.pre_airship) AS current_avl
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
    SELECT area_id, s_type FROM all_areas
    UNION ALL
    SELECT area_id, s_type FROM filtered_res_areas
    UNION ALL
    SELECT area_id, s_type FROM content_areas
)
SELECT area_id
FROM final_combination
CROSS JOIN w
GROUP BY area_id, w.reqs, w.excls
HAVING ARRAY_AGG(s_type) @> w.reqs
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY area_id;




-- name: FilterSublocationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.arg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id,
        sqlc.narg('methods')::text[] AS methods
),
all_sublocations AS (
    SELECT id AS sublocation_id, 'sublocation' AS s_type 
    FROM sublocations
    CROSS JOIN w
    WHERE id = ANY(w.ids)
),
raw_res_sublocations AS (
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

    UNION ALL

    SELECT
        g.sublocation_id,
        'auto-ability'::text AS s_type,
        get_avl_rank(aas.avl_context, w.pre_airship) AS current_avl,
        aas.is_repeatable_loc AS is_rep
    FROM mv_auto_ability_sources aas
    JOIN mv_geography g ON aas.area_id = g.area_id
    CROSS JOIN w
    WHERE aas.area_id = ANY(w.ids)
        AND aas.auto_ability_id = w.auto_ability_id
),
raw_res_sublocations_rep AS (
    SELECT
        sublocation_id,
        s_type,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_res_sublocations
    GROUP BY sublocation_id, s_type, current_avl
),
filtered_res_sublocations AS (
    SELECT s.sublocation_id, s.s_type
    FROM raw_res_sublocations_rep s
    CROSS JOIN w
    GROUP BY s.sublocation_id, s.s_type, w.availability, w.is_repeatable
    HAVING filter_avl_rep(
        MIN(s.current_avl),
        BOOL_OR(s.is_rep = w.is_repeatable),
        MIN(s.current_avl) FILTER (WHERE s.is_rep = w.is_repeatable),
        BOOL_OR(s.is_rep) FILTER (WHERE s.current_avl = ANY(w.availability)),
        MIN(s.current_avl) FILTER (WHERE s.is_rep = FALSE),
        w.availability,
        w.is_repeatable
    )
),
content_sublocations AS (
    SELECT sublocation_id, s_type FROM (
        SELECT
            g.sublocation_id,
            a.s_id as source_id, 
            a.source_type AS s_type,
            get_avl_rank_condition(a.avl_self, a.avl_context, w.pre_airship, (a.source_type = 'shop')) AS current_avl
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
    SELECT sublocation_id, s_type FROM all_sublocations
    UNION ALL
    SELECT sublocation_id, s_type FROM filtered_res_sublocations
    UNION ALL
    SELECT sublocation_id, s_type FROM content_sublocations
)
SELECT sublocation_id
FROM final_combination
CROSS JOIN w
GROUP BY sublocation_id, w.reqs, w.excls
HAVING ARRAY_AGG(s_type) @> w.reqs
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY sublocation_id;



-- name: FilterLocationIDsByAvailability :many
WITH w AS (
    SELECT
        sqlc.arg('ids')::int[] AS ids,
        sqlc.narg('availability')::int[] AS availability,
        sqlc.narg('is_repeatable')::boolean AS is_repeatable,
        sqlc.arg('pre_airship')::boolean AS pre_airship,
        sqlc.arg('required_sources')::text[] AS reqs,
        sqlc.narg('excluded_sources')::text[] AS excls,
        sqlc.narg('monster_id')::int AS monster_id,
        sqlc.narg('key_item_id')::int AS key_item_id,
        sqlc.narg('item_id')::int AS item_id,
        sqlc.narg('auto_ability_id')::int AS auto_ability_id,
        sqlc.narg('methods')::text[] AS methods
),
all_locations AS (
    SELECT id AS location_id, 'location' AS s_type 
    FROM locations
    CROSS JOIN w
    WHERE id = ANY(w.ids)
),
raw_res_locations AS (
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

    UNION ALL

    SELECT
        g.sublocation_id,
        'auto-ability'::text AS s_type,
        get_avl_rank(aas.avl_context, w.pre_airship) AS current_avl,
        aas.is_repeatable_loc AS is_rep
    FROM mv_auto_ability_sources aas
    JOIN mv_geography g ON aas.area_id = g.area_id
    CROSS JOIN w
    WHERE aas.area_id = ANY(w.ids)
        AND aas.auto_ability_id = w.auto_ability_id
),
raw_res_locations_rep AS (
    SELECT
        location_id,
        s_type,
        current_avl,
        BOOL_OR(is_rep) AS is_rep
    FROM raw_res_locations
    GROUP BY location_id, s_type, current_avl
),
filtered_res_locations AS (
    SELECT l.location_id, l.s_type
    FROM raw_res_locations_rep l
    CROSS JOIN w
    GROUP BY l.location_id, l.s_type, w.availability, w.is_repeatable
    HAVING filter_avl_rep(
        MIN(l.current_avl),
        BOOL_OR(l.is_rep = w.is_repeatable),
        MIN(l.current_avl) FILTER (WHERE l.is_rep = w.is_repeatable),
        BOOL_OR(l.is_rep) FILTER (WHERE l.current_avl = ANY(w.availability)),
        MIN(l.current_avl) FILTER (WHERE l.is_rep = FALSE),
        w.availability,
        w.is_repeatable
    )
),
content_locations AS (
    SELECT location_id, s_type FROM (
        SELECT
            g.location_id,
            a.s_id as source_id, 
            a.source_type AS s_type,
            get_avl_rank_condition(a.avl_self, a.avl_context, w.pre_airship, (a.source_type = 'shop')) AS current_avl
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
    SELECT location_id, s_type FROM all_locations
    UNION ALL
    SELECT location_id, s_type FROM filtered_res_locations
    UNION ALL
    SELECT location_id, s_type FROM content_locations
)
SELECT location_id
FROM final_combination
CROSS JOIN w
GROUP BY location_id, w.reqs, w.excls
HAVING ARRAY_AGG(s_type) @> w.reqs
   AND (w.excls IS NULL OR NOT ARRAY_AGG(s_type) && w.excls)
ORDER BY location_id;
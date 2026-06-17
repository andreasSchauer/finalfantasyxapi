-- +goose Up


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_loc_ctx_id(
    w_loc_context_type TEXT, 
    location_id INTEGER,
    sublocation_id INTEGER,
    area_id INTEGER
)
RETURNS INT AS $$
    SELECT CASE
        WHEN w_loc_context_type = 'location' THEN location_id
        WHEN w_loc_context_type = 'sublocation' THEN sublocation_id
        WHEN w_loc_context_type = 'area' THEN area_id
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION select_avl(
    w_avl_type TEXT,
    avl_self availability_type,
    avl_context availability_type,
    avl_context_2 availability_type
)
RETURNS availability_type AS $$
    SELECT CASE
        WHEN w_avl_type = 'self' THEN avl_self
        WHEN w_avl_type = 'context' THEN avl_context
        WHEN w_avl_type = 'context-2' THEN avl_context_2
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION select_shop_equip_avl(
    auto_ability_id INTEGER,
    character_id INTEGER,
    empty_slots_amount INTEGER[],
    avl_acs availability_type,
    avl_ac availability_type,
    avl_as availability_type,
    avl_cs availability_type,
    avl_a availability_type,
    avl_c availability_type,
    avl_s availability_type,
    avl_self availability_type
)
RETURNS availability_type AS $$
    SELECT CASE
        WHEN (auto_ability_id IS NOT NULL AND character_id IS NOT NULL AND empty_slots_amount IS NOT NULL) THEN avl_acs
        WHEN (auto_ability_id IS NOT NULL AND character_id IS NOT NULL) THEN avl_ac
        WHEN (auto_ability_id IS NOT NULL AND empty_slots_amount IS NOT NULL) THEN avl_as
        WHEN (character_id IS NOT NULL AND empty_slots_amount IS NOT NULL) THEN avl_cs
        WHEN (auto_ability_id IS NOT NULL) THEN avl_a
        WHEN (character_id IS NOT NULL) THEN avl_c
        WHEN (empty_slots_amount IS NOT NULL) THEN avl_s
        ELSE avl_self
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION select_loc_avl(
    w_l_ctx_type TEXT,
    avl_loc availability_type,
    avl_subloc availability_type,
    avl_area availability_type
)
RETURNS availability_type AS $$
    SELECT CASE
        WHEN w_l_ctx_type = 'location' THEN avl_loc
        WHEN w_l_ctx_type = 'sublocation' THEN avl_subloc
        WHEN w_l_ctx_type = 'area' THEN avl_area
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank(
    avl availability_type,
    w_pre_airship BOOLEAN
) 
RETURNS INT AS $$
    SELECT CASE 
        WHEN avl = 'always' THEN 1

        WHEN NOT w_pre_airship AND avl = 'post' THEN 2
        WHEN NOT w_pre_airship AND avl = 'pre-story' THEN 3

        WHEN w_pre_airship AND avl = 'pre-story' THEN 2
        WHEN w_pre_airship AND avl = 'post' THEN 3

        WHEN avl = 'post-story' THEN 4
    END;
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank_item(
    source_type TEXT,
    w_avl_type TEXT,
    w_l_ctx_type TEXT,
    w_pre_airship BOOLEAN,
    avl_self availability_type,
    avl_context availability_type,
    avl_context_2 availability_type,
    avl_shop_loc availability_type,
    avl_shop_subloc availability_type,
    avl_shop_area availability_type
)
RETURNS INT AS $$
    SELECT get_avl_rank(
        CASE
            WHEN source_type = 'shop' AND w_avl_type = 'self' THEN avl_context
            WHEN source_type = 'shop' AND w_l_ctx_type IS NOT NULL THEN select_loc_avl(w_l_ctx_type, avl_shop_loc, avl_shop_subloc, avl_shop_area)
            ELSE select_avl(w_avl_type, avl_self, avl_context, avl_context_2)
        END,
        w_pre_airship
    )
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank_auto_ability(
    source_type TEXT,
    w_avl_type TEXT,
    w_l_ctx_type TEXT,
    w_char_id INTEGER,
    w_pre_airship BOOLEAN,
    avl_self availability_type,
    avl_context availability_type,
    avl_context_2 availability_type,
    avl_shop_loc availability_type,
    avl_shop_subloc availability_type,
    avl_shop_area availability_type,
    avl_shop_loc_char availability_type,
    avl_shop_subloc_char availability_type,
    avl_shop_area_char availability_type
)
RETURNS INT AS $$
    SELECT get_avl_rank(
        CASE
            WHEN source_type = 'shop' AND w_avl_type = 'self' THEN avl_context
            WHEN source_type = 'shop' AND w_l_ctx_type IS NOT NULL THEN 
                CASE
                    WHEN w_char_id IS NULL THEN select_loc_avl(w_l_ctx_type, avl_shop_loc, avl_shop_subloc, avl_shop_area)
                    ELSE select_loc_avl(w_l_ctx_type, avl_shop_loc_char, avl_shop_subloc_char, avl_shop_area_char)
                END
            ELSE select_avl(w_avl_type, avl_self, avl_context, avl_context_2)
        END,
        w_pre_airship
    )
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank_condition(
    avl_true availability_type,
    avl_false availability_type,
    w_pre_airship BOOLEAN,
    condition BOOLEAN
)
RETURNS INT AS $$
    SELECT get_avl_rank(
        CASE
            WHEN condition THEN avl_true
            ELSE avl_false
        END,
        w_pre_airship
    )
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_is_rep(
    w_loc_context_id INTEGER,
    is_repeatable BOOLEAN,
    is_repeatable_loc BOOLEAN
)
RETURNS BOOLEAN AS $$
    SELECT CASE
        WHEN w_loc_context_id IS NOT NULL THEN is_repeatable_loc
        ELSE is_repeatable
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION filter_avl_rep(
    avl INTEGER,                    -- MIN(current_avl)
    is_rep BOOLEAN,                 -- BOOL_OR(is_rep = w.is_repeatable)
    avl_matching_rep INTEGER,       -- MIN(current_avl) FILTER (WHERE is_rep = w.is_repeatable)
    is_rep_in_avls BOOLEAN,    -- BOOL_OR(is_rep) FILTER WHERE current_avl = ANY(w.availability)
    avl_not_rep INTEGER,            -- MIN(current_avl) FILTER (WHERE is_rep = FALSE)
    w_avls INTEGER[],               -- w.availability
    w_is_rep BOOLEAN)               -- w.is_repeatable
RETURNS BOOLEAN AS $$
    SELECT CASE
        WHEN w_is_rep IS NULL THEN avl = ANY(w_avls)
        WHEN w_avls IS NULL THEN is_rep
        ELSE avl_matching_rep = ANY(w_avls) AND is_rep_in_avls = w_is_rep
    END
    AND (w_is_rep IS NOT FALSE OR avl_not_rep = avl)
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose Down
DROP FUNCTION IF EXISTS filter_avl_rep(INTEGER, BOOLEAN, INTEGER, BOOLEAN, INTEGER, INTEGER[], BOOLEAN);
DROP FUNCTION IF EXISTS get_is_rep(INTEGER, BOOLEAN, BOOLEAN);
DROP FUNCTION IF EXISTS get_avl_rank_condition(availability_type, availability_type, BOOLEAN, BOOLEAN);
DROP FUNCTION IF EXISTS get_avl_rank_auto_ability(TEXT, TEXT, TEXT, INTEGER, BOOLEAN, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS get_avl_rank_item(TEXT, TEXT, TEXT, BOOLEAN, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS get_avl_rank(availability_type, BOOLEAN);
DROP FUNCTION IF EXISTS select_loc_avl(TEXT, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS select_shop_equip_avl(INTEGER, INTEGER, INTEGER[], availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS select_avl(TEXT, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS get_loc_ctx_id(TEXT, INTEGER, INTEGER, INTEGER);
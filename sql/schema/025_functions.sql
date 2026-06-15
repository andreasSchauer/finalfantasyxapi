-- +goose Up


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_loc_ctx_id(loc_context_type TEXT, location_id INTEGER, sublocation_id INTEGER, area_id INTEGER)
RETURNS INT AS $$
    SELECT CASE
        WHEN loc_context_type = 'location' THEN location_id
        WHEN loc_context_type = 'sublocation' THEN sublocation_id
        WHEN loc_context_type = 'area' THEN area_id
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION select_avl(avl_type TEXT, avl_self availability_type, avl_context availability_type, avl_context_2 availability_type)
RETURNS availability_type AS $$
    SELECT CASE
        WHEN avl_type = 'self' THEN avl_self
        WHEN avl_type = 'context' THEN avl_context
        WHEN avl_type = 'context-2' THEN avl_context_2
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION select_loc_avl(l_ctx_type TEXT, avl_loc availability_type, avl_subloc availability_type, avl_area availability_type)
RETURNS availability_type AS $$
    SELECT CASE
        WHEN l_ctx_type = 'location' THEN avl_loc
        WHEN l_ctx_type = 'sublocation' THEN avl_subloc
        WHEN l_ctx_type = 'area' THEN avl_area
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
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
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank_item(
    avl_type TEXT,
    l_ctx_type TEXT,
    source_type TEXT,
    pre_airship BOOLEAN,
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
            WHEN source_type = 'shop' AND avl_type = 'self' THEN avl_context
            WHEN source_type = 'shop' AND l_ctx_type IS NOT NULL THEN select_loc_avl(l_ctx_type, avl_shop_loc, avl_shop_subloc, avl_shop_area)
            ELSE select_avl(avl_type, avl_self, avl_context, avl_context_2)
        END,
        pre_airship
    )
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank_auto_ability(
    avl_type TEXT,
    l_ctx_type TEXT,
    source_type TEXT,
    char_id INTEGER,
    pre_airship BOOLEAN,
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
            WHEN source_type = 'shop' AND avl_type = 'self' THEN avl_context
            WHEN source_type = 'shop' AND l_ctx_type IS NOT NULL THEN 
                CASE
                    WHEN char_id IS NULL THEN select_loc_avl(l_ctx_type, avl_shop_loc, avl_shop_subloc, avl_shop_area)
                    ELSE select_loc_avl(l_ctx_type, avl_shop_loc_char, avl_shop_subloc_char, avl_shop_area_char)
                END
            ELSE select_avl(avl_type, avl_self, avl_context, avl_context_2)
        END,
        pre_airship
    )
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_avl_rank_condition(avl_true availability_type, avl_false availability_type, pre_airship BOOLEAN, condition BOOLEAN)
RETURNS INT AS $$
    SELECT get_avl_rank(
        CASE
            WHEN condition THEN avl_true
            ELSE avl_false
        END,
        pre_airship
    )
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_is_rep(loc_context_id INTEGER, is_repeatable BOOLEAN, is_repeatable_loc BOOLEAN)
RETURNS BOOLEAN AS $$
    SELECT CASE
        WHEN loc_context_id IS NOT NULL THEN is_repeatable_loc
        ELSE is_repeatable
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION filter_avl_rep(avl INTEGER, is_rep BOOLEAN, avl_filtered INTEGER, avl_input INTEGER[], is_rep_input BOOLEAN)
RETURNS BOOLEAN AS $$
    SELECT CASE
        WHEN is_rep_input IS NULL THEN avl = ANY(avl_input)
        WHEN avl_input IS NULL THEN is_rep
        ELSE avl_filtered = ANY(avl_input)
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose Down
DROP FUNCTION IF EXISTS filter_avl_rep(INTEGER, BOOLEAN, INTEGER, INTEGER[], BOOLEAN);
DROP FUNCTION IF EXISTS get_is_rep(INTEGER, BOOLEAN, BOOLEAN);
DROP FUNCTION IF EXISTS get_avl_rank_condition(availability_type, availability_type, BOOLEAN, BOOLEAN);
DROP FUNCTION IF EXISTS get_avl_rank_auto_ability(TEXT, TEXT, TEXT, INTEGER, BOOLEAN, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS get_avl_rank_item(TEXT, TEXT, TEXT, BOOLEAN, availability_type, availability_type, availability_type, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS get_avl_rank(availability_type, BOOLEAN);
DROP FUNCTION IF EXISTS select_loc_avl(TEXT, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS select_avl(TEXT, availability_type, availability_type, availability_type);
DROP FUNCTION IF EXISTS get_loc_ctx_id(TEXT, INTEGER, INTEGER, INTEGER);
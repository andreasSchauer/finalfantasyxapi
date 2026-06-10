-- +goose Up

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
CREATE OR REPLACE FUNCTION get_avl_rank_item(avl_type TEXT, avl_self availability_type, avl_context availability_type, avl_area availability_type, pre_airship BOOLEAN, is_shop BOOLEAN)
RETURNS INT AS $$
    SELECT get_avl_rank(
        CASE
            WHEN is_shop AND avl_type = 'self' THEN avl_context
            WHEN avl_type = 'self' THEN avl_self
            WHEN avl_type = 'context' THEN avl_context
            WHEN avl_type = 'area' THEN avl_area
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
CREATE OR REPLACE FUNCTION get_loc_ctx_id(loc_context_type TEXT, location_id INTEGER, sublocation_id INTEGER, area_id INTEGER)
RETURNS INT AS $$
    SELECT CASE
        WHEN loc_context_type = 'location' THEN location_id
        WHEN loc_context_type = 'sublocation' THEN sublocation_id
        WHEN loc_context_type = 'area' THEN area_id
    END
$$ LANGUAGE sql IMMUTABLE;
-- +goose StatementEnd


-- +goose Down
DROP FUNCTION IF EXISTS get_loc_ctx_id(TEXT, INTEGER, INTEGER, INTEGER);
DROP FUNCTION IF EXISTS get_is_rep(INTEGER, BOOLEAN, BOOLEAN);
DROP FUNCTION IF EXISTS get_avl_rank_item(TEXT, availability_type, availability_type, availability_type, BOOLEAN, BOOLEAN);
DROP FUNCTION IF EXISTS get_avl_rank(availability_type, BOOLEAN);
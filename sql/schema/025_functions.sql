-- +goose Up
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



-- +goose Down
DROP FUNCTION get_avl_rank;
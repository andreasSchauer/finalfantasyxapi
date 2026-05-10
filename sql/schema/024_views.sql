-- +goose Up
CREATE MATERIALIZED VIEW mv_monster_drops AS
SELECT mi.monster_id, ia.master_item_id
FROM monster_items mi
JOIN item_amounts ia ON ia.id IN (
    mi.steal_common_id,
    mi.steal_rare_id,
    mi.drop_common_id,
    mi.drop_rare_id,
    mi.secondary_drop_common_id,
    mi.secondary_drop_rare_id,
    mi.bribe_id
)

UNION ALL

SELECT mi.monster_id, ia.master_item_id
FROM monster_items mi
JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
JOIN possible_items pi ON pi.id = jmio.possible_item_id
JOIN item_amounts ia ON pi.item_amount_id = ia.id;


CREATE INDEX idx_mv_monster_drops_item ON mv_monster_drops (master_item_id);
CREATE INDEX idx_mv_monster_drops_monster ON mv_monster_drops (monster_id);



-- +goose Down
DROP MATERIALIZED VIEW monster_drops;
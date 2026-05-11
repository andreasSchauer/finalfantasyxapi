-- +goose Up
CREATE MATERIALIZED VIEW mv_monster_item_drops AS
SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'steal_common' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.steal_common_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'steal_rare' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.steal_rare_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'drop_common' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.drop_common_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'drop_rare' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.drop_rare_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'drop_secondary_common' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.secondary_drop_common_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'drop_secondary_rare' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.secondary_drop_rare_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'bribe' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.bribe_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id

UNION ALL

SELECT
    mi.monster_id,
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'other' AS source_type,
    m.availability
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
JOIN possible_items pi ON pi.id = jmio.possible_item_id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id;


CREATE INDEX idx_mv_monster_item_drops_item ON mv_monster_item_drops (item_id);
CREATE INDEX idx_mv_monster_item_drops_master_item ON mv_monster_item_drops (master_item_id);
CREATE INDEX idx_mv_monster_item_drops_monster ON mv_monster_item_drops (monster_id);




CREATE MATERIALIZED VIEW mv_monster_encounters AS
SELECT DISTINCT
    ea.area_id,
    mf.id AS formation_id,
    ma.monster_id,
    ma.amount AS monster_amount
FROM encounter_areas ea
JOIN j_monster_formations_encounter_areas jmfea ON ea.id = jmfea.encounter_area_id
JOIN monster_formations mf ON jmfea.monster_formation_id = mf.id
JOIN j_monster_selections_monsters jmsm ON mf.monster_selection_id = jmsm.monster_selection_id
JOIN monster_amounts ma ON jmsm.monster_amount_id = ma.id;

CREATE INDEX idx_mv_enc_area ON mv_monster_encounters (area_id);
CREATE INDEX idx_mv_enc_monster ON mv_monster_encounters (monster_id);
CREATE INDEX idx_mv_enc_formation ON mv_monster_encounters (formation_id);




CREATE MATERIALIZED VIEW mv_geography AS
SELECT 
    a.id AS area_id,
    s.id AS sublocation_id,
    l.id AS location_id,
    a.name AS area_name,
    s.name AS sublocation_name,
    l.name AS location_name
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id;

CREATE INDEX idx_mv_geo_area ON mv_geography (area_id);
CREATE INDEX idx_mv_geo_sublocation ON mv_geography (sublocation_id);
CREATE INDEX idx_mv_geo_location ON mv_geography (location_id);




CREATE MATERIALIZED VIEW mv_item_sources AS
SELECT
    ia.master_item_id,
    t.id AS source_id,
    ia.amount,
    'treasures' AS source_type,
    t.availability
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id

UNION ALL

-- Shops
SELECT
    i.master_item_id,
    j.shop_id AS source_id,
    1 AS amount,
    'shops' AS source_type,
    CASE j.shop_type 
        WHEN 'pre-airship' THEN 'pre-story'::availability_type
        WHEN 'post-airship' THEN 'post'::availability_type
    END AS availability
FROM j_shops_items j
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id

UNION ALL

-- Quests
SELECT
    ia.master_item_id,
    q.id AS source_id,
    ia.amount,
    'quests' AS source_type,
    q.availability
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id

UNION ALL

SELECT
    ia.master_item_id,
    bi.position_id AS source_id,
    ia.amount,
    'blitzball' AS source_type,
    'always' AS availability
FROM blitzball_items bi
JOIN possible_items pi ON bi.possible_item_id = pi.id
JOIN item_amounts ia ON pi.item_amount_id = ia.id

UNION ALL

SELECT
    master_item_id,
    monster_id AS source_id,
    amount,
    'monsters' AS source_type,
    availability
FROM mv_monster_item_drops;


CREATE INDEX idx_mv_item_sources_master_id ON mv_item_sources (master_item_id);



-- +goose Down
DROP MATERIALIZED VIEW IF EXISTS mv_item_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_geography;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_encounters;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_item_drops;
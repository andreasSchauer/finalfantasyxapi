-- +goose Up
CREATE MATERIALIZED VIEW mv_monster_item_drops AS
SELECT 
    mi.monster_id, 
    i.id AS item_id,
    ia.master_item_id,
    ia.amount,
    'steal_common' AS source_type,
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
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
    m.availability,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
JOIN possible_items pi ON pi.id = jmio.possible_item_id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN items i ON ia.master_item_id = i.master_item_id;


CREATE INDEX idx_mv_monster_item_drops_item ON mv_monster_item_drops (item_id);
CREATE INDEX idx_mv_monster_item_drops_master_item ON mv_monster_item_drops (master_item_id);
CREATE INDEX idx_mv_monster_item_drops_monster ON mv_monster_item_drops (monster_id);



CREATE MATERIALIZED VIEW mv_monster_equipment_drops AS
SELECT DISTINCT
    me.monster_id,
    asl.id AS ability_slots_id,
    aa.id AS attached_abilities_id,
    ed.auto_ability_id,
    m.availability,
    m.is_repeatable,
    j.equipment_drop_id,
    ed.is_forced,
    ed.probability,
    ed.type AS auto_ability_type
FROM monsters m
JOIN monster_equipment me ON me.monster_id = m.id
JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
JOIN monster_equipment_slots asl ON asl.monster_equipment_id = me.id AND asl.type = 'ability-slots'
JOIN monster_equipment_slots aa ON aa.monster_equipment_id = me.id AND aa.type = 'attached-abilities';

CREATE INDEX idx_mv_monster_equipment_drops_monster ON mv_monster_equipment_drops(monster_id);
CREATE INDEX idx_mv_monster_equipment_drops_auto_ability ON mv_monster_equipment_drops(auto_ability_id);
CREATE INDEX idx_mv_monster_equipment_drops_ability_slots ON mv_monster_equipment_drops(ability_slots_id);
CREATE INDEX idx_mv_monster_equipment_drops_attached_abilities ON mv_monster_equipment_drops(attached_abilities_id);
CREATE INDEX idx_mv_monster_equipment_drops_equipment_drops ON mv_monster_equipment_drops(equipment_drop_id);






CREATE MATERIALIZED VIEW mv_monster_encounters AS
SELECT DISTINCT
    mf.id AS formation_id,
    ma.monster_id,
    ma.amount AS monster_amount,
    fd.id AS data_id,
    fd.category,
    fd.availability,
    fd.is_forced_ambush,
    fd.can_escape,
    ea.area_id,
    fbs.song_id,
    jmft.trigger_command_id AS f_trigger_command_id
FROM encounter_areas ea
JOIN j_monster_formations_encounter_areas jmfea ON ea.id = jmfea.encounter_area_id
JOIN monster_formations mf ON jmfea.monster_formation_id = mf.id
JOIN j_monster_selections_monsters jmsm ON mf.monster_selection_id = jmsm.monster_selection_id
JOIN monster_amounts ma ON jmsm.monster_amount_id = ma.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
LEFT JOIN formation_boss_songs fbs ON fd.boss_song_id = fbs.id
LEFT JOIN j_monster_formations_trigger_commands jmft ON jmft.monster_formation_id = mf.id;

CREATE INDEX idx_mv_enc_area ON mv_monster_encounters (area_id);
CREATE INDEX idx_mv_enc_monster ON mv_monster_encounters (monster_id);
CREATE INDEX idx_mv_enc_formation ON mv_monster_encounters (formation_id);
CREATE INDEX idx_mv_enc_song ON mv_monster_encounters (song_id);




CREATE MATERIALIZED VIEW mv_geography AS
SELECT 
    l.id AS location_id,
    s.id AS sublocation_id,
    a.id AS area_id,
    l.name AS location_name,
    s.name AS sublocation_name,
    a.name AS area_name,
    a.version AS area_version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id;

CREATE INDEX idx_mv_geo_area ON mv_geography (area_id);
CREATE INDEX idx_mv_geo_sublocation ON mv_geography (sublocation_id);
CREATE INDEX idx_mv_geo_location ON mv_geography (location_id);


CREATE MATERIALIZED VIEW mv_geography_graph AS
SELECT
    g.location_id AS l_id,
    g.sublocation_id AS s_id,
    g.area_id AS a_id,
    g.location_name AS l_name,
    g.sublocation_name AS s_name,
    g.area_name AS a_name,
    g.area_version AS a_version,
    cg.location_id AS cl_id,
    cg.sublocation_id AS cs_id,
    cg.area_id AS ca_id,
    cg.location_name AS cl_name,
    cg.sublocation_name AS cs_name,
    cg.area_name AS ca_name,
    cg.area_version AS ca_version,
    ac.connection_type,
    ac.is_story_based
FROM mv_geography g
JOIN j_area_connected_areas j ON j.area_id = g.area_id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN mv_geography cg ON ac.area_id = cg.area_id;

CREATE INDEX idx_mv_geo_graph_area ON mv_geography_graph (a_id);
CREATE INDEX idx_mv_geo_graph_sublocation ON mv_geography_graph (s_id);
CREATE INDEX idx_mv_geo_graph_location ON mv_geography_graph (l_id);




CREATE MATERIALIZED VIEW mv_item_sources AS
SELECT
    ia.master_item_id,
    t.id AS source_id,
    t.area_id,
    ia.amount,
    'treasure' AS source_type,
    t.availability
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id

UNION ALL

-- Shops
SELECT
    i.master_item_id,
    j.shop_id AS source_id,
    sh.area_id,
    1 AS amount,
    'shop' AS source_type,
    CASE j.shop_type 
        WHEN 'pre-airship' THEN 'pre-story'::availability_type
        WHEN 'post-airship' THEN 'post'::availability_type
    END AS availability
FROM shops sh
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id

UNION ALL

-- Quests
SELECT
    ia.master_item_id,
    q.id AS source_id,
    ca.area_id,
    ia.amount,
    'quest' AS source_type,
    q.availability
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN completion_areas ca ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id

UNION ALL

SELECT
    ia.master_item_id,
    bi.position_id AS source_id,
    75 AS area_id,
    ia.amount,
    'blitzball' AS source_type,
    'always' AS availability
FROM blitzball_items bi
JOIN possible_items pi ON bi.possible_item_id = pi.id
JOIN item_amounts ia ON pi.item_amount_id = ia.id

UNION ALL

SELECT
    mi.master_item_id,
    mi.monster_id AS source_id,
    me.area_id,
    mi.amount,
    'monster' AS source_type,
    mi.availability
FROM mv_monster_item_drops mi
JOIN mv_monster_encounters me ON mi.monster_id = me.monster_id;


CREATE INDEX idx_mv_item_sources_master_id ON mv_item_sources (master_item_id);
CREATE INDEX idx_mv_item_sources_source_id ON mv_item_sources (source_id);
CREATE INDEX idx_mv_item_sources_area_id ON mv_item_sources (area_id);



-- +goose Down
DROP MATERIALIZED VIEW IF EXISTS mv_item_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_geography_graph;
DROP MATERIALIZED VIEW IF EXISTS mv_geography;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_encounters;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_equipment_drops;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_item_drops;
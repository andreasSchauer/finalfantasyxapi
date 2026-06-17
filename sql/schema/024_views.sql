-- +goose Up
CREATE MATERIALIZED VIEW mv_geography AS
SELECT DISTINCT
    l.id AS location_id,
    s.id AS sublocation_id,
    a.id AS area_id,
    l.name AS location,
    s.name AS sublocation,
    a.name AS area,
    a.version
FROM areas a
JOIN sublocations s ON a.sublocation_id = s.id
JOIN locations l ON s.location_id = l.id;

CREATE INDEX idx_mv_geo_area ON mv_geography (area_id);
CREATE INDEX idx_mv_geo_sublocation ON mv_geography (sublocation_id);
CREATE INDEX idx_mv_geo_location ON mv_geography (location_id);




CREATE MATERIALIZED VIEW mv_geography_graph AS
SELECT DISTINCT
    g.location_id AS l_id,
    g.sublocation_id AS s_id,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS a_version,
    cg.location_id AS cl_id,
    cg.sublocation_id AS cs_id,
    cg.area_id AS ca_id,
    cg.location AS c_location,
    cg.sublocation AS c_sublocation,
    cg.area AS c_area,
    cg.version AS ca_version,
    ac.connection_type,
    ac.is_story_based
FROM mv_geography g
JOIN j_area_connected_areas j ON j.area_id = g.area_id
JOIN area_connections ac ON j.connection_id = ac.id
JOIN mv_geography cg ON ac.area_id = cg.area_id;

CREATE INDEX idx_mv_geo_graph_area ON mv_geography_graph (a_id);
CREATE INDEX idx_mv_geo_graph_sublocation ON mv_geography_graph (s_id);
CREATE INDEX idx_mv_geo_graph_location ON mv_geography_graph (l_id);






CREATE MATERIALIZED VIEW mv_monster_encounters AS
SELECT DISTINCT
    mf.id AS formation_id,
    ma.monster_id,
    m.name AS monster,
    m.version,
    ma.amount AS monster_amount,
    fd.id AS data_id,
    fd.category,
    fd.is_forced_ambush,
    fd.can_escape,
    m.availability AS avl_self,
    fd.availability AS avl_context,
    ea.availability AS avl_context_2,
    m.is_repeatable,
    CASE 
        WHEN m.id IN (299, 300) THEN TRUE -- dark yojimbo
        WHEN fd.category IN (
            'random-encounter'::monster_formation_category, 
            'on-demand-fight'::monster_formation_category
        ) THEN TRUE
        ELSE FALSE
    END AS is_repeatable_loc,
    g.area_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS a_version,
    fbs.song_id,
    jmft.trigger_command_id AS f_trigger_command_id
FROM mv_geography g
JOIN encounter_areas ea ON g.area_id = ea.area_id
JOIN j_monster_formations_encounter_areas jmfea ON ea.id = jmfea.encounter_area_id
JOIN monster_formations mf ON jmfea.monster_formation_id = mf.id
JOIN j_monster_selections_monsters jmsm ON mf.monster_selection_id = jmsm.monster_selection_id
JOIN monster_amounts ma ON jmsm.monster_amount_id = ma.id
JOIN monsters m ON ma.monster_id = m.id
JOIN formation_data fd ON mf.formation_data_id = fd.id
LEFT JOIN formation_boss_songs fbs ON fd.boss_song_id = fbs.id
LEFT JOIN j_monster_formations_trigger_commands jmft ON jmft.monster_formation_id = mf.id;

CREATE INDEX idx_mv_enc_area ON mv_monster_encounters (area_id);
CREATE INDEX idx_mv_enc_monster ON mv_monster_encounters (monster_id);
CREATE INDEX idx_mv_enc_formation ON mv_monster_encounters (formation_id);
CREATE INDEX idx_mv_enc_song ON mv_monster_encounters (song_id);



CREATE MATERIALIZED VIEW mv_monster_item_drops AS
SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'steal_common' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.steal_common_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'steal_rare' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.steal_rare_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'drop_common' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.drop_common_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'drop_rare' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.drop_rare_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'drop_secondary_common' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.secondary_drop_common_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'drop_secondary_rare' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.secondary_drop_rare_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id, 
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'bribe' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN item_amounts ia ON mi.bribe_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id

UNION ALL

SELECT DISTINCT
    mi.monster_id,
    m.name AS monster,
    m.version,
    i.id AS item_id,
    ia.master_item_id,
    mit.name AS item,
    ia.amount,
    'other' AS source_type,
    m.is_repeatable
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
JOIN j_monster_items_other_items jmio ON jmio.monster_items_id = mi.id
JOIN possible_items pi ON pi.id = jmio.possible_item_id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON mit.id = i.id;


CREATE INDEX idx_mv_monster_item_drops_item ON mv_monster_item_drops (item_id);
CREATE INDEX idx_mv_monster_item_drops_master_item ON mv_monster_item_drops (master_item_id);
CREATE INDEX idx_mv_monster_item_drops_monster ON mv_monster_item_drops (monster_id);



CREATE MATERIALIZED VIEW mv_monster_equipment_drops AS
SELECT DISTINCT
    me.monster_id,
    m.name AS monster,
    m.version,
    asl.id AS ability_slots_id,
    aa.id AS attached_abilities_id,
    ed.auto_ability_id,
    au.name AS auto_ability,
    m.is_repeatable,
    j.equipment_drop_id,
    ed.is_forced,
    ed.probability,
    ed.type AS auto_ability_type,
    jedc.character_id
FROM monsters m
JOIN monster_equipment me ON me.monster_id = m.id
JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
LEFT JOIN j_equipment_drops_characters jedc ON jedc.equipment_drop_id = ed.id -- new join
JOIN auto_abilities au ON ed.auto_ability_id = au.id
JOIN monster_equipment_slots asl ON asl.monster_equipment_id = me.id AND asl.type = 'ability-slots'
JOIN monster_equipment_slots aa ON aa.monster_equipment_id = me.id AND aa.type = 'attached-abilities';

CREATE INDEX idx_mv_monster_equipment_drops_monster ON mv_monster_equipment_drops(monster_id);
CREATE INDEX idx_mv_monster_equipment_drops_auto_ability ON mv_monster_equipment_drops(auto_ability_id);
CREATE INDEX idx_mv_monster_equipment_drops_ability_slots ON mv_monster_equipment_drops(ability_slots_id);
CREATE INDEX idx_mv_monster_equipment_drops_attached_abilities ON mv_monster_equipment_drops(attached_abilities_id);
CREATE INDEX idx_mv_monster_equipment_drops_equipment_drops ON mv_monster_equipment_drops(equipment_drop_id);





CREATE MATERIALIZED VIEW mv_item_sources AS
SELECT DISTINCT
    ia.master_item_id,
    mi.name AS item,
    t.id AS source_id,
    'treasure' AS source_type,
    ia.amount,
    t.availability AS avl_self,
    t.availability AS avl_context,
    t.availability AS avl_context_2,
    NULL::availability_type AS avl_shop_area,
    NULL::availability_type AS avl_shop_subloc,
    NULL::availability_type AS avl_shop_loc,
    FALSE AS is_repeatable,
    FALSE AS is_repeatable_loc,
    t.area_id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id

UNION ALL

SELECT DISTINCT
    i.master_item_id,
    mi.name AS item,
    j.shop_id AS source_id,
    'shop' AS source_type,
    1 AS amount,
    sh.availability AS avl_self,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY i.master_item_id, sh.id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY i.master_item_id, sh.id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY i.master_item_id, sh.id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY i.master_item_id, sh.id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context_2,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY i.master_item_id, g.area_id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY i.master_item_id, g.area_id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_area,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY i.master_item_id, g.sublocation_id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY i.master_item_id, g.sublocation_id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_subloc,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY i.master_item_id, g.location_id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY i.master_item_id, g.location_id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_loc,
    TRUE AS is_repeatable,
    TRUE AS is_repeatable_loc,
    sh.area_id
FROM shops sh
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
JOIN master_items mi ON i.master_item_id = mi.id
JOIN mv_geography g ON sh.area_id = g.area_id

UNION ALL

SELECT DISTINCT
    ia.master_item_id,
    mi.name AS item,
    q.id AS source_id,
    'quest' AS source_type,
    ia.amount,
    q.availability AS avl_self,
    q.availability AS avl_context,
    q.availability AS avl_context_2,
    NULL::availability_type AS avl_shop_area,
    NULL::availability_type AS avl_shop_subloc,
    NULL::availability_type AS avl_shop_loc,
    q.is_repeatable,
    q.is_repeatable AS is_repeatable_loc,
    ca.area_id
FROM quests q
JOIN quest_completions qc ON qc.quest_id = q.id
JOIN completion_areas ca ON ca.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id

UNION ALL

SELECT DISTINCT
    ia.master_item_id,
    mi.name AS item,
    bi.position_id AS source_id,
    'blitzball' AS source_type,
    ia.amount,
    'always'::availability_type AS avl_self,
    'always'::availability_type AS avl_context,
    'always'::availability_type AS avl_context_2,
    NULL::availability_type AS avl_shop_area,
    NULL::availability_type AS avl_shop_subloc,
    NULL::availability_type AS avl_shop_loc,
    TRUE AS is_repeatable,
    TRUE AS is_repeatable_loc,
    75 AS area_id
FROM blitzball_items bi
JOIN possible_items pi ON bi.possible_item_id = pi.id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id

UNION ALL

SELECT DISTINCT
    mi.master_item_id,
    mi.item,
    mi.monster_id AS source_id,
    'monster' AS source_type,
    mi.amount,
    me.avl_self,
    me.avl_context,
    me.avl_context_2,
    NULL::availability_type AS avl_shop_area,
    NULL::availability_type AS avl_shop_subloc,
    NULL::availability_type AS avl_shop_loc,
    CASE
        WHEN me.category = 'static-encounter'::monster_formation_category
         AND mi.source_type IN ('steal_common', 'steal_rare')
        THEN TRUE

        ELSE mi.is_repeatable
    END AS is_repeatable,
    CASE 
        WHEN me.monster_id = 299 THEN TRUE -- dark yojimbo
        
        WHEN me.category IN (
            'random-encounter'::monster_formation_category, 
            'on-demand-fight'::monster_formation_category
        ) THEN TRUE

        WHEN me.category = 'static-encounter'::monster_formation_category
         AND mi.source_type IN ('steal_common', 'steal_rare')
        THEN TRUE

        ELSE FALSE
    END AS is_repeatable_loc,
    me.area_id
FROM mv_monster_item_drops mi
JOIN mv_monster_encounters me ON mi.monster_id = me.monster_id;


CREATE INDEX idx_mv_item_sources_master_id ON mv_item_sources (master_item_id);
CREATE INDEX idx_mv_item_sources_lookup ON mv_item_sources (source_id, source_type, area_id);
CREATE INDEX idx_mv_item_sources_area_id ON mv_item_sources (area_id);





CREATE MATERIALIZED VIEW mv_equipment_sources AS
WITH shop_equip_configs AS (
    SELECT
        en.id AS name_id,
        se.id AS shop_equipment_id,
        sh.id AS shop_id,
        se.shop_type,
        en.character_id,
        se.empty_slots_amount,
        ARRAY_AGG(j.auto_ability_id ORDER BY j.auto_ability_id) AS ability_set
    FROM shop_equipment_pieces se
    JOIN shops sh ON se.shop_id = sh.id
    JOIN equipment_names en ON se.equipment_name_id = en.id
    LEFT JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
    GROUP BY en.id, se.id, sh.id, se.shop_type, en.character_id, se.empty_slots_amount
),
shop_equip_avl AS (
    SELECT
        shop_equipment_id,
        CASE
            WHEN BOOL_OR(shop_type = 'pre-airship') OVER (PARTITION BY shop_id, name_id, character_id, empty_slots_amount, ability_set) 
             AND BOOL_OR(shop_type = 'post-airship') OVER (PARTITION BY shop_id, name_id, character_id, empty_slots_amount, ability_set) 
            THEN 'always'::availability_type
            WHEN shop_type = 'pre-airship' THEN 'pre-story'::availability_type
            WHEN shop_type = 'post-airship' THEN 'post'::availability_type
        END AS availability
    FROM shop_equip_configs
)
SELECT DISTINCT
    en.id AS name_id,
    en.name AS name,
    en.character_id,
    t.id AS source_id,
    'treasure' AS source_type,
    te.empty_slots_amount,
    j.auto_ability_id,
    aa.name AS auto_ability,
    t.availability AS avl_self,
    t.availability AS avl_context,
    t.availability AS avl_context_2,
    FALSE AS is_repeatable,
    FALSE AS is_repeatable_loc,
    t.area_id,
    NULL::shop_type AS shop_type
FROM treasures t
JOIN treasure_equipment_pieces te ON te.treasure_id = t.id
LEFT JOIN j_treasure_equipment_abilities j ON j.treasure_equipment_id = te.id
LEFT JOIN auto_abilities aa ON j.auto_ability_id = aa.id
JOIN equipment_names en ON te.equipment_name_id = en.id

UNION ALL

SELECT DISTINCT
    en.id AS name_id,
    en.name AS name,
    en.character_id,
    sh.id AS source_id,
    'shop' AS source_type,
    se.empty_slots_amount,
    j.auto_ability_id,
    aa.name AS auto_ability,
    sh.availability AS avl_self,
    sa.availability AS avl_context,
    sa.availability AS avl_context_2,
    TRUE AS is_repeatable,
    TRUE AS is_repeatable_loc,
    sh.area_id,
    se.shop_type AS shop_type
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN shop_equip_avl sa ON sa.shop_equipment_id = se.id
LEFT JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
LEFT JOIN auto_abilities aa ON j.auto_ability_id = aa.id
JOIN equipment_names en ON se.equipment_name_id = en.id;


CREATE INDEX idx_mv_equipment_sources_name_id ON mv_equipment_sources (name_id);
CREATE INDEX idx_mv_equipment_sources_lookup ON mv_equipment_sources (source_id, source_type, area_id);
CREATE INDEX idx_mv_equipment_sources_area_id ON mv_equipment_sources (area_id);




CREATE MATERIALIZED VIEW mv_auto_ability_sources AS
SELECT DISTINCT
    aa.id AS auto_ability_id,
    aa.name AS auto_ability,
    t.id AS source_id,
    'treasure' AS source_type,
    en.character_id,
    t.availability AS avl_self,
    t.availability AS avl_context,
    t.availability AS avl_context_2,
    NULL::availability_type AS avl_shop_area,
    NULL::availability_type AS avl_shop_subloc,
    NULL::availability_type AS avl_shop_loc,
    NULL::availability_type AS avl_shop_area_char,
    NULL::availability_type AS avl_shop_subloc_char,
    NULL::availability_type AS avl_shop_loc_char,
    FALSE AS is_repeatable,
    FALSE AS is_repeatable_loc,
    t.area_id,
    NULL::shop_type AS shop_type
FROM treasures t
JOIN treasure_equipment_pieces te ON te.treasure_id = t.id
JOIN j_treasure_equipment_abilities j ON j.treasure_equipment_id = te.id
JOIN auto_abilities aa ON j.auto_ability_id = aa.id
JOIN equipment_names en ON te.equipment_name_id = en.id

UNION ALL

SELECT DISTINCT
    aa.id AS auto_ability_id,
    aa.name AS auto_ability,
    sh.id AS source_id,
    'shop' AS source_type,
    en.character_id,
    sh.availability AS avl_self,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, sh.id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, sh.id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, sh.id, en.character_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, sh.id, en.character_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context_2,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, g.area_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, g.area_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_area,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, g.sublocation_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, g.sublocation_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_subloc,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, g.location_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, g.location_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_loc,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, en.character_id, g.area_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, en.character_id, g.area_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_area_char,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, en.character_id, g.sublocation_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, en.character_id, g.sublocation_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_subloc_char,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY aa.id, en.character_id, g.location_id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY aa.id, en.character_id, g.location_id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_shop_loc_char,
    TRUE AS is_repeatable,
    TRUE AS is_repeatable_loc,
    sh.area_id,
    se.shop_type AS shop_type
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
JOIN auto_abilities aa ON j.auto_ability_id = aa.id
JOIN equipment_names en ON se.equipment_name_id = en.id
JOIN mv_geography g ON sh.area_id = g.area_id

UNION ALL

SELECT DISTINCT
    med.auto_ability_id,
    med.auto_ability,
    med.monster_id AS source_id,
    'monster' AS source_type,
    med.character_id,
    me.avl_self,
    me.avl_context,
    me.avl_context_2,
    NULL::availability_type AS avl_shop_area,
    NULL::availability_type AS avl_shop_subloc,
    NULL::availability_type AS avl_shop_loc,
    NULL::availability_type AS avl_shop_area_char,
    NULL::availability_type AS avl_shop_subloc_char,
    NULL::availability_type AS avl_shop_loc_char,
    me.is_repeatable,
    me.is_repeatable_loc,
    me.area_id,
    NULL::shop_type AS shop_type
FROM mv_monster_equipment_drops med
JOIN mv_monster_encounters me ON med.monster_id = me.monster_id;

CREATE INDEX idx_mv_auto_ability_sources_auto_ability_id ON mv_auto_ability_sources (auto_ability_id);
CREATE INDEX idx_mv_auto_ability_sources_lookup ON mv_auto_ability_sources (source_id, source_type, area_id);
CREATE INDEX idx_mv_auto_ability_sources_area_id ON mv_auto_ability_sources (area_id);




CREATE MATERIALIZED VIEW mv_availabilities AS
SELECT DISTINCT
    m.id AS s_id,
    m.name,
    m.version AS v,
    'monster' AS source_type,
    CASE
        WHEN m.category = 'boss' THEN 'boss'
        ELSE 'monster' -- I'm not quite sure how to handle it, probably change GetAreasIDsWithBosses query
    END::text AS sub_type,
    me.avl_self,
    me.avl_context,
    me.avl_context_2,
    m.is_repeatable,
    me.is_repeatable_loc,
    me.area_id AS a_id,
    me.location,
    me.sublocation,
    me.area,
    me.version AS av
FROM mv_monster_encounters me
JOIN monsters m ON me.monster_id = m.id

UNION ALL

SELECT DISTINCT
    me.formation_id AS s_id,
    NULL::text AS name,
    NULL::int AS v,
    'monster-formation' AS source_type,
    NULL::text AS sub_type,
    me.avl_context AS avl_self,
    me.avl_context,
    me.avl_context_2,
    me.is_repeatable,
    me.is_repeatable_loc,
    me.area_id AS a_id,
    me.location,
    me.sublocation,
    me.area,
    me.version AS av
FROM mv_monster_encounters me

UNION ALL

SELECT DISTINCT
    g.area_id AS s_id,
    g.area AS name,
    g.version::int AS v,
    'area' AS source_type,
    NULL::text AS sub_type,
    a.availability AS avl_self,
    a.availability AS avl_context,
    a.availability AS avl_context_2,
    NULL::boolean AS is_repeatable,
    NULL::boolean AS is_repeatable_loc,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS av
FROM areas a
JOIN mv_geography g ON g.area_id = a.id

UNION ALL

SELECT DISTINCT
    t.id AS s_id,
    NULL::text AS name,
    NULL::int AS v,
    'treasure' AS source_type,
    NULL::text AS sub_type,
    t.availability AS avl_self,
    t.availability AS avl_context,
    t.availability AS avl_context_2,
    FALSE AS is_repeatable,
    FALSE AS is_repeatable_loc,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS av
FROM treasures t
JOIN mv_geography g ON t.area_id = g.area_id

UNION ALL

SELECT DISTINCT
    sh.id AS s_id,
    NULL::text AS name,
    NULL::int AS v,
    'shop' AS source_type,
    'item'::text AS sub_type,
    sh.availability AS avl_self,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY sh.id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY sh.id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context,
    CASE
        WHEN BOOL_OR(j.shop_type = 'pre-airship') OVER (PARTITION BY sh.id) 
         AND BOOL_OR(j.shop_type = 'post-airship') OVER (PARTITION BY sh.id) 
        THEN 'always'::availability_type
        WHEN j.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN j.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context_2,
    FALSE AS is_repeatable,
    FALSE AS is_repeatable_loc,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS av
FROM shops sh
JOIN j_shops_items j ON j.shop_id = sh.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
JOIN mv_geography g ON sh.area_id = g.area_id

UNION ALL

SELECT DISTINCT
    sh.id AS s_id,
    NULL::text AS name,
    NULL::int AS v,
    'shop' AS source_type,
    'equip'::text AS sub_type,
    sh.availability AS avl_self,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY sh.id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY sh.id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context,
    CASE
        WHEN BOOL_OR(se.shop_type = 'pre-airship') OVER (PARTITION BY sh.id) 
         AND BOOL_OR(se.shop_type = 'post-airship') OVER (PARTITION BY sh.id) 
        THEN 'always'::availability_type
        WHEN se.shop_type = 'pre-airship' THEN 'pre-story'::availability_type
        WHEN se.shop_type = 'post-airship' THEN 'post'::availability_type
    END AS avl_context_2,
    FALSE AS is_repeatable,
    FALSE AS is_repeatable_loc,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS av
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN mv_geography g ON sh.area_id = g.area_id

UNION ALL

SELECT DISTINCT
    q.id AS s_id,
    q.name,
    NULL::int AS v,
    'quest' AS source_type,
    NULL::text AS sub_type,
    q.availability AS avl_self,
    q.availability AS avl_context,
    q.availability AS avl_context_2,
    q.is_repeatable,
    q.is_repeatable AS is_repeatable_loc,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS av
FROM quests q
LEFT JOIN quest_completions qc ON qc.quest_id = q.id
LEFT JOIN completion_areas ca ON ca.completion_id = qc.id
LEFT JOIN mv_geography g ON ca.area_id = g.area_id

UNION ALL

SELECT DISTINCT
    bp.id AS s_id,
    NULL::text AS name,
    NULL::int AS v,
    'blitzball' AS source_type,
    NULL::text AS sub_type,
    'always'::availability_type AS avl_self,
    'always'::availability_type AS avl_context,
    'always'::availability_type AS avl_context_2,
    TRUE AS is_repeatable,
    TRUE AS is_repeatable_loc,
    g.area_id AS a_id,
    g.location,
    g.sublocation,
    g.area,
    g.version AS av
FROM blitzball_positions bp
JOIN mv_geography g ON 75 = g.area_id;

CREATE INDEX idx_mv_availabilities_lookup ON mv_availabilities (s_id, source_type, sub_type);
CREATE INDEX idx_mv_availabilities_area_id ON mv_availabilities (a_id);




CREATE MATERIALIZED VIEW mv_abilities AS
SELECT DISTINCT
    a.id AS ability_id,
    a.name,
    a.version,
    pa.id AS typed_id,
    a.type,
    aa.rank,
    aa.can_copycat,
    aa.appears_in_help_bar,
    bi.id AS battle_int_id,
    bi.target,
    bi.based_on_user_attack,
    bi.range,
    bi.inflicted_delay_id,
    d.critical,
    d.break_dmg_limit,
    d.element_id,
    ad.attack_type,
    ad.stat_id,
    ad.damage_type,
    ad.damage_formula,
    ad.damage_constant,
    idl.ctb_attack_type,
    jab.status_condition_id AS affected_status_id,
    jis.inflicted_status_id,
    jrs.status_condition_id AS removed_status_id,
    jsc.stat_change_id,
    jmc.modifier_change_id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
JOIN player_abilities pa ON pa.ability_id = a.id
LEFT JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
LEFT JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_damage jd ON jd.ability_id = a.id AND jd.battle_interaction_id = bi.id
LEFT JOIN damages d ON jd.damage_id = d.id
LEFT JOIN j_damages_damage_calc jdc ON jdc.ability_id = a.id AND jdc.battle_interaction_id = bi.id AND jdc.damage_id = d.id
LEFT JOIN ability_damages ad ON jdc.ability_damage_id = ad.id
LEFT JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
LEFT JOIN j_battle_interactions_affected_by jab ON jab.ability_id = a.id AND jab.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_inflicted_status_conditions jis ON jis.ability_id = a.id AND jis.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_removed_status_conditions jrs ON jrs.ability_id = a.id AND jrs.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_stat_changes jsc ON jsc.ability_id = a.id AND jsc.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_modifier_changes jmc ON jmc.ability_id = a.id AND jmc.battle_interaction_id = bi.id

UNION ALL

SELECT DISTINCT
    a.id AS ability_id,
    a.name,
    a.version,
    oa.id AS typed_id,
    a.type,
    aa.rank,
    aa.can_copycat,
    aa.appears_in_help_bar,
    bi.id AS battle_int_id,
    bi.target,
    bi.based_on_user_attack,
    bi.range,
    bi.inflicted_delay_id,
    d.critical,
    d.break_dmg_limit,
    d.element_id,
    ad.attack_type,
    ad.stat_id,
    ad.damage_type,
    ad.damage_formula,
    ad.damage_constant,
    idl.ctb_attack_type,
    jab.status_condition_id AS affected_status_id,
    jis.inflicted_status_id,
    jrs.status_condition_id AS removed_status_id,
    jsc.stat_change_id,
    jmc.modifier_change_id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
JOIN overdrive_abilities oa ON oa.ability_id = a.id
LEFT JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
LEFT JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_damage jd ON jd.ability_id = a.id AND jd.battle_interaction_id = bi.id
LEFT JOIN damages d ON jd.damage_id = d.id
LEFT JOIN j_damages_damage_calc jdc ON jdc.ability_id = a.id AND jdc.battle_interaction_id = bi.id AND jdc.damage_id = d.id
LEFT JOIN ability_damages ad ON jdc.ability_damage_id = ad.id
LEFT JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
LEFT JOIN j_battle_interactions_affected_by jab ON jab.ability_id = a.id AND jab.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_inflicted_status_conditions jis ON jis.ability_id = a.id AND jis.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_removed_status_conditions jrs ON jrs.ability_id = a.id AND jrs.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_stat_changes jsc ON jsc.ability_id = a.id AND jsc.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_modifier_changes jmc ON jmc.ability_id = a.id AND jmc.battle_interaction_id = bi.id

UNION ALL

SELECT DISTINCT
    a.id AS ability_id,
    a.name,
    a.version,
    ia.id AS typed_id,
    a.type,
    aa.rank,
    aa.can_copycat,
    aa.appears_in_help_bar,
    bi.id AS battle_int_id,
    bi.target,
    bi.based_on_user_attack,
    bi.range,
    bi.inflicted_delay_id,
    d.critical,
    d.break_dmg_limit,
    d.element_id,
    ad.attack_type,
    ad.stat_id,
    ad.damage_type,
    ad.damage_formula,
    ad.damage_constant,
    idl.ctb_attack_type,
    jab.status_condition_id AS affected_status_id,
    jis.inflicted_status_id,
    jrs.status_condition_id AS removed_status_id,
    jsc.stat_change_id,
    jmc.modifier_change_id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
JOIN item_abilities ia ON ia.ability_id = a.id
LEFT JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
LEFT JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_damage jd ON jd.ability_id = a.id AND jd.battle_interaction_id = bi.id
LEFT JOIN damages d ON jd.damage_id = d.id
LEFT JOIN j_damages_damage_calc jdc ON jdc.ability_id = a.id AND jdc.battle_interaction_id = bi.id AND jdc.damage_id = d.id
LEFT JOIN ability_damages ad ON jdc.ability_damage_id = ad.id
LEFT JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
LEFT JOIN j_battle_interactions_affected_by jab ON jab.ability_id = a.id AND jab.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_inflicted_status_conditions jis ON jis.ability_id = a.id AND jis.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_removed_status_conditions jrs ON jrs.ability_id = a.id AND jrs.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_stat_changes jsc ON jsc.ability_id = a.id AND jsc.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_modifier_changes jmc ON jmc.ability_id = a.id AND jmc.battle_interaction_id = bi.id

UNION ALL

SELECT DISTINCT
    a.id AS ability_id,
    a.name,
    a.version,
    tc.id AS typed_id,
    a.type,
    aa.rank,
    aa.can_copycat,
    aa.appears_in_help_bar,
    bi.id AS battle_int_id,
    bi.target,
    bi.based_on_user_attack,
    bi.range,
    bi.inflicted_delay_id,
    d.critical,
    d.break_dmg_limit,
    d.element_id,
    ad.attack_type,
    ad.stat_id,
    ad.damage_type,
    ad.damage_formula,
    ad.damage_constant,
    idl.ctb_attack_type,
    jab.status_condition_id AS affected_status_id,
    jis.inflicted_status_id,
    jrs.status_condition_id AS removed_status_id,
    jsc.stat_change_id,
    jmc.modifier_change_id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
JOIN trigger_commands tc ON tc.ability_id = a.id
LEFT JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
LEFT JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_damage jd ON jd.ability_id = a.id AND jd.battle_interaction_id = bi.id
LEFT JOIN damages d ON jd.damage_id = d.id
LEFT JOIN j_damages_damage_calc jdc ON jdc.ability_id = a.id AND jdc.battle_interaction_id = bi.id AND jdc.damage_id = d.id
LEFT JOIN ability_damages ad ON jdc.ability_damage_id = ad.id
LEFT JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
LEFT JOIN j_battle_interactions_affected_by jab ON jab.ability_id = a.id AND jab.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_inflicted_status_conditions jis ON jis.ability_id = a.id AND jis.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_removed_status_conditions jrs ON jrs.ability_id = a.id AND jrs.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_stat_changes jsc ON jsc.ability_id = a.id AND jsc.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_modifier_changes jmc ON jmc.ability_id = a.id AND jmc.battle_interaction_id = bi.id

UNION ALL

SELECT DISTINCT
    a.id AS ability_id,
    a.name,
    a.version,
    ma.id AS typed_id,
    a.type,
    aa.rank,
    aa.can_copycat,
    aa.appears_in_help_bar,
    bi.id AS battle_int_id,
    bi.target,
    bi.based_on_user_attack,
    bi.range,
    bi.inflicted_delay_id,
    d.critical,
    d.break_dmg_limit,
    d.element_id,
    ad.attack_type,
    ad.stat_id,
    ad.damage_type,
    ad.damage_formula,
    ad.damage_constant,
    idl.ctb_attack_type,
    jab.status_condition_id AS affected_status_id,
    jis.inflicted_status_id,
    jrs.status_condition_id AS removed_status_id,
    jsc.stat_change_id,
    jmc.modifier_change_id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
JOIN misc_abilities ma ON ma.ability_id = a.id
LEFT JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
LEFT JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_damage jd ON jd.ability_id = a.id AND jd.battle_interaction_id = bi.id
LEFT JOIN damages d ON jd.damage_id = d.id
LEFT JOIN j_damages_damage_calc jdc ON jdc.ability_id = a.id AND jdc.battle_interaction_id = bi.id AND jdc.damage_id = d.id
LEFT JOIN ability_damages ad ON jdc.ability_damage_id = ad.id
LEFT JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
LEFT JOIN j_battle_interactions_affected_by jab ON jab.ability_id = a.id AND jab.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_inflicted_status_conditions jis ON jis.ability_id = a.id AND jis.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_removed_status_conditions jrs ON jrs.ability_id = a.id AND jrs.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_stat_changes jsc ON jsc.ability_id = a.id AND jsc.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_modifier_changes jmc ON jmc.ability_id = a.id AND jmc.battle_interaction_id = bi.id

UNION ALL

SELECT DISTINCT
    a.id AS ability_id,
    a.name,
    a.version,
    ea.id AS typed_id,
    a.type,
    aa.rank,
    aa.can_copycat,
    aa.appears_in_help_bar,
    bi.id AS battle_int_id,
    bi.target,
    bi.based_on_user_attack,
    bi.range,
    bi.inflicted_delay_id,
    d.critical,
    d.break_dmg_limit,
    d.element_id,
    ad.attack_type,
    ad.stat_id,
    ad.damage_type,
    ad.damage_formula,
    ad.damage_constant,
    idl.ctb_attack_type,
    jab.status_condition_id AS affected_status_id,
    jis.inflicted_status_id,
    jrs.status_condition_id AS removed_status_id,
    jsc.stat_change_id,
    jmc.modifier_change_id
FROM abilities a
JOIN ability_attributes aa ON a.attributes_id = aa.id
JOIN enemy_abilities ea ON ea.ability_id = a.id
LEFT JOIN j_abilities_battle_interactions j ON j.ability_id = a.id
LEFT JOIN battle_interactions bi ON j.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_damage jd ON jd.ability_id = a.id AND jd.battle_interaction_id = bi.id
LEFT JOIN damages d ON jd.damage_id = d.id
LEFT JOIN j_damages_damage_calc jdc ON jdc.ability_id = a.id AND jdc.battle_interaction_id = bi.id AND jdc.damage_id = d.id
LEFT JOIN ability_damages ad ON jdc.ability_damage_id = ad.id
LEFT JOIN inflicted_delays idl ON bi.inflicted_delay_id = idl.id
LEFT JOIN j_battle_interactions_affected_by jab ON jab.ability_id = a.id AND jab.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_inflicted_status_conditions jis ON jis.ability_id = a.id AND jis.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_removed_status_conditions jrs ON jrs.ability_id = a.id AND jrs.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_stat_changes jsc ON jsc.ability_id = a.id AND jsc.battle_interaction_id = bi.id
LEFT JOIN j_battle_interactions_modifier_changes jmc ON jmc.ability_id = a.id AND jmc.battle_interaction_id = bi.id
ORDER BY ability_id;


CREATE INDEX idx_mv_abilities_id ON mv_abilities (ability_id);
CREATE INDEX idx_mv_abilities_typed_id ON mv_abilities (typed_id);
CREATE INDEX idx_mv_abilities_type ON mv_abilities (type);
CREATE INDEX idx_mv_abilities_inflicted_status_id ON mv_abilities (inflicted_status_id);
CREATE INDEX idx_mv_abilities_removed_status_id ON mv_abilities (removed_status_id);
CREATE INDEX idx_mv_abilities_element_id ON mv_abilities (element_id);



-- +goose Down
DROP MATERIALIZED VIEW IF EXISTS mv_abilities;
DROP MATERIALIZED VIEW IF EXISTS mv_availabilities;
DROP MATERIALIZED VIEW IF EXISTS mv_auto_ability_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_equipment_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_item_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_equipment_drops;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_item_drops;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_encounters;
DROP MATERIALIZED VIEW IF EXISTS mv_geography_graph;
DROP MATERIALIZED VIEW IF EXISTS mv_geography;
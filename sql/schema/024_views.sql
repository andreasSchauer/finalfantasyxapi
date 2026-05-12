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


CREATE MATERIALIZED VIEW mv_equipment_sources AS
SELECT
    en.id AS name_id,
    en.name AS name,
    t.id AS source_id,
    t.area_id,
    'treasure' AS source_type,
    te.empty_slots_amount,
    j.auto_ability_id,
    aa.name AS auto_ability,
    t.availability,
    NULL::shop_type AS shop_type
FROM treasures t
JOIN treasure_equipment_pieces te ON te.treasure_id = t.id
LEFT JOIN j_treasure_equipment_abilities j ON j.treasure_equipment_id = te.id
LEFT JOIN auto_abilities aa ON j.auto_ability_id = aa.id
JOIN equipment_names en ON te.equipment_name_id = en.id

UNION ALL

SELECT
    en.id AS name_id,
    en.name AS name,
    sh.id AS source_id,
    sh.area_id,
    'shop' AS source_type,
    se.empty_slots_amount,
    j.auto_ability_id,
    aa.name AS auto_ability,
    CASE se.shop_type 
        WHEN 'pre-airship' THEN 'pre-story'::availability_type
        WHEN 'post-airship' THEN 'post'::availability_type
    END AS availability,
    se.shop_type::shop_type AS shop_type
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
LEFT JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
LEFT JOIN auto_abilities aa ON j.auto_ability_id = aa.id
JOIN equipment_names en ON se.equipment_name_id = en.id;


CREATE INDEX idx_mv_equipment_sources_name_id ON mv_equipment_sources (name_id);
CREATE INDEX idx_mv_equipment_sources_source_id ON mv_equipment_sources (source_id);
CREATE INDEX idx_mv_equipment_sources_area_id ON mv_equipment_sources (area_id);





CREATE MATERIALIZED VIEW mv_abilities AS
SELECT
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

SELECT
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

SELECT
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

SELECT
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

SELECT
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

SELECT
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
DROP MATERIALIZED VIEW IF EXISTS mv_equipment_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_item_sources;
DROP MATERIALIZED VIEW IF EXISTS mv_geography_graph;
DROP MATERIALIZED VIEW IF EXISTS mv_geography;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_encounters;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_equipment_drops;
DROP MATERIALIZED VIEW IF EXISTS mv_monster_item_drops;
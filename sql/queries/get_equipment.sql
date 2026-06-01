-- name: GetAutoAbilityItemMonsterIDs :many
WITH w AS (
    SELECT 
        sqlc.narg('repeatable')::BOOLEAN AS repeatable,
        sqlc.narg('availability')::availability_type[] AS availability,
        (
            SELECT mit.id AS master_item_id
            FROM auto_abilities aa
            JOIN item_amounts ia_req ON aa.required_item_amount_id = ia_req.id
            JOIN master_items mit ON ia_req.master_item_id = mit.id
            WHERE aa.id = sqlc.arg('auto_ability_id')::int
        )::int AS target_master_id
)
SELECT DISTINCT md.monster_id
FROM mv_monster_item_drops md
JOIN monsters m ON md.monster_id = m.id
CROSS JOIN w
WHERE md.master_item_id = w.target_master_id
  AND (w.repeatable IS NULL OR md.is_repeatable = w.repeatable)
  AND (w.availability IS NULL OR m.availability = ANY(w.availability))
ORDER BY md.monster_id;


-- name: GetAutoAbilityMonsterIDs :many
WITH w AS (
    SELECT 
        sqlc.arg('auto_ability_id')::int AS auto_ability_id,
        sqlc.narg('repeatable')::BOOLEAN AS repeatable,
        sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT m.id
FROM monsters m
JOIN mv_monster_equipment_drops me ON me.monster_id = m.id
CROSS JOIN w
WHERE me.auto_ability_id = w.auto_ability_id
  AND (w.repeatable IS NULL OR m.is_repeatable = w.repeatable)
  AND (w.availability IS NULL OR m.availability = ANY(w.availability))
ORDER BY m.id;



-- name: GetAutoAbilityTreasureIDs :many
WITH w as (
  SELECT sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT es.source_id
FROM mv_equipment_sources es
JOIN treasures t ON es.source_id = t.id AND es.source_type = 'treasure' AND es.area_id = t.area_id
CROSS JOIN w
WHERE es.auto_ability_id = sqlc.arg('auto_ability_id')::int
  AND (w.availability IS NULL OR t.availability = ANY(w.availability))
ORDER BY es.source_id;


-- name: GetAutoAbilityEquipmentTableIDs :many
SELECT j.equipment_table_id
FROM j_equipment_tables_required_auto_abilities j
WHERE j.auto_ability_id = $1

UNION

SELECT ap.equipment_table_id
FROM ability_pools ap
JOIN j_ability_pools_auto_abilities j ON j.ability_pool_id = ap.id
WHERE j.auto_ability_id = $1
ORDER BY equipment_table_id;


-- name: GetAutoAbilityShopIDsPre :many
WITH w as (
  SELECT
    sqlc.arg('auto_ability_id')::int AS auto_ability_id,
    sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT es.source_id
FROM mv_equipment_sources es
CROSS JOIN w
WHERE es.auto_ability_id = w.auto_ability_id
  AND es.source_type = 'shop'
  AND es.shop_type = 'pre-airship'
  AND (w.availability IS NULL OR es.avl_context = ANY(w.availability))
ORDER BY es.source_id;


-- name: GetAutoAbilityShopIDsPost :many
WITH w as (
  SELECT
    sqlc.arg('auto_ability_id')::int AS auto_ability_id,
    sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT es.source_id
FROM mv_equipment_sources es
CROSS JOIN w
WHERE es.auto_ability_id = w.auto_ability_id
  AND es.source_type = 'shop'
  AND es.shop_type = 'post-airship'
  AND (w.availability IS NULL OR es.avl_context = ANY(w.availability))
ORDER BY es.source_id;


-- name: GetAutoAbilityIDs :many
SELECT id FROM auto_abilities ORDER BY id;


-- name: GetAutoAbilityIDsByCategory :many
SELECT id FROM auto_abilities WHERE category = ANY(sqlc.narg('auto_ability_category')::auto_ability_category[]) ORDER BY id;


-- name: GetAutoAbilityIDsByEquipType :many
SELECT id FROM auto_abilities WHERE type = $1 ORDER BY id;


-- name: GetAutoAbilityIDsByMonster :many
WITH w AS (
  SELECT
    sqlc.arg('monster_id')::int AS monster_id,
    sqlc.narg('character_id')::int AS character_id
)
SELECT DISTINCT a.auto_ability_id
FROM mv_auto_ability_sources a
CROSS JOIN w
WHERE a.source_id = w.monster_id
  AND a.source_type = 'monster'
  AND (w.character_id IS NULL OR a.character_id = w.character_id OR a.character_id IS NULL)
ORDER BY a.auto_ability_id;


-- name: GetAutoAbilityIDsByMonsterItems :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN item_amounts ia_req ON aa.required_item_amount_id = ia_req.id
JOIN mv_monster_item_drops md ON ia_req.master_item_id = md.master_item_id
WHERE md.monster_id = $1
ORDER BY aa.id;


-- name: GetAutoAbilityIDsByShop :many
WITH w AS (
  SELECT
    sqlc.arg('shop_id')::int AS shop_id,
    sqlc.narg('character_id')::int AS character_id
)
SELECT DISTINCT a.auto_ability_id
FROM mv_auto_ability_sources a
CROSS JOIN w
WHERE a.source_id = w.shop_id
  AND a.source_type = 'shop'
  AND (w.character_id IS NULL OR a.character_id = w.character_id OR a.character_id IS NULL)
ORDER BY a.auto_ability_id;


-- name: GetAutoAbilityIDsByLocation :many
SELECT DISTINCT a.auto_ability_id
FROM mv_auto_ability_sources a
JOIN mv_geography g ON a.area_id = g.area_id
WHERE g.location_id = $1
ORDER BY a.auto_ability_id;


-- name: GetAutoAbilityIDsBySublocation :many
SELECT DISTINCT a.auto_ability_id
FROM mv_auto_ability_sources a
JOIN mv_geography g ON a.area_id = g.area_id
WHERE g.sublocation_id = $1
ORDER BY a.auto_ability_id;


-- name: GetAutoAbilityIDsByArea :many
SELECT DISTINCT auto_ability_id
FROM mv_auto_ability_sources
WHERE area_id = $1
ORDER BY auto_ability_id;






-- name: GetEquipmentTableCelestialWeaponID :one
SELECT celestial_weapon_id::int FROM j_equipment_tables_names WHERE celestial_weapon_id IS NOT NULL AND equipment_table_id = $1;


-- name: GetEquipmentTableIDs :many
SELECT id FROM equipment_tables ORDER BY id;


-- name: GetEquipmentTableIDsByAutoAbility :many
WITH all_matches AS (
    SELECT equipment_table_id, auto_ability_id
    FROM j_equipment_tables_required_auto_abilities

    UNION

    SELECT ap.equipment_table_id, jpool.auto_ability_id
    FROM ability_pools ap
    JOIN j_ability_pools_auto_abilities jpool ON jpool.ability_pool_id = ap.id
)
SELECT m.equipment_table_id
FROM all_matches m
CROSS JOIN (SELECT sqlc.arg('auto_ability_ids')::int[] AS ids) w
WHERE m.auto_ability_id = ANY(w.ids)
GROUP BY m.equipment_table_id, w.ids
HAVING COUNT(DISTINCT m.auto_ability_id) = cardinality(w.ids)
ORDER BY m.equipment_table_id;


-- name: GetEquipmentTableIDsEquipType :many
SELECT id FROM equipment_tables WHERE type = $1 ORDER BY id;


-- name: GetEquipmentTableIDsCelestialWeapon :many
SELECT id FROM equipment_tables WHERE classification = 'celestial-weapon' ORDER BY id;





-- name: GetEquipmentEquipmentTableIDs :many
SELECT equipment_table_id FROM j_equipment_tables_names WHERE equipment_name_id = $1 ORDER BY equipment_table_id;


-- name: GetEquipmentTreasureIDs :many
WITH w AS (
  SELECT
    sqlc.arg('equipment_id')::int AS equipment_name_id,
    sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT t.id
FROM treasures t
JOIN mv_equipment_sources es ON es.source_id = t.id AND es.source_type = 'treasure' AND es.area_id = t.area_id
CROSS JOIN w
WHERE es.name_id = w.equipment_name_id
  AND (w.availability IS NULL OR t.availability = ANY(w.availability))
ORDER BY t.id;


-- name: GetEquipmentShopIDs :many
WITH w AS (
  SELECT
    sqlc.arg('equipment_id')::int AS equipment_name_id,
    sqlc.narg('availability')::availability_type[] AS availability
)
SELECT DISTINCT sh.id
FROM shops sh
JOIN mv_equipment_sources es ON es.source_id = sh.id AND es.source_type = 'shop' AND es.area_id = sh.area_id
CROSS JOIN w
WHERE es.name_id = w.equipment_name_id
  AND (w.availability IS NULL OR sh.availability = ANY(w.availability))
ORDER BY sh.id;


-- name: GetEquipmentIDs :many
SELECT id FROM equipment_names ORDER BY id;


-- name: GetEquipmentIDsByCharacter :many
SELECT id FROM equipment_names WHERE character_id = $1 ORDER BY id;


-- name: GetEquipmentIDsByEquipType :many
SELECT DISTINCT j.equipment_name_id
FROM j_equipment_tables_names j
JOIN equipment_tables et ON j.equipment_table_id = et.id
WHERE et.type = $1
ORDER BY j.equipment_name_id;


-- name: GetEquipmentIDsCelestialWeapon :many
SELECT DISTINCT j.equipment_name_id
FROM j_equipment_tables_names j
JOIN equipment_tables et ON j.equipment_table_id = et.id
WHERE et.classification = 'celestial-weapon'
ORDER BY j.equipment_name_id;


-- name: GetEquipmentIDsByAutoAbilty :many
WITH all_matches AS (
    SELECT en.id AS equipment_name_id, jreq.auto_ability_id
    FROM equipment_names en
    JOIN j_equipment_tables_names j ON j.equipment_name_id = en.id
    JOIN equipment_tables et ON j.equipment_table_id = et.id
    JOIN j_equipment_tables_required_auto_abilities jreq ON jreq.equipment_table_id = et.id
                
    UNION                                  
                         
    SELECT en.id AS equipment_name_id, jpool.auto_ability_id
    FROM equipment_names en
    JOIN j_equipment_tables_names j ON j.equipment_name_id = en.id
    JOIN equipment_tables et ON j.equipment_table_id = et.id
    JOIN ability_pools ap ON ap.equipment_table_id = et.id
    JOIN j_ability_pools_auto_abilities jpool ON jpool.ability_pool_id = ap.id
)
SELECT m.equipment_name_id
FROM all_matches m
CROSS JOIN (SELECT sqlc.arg('auto_ability_ids')::int[] AS ids) w
WHERE m.auto_ability_id = ANY(w.ids)
GROUP BY m.equipment_name_id, w.ids
HAVING COUNT(DISTINCT m.auto_ability_id) = cardinality(w.ids)
ORDER BY m.equipment_name_id;





-- name: GetCelestialWeaponTreasureID :one
SELECT DISTINCT es.source_id
FROM mv_equipment_sources es
JOIN j_equipment_tables_names j ON j.equipment_name_id = es.name_id
WHERE j.celestial_weapon_id = $1
  AND es.source_type = 'treasure';


-- name: GetCelestialWeaponAutoAbilityIDs :many
SELECT DISTINCT j1.auto_ability_id
FROM j_equipment_tables_required_auto_abilities j1
JOIN j_equipment_tables_names j2 ON j1.equipment_table_id = j2.equipment_table_id
WHERE j2.celestial_weapon_id = sqlc.arg('celestial_weapon_id')::int
  AND j2.equipment_table_id = sqlc.arg('equipment_table_id')::int
ORDER BY j1.auto_ability_id;


-- name: GetCelestialWeaponIDs :many
SELECT id FROM celestial_weapons ORDER BY id;


-- name: GetCelestialWeaponIDsByFormula :many
SELECT id FROM celestial_weapons WHERE formula = $1 ORDER BY id;
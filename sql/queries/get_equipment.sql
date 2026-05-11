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
            WHERE aa.id = sqlc.arg(auto_ability_id)
        )::int AS target_master_id
)
SELECT DISTINCT md.monster_id
FROM mv_monster_item_drops md
CROSS JOIN w
WHERE md.master_item_id = w.target_master_id
  AND (w.repeatable IS NULL OR md.is_repeatable = w.repeatable)
  AND (w.availability IS NULL OR md.availability = ANY(w.availability))
ORDER BY md.monster_id;


-- name: GetAutoAbilityMonsterIDs :many
WITH w AS (
    SELECT 
        sqlc.narg('repeatable')::BOOLEAN AS repeatable,
        sqlc.narg('availability')::availability_type[] AS availability
),
filtered_monsters AS (
    SELECT m.id
    FROM monsters m
    CROSS JOIN w
    WHERE (w.repeatable IS NULL OR m.is_repeatable = w.repeatable)
      AND (w.availability IS NULL OR m.availability = ANY(w.availability))
)
SELECT DISTINCT fm.id
FROM filtered_monsters fm
JOIN mv_monster_equipment_drops me ON me.monster_id = fm.id
WHERE me.auto_ability_id = sqlc.arg('auto_ability_id')
ORDER BY fm.id;


-- name: GetAutoAbilityTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN treasure_equipment_pieces te ON te.treasure_id = t.id
JOIN j_treasure_equipment_abilities j ON j.treasure_equipment_id = te.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
WHERE (w.availability IS NULL OR t.availability = ANY(w.availability))
  AND j.auto_ability_id = sqlc.arg(auto_ability_id)
ORDER BY t.id;


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
SELECT DISTINCT sh.id
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
WHERE (w.availability IS NULL OR sh.availability = ANY(w.availability))
  AND se.shop_type = 'pre-airship'
  AND j.auto_ability_id = $1
ORDER BY sh.id;


-- name: GetAutoAbilityShopIDsPost :many
SELECT DISTINCT sh.id
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
WHERE( w.availability IS NULL OR sh.availability = ANY(w.availability))
  AND se.shop_type = 'post-airship'
  AND j.auto_ability_id = $1
ORDER BY sh.id;


-- name: GetAutoAbilityIDs :many
SELECT id FROM auto_abilities ORDER BY id;


-- name: GetAutoAbilityIDsByCategory :many
SELECT id FROM auto_abilities WHERE category = ANY(sqlc.narg('auto_ability_category')::auto_ability_category[]) ORDER BY id;


-- name: GetAutoAbilityIDsByEquipType :many
SELECT id FROM auto_abilities WHERE type = $1 ORDER BY id;


-- name: GetAutoAbilityIDsByMonster :many
SELECT DISTINCT auto_ability_id FROM mv_monster_equipment_drops WHERE monster_id = $1 ORDER BY auto_ability_id;


-- name: GetAutoAbilityIDsByMonsterItems :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN item_amounts ia_req ON aa.required_item_amount_id = ia_req.id
JOIN mv_monster_item_drops md ON ia_req.master_item_id = md.master_item_id
WHERE md.monster_id = $1
ORDER BY aa.id;




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
SELECT DISTINCT t.id
FROM treasures t
JOIN treasure_equipment_pieces te ON te.treasure_id = t.id
JOIN equipment_names en ON te.equipment_name_id = en.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
WHERE en.id = sqlc.arg('equipment_id')::int
  AND (w.availability IS NULL OR t.availability = ANY(w.availability))
ORDER BY t.id;


-- name: GetEquipmentShopIDs :many
SELECT DISTINCT se.shop_id
FROM shop_equipment_pieces se
JOIN equipment_names en ON se.equipment_name_id = en.id
CROSS JOIN (SELECT sqlc.narg('availability')::availability_type[] AS availability) w
CROSS JOIN LATERAL (
    SELECT CASE se.shop_type
        WHEN 'pre-airship' THEN 'pre-story'::availability_type
        WHEN 'post-airship' THEN 'post'::availability_type
    END AS shop_availability
) calc
WHERE en.id = sqlc.arg('equipment_id')::int
  AND (w.availability IS NULL OR calc.shop_availability = ANY(w.availability))
ORDER BY se.shop_id;


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
SELECT te.treasure_id
FROM treasure_equipment_pieces te
JOIN equipment_names en ON te.equipment_name_id = en.id
JOIN j_equipment_tables_names j ON j.equipment_name_id = en.id
WHERE j.celestial_weapon_id = $1;


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
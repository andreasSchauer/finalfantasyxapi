-- name: GetAutoAbilityItemMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
WHERE (sqlc.narg('repeatable')::BOOLEAN IS NULL OR m.is_repeatable = sqlc.narg('repeatable')::BOOLEAN)
  AND (sqlc.narg('availability')::availability_type[] IS NULL OR m.availability = ANY(sqlc.narg('availability')::availability_type[]))
  AND (
    EXISTS (
      SELECT 1
      FROM item_amounts ia
      JOIN master_items mit ON ia.master_item_id = mit.id
      JOIN item_amounts ia2 ON ia2.master_item_id = mit.id
      JOIN auto_abilities aa ON aa.required_item_amount_id = ia2.id
      WHERE ia.id IN (
          mi.steal_common_id,
          mi.steal_rare_id,
          mi.drop_common_id,
          mi.drop_rare_id,
          mi.secondary_drop_common_id,
          mi.secondary_drop_rare_id,
          mi.bribe_id
      )
      AND aa.id = sqlc.arg(auto_ability_id)
  )
    OR EXISTS (
      SELECT 1
      FROM j_monster_items_other_items jmio
      JOIN possible_items pi ON pi.id = jmio.possible_item_id
      JOIN item_amounts ia ON pi.item_amount_id = ia.id
      JOIN master_items mit ON ia.master_item_id = mit.id
      JOIN item_amounts ia2 ON ia2.master_item_id = mit.id
      JOIN auto_abilities aa ON aa.required_item_amount_id = ia2.id
      WHERE jmio.monster_items_id = mi.id
        AND aa.id = sqlc.arg(auto_ability_id)
    )
  )
ORDER BY m.id;


-- name: GetAutoAbilityMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_equipment me ON me.monster_id = m.id
JOIN j_monster_equipment_abilities j ON j.monster_equipment_id = me.id
JOIN equipment_drops ed ON j.equipment_drop_id = ed.id
JOIN auto_abilities aa ON ed.auto_ability_id = aa.id
WHERE (sqlc.narg('repeatable')::BOOLEAN IS NULL OR m.is_repeatable = sqlc.narg('repeatable')::BOOLEAN)
  AND (sqlc.narg('availability')::availability_type[] IS NULL OR m.availability = ANY(sqlc.narg('availability')::availability_type[]))
  AND aa.id = sqlc.arg(auto_ability_id)
ORDER BY m.id;


-- name: GetAutoAbilityTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN treasure_equipment_pieces te ON te.treasure_id = t.id
JOIN j_treasure_equipment_abilities j ON j.treasure_equipment_id = te.id
JOIN auto_abilities aa ON j.auto_ability_id = aa.id
WHERE (sqlc.narg('availability')::availability_type[] IS NULL OR t.availability = ANY(sqlc.narg('availability')::availability_type[]))
  AND aa.id = sqlc.arg(auto_ability_id)
ORDER BY t.id;


-- name: GetAutoAbilityEquipmentTableIDs :many
SELECT DISTINCT et.id
FROM equipment_tables et
WHERE EXISTS (
    SELECT 1
    FROM j_equipment_tables_required_auto_abilities j
    JOIN auto_abilities aa ON j.auto_ability_id = aa.id
    WHERE j.equipment_table_id = et.id
        AND aa.id = $1
)
OR EXISTS (
    SELECT 1
    FROM ability_pools ap
    JOIN j_ability_pools_auto_abilities j ON j.ability_pool_id = ap.id
    JOIN auto_abilities aa ON j.auto_ability_id = aa.id
    WHERE ap.equipment_table_id = et.id
        AND aa.id = $1
)
ORDER BY et.id;


-- name: GetAutoAbilityShopIDsPre :many
SELECT DISTINCT sh.id
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
WHERE
    (
        sqlc.narg('availability')::availability_type[] IS NULL
        OR
        sh.availability = ANY(sqlc.narg('availability')::availability_type[])
    )
    AND se.shop_type = 'pre-airship'
    AND j.auto_ability_id = $1
ORDER BY sh.id;



-- name: GetAutoAbilityShopIDsPost :many
SELECT DISTINCT sh.id
FROM shops sh
JOIN shop_equipment_pieces se ON se.shop_id = sh.id
JOIN j_shop_equipment_abilities j ON j.shop_equipment_id = se.id
WHERE
    (
        sqlc.narg('availability')::availability_type[] IS NULL
        OR
        sh.availability = ANY(sqlc.narg('availability')::availability_type[])
    )
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
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN equipment_drops ed ON ed.auto_ability_id = aa.id
JOIN j_monster_equipment_abilities j ON j.equipment_drop_id = ed.id
JOIN monster_equipment me ON j.monster_equipment_id = me.id
JOIN monsters m ON me.monster_id = m.id
WHERE m.id = $1
ORDER BY aa.id;


-- name: GetAutoAbilityIDsByMonsterItems :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN item_amounts ia1 ON aa.required_item_amount_id = ia1.id
JOIN master_items mit ON ia1.master_item_id = mit.id
JOIN item_amounts ia2 ON ia2.master_item_id = mit.id
WHERE EXISTS (
    SELECT 1
    FROM monster_items mi
    JOIN monsters m ON mi.monster_id = m.id
    WHERE ia2.id IN (
        mi.steal_common_id,
        mi.steal_rare_id,
        mi.drop_common_id,
        mi.drop_rare_id,
        mi.secondary_drop_common_id,
        mi.secondary_drop_rare_id,
        mi.bribe_id
    )
    AND m.id = $1
)
OR EXISTS (
    SELECT 1
    FROM possible_items pi
    JOIN j_monster_items_other_items jmio ON pi.id = jmio.possible_item_id
    JOIN monster_items mi ON jmio.monster_items_id = mi.id
    JOIN monsters m ON mi.monster_id = m.id
    WHERE pi.item_amount_id = ia2.id
      AND m.id = $1
)
ORDER BY aa.id;



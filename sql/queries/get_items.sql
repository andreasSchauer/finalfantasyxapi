-- name: GetItemIDs :many
SELECT id FROM items ORDER BY id;


-- name: GetItemMonsterIDs :many
SELECT DISTINCT m.id
FROM monsters m
JOIN monster_items mi ON mi.monster_id = m.id
LEFT JOIN j_monster_items_other_items jmio
  ON jmio.monster_items_id = mi.id
LEFT JOIN possible_items pi
  ON pi.id = jmio.possible_item_id
JOIN item_amounts ia
  ON ia.id IN (
      mi.steal_common_id,
      mi.steal_rare_id,
      mi.drop_common_id,
      mi.drop_rare_id,
      mi.secondary_drop_common_id,
      mi.secondary_drop_rare_id,
      mi.bribe_id,
      pi.item_amount_id
  )
JOIN master_items mit ON ia.master_item_id = mit.id
JOIN items i ON i.master_item_id = mit.id
WHERE
    i.id = sqlc.arg(item_id)
    AND (sqlc.narg('story_based')::BOOLEAN IS NULL OR m.is_story_based = sqlc.narg('story_based')::BOOLEAN)
    AND (sqlc.narg('repeatable')::BOOLEAN IS NULL OR m.is_repeatable = sqlc.narg('repeatable')::BOOLEAN)
    AND (
        sqlc.narg('post_airship')::boolean IS NULL

        OR (
            sqlc.narg('post_airship')::boolean = TRUE -- at least one post_airship = true
            AND NOT EXISTS (
                SELECT 1
                FROM monster_amounts ma2
                JOIN j_monster_selections_monsters j2 ON j2.monster_amount_id = ma2.id
                JOIN monster_selections ms2 ON ms2.id = j2.monster_selection_id
                JOIN monster_formations mf2 ON mf2.monster_selection_id = ms2.id
                JOIN formation_data fd2 ON fd2.id = mf2.formation_data_id
                WHERE
                    ma2.monster_id = m.id
                    AND fd2.is_post_airship = FALSE
            )  
        )

        OR (
            sqlc.narg('post_airship')::boolean = FALSE -- all post_airship = false
            AND EXISTS (
                SELECT 1
                FROM monster_amounts ma2
                JOIN j_monster_selections_monsters j2 ON j2.monster_amount_id = ma2.id
                JOIN monster_selections ms2 ON ms2.id = j2.monster_selection_id
                JOIN monster_formations mf2 ON mf2.monster_selection_id = ms2.id
                JOIN formation_data fd2 ON fd2.id = mf2.formation_data_id
                WHERE
                    ma2.monster_id = m.id
                    AND fd2.is_post_airship = FALSE
            )
        )
    )
ORDER BY m.id;


-- name: GetItemTreasureIDs :many
SELECT DISTINCT t.id
FROM treasures t
JOIN j_treasures_items j ON j.treasure_id = t.id
JOIN item_amounts ia ON j.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY t.id;


-- name: GetItemShopIDs :many
SELECT DISTINCT s.id
FROM shops s
JOIN j_shops_items j ON j.shop_id = s.id
JOIN shop_items si ON j.shop_item_id = si.id
JOIN items i ON si.item_id = i.id
WHERE i.id = $1
ORDER BY s.id;


-- name: GetItemQuestIDs :many
SELECT DISTINCT q.id
FROM quests q
JOIN quest_completions qc ON q.completion_id = qc.id
JOIN item_amounts ia ON qc.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY q.id;


-- name: GetItemBlitzballPrizeIDs :many
SELECT DISTINCT bp.id
FROM blitzball_positions bp
JOIN blitzball_items bi ON bi.position_id = bp.id
JOIN possible_items pi ON bi.possible_item_id = pi.id
JOIN item_amounts ia ON pi.item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY bp.id;


-- name: GetItemPlayerAbilityIDs :many
SELECT DISTINCT pa.id
FROM player_abilities pa
JOIN item_amounts ia ON pa.aeon_learn_item_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY pa.id;


-- name: GetItemAutoAbilityIDs :many
SELECT DISTINCT aa.id
FROM auto_abilities aa
JOIN item_amounts ia ON aa.required_item_amount_id = ia.id
JOIN master_items mi ON ia.master_item_id = mi.id
JOIN items i ON i.master_item_id = mi.id
WHERE i.id = $1
ORDER BY aa.id;


-- name: GetItemMixIDs :many
SELECT DISTINCT m.id
FROM mixes m
JOIN mix_combinations mc ON mc.mix_id = m.id
JOIN items i ON mc.first_item_id = i.id OR mc.second_item_id = i.id
WHERE i.id = $1
ORDER BY m.id;
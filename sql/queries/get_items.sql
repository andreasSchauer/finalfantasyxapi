-- name: GetItemIDs :many
SELECT id FROM items ORDER BY id;


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
JOIN quest_completions qc ON qc.quest_id = q.id
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
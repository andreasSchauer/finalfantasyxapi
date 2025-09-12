-- name: CreateCharacter :exec
INSERT INTO characters (data_hash, name, weapon_type, armor_type, physical_attack_range, can_fight_underwater)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAeon :exec
INSERT INTO aeons (data_hash, name, unlock_condition, category, is_optional, battles_to_regenerate, phys_atk_damage_constant, phys_atk_range, phys_atk_shatter_rate, phys_atk_acc_source, phys_atk_hit_chance, phys_atk_acc_modifier)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateDefaultAbilitesEntry :exec
INSERT INTO default_abilities (data_hash, name)
VALUES ($1, $2)
ON CONFLICT(data_hash) DO NOTHING;
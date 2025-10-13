-- name: CreateCelestialWeapon :one
INSERT INTO celestial_weapons (data_hash, name, key_item_base, formula)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO UPDATE SET data_hash = celestial_weapons.data_hash
RETURNING *;


-- name: UpdateCelestialWeapon :exec
UPDATE celestial_weapons
SET data_hash = $1,
    name = $2,
    key_item_base = $3,
    formula = $4,
    character_id = $5,
    aeon_id = $6
WHERE id = $7;


-- name: CreateAutoAbility :exec
INSERT INTO auto_abilities (data_hash, name, description, effect, type, category, ability_value, activation_condition, counter, gradual_recovery, on_hit_element)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateEquipmentAbility :exec
INSERT INTO equipment_abilities (data_hash, type, classification, specific_character_id, version, priority, pool_1_amt, pool_2_amt, empty_slots_amt)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(data_hash) DO NOTHING;



-- name: CreateCelestialWeapon :exec
INSERT INTO celestial_weapons (data_hash, name, key_item_base, formula)
VALUES ($1, $2, $3, $4)
ON CONFLICT(data_hash) DO NOTHING;


-- name: CreateAutoAbility :exec
INSERT INTO auto_abilities (data_hash, name, description, effect, type, category, ability_value, activation_condition, counter, gradual_recovery, on_hit_element, conversion_from, conversion_to)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
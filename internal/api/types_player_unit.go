package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type PlayerUnit struct {
	ID               int32              `json:"id"`
	Name             string             `json:"name"`
	Type             database.UnitType  `json:"type"`
	TypedUnit        NamedAPIResource   `json:"typed_unit"`
	Area             AreaAPIResource    `json:"area"`
	CelestialWeapon  *NamedAPIResource  `json:"celestial_weapon"`
	CharacterClasses []NamedAPIResource `json:"character_classes"`
}

package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerUnit struct {
	ID               int32                          `json:"id"`
	Name             string                         `json:"name"`
	Type             database.UnitType              `json:"type"`
	TypedUnit        apiresource.NamedAPIResource   `json:"typed_unit"`
	Area             apiresource.AreaAPIResource    `json:"area"`
	CelestialWeapon  *apiresource.NamedAPIResource  `json:"celestial_weapon"`
	CharacterClasses []apiresource.NamedAPIResource `json:"character_classes"`
}

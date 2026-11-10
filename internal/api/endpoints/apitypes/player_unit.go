package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerUnit struct {
	ID               int32               `json:"id"`
	Name             string              `json:"name"`
	Type             database.UnitType   `json:"type"`
	TypedUnit        apiresource.Named   `json:"typed_unit"`
	Area             apiresource.Area    `json:"area"`
	CelestialWeapon  *apiresource.Named  `json:"celestial_weapon"`
	CharacterClasses []apiresource.Named `json:"character_classes"`
}

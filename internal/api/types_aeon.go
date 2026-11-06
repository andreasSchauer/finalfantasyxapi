package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertAeonEquipment(ae seedtypes.AeonEquipment) apitypes.AeonEquipment {
	return apitypes.AeonEquipment{
		AutoAbility:     nameToNamedAPIResource(E.AutoAbilities, ae.AutoAbility, nil),
		CelestialWeapon: ae.CelestialWeapon,
	}
}

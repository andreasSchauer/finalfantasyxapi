package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertAeonEquipment(ae seedtypes.AeonEquipment) apitypes.AeonEquipment {
	return apitypes.AeonEquipment{
		AutoAbility:     handlerinput.NameToNamedAPIResource(apilookups.E.AutoAbilities, ae.AutoAbility, nil),
		CelestialWeapon: ae.CelestialWeapon,
	}
}

package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func AeonEquipment(ae seedtypes.AeonEquipment) apitypes.AeonEquipment {
	return apitypes.AeonEquipment{
		AutoAbility:     handlerinput.NameToNamedAPIResource(apilookups.E.AutoAbilities, ae.AutoAbility, nil),
		CelestialWeapon: ae.CelestialWeapon,
	}
}

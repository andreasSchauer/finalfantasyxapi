package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func AbilityPool(p seedtypes.AbilityPool) apitypes.AbilityPool {
	return apitypes.AbilityPool{
		AutoAbilities: handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func EquipmentName(n seedtypes.EquipmentName) apiresource.Named {
	return handlerinput.NameToNamedAPIResource(apilookups.E.Equipment, n.Name, nil)
}

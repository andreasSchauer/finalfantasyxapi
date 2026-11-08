package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertAbilityPool(p seedtypes.AbilityPool) apitypes.AbilityPool {
	return apitypes.AbilityPool{
		AutoAbilities: handlerinput.NamesToNamedAPIResources(E.AutoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func convertEquipmentName(n seedtypes.EquipmentName) apiresource.Named {
	return handlerinput.NameToNamedAPIResource(E.Equipment, n.Name, nil)
}

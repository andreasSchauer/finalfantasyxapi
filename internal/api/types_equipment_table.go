package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertAbilityPool(p seedtypes.AbilityPool) apitypes.AbilityPool {
	return apitypes.AbilityPool{
		AutoAbilities: namesToNamedAPIResources(E.AutoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func convertEquipmentName(n seedtypes.EquipmentName) apiresource.NamedAPIResource {
	return nameToNamedAPIResource(E.Equipment, n.Name, nil)
}

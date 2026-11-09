package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertAbilityPool(p seedtypes.AbilityPool) apitypes.AbilityPool {
	return apitypes.AbilityPool{
		AutoAbilities: handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func ConvertEquipmentName(n seedtypes.EquipmentName) apiresource.Named {
	return handlerinput.NameToNamedAPIResource(apilookups.E.Equipment, n.Name, nil)
}

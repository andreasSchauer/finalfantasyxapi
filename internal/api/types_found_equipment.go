package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertFoundEquipment(fe seedtypes.TreasureEquipment) apitypes.FoundEquipment {
	return apitypes.FoundEquipment{
		EquipmentName:    handlerinput.NameToNamedAPIResource(E.Equipment, fe.Name, nil),
		Abilities:        handlerinput.NamesToNamedAPIResources(E.AutoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
}

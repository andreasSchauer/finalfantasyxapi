package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertFoundEquipment(fe seedtypes.TreasureEquipment) apitypes.FoundEquipment {
	return apitypes.FoundEquipment{
		EquipmentName:    handlerinput.NameToNamedAPIResource(apilookups.E.Equipment, fe.Name, nil),
		Abilities:        handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
}

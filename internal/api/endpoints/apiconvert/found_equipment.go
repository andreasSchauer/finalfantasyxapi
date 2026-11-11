package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func FoundEquipment(fe seedtypes.TreasureEquipment) apitypes.FoundEquipment {
	return apitypes.FoundEquipment{
		EquipmentName:    handlerinput.NameToNamedAPIResource(apilookups.E.Equipment, fe.Name, nil),
		Abilities:        handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
}

package dropchance

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func getWantedEquipment(params servicetypes.DropChanceParams, wantedAbilities []seedtypes.AutoAbility, slotsTable []seedtypes.EquipmentSlotsChance) servicetypes.WantedEquipment {
	requiredSlots := params.GetRequiredSlots()

	we := servicetypes.WantedEquipment{
		MinSlots:         slotsTable[0].Amount,
		MaxSlots:         slotsTable[len(slotsTable)-1].Amount,
		ReqAbilities:     getReqAbilities(wantedAbilities),
		ReqEmptySlots:    params.MinEmptySlots,
		OptionalSlotsAre: getOptionalSlotDef(params),
	}

	we.MinOptionalSlots = we.MinSlots - requiredSlots
	we.MaxOptionalSlots = we.MaxSlots - requiredSlots

	return we
}

func getReqAbilities(wantedAbilities []seedtypes.AutoAbility) []string {
	reqAbilities := []string{}

	for _, ability := range wantedAbilities {
		reqAbilities = append(reqAbilities, ability.Name)
	}

	return reqAbilities
}

func getOptionalSlotDef(params servicetypes.DropChanceParams) string {
	if params.OptionalAbilities {
		return string(servicetypes.OptionalSlotAbilityOrEmpty)
	}

	return string(servicetypes.OptionalSlotAllEmpty)
}

package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type WantedEquipment struct {
	MinSlots         int32    `json:"min_slots"`
	MaxSlots         int32    `json:"max_slots"`
	ReqAbilities     []string `json:"req_abilities"`
	ReqEmptySlots    *int32   `json:"req_empty_slots"`
	MinOptionalSlots int32    `json:"min_optional_slots"`
	MaxOptionalSlots int32    `json:"max_optional_slots"`
	OptionalSlotsAre string   `json:"optional_slots_are"`
}

type OptionalSlotDef string

const (
	OptionalSlotAllEmpty       OptionalSlotDef = "all-empty"
	OptionalSlotAbilityOrEmpty OptionalSlotDef = "ability-or-empty"
)

func getWantedEquipment(params DropChanceParams, wantedAbilities []seeding.AutoAbility, slotsTable []seeding.EquipmentSlotsChance) WantedEquipment {
	requiredSlots := params.GetRequiredSlots()

	we := WantedEquipment{
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

func getReqAbilities(wantedAbilities []seeding.AutoAbility) []string {
	reqAbilities := []string{}

	for _, ability := range wantedAbilities {
		reqAbilities = append(reqAbilities, ability.Name)
	}

	return reqAbilities
}

func getOptionalSlotDef(params DropChanceParams) string {
	if params.OptionalAbilities {
		return string(OptionalSlotAbilityOrEmpty)
	}

	return string(OptionalSlotAllEmpty)
}

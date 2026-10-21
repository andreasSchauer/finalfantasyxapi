package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getDropChanceParamsSamples() []SampleParams[DropChanceParams] {
	return []SampleParams[DropChanceParams]{
		{
			Description: "Calculate the chance of Iron Giant dropping an armor for Kimahri with only 'SOS Regen' on it. If other slots exist, they must be empty.",
			Params: DropChanceParams{
				Monster: 75,
				AutoAbilities: []int32{97},
				CharacterID: h.GetInt32Ptr(5),
			},
		},
		{
			Description: "Calculate the chance of Iron Giant dropping an armor for Kimahri with 'SOS Regen' on it. The armor can also bear other auto-abilities.",
			Params: DropChanceParams{
				Monster: 75,
				AutoAbilities: []int32{97},
				LenientAbilities: true,
				CharacterID: h.GetInt32Ptr(5),
			},
		},
		{
			Description: "Calculate the chance of Iron Giant dropping an armor with 'SOS Regen' on it and at least 1 empty slot. The armor can also bear other auto-abilities, if it has 3 slots (the maximum slots for this monster).",
			Params: DropChanceParams{
				Monster: 75,
				AutoAbilities: []int32{97},
				MinEmptySlots: h.GetInt32Ptr(1),
				LenientAbilities: true,
			},
		},
		{
			Description: "Calculate the chance of Iron Giant dropping an armor with 'SOS Regen' on it and at least 2 empty slots. The armor could also bear other auto-abilities, but since 3 slots is the maximum in this specific case, only armor with SOS Regen and 2 empty slots will pass. In this case 'lenient_abilities' won't have any effect.",
			Params: DropChanceParams{
				Monster: 75,
				AutoAbilities: []int32{97},
				MinEmptySlots: h.GetInt32Ptr(2),
				LenientAbilities: true,
			},
		},
		{
			Description: "Calculate the chance of Iron Giant dropping an armor with 3 slots and with 'SOS Regen' on it. The armor can also bear other auto-abilities.",
			Params: DropChanceParams{
				Monster: 75,
				AutoAbilities: []int32{97},
				TotalSlots: h.GetInt32Ptr(3),
				LenientAbilities: true,
			},
		},
		{
			Description: "Calculate the chance of Dark Yojimbo dropping an armor with 4 slots and with only 'Break HP Limit' and 'Ribbon' on it. The armor must have exactly these two auto-abilities.",
			Params: DropChanceParams{
				Monster: 299,
				AutoAbilities: []int32{128,129},
				TotalSlots: h.GetInt32Ptr(4),
				LenientAbilities: false,
			},
		},
		{
			Description: "Calculate the chance of YAT-97 dropping a weapon with at least 3 empty slots for any character (meaning 3 or 4 empty slots).",
			Params: DropChanceParams{
				Monster: 191,
				MinEmptySlots: h.GetInt32Ptr(3),
				EquipType: h.GetStrPtr(string(database.EquipTypeWeapon)),
			},
		},
		{
			Description: "Calculate the chance of YAT-97 dropping a weapon with 4 empty slots for any character.",
			Params: DropChanceParams{
				Monster: 191,
				MinEmptySlots: h.GetInt32Ptr(4),
				EquipType: h.GetStrPtr(string(database.EquipTypeWeapon)),
			},
		},
		{
			Description: "Another, more implicit method to calculate the chance of YAT-97 dropping a weapon with 4 empty slots for any character. That is, because no auto-abilities were given and 'lenient_abilities' being false enforces zero abilities.",
			Params: DropChanceParams{
				Monster: 191,
				TotalSlots: h.GetInt32Ptr(4),
				LenientAbilities: false,
				EquipType: h.GetStrPtr(string(database.EquipTypeWeapon)),
			},
		},
	}
}
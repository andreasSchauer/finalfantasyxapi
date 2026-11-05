package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type MonsterItemsSimple struct {
	StealCommon         *string  `json:"steal_common,omitempty"`
	StealRare           *string  `json:"steal_rare,omitempty"`
	DropCommon          *string  `json:"drop_common,omitempty"`
	DropRare            *string  `json:"drop_rare,omitempty"`
	SecondaryDropCommon *string  `json:"secondary_drop_common,omitempty"`
	SecondaryDropRare   *string  `json:"secondary_drop_rare,omitempty"`
	Bribe               *string  `json:"bribe,omitempty"`
	OtherItems          []string `json:"other_items,omitempty"`
}

func convertMonsterItemsSimple(items seedtypes.MonsterItems) MonsterItemsSimple {
	return MonsterItemsSimple{
		StealCommon:         ah.ConvertObjPtr(items.StealCommon, convertItemAmountSimple),
		StealRare:           ah.ConvertObjPtr(items.StealRare, convertItemAmountSimple),
		DropCommon:          ah.ConvertObjPtr(items.DropCommon, convertItemAmountSimple),
		DropRare:            ah.ConvertObjPtr(items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: ah.ConvertObjPtr(items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   ah.ConvertObjPtr(items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               ah.ConvertObjPtr(items.Bribe, convertItemAmountSimple),
		OtherItems:          ah.ConvertObjSliceOrNil(items.OtherItems, posItemToItemAmtSimple),
	}
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterEquipmentSimple(equipment seedtypes.MonsterEquipment) MonsterEquipmentSimple {
	return MonsterEquipmentSimple{
		WeaponAbilities: ah.ConvertObjSlice(equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  ah.ConvertObjSlice(equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}

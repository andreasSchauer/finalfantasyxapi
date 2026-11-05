package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

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
		StealCommon:         convertObjPtr(items.StealCommon, convertItemAmountSimple),
		StealRare:           convertObjPtr(items.StealRare, convertItemAmountSimple),
		DropCommon:          convertObjPtr(items.DropCommon, convertItemAmountSimple),
		DropRare:            convertObjPtr(items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: convertObjPtr(items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   convertObjPtr(items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               convertObjPtr(items.Bribe, convertItemAmountSimple),
		OtherItems:          convertObjSliceOrNil(items.OtherItems, posItemToItemAmtSimple),
	}
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterEquipmentSimple(equipment seedtypes.MonsterEquipment) MonsterEquipmentSimple {
	return MonsterEquipmentSimple{
		WeaponAbilities: convertObjSlice(equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  convertObjSlice(equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}

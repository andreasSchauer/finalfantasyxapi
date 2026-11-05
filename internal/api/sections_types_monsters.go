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
		StealCommon:         ConvertObjPtr(items.StealCommon, convertItemAmountSimple),
		StealRare:           ConvertObjPtr(items.StealRare, convertItemAmountSimple),
		DropCommon:          ConvertObjPtr(items.DropCommon, convertItemAmountSimple),
		DropRare:            ConvertObjPtr(items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: ConvertObjPtr(items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   ConvertObjPtr(items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               ConvertObjPtr(items.Bribe, convertItemAmountSimple),
		OtherItems:          ConvertObjSliceOrNil(items.OtherItems, posItemToItemAmtSimple),
	}
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterEquipmentSimple(equipment seedtypes.MonsterEquipment) MonsterEquipmentSimple {
	return MonsterEquipmentSimple{
		WeaponAbilities: ConvertObjSlice(equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  ConvertObjSlice(equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}

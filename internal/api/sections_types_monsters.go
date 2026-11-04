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

func convertMonsterItemsSimple(server *Server, items seedtypes.MonsterItems) MonsterItemsSimple {
	return MonsterItemsSimple{
		StealCommon:         convertObjPtr(server, items.StealCommon, convertItemAmountSimple),
		StealRare:           convertObjPtr(server, items.StealRare, convertItemAmountSimple),
		DropCommon:          convertObjPtr(server, items.DropCommon, convertItemAmountSimple),
		DropRare:            convertObjPtr(server, items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: convertObjPtr(server, items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   convertObjPtr(server, items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               convertObjPtr(server, items.Bribe, convertItemAmountSimple),
		OtherItems:          convertObjSliceOrNil(server, items.OtherItems, posItemToItemAmtSimple),
	}
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterEquipmentSimple(server *Server, equipment seedtypes.MonsterEquipment) MonsterEquipmentSimple {
	return MonsterEquipmentSimple{
		WeaponAbilities: convertObjSlice(server, equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  convertObjSlice(server, equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}

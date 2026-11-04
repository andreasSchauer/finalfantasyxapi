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

func convertMonsterItemsSimple(srv *Server, items seedtypes.MonsterItems) MonsterItemsSimple {
	return MonsterItemsSimple{
		StealCommon:         convertObjPtr(srv, items.StealCommon, convertItemAmountSimple),
		StealRare:           convertObjPtr(srv, items.StealRare, convertItemAmountSimple),
		DropCommon:          convertObjPtr(srv, items.DropCommon, convertItemAmountSimple),
		DropRare:            convertObjPtr(srv, items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: convertObjPtr(srv, items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   convertObjPtr(srv, items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               convertObjPtr(srv, items.Bribe, convertItemAmountSimple),
		OtherItems:          convertObjSliceOrNil(srv, items.OtherItems, posItemToItemAmtSimple),
	}
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterEquipmentSimple(srv *Server, equipment seedtypes.MonsterEquipment) MonsterEquipmentSimple {
	return MonsterEquipmentSimple{
		WeaponAbilities: convertObjSlice(srv, equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  convertObjSlice(srv, equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}
